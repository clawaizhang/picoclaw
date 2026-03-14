// Package ext 定义了第三方嵌入 Channel 的通信协议
package ext

import "encoding/json"

// request PicoClaw 向子进程发送的请求
type request struct {
	ID     string          `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

// response 子进程返回的响应
type response struct {
	ID     string          `json:"id"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *rpcError       `json:"error,omitempty"`
}

// rpcError 错误信息
type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// notification 子进程主动向 PicoClaw 发送的通知
type notification struct {
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

// 方法名常量
const (
	// PicoClaw -> 子进程
	methodInitialize = "initialize"
	methodStart      = "start"
	methodStop       = "stop"
	methodSend       = "send"
	methodTyping     = "typing"

	// 子进程 -> PicoClaw
	notifyInbound   = "inbound"
	notifyHeartbeat = "heartbeat"
	notifyError     = "error"
)

// initParams 初始化参数
type initParams struct {
	ChannelID string          `json:"channel_id"` // Channel 标识
	Config    json.RawMessage `json:"config"`     // 原始配置（原样传递）
	LogLevel  string          `json:"log_level"`  // 日志级别
}

// initResult 初始化结果
type initResult struct {
	Name         string       `json:"name"`         // 通道名称
	Version      string       `json:"version"`      // 版本
	Capabilities capabilities `json:"capabilities"` // 能力声明
}

// capabilities 子进程能力
type capabilities struct {
	SendMessage     bool `json:"send_message"`     // 能发送消息
	UpdateMessage   bool `json:"update_message"`   // 能更新消息（流式）
	TypingIndicator bool `json:"typing_indicator"` // 能显示输入中
	ReceiveMessage  bool `json:"receive_message"`  // 能接收消息
}

// inboundPayload 入站消息（子进程 -> PicoClaw）
type inboundPayload struct {
	// 消息标识
	MessageID string `json:"message_id"` // 平台消息ID
	ChatID    string `json:"chat_id"`    // 平台会话ID
	ChatType  string `json:"chat_type"`  // "direct" | "group"

	// 发送者
	SenderID   string `json:"sender_id"`
	SenderName string `json:"sender_name,omitempty"`

	// 内容
	Content      string `json:"content"`
	WasMentioned bool   `json:"was_mentioned,omitempty"` // 是否被@

	// 上下文（透传，回复时原样返回）
	Context json.RawMessage `json:"context,omitempty"`
}

// outboundPayload 出站消息（PicoClaw -> 子进程）
type outboundPayload struct {
	ChatID  string          `json:"chat_id"`
	Content string          `json:"content"`          // AI 回复内容
	Context json.RawMessage `json:"context,omitempty"` // 原始上下文
}

// typingPayload 输入指示器
type typingPayload struct {
	ChatID string `json:"chat_id"`
	Action string `json:"action"` // "start" | "stop"
}

// heartbeatPayload 心跳
type heartbeatPayload struct {
	Status string `json:"status"` // "ok" | "error"
}
