# loop.go 冲突详细分析

## 概述

loop.go 有 11 处冲突，涉及核心架构差异：
- **HEAD (feat/subturn-poc)**: 基于 context 的 SubTurn 层级管理，使用 `activeTurnStates` map 支持并发
- **Incoming (refactor/agent)**: 事件驱动架构，使用 `EventBus`、`HookManager`，单个 `activeTurn` **不支持并发 turn**

## 关键发现：Incoming 的并发限制

**重要**: Incoming 分支的 `activeTurn` 设计**不支持并发 turn 执行**！

```go
// Incoming 的实现
func (al *AgentLoop) runTurn(ctx context.Context, ts *turnState) (turnResult, error) {
    al.registerActiveTurn(ts)      // 设置 al.activeTurn = ts
    defer al.clearActiveTurn(ts)   // 清除 al.activeTurn = nil
    // ...
}

func (al *AgentLoop) registerActiveTurn(ts *turnState) {
    al.activeTurnMu.Lock()
    defer al.activeTurnMu.Unlock()
    al.activeTurn = ts  // 单例！后面的会覆盖前面的
}
```

**问题**:
1. 如果两个 session 同时调用 `runAgentLoop`，第二个会覆盖第一个的 `activeTurn`
2. `GetActiveTurn()` 只能返回最后一个注册的 turn
3. 中断操作 (`InterruptGraceful`, `InterruptHard`) 只能影响当前的 `activeTurn`

**HEAD 的优势**:
```go
// HEAD 的实现
activeTurnStates sync.Map  // 支持多个并发 turn
// key: sessionKey, value: *turnState

// 每个 session 有独立的 turnState
al.activeTurnStates.Store(opts.SessionKey, rootTS)
```

## 架构决策的影响

如果采用 Incoming 的架构（方案 B），我们会**失去并发 turn 的能力**！

### 选项分析

**选项 1: 完全采用 Incoming（会失去并发）**
- ✅ 获得事件驱动架构
- ✅ 获得 Hook 系统
- ❌ **失去并发 turn 支持**
- ❌ **失去 SubTurn 并发支持**
- ❌ 多个 session 无法同时处理

**选项 2: 混合方案（推荐）**
- ✅ 保留 HEAD 的 `activeTurnStates sync.Map`
- ✅ 采用 Incoming 的 `EventBus` 和 `HookManager`
- ✅ 保持并发能力
- ⚠️ 需要调整 `GetActiveTurn()` 等 API

**选项 3: 改造 Incoming 支持并发**
- 将 `activeTurn *turnState` 改为 `activeTurns sync.Map`
- 修改所有相关方法支持 sessionKey 参数
- 工作量大，但架构更清晰

## 推荐方案：选项 2（混合方案）

### AgentLoop 结构体设计

```go
type AgentLoop struct {
    // Incoming 的字段
    bus            *bus.MessageBus
    cfg            *config.Config
    registry       *AgentRegistry
    state          *state.Manager
    eventBus       *EventBus        // ✅ 保留
    hooks          *HookManager     // ✅ 保留
    hookRuntime    hookRuntime      // ✅ 保留
    running        atomic.Bool
    summarizing    sync.Map
    fallback       *providers.FallbackChain
    channelManager *channels.Manager
    mediaStore     media.MediaStore
    transcriber    voice.Transcriber
    cmdRegistry    *commands.Registry
    mcp            mcpRuntime
    steering       *steeringQueue
    mu             sync.RWMutex

    // HEAD 的并发支持（保留）
    activeTurnStates sync.Map     // ✅ 保留：支持并发 turn
    subTurnCounter   atomic.Int64 // ✅ 保留：SubTurn ID 生成

    // Incoming 的字段（调整）
    turnSeq          atomic.Uint64 // ✅ 保留：全局 turn 序列号
    activeRequests   sync.WaitGroup // ✅ 保留：请求跟踪

    reloadFunc       func() error
}
```

### 关键方法调整

1. **GetActiveTurn()**: 需要接受 sessionKey 参数
2. **InterruptGraceful/Hard()**: 需要接受 sessionKey 参数
3. **runAgentLoop()**: 使用 `activeTurnStates` 而不是单个 `activeTurn`

## 冲突详情

### 冲突 1: AgentLoop 结构体 (38-77 行)

**HEAD 新增字段**:
```go
activeTurnStates sync.Map     // key: sessionKey (string), value: *turnState
subTurnCounter   atomic.Int64 // Counter for generating unique SubTurn IDs
```

**Incoming 新增字段**:
```go
eventBus       *EventBus
hooks          *HookManager
hookRuntime    hookRuntime
activeTurnMu   sync.RWMutex
activeTurn     *turnState
turnSeq        atomic.Uint64
activeRequests sync.WaitGroup
```

