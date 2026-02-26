package main

import (
	"context"
	"flag"
	"os"

	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xplugin"

	_ "github.com/zzy-rabbit/bp/protocol/http"
	_ "github.com/zzy-rabbit/bp/protocol/report"
	_ "github.com/zzy-rabbit/bp/protocol/websocket"
	_ "github.com/zzy-rabbit/bp/tool/encrypt"
	_ "github.com/zzy-rabbit/bp/tool/log"
	_ "github.com/zzy-rabbit/bp/tool/uniform"

	_ "github.com/zzy-rabbit/patrol/data/dao"
	_ "github.com/zzy-rabbit/patrol/hardware/nfc"
	_ "github.com/zzy-rabbit/patrol/hardware/qrcode"
	_ "github.com/zzy-rabbit/patrol/logic/config"
	_ "github.com/zzy-rabbit/patrol/logic/executor"
	_ "github.com/zzy-rabbit/patrol/logic/trigger"
	_ "github.com/zzy-rabbit/patrol/protocol/http"
)

func ParseStartParams(ctx context.Context) xerror.IError {
	var configPath string
	flag.StringVar(&configPath, "config", "config/plugin.json", "config file path")
	flag.Parse()

	content, err := os.ReadFile(configPath)
	if xerror.Error(err) {
		return xerror.Extend(xerror.ErrFileOperationFail, err.Error())
	}
	xerr := xplugin.ParseConfig(ctx, content)
	if xerror.Error(xerr) {
		return xerr
	}
	return nil
}
