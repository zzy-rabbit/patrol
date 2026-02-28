package api

import (
	"context"
	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.logic.trigger"
)

type IPlugin interface {
	xplugin.IPlugin
	AddDepartment(ctx context.Context, department string) xerror.IError
	DeleteDepartment(ctx context.Context, department string) xerror.IError
}
