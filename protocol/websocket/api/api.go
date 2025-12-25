package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.protocol.websocket"
)

type IPlugin interface {
	xplugin.IPlugin
}
