package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
)

func (s *service) AddPoint(ctx context.Context, point model.Point) (int, xerror.IError) {
	id, err := s.IDao.GetDB(ctx).AddPoint(ctx, point)
	if xerror.Error(err) {
		s.ILogger.Error(ctx, "add point error", err)
		return 0, err
	}
	//err = s.IReport.Broadcast(ctx, 0, point)
	//if xerror.Error(err) {
	//	s.ILogger.Error(ctx, "broadcast error", err)
	//}
	return id, nil
}

func (s *service) UpdatePoint(ctx context.Context, point model.Point) xerror.IError {
	return s.IDao.GetDB(ctx).UpdatePoint(ctx, point)
}

func (s *service) DeletePoints(ctx context.Context, points ...model.Identify) xerror.IError {
	return s.IDao.GetDB(ctx).DeletePoints(ctx, points...)
}

func (s *service) GetPoints(ctx context.Context, condition model.PointCondition) ([]model.Point, model.PageInfo, xerror.IError) {
	return s.IDao.GetDB(ctx).GetPoints(ctx, condition)
}

func (s *service) AddRouter(ctx context.Context, router model.Router) (int, xerror.IError) {
	return s.IDao.GetDB(ctx).AddRouter(ctx, router)
}

func (s *service) UpdateRouter(ctx context.Context, router model.Router) xerror.IError {
	return s.IDao.GetDB(ctx).UpdateRouter(ctx, router)
}

func (s *service) DeleteRouters(ctx context.Context, routers ...model.Identify) xerror.IError {
	return s.IDao.GetDB(ctx).DeleteRouters(ctx, routers...)
}

func (s *service) GetRouters(ctx context.Context, condition model.RouterCondition) ([]model.Router, model.PageInfo, xerror.IError) {
	return s.IDao.GetDB(ctx).GetRouters(ctx, condition)
}

func (s *service) AddPlan(ctx context.Context, plan model.Plan) (int, xerror.IError) {
	return s.IDao.GetDB(ctx).AddPlan(ctx, plan)
}

func (s *service) UpdatePlan(ctx context.Context, plan model.Plan) xerror.IError {
	return s.IDao.GetDB(ctx).UpdatePlan(ctx, plan)
}

func (s *service) DeletePlans(ctx context.Context, plans ...model.Identify) xerror.IError {
	return s.IDao.GetDB(ctx).DeletePlans(ctx, plans...)
}

func (s *service) GetPlans(ctx context.Context, condition model.PlanCondition) ([]model.Plan, model.PageInfo, xerror.IError) {
	return s.IDao.GetDB(ctx).GetPlans(ctx, condition)
}
