package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.protocol.report"
)

type IPlugin interface {
	xplugin.IPlugin
}
