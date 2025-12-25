package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.protocol.http"
)

type IPlugin interface {
	xplugin.IPlugin
}
