package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/protocol/report/api"
	websocketApi "github.com/zzy-rabbit/patrol/protocol/websocket/api"
	logApi "github.com/zzy-rabbit/patrol/utils/log/api"
)

type service struct {
	ILogger    logApi.IPlugin       `xplugin:"patrol.utils.log"`
	IWebsocket websocketApi.IPlugin `xplugin:"patrol.protocol.websocket"`
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
