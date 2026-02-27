package internal

import (
	"context"
	logApi "github.com/zzy-rabbit/bp/tool/log/api"
	"github.com/zzy-rabbit/patrol/model"
	"time"
)

type auxiliary struct {
	pointMap          map[string][]time.Time
	checkPointMap     map[string][]int // 根据物理设备序列号，对应的打卡记录，value为打卡记录的下标
	checkPointUsedMap map[int]bool     // 打卡记录是否被使用，key为打卡记录的下标
	userMap           map[string]bool
}

type Judge struct {
	ILogger   logApi.IPlugin `xplugin:"bp.tool.log"`
	params    model.ExecutorParams
	auxiliary auxiliary
	result    model.ExecuteResult
}

func (s *service) NewJudge(_ context.Context, param model.ExecutorParams) *Judge {
	aux := auxiliary{
		pointMap:          make(map[string][]time.Time, len(param.Points)),
		checkPointMap:     make(map[string][]int, len(param.CheckPoints)),
		checkPointUsedMap: make(map[int]bool, len(param.CheckPoints)),
		userMap:           make(map[string]bool, len(param.Plan.Users)),
	}

	// 打卡记录
	for i, checkPoint := range param.CheckPoints {
		aux.checkPointMap[checkPoint.Serial] = append(aux.checkPointMap[checkPoint.Serial], i)
	}
	for _, user := range param.Plan.Users {
		aux.userMap[user] = true
	}

	return &Judge{
		params:    param,
		ILogger:   s.ILogger,
		auxiliary: aux,
		result: model.ExecuteResult{
			ID:     0,
			Status: model.ExecuteStatusWaiting,
			Start:  param.StartDate.Add(param.Plan.Start.Sub(time.Time{})),
			End:    param.StartDate.Add(param.Plan.End.Sub(time.Time{})),
			Points: make([]model.ExecutorPoint, 0, len(param.Points)),
		},
	}
}

func (j *Judge) Exec(ctx context.Context) {
	start := j.result.Start.Add(-1 * time.Second * time.Duration(j.params.Router.Deviation))
	end := j.result.End.Add(time.Second * time.Duration(j.params.Router.Deviation))

POINT:
	for _, point := range j.params.Points {
		// 获取该点位下的所有打卡记录，这里获取到的是打卡记录的下标列表
		checkPointIndexes := j.auxiliary.checkPointMap[point.Serial]

		// 遍历该点位下的所有打卡记录
		for _, checkPointIndex := range checkPointIndexes {
			// 该记录已被使用，跳过
			if j.auxiliary.checkPointUsedMap[checkPointIndex] {
				continue
			}
			checkPoint := j.params.CheckPoints[checkPointIndex]
			// 用户符合、时间符合，认为该记录有效，正常
			if j.auxiliary.userMap[checkPoint.User] && checkPoint.Time.After(start) && checkPoint.Time.Before(end) {
				j.auxiliary.checkPointUsedMap[checkPointIndex] = true
				j.result.Points = append(j.result.Points, model.ExecutorPoint{
					Point:  point.Name,
					Status: model.ExecuteStatusNormal,
					Time:   checkPoint.Time,
				})
				continue POINT
			}
		}

		// 该点位下没有符合要求的打卡记录，认为该点位没有被正常打卡
		j.result.Points = append(j.result.Points, model.ExecutorPoint{
			Point:  point.Name,
			Status: model.ExecuteStatusAbnormal,
			Time:   time.Time{},
		})
		j.result.Status = model.ExecuteStatusAbnormal
	}

	if j.result.Status == model.ExecuteStatusRunning {
		j.result.Status = model.ExecuteStatusNormal
	}
}
