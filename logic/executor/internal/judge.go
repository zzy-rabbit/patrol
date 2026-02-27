package internal

import (
	"context"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	"github.com/zzy-rabbit/patrol/model"
	"time"
)

type auxiliary struct {
	pointMap      map[string][]time.Time
	checkPointMap map[string][]model.CheckPoint
}

type Judge struct {
	ILogger   logApi.IPlugin `xplugin:"bp.tool.log"`
	params    model.ExecutorParams
	auxiliary auxiliary
}

func (s *service) NewJudge(_ context.Context, param model.ExecutorParams) *Judge {
	aux := auxiliary{
		pointMap:      make(map[string][]time.Time, len(param.Points)),
		checkPointMap: make(map[string][]model.CheckPoint, len(param.CheckPoints)),
	}
	// 路线+计划要求的点位信息

	// 打卡记录
	for _, checkPoint := range param.CheckPoints {
		aux.checkPointMap[checkPoint.Serial] = append(aux.checkPointMap[checkPoint.Serial], checkPoint)
	}

	return &Judge{
		params:    param,
		ILogger:   s.ILogger,
		auxiliary: aux,
	}
}

func (j *Judge) Exec(ctx context.Context) {

}
