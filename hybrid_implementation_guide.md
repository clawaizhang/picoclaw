# 混合方案落地指南

## 目标

结合 Incoming 的事件驱动架构和 HEAD 的并发能力，实现：
- ✅ 保留 `activeTurnStates` map（支持并发 Session）
- ✅ 采用 `EventBus` 和 `HookManager`（事件驱动 + 扩展性）
- ✅ 保留 SubTurn 并发支持
- ✅ 统一使用 `runTurn` 函数（简化代码）

---

## 实施步骤

### 步骤 1: 合并 AgentLoop 结构体 (30 分钟)

**目标**: 结合两边的字段

```go
type AgentLoop struct {
    // ===== Incoming 的字段 (保留) =====
    bus            *bus.MessageBus
    cfg            *config.Config
    registry       *AgentRegistry
    state          *state.Manager
    eventBus       *EventBus        // ✅ 新增：事件系统
    hooks          *HookManager     // ✅ 新增：Hook 系统
    running        atomic.Bool
    summarizing    sync.Map
    fallback       *providers.FallbackChain
    channelManager *channels.Manager
    mediaStore     media.MediaStore
    transcriber    voice.Transcriber
    cmdRegistry    *commands.Registry
    mcp            mcpRuntime
    hookRuntime    hookRuntime      // ✅ 新增：Hook 运行时
    steering       *steeringQueue
    mu             sync.RWMutex

    // ===== HEAD 的字段 (保留) =====
    activeTurnStates sync.Map       // ✅ 保留：支持并发 Session
    subTurnCounter   atomic.Int64   // ✅ 保留：SubTurn ID 生成

    // ===== Incoming 的字段 (调整) =====
    turnSeq          atomic.Uint64  // ✅ 保留：全局 Turn 序列号
    activeRequests   sync.WaitGroup // ✅ 保留：请求跟踪

    reloadFunc       func() error
}
```

**操作**:
1. 找到 AgentLoop 结构体定义（38-77 行的冲突）
2. 采用上面的合并版本
3. 删除 Incoming 的 `activeTurn *turnState` 和 `activeTurnMu`（不需要了）

---

### 步骤 2: 合并 processOptions 结构体 (10 分钟)

**目标**: 采用 Incoming 的版本，移除 HEAD 的 `SkipAddUserMessage`

```go
type processOptions struct {
    SessionKey              string
    Channel                 string
    ChatID                  string
    SenderID                string
    SenderDisplayName       string
    UserMessage             string
    SystemPromptOverride    string
    Media                   []string
    InitialSteeringMessages []providers.Message  // ✅ Incoming 的方式
    DefaultResponse         string
    EnableSummary           bool
    SendResponse            bool
    NoHistory               bool
    SkipInitialSteeringPoll bool
}

type continuationTarget struct {
    SessionKey string
    Channel    string
    ChatID     string
}
```

**操作**:
1. 找到 processOptions 结构体（92-112 行的冲突）
2. 采用上面的版本
3. 添加 `continuationTarget` 结构体

---

### 步骤 3: 更新 turnState 结构体 (20 分钟)

**目标**: 在 Incoming 的 turnState 基础上添加 SubTurn 支持

需要检查 `turn.go` 或 `turn_state.go` 文件，确保 turnState 有这些字段：

```go
type turnState struct {
    mu sync.RWMutex

    // ===== Incoming 的字段 (保留) =====
    agent *AgentInstance
    opts  processOptions
    scope turnEventScope

    turnID     string
    agentID    string
    sessionKey string
    channel    string
    chatID     string
    userMessage string
    media       []string

    phase        TurnPhase
    iteration    int
    startedAt    time.Time
    finalContent string
    followUps    []bus.InboundMessage

    gracefulInterrupt     bool
    gracefulInterruptHint string
    gracefulTerminalUsed  bool
    hardAbort             bool
    providerCancel        context.CancelFunc
    turnCancel            context.CancelFunc

    restorePointHistory []providers.Message
    restorePointSummary string
    persistedMessages   []providers.Message

    // ===== HEAD 的字段 (新增：SubTurn 支持) =====
    depth                int                      // ✅ SubTurn 深度
    parentTurnID         string                   // ✅ 父 Turn ID
    childTurnIDs         []string                 // ✅ 子 Turn IDs
    pendingResults       chan *tools.ToolResult   // ✅ SubTurn 结果 channel
    concurrencySem       chan struct{}            // ✅ 并发信号量
    isFinished           atomic.Bool              // ✅ 是否已完成
}
```