**关键差异**:
- HEAD: 使用 `sync.Map` 管理多个并发 turn (`activeTurnStates`)
- Incoming: 使用单个 `activeTurn` + 锁 (`activeTurnMu`)
- HEAD: SubTurn 计数器 (`subTurnCounter`)
- Incoming: Turn 序列号 (`turnSeq`)
- Incoming: 新增事件系统 (`eventBus`, `hooks`, `hookRuntime`)

**解决方案**: 采用 Incoming 的结构，但需要考虑如何在新架构中实现 SubTurn 的并发管理。

---

### 冲突 2: processOptions 结构体 (92-112 行)

**HEAD**:
```go
SkipAddUserMessage      bool     // If true, skip adding UserMessage to session history
```

**Incoming**:
```go
InitialSteeringMessages []providers.Message

// 新增结构体
type continuationTarget struct {
	SessionKey string
	Channel    string
	ChatID     string
}
```

**关键差异**:
- HEAD: 使用 `SkipAddUserMessage` 标志
- Incoming: 使用 `InitialSteeringMessages` 数组 + 新的 `continuationTarget` 结构体

**解决方案**: 采用 Incoming 的实现，`InitialSteeringMessages` 提供更灵活的 steering 消息处理。

---

### 冲突 3: runAgentLoop 函数 (1439-1581 行)

这是最大的冲突，涉及核心执行逻辑。

**HEAD 的实现**:
1. 检查是否在 SubTurn 中 (`turnStateFromContext`)
2. 如果是 SubTurn，复用现有 turnState
3. 如果是根 turn，创建新的 rootTS
4. 使用 `activeTurnStates.Store` 注册 turn
5. 调用 `runLLMIteration` 执行 LLM 循环

**Incoming 的实现**:
1. 记录 last channel
2. 调用 `newTurnState` 创建 turn state
3. 调用 `al.runTurn(ctx, ts)` 执行 turn
4. 处理 follow-up 消息
5. 发布响应

**关键差异**:
- HEAD: 复杂的 SubTurn 层级管理，支持嵌套
- Incoming: 简化的 turn 管理，通过 `newTurnState` 和 `runTurn`
- HEAD: 使用 `runLLMIteration` 函数
- Incoming: 使用 `runTurn` 函数
- Incoming: 新增 follow-up 消息处理机制

**解决方案**: 采用 Incoming 的简化架构，但需要在 `runTurn` 中添加 SubTurn 支持。

---

### 冲突 4: runLLMIteration vs runTurn (1672-1689 行)

**HEAD**: 有独立的 `runLLMIteration` 函数
**Incoming**: 使用 `runTurn` 函数

需要查看具体实现来决定如何合并。

---

### 冲突 5-11: 其他冲突点

剩余冲突主要涉及：
- 工具执行逻辑
- Steering 消息处理
- 中断处理
- 变量命名差异（`agent` vs `ts.agent`）

## 架构决策

根据方案 B（采用重构架构），需要：

1. **采用 Incoming 的 AgentLoop 结构**
   - 使用 `eventBus`, `hooks`, `hookRuntime`
   - 使用单个 `activeTurn` + `activeTurnMu`
   - 保留 `turnSeq`

2. **SubTurn 支持策略**
   - 选项 A: 在 `turnState` 中添加父子关系字段
   - 选项 B: 使用 context 传递 SubTurn 信息
   - 选项 C: 在 EventBus 中管理 SubTurn 层级

3. **函数迁移顺序**
   - 先采用 Incoming 的结构体定义
   - 更新 `newTurnState` 函数
   - 采用 `runTurn` 函数
   - 在 `runTurn` 中集成 SubTurn 逻辑

## 推荐实施步骤

### 步骤 1: 结构体定义 (30 分钟)
- 采用 Incoming 的 `AgentLoop` 结构体
- 采用 Incoming 的 `processOptions` 结构体
- 添加 `continuationTarget` 结构体

### 步骤 2: 辅助函数 (30 分钟)
- 更新 `NewAgentLoop` 初始化函数
- 确保 EventBus、Hook 正确初始化

### 步骤 3: runAgentLoop 函数 (1-2 小时)
- 采用 Incoming 的简化实现
- 保留 channel 记录逻辑
- 调用 `newTurnState` 和 `runTurn`
- 处理 follow-up 消息

### 步骤 4: runTurn 函数 (2-3 小时)
- 采用 Incoming 的 `runTurn` 实现
- 在其中添加 SubTurn 检测和处理逻辑
- 集成 SubTurn 结果回传机制

### 步骤 5: 其他冲突点 (1-2 小时)
- 逐个解决剩余 7 个冲突
- 确保变量命名一致
- 更新工具执行和 steering 逻辑

## 风险和注意事项

1. **SubTurn 语义变化**: 新架构中 SubTurn 的实现方式可能不同
2. **并发安全**: 从 `sync.Map` 迁移到单个 `activeTurn` + 锁
3. **事件系统集成**: 需要确保 SubTurn 事件正确触发
4. **测试覆盖**: 原有 SubTurn 测试需要更新

## 下一步

建议先实现步骤 1-2（结构体定义和初始化），然后再处理复杂的执行逻辑。
