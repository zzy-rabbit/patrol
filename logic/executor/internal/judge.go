package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	logApi "github.com/zzy-rabbit/xtools/plugins/log/api"
)

type auxiliary struct {
	checkPointMap map[string][]model.CheckPoint
}

type Judge struct {
	ILogger   logApi.IPlugin `xplugin:"xtools.plugins.log"`
	params    model.ExecutorParams
	auxiliary auxiliary
}

func (s *service) NewJudge(params model.ExecutorParams) *Judge {
	return &Judge{
		params:  params,
		ILogger: s.ILogger,
	}
}

func (j *Judge) Exec(ctx context.Context) {

}