**操作**:
1. 查找 `turnState` 结构体定义
2. 如果有冲突，采用 Incoming 的基础版本
3. 添加 SubTurn 相关字段（depth, parentTurnID 等）

---

### 步骤 4: 重写 runAgentLoop 函数 (1 小时)

**目标**: 简化为调用 runTurn，但保留 SubTurn 检测

```go
func (al *AgentLoop) runAgentLoop(
    ctx context.Context,
    agent *AgentInstance,
    opts processOptions,
) (string, error) {
    // 1. 检查是否在 SubTurn 中
    existingTS := turnStateFromContext(ctx)
    var ts *turnState
    var isRootTurn bool

    if existingTS != nil {
        // 在 SubTurn 中 - 创建子 turnState
        ts = newSubTurnState(agent, opts, existingTS, al.newTurnEventScope(agent.ID, opts.SessionKey))
        isRootTurn = false
    } else {
        // 根 Turn - 创建新的 turnState
        ts = newTurnState(agent, opts, al.newTurnEventScope(agent.ID, opts.SessionKey))
        isRootTurn = true

        // 注册到 activeTurnStates（支持并发）
        al.activeTurnStates.Store(opts.SessionKey, ts)
        defer al.activeTurnStates.Delete(opts.SessionKey)
    }

    // 2. 记录 last channel
    if opts.Channel != "" && opts.ChatID != "" && !constants.IsInternalChannel(opts.Channel) {
        channelKey := fmt.Sprintf("%s:%s", opts.Channel, opts.ChatID)
        if err := al.RecordLastChannel(channelKey); err != nil {
            logger.WarnCF("agent", "Failed to record last channel",
                map[string]any{"error": err.Error()})
        }
    }

    // 3. 执行 Turn（带事件和 Hook）
    result, err := al.runTurn(ctx, ts)
    if err != nil {
        return "", err
    }
    if result.status == TurnEndStatusAborted {
        return "", nil
    }

    // 4. 处理 SubTurn 结果（仅根 Turn）
    if isRootTurn && ts.pendingResults != nil {
        finalResults := al.drainPendingSubTurnResults(ts)
        for _, r := range finalResults {
            if r != nil && r.ForLLM != "" {
                result.finalContent += fmt.Sprintf("\n\n[SubTurn Result] %s", r.ForLLM)
            }
        }
    }

    // 5. 处理 follow-up 消息
    for _, followUp := range result.followUps {
        if pubErr := al.bus.PublishInbound(ctx, followUp); pubErr != nil {
            logger.WarnCF("agent", "Failed to publish follow-up after turn",
                map[string]any{"turn_id": ts.turnID, "error": pubErr.Error()})
        }
    }

    // 6. 发送响应
    if opts.SendResponse && result.finalContent != "" {
        al.bus.PublishOutbound(ctx, bus.OutboundMessage{
            Channel: opts.Channel,
            ChatID:  opts.ChatID,
            Content: result.finalContent,
        })
    }

    return result.finalContent, nil
}
```

**操作**:
1. 找到 runAgentLoop 函数（1439-1581 行的冲突）
2. 替换为上面的简化版本
3. 保留 SubTurn 检测逻辑（`turnStateFromContext`）
4. 保留 `activeTurnStates` 注册逻辑

---

### 步骤 5: 采用 Incoming 的 runTurn 函数 (30 分钟)

**目标**: 使用 Incoming 的 runTurn，但添加 SubTurn 结果轮询

