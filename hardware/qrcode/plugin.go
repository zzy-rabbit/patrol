package qrcode

import (
	"github.com/zzy-rabbit/patrol/hardware/qrcode/internal"
	"github.com/zzy-rabbit/xtools/xcontext"
	"github.com/zzy-rabbit/xtools/xplugin"
)

func init() {
	ctx := xcontext.Background()
	plugin := internal.New(ctx)
	err := xplugin.Register(ctx, plugin)
	if err != nil {
		panic(err)
	}
}
