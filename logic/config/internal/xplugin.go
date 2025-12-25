package internal

import (
	"context"
	daoApi "github.com/zzy-rabbit/patrol/data/dao/api"
	"github.com/zzy-rabbit/patrol/logic/config/api"
)

type service struct {
	IDao daoApi.IPlugin `xplugin:"patrol.data.dao"`
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
