package api

import (
	"context"
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.plugins.data.dao"
)

type ISession interface {
}

type ITransaction interface {
	ISession
	Commit()
	Rollback()
}

type IPlugin interface {
	xplugin.IPlugin
	GetDB(ctx context.Context) ISession
	GetTransaction(ctx context.Context) ITransaction
}
