package internal

import (
	"context"
	configApi "github.com/zzy-rabbit/patrol/plugins/logic/config/api"
	"github.com/zzy-rabbit/patrol/protocol/http/api"
)

type service struct {
	IConfig configApi.IPlugin `xplugin:"patrol.plugins.logic.config"`
}

func New(ctx context.Context) api.IPlugin {
	return &service{}
}

func (s *service) GetName(ctx context.Context) string {
	return api.PluginName
}

func (s *service) Init(ctx context.Context, initParam string) error {
	return nil
}

func (s *service) Run(ctx context.Context, runParam string) error {
	return nil
}

func (s *service) Stop(ctx context.Context, stopParam string) error {
	return nil
}