```go
func (al *AgentLoop) runTurn(ctx context.Context, ts *turnState) (turnResult, error) {
    turnCtx, turnCancel := context.WithCancel(ctx)
    defer turnCancel()
    ts.setTurnCancel(turnCancel)

    // ===== 不使用单例 activeTurn，因为我们有 activeTurnStates =====
    // al.registerActiveTurn(ts)  ← 删除这行
    // defer al.clearActiveTurn(ts) ← 删除这行

    turnStatus := TurnEndStatusCompleted
    defer func() {
        al.emitEvent(
            EventKindTurnEnd,
            ts.eventMeta("runTurn", "turn.end"),
            TurnEndPayload{
                Status:          turnStatus,
                Iterations:      ts.currentIteration(),
                Duration:        time.Since(ts.startedAt),
                FinalContentLen: ts.finalContentLen(),
            },
        )
    }()

    al.emitEvent(
        EventKindTurnStart,
        ts.eventMeta("runTurn", "turn.start"),
        TurnStartPayload{
            Channel:     ts.channel,
            ChatID:      ts.chatID,
            UserMessage: ts.userMessage,
            MediaCount:  len(ts.media),
        },
    )

    // ... 保留 Incoming 的其余逻辑 ...

    // ===== 在 Turn Loop 中添加 SubTurn 结果轮询 =====
turnLoop:
    for ts.currentIteration() < ts.agent.MaxIterations || len(pendingMessages) > 0 {
        // ... LLM 调用 ...
        // ... Tool 执行 ...

        // ✅ 新增：轮询 SubTurn 结果
        if ts.pendingResults != nil {
            subTurnResults := al.pollSubTurnResults(ts)
            for _, result := range subTurnResults {
                if result.ForLLM != "" {
                    // 将 SubTurn 结果作为 steering message 注入
                    pendingMessages = append(pendingMessages, providers.Message{
                        Role:    "user",
                        Content: fmt.Sprintf("[SubTurn Result] %s", result.ForLLM),
                    })
                }
            }
        }

        // ... 继续迭代 ...
    }

    // ... 返回结果 ...
}
```

**操作**:
1. 找到 runTurn 函数（1672-1689 行开始的冲突）
2. 采用 Incoming 的完整实现
3. 删除 `registerActiveTurn` 和 `clearActiveTurn` 调用
4. 在 Turn Loop 中添加 SubTurn 结果轮询逻辑

---

### 步骤 6: 实现辅助函数 (30 分钟)

需要实现以下辅助函数：

#### 6.1 newSubTurnState
```go
func newSubTurnState(
    agent *AgentInstance,
    opts processOptions,
    parent *turnState,
    scope turnEventScope,
) *turnState {
    ts := newTurnState(agent, opts, scope)

    // 设置 SubTurn 关系
    ts.depth = parent.depth + 1
    ts.parentTurnID = parent.turnID
    ts.pendingResults = parent.pendingResults  // 共享结果 channel
    ts.concurrencySem = parent.concurrencySem  // 共享信号量

    // 记录父子关系
    parent.mu.Lock()
    parent.childTurnIDs = append(parent.childTurnIDs, ts.turnID)
    parent.mu.Unlock()

    return ts
}
```

#### 6.2 pollSubTurnResults
```go
func (al *AgentLoop) pollSubTurnResults(ts *turnState) []*tools.ToolResult {
    if ts.pendingResults == nil {
        return nil
    }

    var results []*tools.ToolResult
    for {
        select {
        case result := <-ts.pendingResults:
            results = append(results, result)
        default:
            return results
        }
    }
}
```

#### 6.3 drainPendingSubTurnResults
```go
func (al *AgentLoop) drainPendingSubTurnResults(ts *turnState) []*tools.ToolResult {
    if ts.pendingResults == nil {
        return nil
    }

    // 等待一小段时间，确保所有 SubTurn 结果都到达
    time.Sleep(100 * time.Millisecond)

    return al.pollSubTurnResults(ts)
}
```

#### 6.4 更新 GetActiveTurn
```go
func (al *AgentLoop) GetActiveTurn(sessionKey string) *ActiveTurnInfo {
    val, ok := al.activeTurnStates.Load(sessionKey)
    if !ok {
        return nil
    }

    ts, ok := val.(*turnState)
    if !ok {
        return nil
    }

    info := ts.snapshot()
    return &info
}
```

---

### 步骤 7: 更新 SpawnSubTurn 实现 (30 分钟)

确保 spawn tool 能正确创建 SubTurn：

