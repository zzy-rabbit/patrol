package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.logic.executor"
)

type IPlugin interface {
	xplugin.IPlugin
}
