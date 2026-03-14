// Package ext 实现了第三方嵌入 Channel
package ext

import (
	"github.com/sipeed/picoclaw/pkg/bus"
	"github.com/sipeed/picoclaw/pkg/channels"
	"github.com/sipeed/picoclaw/pkg/config"
)

func init() {
	// 注册 ext channel 工厂
	channels.RegisterFactory("ext", func(cfg *config.Config, bus *bus.MessageBus) (channels.Channel, error) {
		return NewChannel(cfg.Channels.Ext, bus)
	})
}
