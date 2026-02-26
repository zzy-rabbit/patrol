package internal

import (
	"context"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	"github.com/zzy-rabbit/patrol/model"
)

type auxiliary struct {
	checkPointMap map[string][]model.CheckPoint
}

type Judge struct {
	ILogger   logApi.IPlugin `xplugin:"bp.tool.log"`
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
