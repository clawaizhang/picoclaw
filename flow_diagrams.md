# Agent Loop 流程图对比

## 1. Incoming (refactor/agent) 流程

### 整体架构
```
User Message
    ↓
Message Bus (串行队列)
    ↓
processMessage()
    ↓
runAgentLoop()
    ↓
newTurnState() → 创建 turnState
    ↓
runTurn()
    ↓
registerActiveTurn(ts)  ← 设置 al.activeTurn = ts (单例)
    ↓
[Turn 执行循环]
    ↓
clearActiveTurn(ts)     ← 清除 al.activeTurn = nil
```

### runTurn() 详细流程
```
┌─────────────────────────────────────────┐
│ runTurn(ctx, turnState)                 │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 1. 注册 activeTurn (单例)               │
│    al.registerActiveTurn(ts)            │
│    defer al.clearActiveTurn(ts)         │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 2. 发送 TurnStart 事件                  │
│    al.emitEvent(EventKindTurnStart)     │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 3. 加载 Session History & Summary       │
│    history = Sessions.GetHistory()      │
│    summary = Sessions.GetSummary()      │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 4. 构建消息                             │
│    messages = BuildMessages(...)        │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 5. 检查 Context Budget                  │
│    if isOverContextBudget() {           │
│        forceCompression()               │
│        emitEvent(ContextCompress)       │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 6. 保存用户消息到 Session               │
│    Sessions.AddMessage("user", ...)     │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 7. Turn Loop (迭代执行)                 │
│    for iteration < MaxIterations {      │
│        ┌─────────────────────────────┐  │
│        │ 7.1 调用 LLM                │  │
│        │     callLLM()               │  │
│        │     emitEvent(LLMStart)     │  │
│        └─────────────────────────────┘  │
│                  ↓                       │
│        ┌─────────────────────────────┐  │
│        │ 7.2 处理 Tool Calls         │  │
│        │     for each toolCall {     │  │
│        │         emitEvent(ToolStart)│  │
│        │         executeTool()       │  │
│        │         emitEvent(ToolEnd)  │  │
│        │     }                       │  │
│        └─────────────────────────────┘  │
│                  ↓                       │
│        ┌─────────────────────────────┐  │
│        │ 7.3 检查中断                │  │
│        │     if gracefulInterrupt {  │  │
│        │         break               │  │
│        │     }                       │  │
│        └─────────────────────────────┘  │
│                  ↓                       │
│        ┌─────────────────────────────┐  │
│        │ 7.4 处理 Steering Messages  │  │
│        │     pollSteering()          │  │
│        └─────────────────────────────┘  │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 8. 保存最终响应到 Session               │
│    Sessions.AddMessage("assistant", ...) │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 9. 发送 TurnEnd 事件                    │
│    al.emitEvent(EventKindTurnEnd)       │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 10. 返回 turnResult                     │
│     {finalContent, status, followUps}   │
└─────────────────────────────────────────┘
```

### 关键特点
- ✅ **事件驱动**: 每个阶段都发送事件到 EventBus
- ✅ **Hook 集成**: 在 before_llm, after_llm, before_tool, after_tool 触发 Hook
- ✅ **单 Turn**: 使用 `activeTurn` 单例，同一时间只有一个 turn
- ❌ **无并发**: 不支持多个 session 同时执行 turn

---

## 2. HEAD (feat/subturn-poc) 流程

### 整体架构
```
User Message
    ↓
Message Bus
    ↓
processMessage()
    ↓
runAgentLoop()
    ↓
检查 Context 中是否有 turnState
    ├─ 有 → 复用 (SubTurn 场景)
    └─ 无 → 创建新的 rootTS
         ↓
         存储到 activeTurnStates[sessionKey]
         ↓
         runLLMIteration()
         ↓
         [并发 SubTurn 支持]
```

### runAgentLoop() 详细流程
```
┌─────────────────────────────────────────┐
│ runAgentLoop(ctx, agent, opts)          │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 1. 检查是否在 SubTurn 中                │
│    existingTS = turnStateFromContext()  │
│    if existingTS != nil {               │
│        rootTS = existingTS  (复用)      │
│        isRootTurn = false               │
│    } else {                             │
│        rootTS = new turnState           │
│        isRootTurn = true                │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 2. 注册 Turn State (支持并发)           │
│    if isRootTurn {                      │
│        al.activeTurnStates.Store(       │
│            sessionKey, rootTS)          │
│        defer activeTurnStates.Delete()  │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 3. 记录 Last Channel                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 4. 构建消息                             │
│    messages = BuildMessages(...)        │
│    messages = resolveMediaRefs(...)     │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 5. 覆盖 System Prompt (如果需要)        │
│    if opts.SystemPromptOverride != "" { │
│        // 用于 SubTurn 的特殊 prompt   │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 6. 保存用户消息                         │
│    if !opts.SkipAddUserMessage {        │
│        Sessions.AddMessage(...)         │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 7. 执行 LLM 迭代                        │
│    finalContent, iteration, err =       │
│        runLLMIteration(ctx, agent, ...) │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 8. 轮询 SubTurn 结果 (如果是根 turn)    │
│    if isRootTurn {                      │
│        results =                        │
│            dequeuePendingSubTurnResults()│
│        // 将结果注入到最终响应          │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 9. 处理空响应                           │
│    if finalContent == "" {              │
│        finalContent = DefaultResponse   │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 10. 保存助手响应                        │
│     Sessions.AddMessage("assistant"...) │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 11. 发送响应 (如果需要)                 │
│     if opts.SendResponse {              │
│         bus.PublishOutbound(...)        │
│     }                                   │
└─────────────────────────────────────────┘
```

