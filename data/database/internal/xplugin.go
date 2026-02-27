package internal

import (
	"context"
	"encoding/json"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	daoApi "github.com/zzy-rabbit/patrol/data/dao/api"
	"github.com/zzy-rabbit/patrol/data/database/api"
	"github.com/zzy-rabbit/xtools/xerror"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type service struct {
	config      api.Config
	ILogger     logApi.IPlugin `xplugin:"bp.tool.log"`
	IDao        daoApi.IPlugin `xplugin:"patrol.data.dao"`
	mutex       sync.RWMutex
	databaseMap map[string]daoApi.IDatabase
}

func New(ctx context.Context) api.IPlugin {
	return &service{
		databaseMap: make(map[string]daoApi.IDatabase),
	}
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
	s.ILogger.Info(ctx, "plugin %s init by config %+v success", s.GetName(ctx), s.config)
	return nil
}

func (s *service) Run(ctx context.Context, runParam string) error {
	err := filepath.Walk(s.config.Path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".db" {
			return nil
		}
		if xerror.Error(err) {
			s.ILogger.Error(ctx, "filepath walk %s fail %v", path, err)
			return err
		}
		filename := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		_, err = s.New(ctx, filename)
		if xerror.Error(err) {
			s.ILogger.Error(ctx, "new database connect %s fail %v", path, err)
			return err
		}
		return nil
	})
	if xerror.Error(err) {
		s.ILogger.Fatal(ctx, "filepath walk database fail %v", err)
		return err
	}
	s.ILogger.Info(ctx, "plugin %s run success", s.GetName(ctx))
	return nil
}

func (s *service) Stop(ctx context.Context, stopParam string) error {
	s.ILogger.Info(ctx, "plugin %s stop success", s.GetName(ctx))
	return nil
}
