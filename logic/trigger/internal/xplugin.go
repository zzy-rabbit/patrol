package internal

import (
	"context"
	configApi "github.com/zzy-rabbit/patrol/logic/config/api"
	executorApi "github.com/zzy-rabbit/patrol/logic/executor/api"
	"github.com/zzy-rabbit/patrol/logic/trigger/api"
	logApi "github.com/zzy-rabbit/bp/log/api"
)

type service struct {
	IConfig   configApi.IPlugin   `xplugin:"patrol.logic.config"`
	IExecutor executorApi.IPlugin `xplugin:"patrol.logic.executor"`
	ILogger   logApi.IPlugin      `xplugin:"bp.log"`
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
