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
	SyncDo(ctx context.Context, param model.ExecutorParams) (model.ExecuteResult, xerror.IError)
	ASyncDo(ctx context.Context, param model.ExecutorParams, after func(model.ExecuteResult, xerror.IError)) xerror.IError
}
