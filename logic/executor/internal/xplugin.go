package internal

import (
	"context"
	"encoding/json"
	configApi "github.com/zzy-rabbit/patrol/logic/config/api"
	"github.com/zzy-rabbit/patrol/logic/executor/api"
	logApi "github.com/zzy-rabbit/bp/log/api"
	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xthread"
)

type service struct {
	config      api.Config
	IConfig     configApi.IPlugin `xplugin:"patrol.logic.config"`
	ILogger     logApi.IPlugin    `xplugin:"bp.log"`
	IThreadPool xthread.IPool
}

func New(ctx context.Context) api.IPlugin {
	return &service{}
}

func (s *service) GetName(ctx context.Context) string {
	return api.PluginName
}

func (s *service) Init(ctx context.Context, initParam string) error {
	err := json.Unmarshal([]byte(initParam), &s.config)
	if xerror.Error(err) {
		s.ILogger.Error(ctx, "plugin %s init error: %s", s.GetName(ctx), err.Error())
		return err
	}
	s.ILogger.Info(ctx, "plugin %s init success", s.GetName(ctx))
	return nil
}

func (s *service) Run(ctx context.Context, runParam string) error {
	s.IThreadPool = xthread.New(ctx, s.config.CoreSize, s.config.TaskSize)
	s.ILogger.Info(ctx, "plugin %s run success", s.GetName(ctx))
	return nil
}

func (s *service) Stop(ctx context.Context, stopParam string) error {
	s.IThreadPool.Close(ctx)
	s.ILogger.Info(ctx, "plugin %s stop success", s.GetName(ctx))
	return nil
}
