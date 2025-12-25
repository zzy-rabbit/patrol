package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.plugins.logic.config"
)

type IPlugin interface {
	xplugin.IPlugin
}
