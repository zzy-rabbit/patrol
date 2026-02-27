package api

import (
	"context"
	daoApi "github.com/zzy-rabbit/patrol/data/dao/api"
	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.data.database"
)

type Config struct {
	Path string `json:"path"`
}

type IPlugin interface {
	xplugin.IPlugin
	New(ctx context.Context, unique string) (daoApi.IDatabase, xerror.IError)
	Get(ctx context.Context, unique string) (daoApi.IDatabase, bool)
	GetAll(ctx context.Context) []daoApi.IDatabase
	Delete(ctx context.Context, unique string) xerror.IError
}
