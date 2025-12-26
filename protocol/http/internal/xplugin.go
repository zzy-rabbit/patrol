package internal

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	configApi "github.com/zzy-rabbit/patrol/logic/config/api"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/patrol/protocol/http/api"
)

type service struct {
	network  model.Network
	fiberApp *fiber.App
	IConfig  configApi.IPlugin `xplugin:"patrol.logic.config"`
}

func New(ctx context.Context) api.IPlugin {
	return &service{}
}

func (s *service) GetName(ctx context.Context) string {
	return api.PluginName
}

func (s *service) Init(ctx context.Context, initParam string) error {
	var network model.Network
	err := json.Unmarshal([]byte(initParam), &network)
	if err != nil {
		return err
	}
	s.fiberApp = fiber.New()
	s.network = network
	s.registerRouter()
	return nil
}

func (s *service) Run(ctx context.Context, runParam string) error {
	addr := fmt.Sprintf("%s:%d", s.network.Host, s.network.Port)
	err := s.fiberApp.Listen(addr)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Stop(ctx context.Context, stopParam string) error {
	err := s.fiberApp.Shutdown()
	if err != nil {
		return err
	}
	return nil
}
