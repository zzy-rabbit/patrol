package main

import (
	"github.com/zzy-rabbit/xtools/xcontext"
	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xplugin"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := xcontext.Background()

	err := ParseStartParams(ctx)
	if xerror.Error(err) {
		panic(err)
	}

	err = xplugin.Init(ctx)
	if xerror.Error(err) {
		panic(err)
	}

	err = xplugin.Run(ctx)
	if xerror.Error(err) {
		panic(err)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL)
	<-exit

	err = xplugin.Stop(ctx)
	if xerror.Error(err) {
		panic(err)
	}
}
