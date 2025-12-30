package api

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.logic.executor"
)

type IPlugin interface {
	xplugin.IPlugin
	Do(ctx context.Context, param model.ExecutorParams) (model.ExecuteResult, xerror.IError)
}
