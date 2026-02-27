package internal

import (
	"context"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	"github.com/zzy-rabbit/patrol/data/dao/api"
)

type service struct {
	ILogger logApi.IPlugin `xplugin:"bp.tool.log"`
}

func New(ctx context.Context) api.IPlugin {
	return &service{}
}

func (s *service) GetName(ctx context.Context) string {
	return api.PluginName
}

func (s *service) Init(ctx context.Context, initParam string) error {
	//err := json.Unmarshal([]byte(initParam), &s.config)
	//if xerror.Error(err) {
	//	s.ILogger.Error(ctx, "plugin %s init error: %s", s.GetName(ctx), err.Error())
	//	return err
	//}
	s.ILogger.Info(ctx, "plugin %s init success", s.GetName(ctx))
	return nil
}

func (s *service) Run(ctx context.Context, runParam string) error {
	//db, err := gorm.Open(sqlite.Open(s.config.Sqlite.File), &gorm.Config{})
	//if err != nil {
	//	s.ILogger.Error(ctx, "plugin %s open database by config %+v fail %v", s.GetName(ctx), s.config, err)
	//	return err
	//}
	//s.db = db
	s.ILogger.Info(ctx, "plugin %s run success", s.GetName(ctx))
	return nil
}

func (s *service) Stop(ctx context.Context, stopParam string) error {
	s.ILogger.Info(ctx, "plugin %s stop success", s.GetName(ctx))
	return nil
}