### SubTurn 执行流程
```
┌─────────────────────────────────────────┐
│ Tool: spawn                             │
│   args: {task: "...", label: "..."}    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ SpawnTool.Execute()                     │
│   if spawner != nil {                   │
│       // 直接 SubTurn 路径             │
│   } else {                              │
│       // SubagentManager 路径          │
│   }                                     │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ spawner.SpawnSubTurn()                  │
│   ┌─────────────────────────────────┐   │
│   │ 1. 生成 SubTurn ID              │   │
│   │    subTurnID = atomic.Add()     │   │
│   └─────────────────────────────────┘   │
│              ↓                          │
│   ┌─────────────────────────────────┐   │
│   │ 2. 创建 SubTurn Context         │   │
│   │    subCtx = withTurnState(...)  │   │
│   │    // 继承父 turnState          │   │
│   └─────────────────────────────────┘   │
│              ↓                          │
│   ┌─────────────────────────────────┐   │
│   │ 3. 获取并发信号量               │   │
│   │    <-rootTS.concurrencySem      │   │
│   │    defer release                │   │
│   └─────────────────────────────────┘   │
│              ↓                          │
│   ┌─────────────────────────────────┐   │
│   │ 4. 启动 Goroutine               │   │
│   │    go func() {                  │   │
│   │        result = runAgentLoop(   │   │
│   │            subCtx, ...)         │   │
│   │        // 将结果发送到 channel  │   │
│   │        rootTS.pendingResults <- │   │
│   │    }()                          │   │
│   └─────────────────────────────────┘   │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 父 Turn 继续执行                        │
│   - 不等待 SubTurn 完成                 │
│   - SubTurn 异步执行                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 父 Turn 轮询 SubTurn 结果               │
│   results = dequeuePendingSubTurnResults│
│   for each result {                     │
│       // 注入到响应或下一次迭代         │
│   }                                     │
└─────────────────────────────────────────┘
```

### SubTurn 层级结构
```
Root Turn (Session A)
  ├─ turnState (depth=0)
  │   ├─ turnID: "session-a"
  │   ├─ pendingResults: chan
  │   └─ concurrencySem: chan (限制并发数)
  │
  ├─ SubTurn 1 (depth=1)
  │   ├─ turnState (继承父 context)
  │   ├─ parentTurnID: "session-a"
  │   └─ 独立的 goroutine
  │
  ├─ SubTurn 2 (depth=1)
  │   ├─ turnState (继承父 context)
  │   ├─ parentTurnID: "session-a"
  │   └─ 独立的 goroutine
  │
  └─ SubTurn 3 (depth=1)
      └─ SubTurn 3.1 (depth=2)  ← 嵌套 SubTurn
          └─ ...

Root Turn (Session B) - 并发执行
  ├─ turnState (depth=0)
  └─ ...
```

### 关键特点
- ✅ **并发支持**: `activeTurnStates` map 支持多个 session 并发
- ✅ **SubTurn 层级**: 通过 context 传递 turnState，支持嵌套
- ✅ **并发控制**: `concurrencySem` 限制 SubTurn 并发数
- ✅ **异步执行**: SubTurn 在独立 goroutine 中执行
- ✅ **结果回传**: 通过 `pendingResults` channel 传递结果
- ❌ **无事件系统**: 没有 EventBus 和 Hook 集成

---

## 3. 对比总结

| 特性 | Incoming (refactor/agent) | HEAD (feat/subturn-poc) |
|------|---------------------------|-------------------------|
| **并发模型** | 单 Turn (串行) | 多 Turn (并发) |
| **Turn 管理** | `activeTurn` (单例) | `activeTurnStates` (map) |
| **事件系统** | ✅ EventBus | ❌ 无 |
| **Hook 系统** | ✅ HookManager | ❌ 无 |
| **SubTurn** | ❓ 未实现或不同方式 | ✅ 完整实现 |
| **并发 Session** | ❌ 不支持 | ✅ 支持 |
| **嵌套 SubTurn** | ❌ 不支持 | ✅ 支持 |
| **架构复杂度** | 简单 | 复杂 |
| **可扩展性** | 高 (Hook) | 低 |
| **调试难度** | 低 | 高 (并发) |

---

## 4. 混合方案流程

结合两者优点的混合方案：

```
┌─────────────────────────────────────────┐
│ runAgentLoop(ctx, agent, opts)          │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 1. 检查 SubTurn Context                 │
│    existingTS = turnStateFromContext()  │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 2. 创建/复用 turnState                  │
│    ts = newTurnState(agent, opts, ...)  │
│    if isRootTurn {                      │
│        activeTurnStates.Store(key, ts)  │
│    }                                    │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 3. 执行 Turn (带事件和 Hook)            │
│    result = runTurn(ctx, ts)            │
│    ├─ emitEvent(TurnStart)              │
│    ├─ Hook: before_llm                  │
│    ├─ callLLM()                         │
│    ├─ Hook: after_llm                   │
│    ├─ Hook: before_tool                 │
│    ├─ executeTool()                     │
│    │   └─ 如果是 spawn → SpawnSubTurn   │
│    ├─ Hook: after_tool                  │
│    └─ emitEvent(TurnEnd)                │
└─────────────────────────────────────────┘
              ↓
┌─────────────────────────────────────────┐
│ 4. 处理 SubTurn 结果                    │
│    if isRootTurn {                      │
│        pollSubTurnResults()             │
│    }                                    │
└─────────────────────────────────────────┘
```

### 混合方案优势
- ✅ 保留并发能力 (`activeTurnStates`)
- ✅ 获得事件系统 (`EventBus`)
- ✅ 获得扩展能力 (`HookManager`)
- ✅ 支持 SubTurn 并发
- ✅ 支持多 Session 并发
