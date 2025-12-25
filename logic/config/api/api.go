package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.logic.config"
)

type IPlugin interface {
	xplugin.IPlugin
}
