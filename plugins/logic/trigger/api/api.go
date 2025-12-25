package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.plugins.logic.trigger"
)

type IPlugin interface {
	xplugin.IPlugin
}
