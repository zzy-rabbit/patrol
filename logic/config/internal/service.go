package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
)

func (s *service) AddPoint(ctx context.Context, department string, point model.Point) (int, xerror.IError) {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return 0, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).AddPoint(ctx, point)
}

func (s *service) UpdatePoint(ctx context.Context, department string, point model.Point) xerror.IError {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).UpdatePoint(ctx, point)
}

func (s *service) DeletePoints(ctx context.Context, department string, points ...model.Identify) xerror.IError {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).DeletePoints(ctx, points...)
}

func (s *service) GetPoints(ctx context.Context, department string, condition model.PointCondition) ([]model.Point, model.PageInfo, xerror.IError) {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return nil, model.PageInfo{}, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).GetPoints(ctx, condition)
}

func (s *service) AddRouter(ctx context.Context, department string, router model.Router) (int, xerror.IError) {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return 0, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).AddRouter(ctx, router)
}

func (s *service) UpdateRouter(ctx context.Context, department string, router model.Router) xerror.IError {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).UpdateRouter(ctx, router)
}

func (s *service) DeleteRouters(ctx context.Context, department string, routers ...model.Identify) xerror.IError {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).DeleteRouters(ctx, routers...)
}

func (s *service) GetRouters(ctx context.Context, department string, condition model.RouterCondition) ([]model.Router, model.PageInfo, xerror.IError) {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return nil, model.PageInfo{}, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).GetRouters(ctx, condition)
}

func (s *service) AddPlan(ctx context.Context, department string, plan model.Plan) (int, xerror.IError) {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return 0, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).AddPlan(ctx, plan)
}

func (s *service) UpdatePlan(ctx context.Context, department string, plan model.Plan) xerror.IError {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).UpdatePlan(ctx, plan)
}

func (s *service) DeletePlans(ctx context.Context, department string, plans ...model.Identify) xerror.IError {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).DeletePlans(ctx, plans...)
}

func (s *service) GetPlans(ctx context.Context, department string, condition model.PlanCondition) ([]model.Plan, model.PageInfo, xerror.IError) {
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return nil, model.PageInfo{}, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).GetPlans(ctx, condition)
}
