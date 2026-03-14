// Package ext 实现了第三方嵌入 Channel
package ext

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/sipeed/picoclaw/pkg/bus"
	"github.com/sipeed/picoclaw/pkg/channels"
	"github.com/sipeed/picoclaw/pkg/config"
	"github.com/sipeed/picoclaw/pkg/logger"
)

// Channel 实现 channels.Channel 接口
type Channel struct {
	*channels.BaseChannel

	channelID string
	command   string
	rawConfig json.RawMessage // 原始配置（传给子进程）

	// 子进程
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
	stderr io.ReadCloser

	// 通信
	encoder *json.Encoder
	decoder *json.Decoder
	writeMu sync.Mutex

	// 请求追踪
	pending   map[string]chan *response
	pendingMu sync.Mutex
	reqSeq    int64

	// 状态
	caps    capabilities
	running atomic.Bool

	// 上下文存储（用于回复时恢复）
	contexts   map[string]json.RawMessage // ChatID -> Context
	contextsMu sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewChannel 创建 Channel
func NewChannel(cfg config.ExtConfig, bus *bus.MessageBus) (*Channel, error) {
	if cfg.Command == "" {
		return nil, errors.New("ext channel requires command")
	}

	rawConfig, err := json.Marshal(cfg)
	if err != nil {
		return nil, fmt.Errorf("marshal ext config: %w", err)
	}

	base := channels.NewBaseChannel("ext", nil, bus, nil)

	ch := &Channel{
		BaseChannel: base,
		channelID:   "ext",
		command:     cfg.Command,
		rawConfig:   rawConfig,
		pending:     make(map[string]chan *response),
		contexts:    make(map[string]json.RawMessage),
	}

	ch.SetOwner(ch)
	return ch, nil
}

// Name 返回通道名称
func (c *Channel) Name() string {
	return c.channelID
}

// Start 启动 Channel
func (c *Channel) Start(ctx context.Context) error {
	logger.InfoCF("ext", "Starting ext channel", map[string]any{
		"channel_id": c.channelID,
		"command":    c.command,
	})

	c.ctx, c.cancel = context.WithCancel(ctx)

	// 启动子进程
	if err := c.startProcess(); err != nil {
		return fmt.Errorf("start process: %w", err)
	}

	// 启动读取协程
	c.wg.Add(1)
	go c.readLoop()

	// 初始化
	initParams := initParams{
		ChannelID: c.channelID,
		Config:    c.rawConfig,
		LogLevel:  logger.GetLevel().String(),
	}

	resp, err := c.call(c.ctx, methodInitialize, initParams)
	if err != nil {
		c.stopProcess()
		return fmt.Errorf("initialize: %w", err)
	}

	var initResult initResult
	if err := json.Unmarshal(resp.Result, &initResult); err != nil {
		c.stopProcess()
		return fmt.Errorf("parse init result: %w", err)
	}

	c.caps = initResult.Capabilities
	logger.InfoCF("ext", "Ext channel initialized", map[string]any{
		"channel_id":   c.channelID,
		"name":         initResult.Name,
		"version":      initResult.Version,
		"capabilities": c.caps,
	})

	// 启动
	if _, err := c.call(c.ctx, methodStart, nil); err != nil {
		c.stopProcess()
		return fmt.Errorf("start: %w", err)
	}

	c.running.Store(true)

	logger.InfoCF("ext", "Ext channel started", map[string]any{
		"channel_id": c.channelID,
	})

	return nil
}

// Stop 停止 Channel
func (c *Channel) Stop(ctx context.Context) error {
	logger.InfoCF("ext", "Stopping ext channel", map[string]any{
		"channel_id": c.channelID,
	})

	c.running.Store(false)

	// 1. 先尝试优雅关闭子进程（使用传入的 ctx）
	if c.cmd != nil && c.cmd.Process != nil {
		stopCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		_, _ = c.call(stopCtx, methodStop, nil)
	}

	// 2. 关闭内部资源
	if c.cancel != nil {
		c.cancel()
	}

	c.stopProcess()

	// 等待读取协程
	done := make(chan struct{})
	go func() {
		c.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		logger.WarnCF("ext", "Timeout waiting for stop", map[string]any{
			"channel_id": c.channelID,
		})
	}

	return nil
}

// Send 发送 AI 回复
func (c *Channel) Send(ctx context.Context, msg bus.OutboundMessage) error {
	if !c.running.Load() {
		return channels.ErrNotRunning
	}
	if !c.caps.SendMessage {
		return errors.New("channel does not support sending")
	}

	// 从 contexts map 恢复上下文
	c.contextsMu.RLock()
	context := c.contexts[msg.ChatID]
	c.contextsMu.RUnlock()

	params := outboundPayload{
		ChatID:  msg.ChatID,
		Content: msg.Content,
		Context: context,
	}

	_, err := c.call(ctx, methodSend, params)
	return err
}

// IsAllowed 始终返回 true
func (c *Channel) IsAllowed(senderID string) bool {
	return true
}

// IsAllowedSender 始终返回 true
func (c *Channel) IsAllowedSender(sender bus.SenderInfo) bool {
	return true
}

// ===== 内部方法 =====

func (c *Channel) startProcess() error {
	cmd := exec.CommandContext(c.ctx, c.command)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	c.cmd = cmd
	c.stdin = stdin
	c.stdout = stdout
	c.stderr = stderr
	c.encoder = json.NewEncoder(stdin)
	c.decoder = json.NewDecoder(bufio.NewReader(stdout))

	// 启动 stderr 日志
	go c.stderrLoop()

	return nil
}

func (c *Channel) stopProcess() {
	if c.cmd == nil || c.cmd.Process == nil {
		return
	}

	_ = c.cmd.Process.Signal(syscall.SIGTERM)
	time.Sleep(100 * time.Millisecond)

	if c.cmd.Process != nil {
		_ = c.cmd.Process.Kill()
	}

	_, _ = c.cmd.Process.Wait()
	c.cmd = nil
}

func (c *Channel) readLoop() {
	defer c.wg.Done()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
		}

		var rawMsg json.RawMessage
		if err := c.decoder.Decode(&rawMsg); err != nil {
			if c.ctx.Err() != nil {
				return
			}
			logger.ErrorCF("ext", "Decode error", map[string]any{
				"channel_id": c.channelID,
				"error":      err.Error(),
			})
			continue
		}

		// 尝试解析为 response
		var resp response
		if err := json.Unmarshal(rawMsg, &resp); err == nil && resp.ID != "" {
			c.handleResponse(&resp)
			continue
		}

		// 尝试解析为 notification
		var notif notification
		if err := json.Unmarshal(rawMsg, &notif); err == nil && notif.Method != "" {
			c.handleNotification(&notif)
			continue
		}

		logger.WarnCF("ext", "Unknown message", map[string]any{
			"channel_id": c.channelID,
			"msg":        string(rawMsg),
		})
	}
}

