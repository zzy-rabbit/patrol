package internal

import (
	"context"
	bpModel "github.com/zzy-rabbit/bp/model"
	httpApi "github.com/zzy-rabbit/bp/protocol/http/api"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	configApi "github.com/zzy-rabbit/patrol/logic/config/api"
	"github.com/zzy-rabbit/patrol/protocol/http/api"
)

type service struct {
	network bpModel.Network
	IConfig configApi.IPlugin `xplugin:"patrol.logic.config"`
	ILogger logApi.IPlugin    `xplugin:"bp.tool.log"`
	IHttp   httpApi.IPlugin   `xplugin:"bp.protocol.http"`
}

func New(ctx context.Context) api.IPlugin {
	return &service{}
}

func (s *service) GetName(ctx context.Context) string {
	return api.PluginName
}

func (s *service) Init(ctx context.Context, initParam string) error {
	s.IHttp.Register(s.registerRouter)
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
