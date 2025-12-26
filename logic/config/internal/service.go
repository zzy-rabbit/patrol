package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
)

func (s *service) AddPoint(ctx context.Context, point model.Point) (int, xerror.IError) {
	return 0, nil
}

func (s *service) UpdatePoint(ctx context.Context, point model.Point) xerror.IError {
	return nil
}

func (s *service) DeletePoints(ctx context.Context, points ...model.Identify) xerror.IError {
	return nil
}

func (s *service) GetPoints(ctx context.Context, condition model.PointCondition) ([]model.Point, model.PageInfo, xerror.IError) {
	return nil, model.PageInfo{}, nil
}

func (s *service) AddRouter(ctx context.Context, router model.Router) (int, xerror.IError) {
	return 0, nil
}

func (s *service) UpdateRouter(ctx context.Context, router model.Router) xerror.IError {
	return nil
}

func (s *service) DeleteRouters(ctx context.Context, routers ...model.Identify) xerror.IError {
	return nil
}

func (s *service) GetRouters(ctx context.Context, condition model.RouterCondition) ([]model.Router, model.PageInfo, xerror.IError) {
	return nil, model.PageInfo{}, nil
}

func (s *service) AddPlan(ctx context.Context, plan model.Plan) (int, xerror.IError) {
	return 0, nil
}

func (s *service) UpdatePlan(ctx context.Context, plan model.Plan) xerror.IError {
	return nil
}

func (s *service) DeletePlans(ctx context.Context, plans ...model.Identify) xerror.IError {
	return nil
}

func (s *service) GetPlans(ctx context.Context, condition model.PlanCondition) ([]model.Plan, model.PageInfo, xerror.IError) {
	return nil, model.PageInfo{}, nil
}
