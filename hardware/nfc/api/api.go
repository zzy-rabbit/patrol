package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.hardware.nfc"
)

type IPlugin interface {
	xplugin.IPlugin
}