func (c *Channel) handleResponse(resp *response) {
	c.pendingMu.Lock()
	ch, ok := c.pending[resp.ID]
	if ok {
		delete(c.pending, resp.ID)
	}
	c.pendingMu.Unlock()

	if !ok {
		return
	}

	select {
	case ch <- resp:
	default:
	}
}

func (c *Channel) handleNotification(notif *notification) {
	switch notif.Method {
	case notifyInbound:
		c.handleInbound(notif.Params)
	case notifyHeartbeat:
		logger.DebugCF("ext", "Heartbeat", map[string]any{
			"channel_id": c.channelID,
		})
	case notifyError:
		var err rpcError
		json.Unmarshal(notif.Params, &err)
		logger.ErrorCF("ext", "Error from sub-process", map[string]any{
			"channel_id": c.channelID,
			"code":       err.Code,
			"message":    err.Message,
		})
	default:
		logger.WarnCF("ext", "Unknown notification", map[string]any{
			"channel_id": c.channelID,
			"method":     notif.Method,
		})
	}
}

func (c *Channel) handleInbound(params json.RawMessage) {
	var msg inboundPayload
	if err := json.Unmarshal(params, &msg); err != nil {
		logger.ErrorCF("ext", "Failed to parse inbound", map[string]any{
			"channel_id": c.channelID,
			"error":      err.Error(),
		})
		return
	}

	// 转换为 bus.InboundMessage
	sender := bus.SenderInfo{
		Platform:    c.channelID,
		PlatformID:  msg.SenderID,
		CanonicalID: fmt.Sprintf("%s:%s", c.channelID, msg.SenderID),
		DisplayName: msg.SenderName,
	}

	peer := bus.Peer{
		Kind: msg.ChatType,
		ID:   msg.ChatID,
	}

	// 保存 Context（使用 Lock）
	if len(msg.Context) > 0 {
		c.contextsMu.Lock()
		c.contexts[msg.ChatID] = msg.Context
		c.contextsMu.Unlock()
	}

	// 交给 PicoClaw 处理
	c.HandleMessage(c.ctx, peer, msg.MessageID, msg.SenderID, msg.ChatID,
		msg.Content, nil, nil, sender)
}

// call 发送同步请求
func (c *Channel) call(ctx context.Context, method string, params interface{}) (*response, error) {
	c.writeMu.Lock()
	defer c.writeMu.Unlock()

	// 生成请求 ID
	c.reqSeq++
	reqID := fmt.Sprintf("%s-%d", c.channelID, c.reqSeq)

	// 注册等待
	respCh := make(chan *response, 1)
	c.pendingMu.Lock()
	c.pending[reqID] = respCh
	c.pendingMu.Unlock()

	defer func() {
		c.pendingMu.Lock()
		delete(c.pending, reqID)
		c.pendingMu.Unlock()
	}()

	// 发送请求
	paramsJSON, _ := json.Marshal(params)
	req := request{
		ID:     reqID,
		Method: method,
		Params: paramsJSON,
	}

	if err := c.encoder.Encode(req); err != nil {
		return nil, err
	}

	// 等待响应
	select {
	case resp := <-respCh:
		if resp == nil {
			return nil, errors.New("nil response")
		}
		if resp.Error != nil {
			return nil, fmt.Errorf("rpc error %d: %s", resp.Error.Code, resp.Error.Message)
		}
		return resp, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-c.ctx.Done():
		return nil, errors.New("channel closed")
	}
}

func (c *Channel) stderrLoop() {
	scanner := bufio.NewScanner(c.stderr)
	for scanner.Scan() {
		line := scanner.Text()
		logger.InfoCF("ext-stderr", line, map[string]any{
			"channel_id": c.channelID,
		})
	}
}

// 确保实现 Channel 接口
var _ channels.Channel = (*Channel)(nil)
