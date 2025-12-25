package dao

import (
	"github.com/zzy-rabbit/patrol/plugins/data/dao/internal"
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