```go
func (spawner *subTurnSpawner) SpawnSubTurn(
    ctx context.Context,
    config SubTurnConfig,
) (*tools.ToolResult, error) {
    // 1. 获取父 turnState
    parentTS := turnStateFromContext(ctx)
    if parentTS == nil {
        return nil, fmt.Errorf("no parent turn state in context")
    }

    // 2. 检查深度限制
    maxDepth := spawner.loop.getSubTurnConfig().maxDepth
    if parentTS.depth >= maxDepth {
        return tools.ErrorResult(fmt.Sprintf(
            "SubTurn depth limit reached (%d)", maxDepth)), nil
    }

    // 3. 获取并发信号量
    select {
    case <-parentTS.concurrencySem:
        defer func() { parentTS.concurrencySem <- struct{}{} }()
    case <-ctx.Done():
        return tools.ErrorResult("SubTurn cancelled"), nil
    }

    // 4. 生成 SubTurn ID
    subTurnID := spawner.loop.subTurnCounter.Add(1)
    turnID := fmt.Sprintf("%s-sub-%d", parentTS.turnID, subTurnID)

    // 5. 创建 SubTurn context
    subCtx := withTurnState(ctx, parentTS)  // 继承父 context

    // 6. 启动 SubTurn goroutine
    go func() {
        opts := processOptions{
            SessionKey:           parentTS.sessionKey,
            Channel:              parentTS.channel,
            ChatID:               parentTS.chatID,
            UserMessage:          config.SystemPrompt,
            SystemPromptOverride: config.SystemPrompt,
            NoHistory:            true,  // SubTurn 不加载历史
            SendResponse:         false, // SubTurn 不发送响应
        }

        result, err := spawner.loop.runAgentLoop(subCtx, spawner.agent, opts)

        // 7. 发送结果到父 Turn
        toolResult := &tools.ToolResult{
            ForLLM: result,
            Error:  err,
        }

        select {
        case parentTS.pendingResults <- toolResult:
        case <-subCtx.Done():
        }
    }()

    // 8. 立即返回（异步执行）
    return tools.AsyncResult(fmt.Sprintf("SubTurn %d started", subTurnID)), nil
}
```

---

### 步骤 8: 解决其他小冲突 (1 小时)

处理剩余的 7 个冲突点：

1. **变量命名冲突** (2179-2183 行等)
   - 统一使用 `ts.channel`, `ts.chatID` 而不是 `opts.Channel`

2. **Tool feedback** (2469-2494 行)
   - 采用 HEAD 的实现（发送 tool feedback 到 chat）

3. **其他小差异**
   - 逐个检查，优先采用 Incoming 的实现
   - 确保 EventBus 事件正确触发

---

## 验证步骤

### 1. 编译验证
```bash
go build ./pkg/agent/
```

### 2. 单元测试
```bash
go test ./pkg/agent/ -v
```

### 3. 功能测试

创建测试用例验证：

```go
func TestMixedArchitecture_ConcurrentSessions(t *testing.T) {
    // 测试多个 session 并发执行
    var wg sync.WaitGroup
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            sessionKey := fmt.Sprintf("session-%d", id)
            // 执行 agent loop
        }(i)
    }
    wg.Wait()
}

func TestMixedArchitecture_SubTurnExecution(t *testing.T) {
    // 测试 SubTurn 执行
    // 1. 启动主 Turn
    // 2. 调用 spawn tool
    // 3. 验证 SubTurn 结果返回
}

func TestMixedArchitecture_EventBusIntegration(t *testing.T) {
    // 测试事件系统
    // 1. 订阅事件
    // 2. 执行 Turn
    // 3. 验证事件触发
}
```

---

## 预期结果

完成后，系统应该：

✅ 支持多个 Session 并发执行
✅ 支持 SubTurn 并发和嵌套
✅ 所有操作都触发 EventBus 事件
✅ Hook 系统正常工作
✅ 代码结构清晰，易于维护

---

## 时间估算

- 步骤 1-2: 结构体合并 (40 分钟)
- 步骤 3: turnState 更新 (20 分钟)
- 步骤 4: runAgentLoop 重写 (1 小时)
- 步骤 5: runTurn 调整 (30 分钟)
- 步骤 6: 辅助函数 (30 分钟)
- 步骤 7: SpawnSubTurn (30 分钟)
- 步骤 8: 其他冲突 (1 小时)
- 测试验证 (1 小时)

**总计: 约 5-6 小时**

---

## 风险和注意事项

1. **Context 传递**: 确保 SubTurn 的 context 正确继承父 context
2. **Channel 关闭**: 确保 `pendingResults` channel 在合适的时机关闭
3. **并发安全**: 所有对 turnState 的访问都要加锁
4. **事件顺序**: 确保事件按正确顺序触发
5. **测试覆盖**: 重点测试并发场景和 SubTurn 场景
