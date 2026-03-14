# 第三方嵌入通道 (Ext Channel)

Ext Channel 允许 PicoClaw 通过外部进程接入任意消息平台。

## 架构

```
PicoClaw (父进程)          外部进程 (胶水层)
     │                           │
     │── fork/exec ─────────────▶│
     │                           │
     │── JSON-RPC (stdin) ──────▶│  启动飞书 WS
     │◀──────── JSON-RPC ───────│  连接官方插件
     │                           │
     │◀── notify/inbound ───────│  收到消息
     │── send ─────────────────▶│  发送回复
```

## 协议

通信使用 **JSON-RPC over stdio**，每行一条 JSON 消息。

### 方法 (PicoClaw -> 子进程)

| 方法 | 说明 |
|------|------|
| `initialize` | 初始化，传递配置 |
| `start` | 开始运行 |
| `stop` | 停止运行 |
| `send` | 发送 AI 回复到平台 |
| `typing` | 输入指示器控制 |

### 通知 (子进程 -> PicoClaw)

| 方法 | 说明 |
|------|------|
| `inbound` | 收到平台消息 |
| `heartbeat` | 心跳 |
| `error` | 错误报告 |

### 数据结构

**初始化请求:**
```json
{
  "id": "1",
  "method": "initialize",
  "params": {
    "channel_id": "feishu-bot",
    "working_dir": "/var/lib/feishu"
  }
}
```

**初始化响应:**
```json
{
  "id": "1",
  "result": {
    "name": "Feishu Bot",
    "version": "1.0.0",
    "capabilities": {
      "send_message": true,
      "update_message": false,
      "typing_indicator": true,
      "receive_message": true
    }
  }
}
```

**入站消息 (子进程 -> PicoClaw):**
```json
{
  "method": "inbound",
  "params": {
    "message_id": "om_xxx",
    "chat_id": "oc_xxx",
    "chat_type": "group",
    "sender_id": "ou_xxx",
    "sender_name": "张三",
    "content": "你好",
    "was_mentioned": true,
    "context": {
      "thread_id": "omt_xxx",
      "reply_to": "om_yyy"
    }
  }
}
```

**发送消息 (PicoClaw -> 子进程):**
```json
{
  "id": "5",
  "method": "send",
  "params": {
    "chat_id": "oc_xxx",
    "content": "你好！有什么可以帮你的？",
    "context": {
      "thread_id": "omt_xxx",
      "reply_to": "om_yyy"
    }
  }
}
```

## 配置

在 `config.json` 中添加:

```json
{
  "channels": {
    "ext": {
      "enabled": true,
      "command": "/path/to/your-adapter",
      "your_config_key": "your_value",
      "another_config": "another_value"
    }
  }
}
```

**说明:**
- `enabled` - 是否启用（必须）
- `command` - 胶水层可执行文件路径（必须）
- **其他所有字段原样传递给子进程**，由子进程自己解析使用

PicoClaw 只负责启动 `command`，所有配置通过 `initialize` 方法传给子进程：

```json
{
  "method": "initialize",
  "params": {
    "channel_id": "ext",
    "config": {
      "enabled": true,
      "command": "/path/to/your-adapter",
      "your_config_key": "your_value",
      "another_config": "another_value"
    }
  }
}
```
```

## 开发胶水层

胶水层需要实现:

1. **读取 stdin** 接收 PicoClaw 的请求
2. **写入 stdout** 返回响应和通知
3. **处理以下方法:**
   - `initialize` - 返回能力声明（会收到 config 参数）
   - `start` - 启动平台连接
   - `stop` - 关闭连接
   - `send` - 发送消息到平台

4. **发送以下通知:**
   - `inbound` - 收到平台消息时

### 最小示例 (Go)

```go
package main

import (
    "bufio"
    "encoding/json"
    "os"
)

type Request struct {
    ID     string          `json:"id"`
    Method string          `json:"method"`
    Params json.RawMessage `json:"params"`
}

type Response struct {
    ID     string          `json:"id"`
    Result json.RawMessage `json:"result,omitempty"`
}

type Notification struct {
    Method string          `json:"method"`
    Params json.RawMessage `json:"params"`
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    encoder := json.NewEncoder(os.Stdout)

    for scanner.Scan() {
        var req Request
        json.Unmarshal(scanner.Bytes(), &req)

        switch req.Method {
        case "initialize":
            // 解析 config，提取你的配置
            var initParams struct {
                ChannelID string          `json:"channel_id"`
                Config    json.RawMessage `json:"config"`
            }
            json.Unmarshal(req.Params, &initParams)
            
            // 保存 config 供后续使用...
            saveConfig(initParams.Config)

            resp := Response{
                ID: req.ID,
                Result: mustJSON(map[string]interface{}{
                    "name":    "My Adapter",
                    "version": "1.0.0",
                    "capabilities": map[string]bool{
                        "send_message":     true,
                        "receive_message":  true,
                        "typing_indicator": false,
                    },
                }),
            }
            encoder.Encode(resp)

        case "start":
            // 启动平台连接...
            go runPlatformAdapter()

            encoder.Encode(Response{
                ID:     req.ID,
                Result: mustJSON(map[string]string{"status": "ok"}),
            })

        case "send":
            var params struct {
                ChatID  string `json:"chat_id"`
                Content string `json:"content"`
            }
            json.Unmarshal(req.Params, &params)
            
            // 发送到平台...
            sendToPlatform(params.ChatID, params.Content)

            encoder.Encode(Response{
                ID:     req.ID,
                Result: mustJSON(map[string]bool{"ok": true}),
            })
        }
    }
}

func runPlatformAdapter() {
    // 连接平台 API/WebSocket
    // 收到消息时发送通知:
    notif := Notification{
        Method: "inbound",
        Params: mustJSON(map[string]interface{}{
            "message_id": "msg_xxx",
            "chat_id":    "chat_xxx",
            "chat_type":  "group",
            "sender_id":  "user_xxx",
            "content":    "你好",
        }),
    }
    json.NewEncoder(os.Stdout).Encode(notif)
}

func mustJSON(v interface{}) json.RawMessage {
    b, _ := json.Marshal(v)
    return b
}
```

## 注意事项

1. **一行一条 JSON** - 使用 `\n` 分隔
2. **stderr 用于日志** - PicoClaw 会捕获并记录
3. **同步调用** - Request/Response 通过 `id` 匹配
4. **通知无响应** - Notification 不需要回复
5. **Context 透传** - `inbound` 的 `context` 会在 `send` 时原样返回
