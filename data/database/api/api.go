package api

import (
	"context"
	daoApi "github.com/zzy-rabbit/patrol/data/dao/api"
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
	New(ctx context.Context, unique string) (daoApi.IDatabase, error)
	Get(ctx context.Context, unique string) (daoApi.IDatabase, bool)
	Delete(ctx context.Context, unique string) error
}
