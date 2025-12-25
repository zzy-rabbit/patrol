package api

import (
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.hardware.qrcode"
)

type IPlugin interface {
	xplugin.IPlugin
}
