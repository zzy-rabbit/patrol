package internal

import (
	"context"
	reportApi "github.com/zzy-rabbit/bp/protocol/report/api"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	daoApi "github.com/zzy-rabbit/patrol/data/dao/api"
	"github.com/zzy-rabbit/patrol/logic/config/api"
)

type service struct {
	IDao    daoApi.IPlugin    `xplugin:"patrol.data.dao"`
	IReport reportApi.IPlugin `xplugin:"bp.protocol.report"`
	ILogger logApi.IPlugin    `xplugin:"bp.tool.log"`
}

func New(ctx context.Context) api.IPlugin {
	return &service{}
}

func (s *service) GetName(ctx context.Context) string {
	return api.PluginName
}

func (s *service) Init(ctx context.Context, initParam string) error {
	s.ILogger.Info(ctx, "plugin %s init success", s.GetName(ctx))
	return nil
}

func (s *service) Run(ctx context.Context, runParam string) error {
	s.ILogger.Info(ctx, "plugin %s run success", s.GetName(ctx))
	return nil
}

func (s *service) Stop(ctx context.Context, stopParam string) error {
	s.ILogger.Info(ctx, "plugin %s stop success", s.GetName(ctx))
	return nil
}
