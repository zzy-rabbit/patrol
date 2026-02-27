package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xtrace"
	"strings"
)

func (s *service) AddDepartment(ctx context.Context, department model.Department) xerror.IError {
	database, err := s.IDatabase.New(ctx, department.ID)
	if xerror.Error(err) {
		s.ILogger.Error(ctx, "new department database %s fail", department.ID)
		return err
	}
	err = database.GetDB(ctx).SetDepartment(ctx, department)
	if xerror.Error(err) {
		s.ILogger.Error(ctx, "set department %+v fail", department)
		return err
	}
	return nil
}

func (s *service) UpdateDepartment(ctx context.Context, department model.Department) xerror.IError {
	database, ok := s.IDatabase.Get(ctx, department.ID)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department.ID)
		return xerror.Extend(xerror.ErrNotFound, "department "+department.ID)
	}
	err := database.GetDB(ctx).SetDepartment(ctx, department)
	if xerror.Error(err) {
		s.ILogger.Error(ctx, "set department %+v fail", department)
		return err
	}
	return nil
}

func (s *service) DeleteDepartments(ctx context.Context, departments ...model.Identify) xerror.IError {
	for _, department := range departments {
		err := s.IDatabase.Delete(ctx, department.ID)
		if xerror.Error(err) {
			s.ILogger.Error(ctx, "delete department %s fail", department.ID)
			return err
		}
	}
	return nil
}

func (s *service) GetDepartments(ctx context.Context, condition model.DepartmentCondition) ([]model.Department, model.PageInfo, xerror.IError) {
	defer xtrace.Trace(ctx)(condition)

	databases := s.IDatabase.GetAll(ctx)

	conditionIDMap := make(map[string]bool, len(condition.IDs))
	for _, id := range condition.IDs {
		conditionIDMap[id] = true
	}

	departments := make([]model.Department, 0, len(databases))
	for _, database := range databases {
		department, err := database.GetDB(ctx).GetDepartment(ctx)
		if xerror.Error(err) {
			s.ILogger.Error(ctx, "get department %s fail", department.ID)
			return nil, model.PageInfo{}, err
		}

		if len(condition.IDs) > 0 && !conditionIDMap[department.ID] {
			continue
		}
		if condition.Name != "" && !strings.Contains(department.Name, condition.Name) {
			continue
		}
		departments = append(departments, department)
	}

	s.ILogger.Info(ctx, "get departments %+v", departments)
	results, page := model.Paginate(departments, condition.PageQuery)
	return results, page, nil
}

func (s *service) AddPoint(ctx context.Context, department string, point model.Point) (int, xerror.IError) {
	defer xtrace.Trace(ctx)(department, point)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return 0, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).AddPoint(ctx, point)
}

func (s *service) UpdatePoint(ctx context.Context, department string, point model.Point) xerror.IError {
	defer xtrace.Trace(ctx)(department, point)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).UpdatePoint(ctx, point)
}

func (s *service) DeletePoints(ctx context.Context, department string, points ...model.Identify) xerror.IError {
	defer xtrace.Trace(ctx)(department, points)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).DeletePoints(ctx, points...)
}

func (s *service) GetPoints(ctx context.Context, department string, condition model.PointCondition) ([]model.Point, model.PageInfo, xerror.IError) {
	defer xtrace.Trace(ctx)(department, condition)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return nil, model.PageInfo{}, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).GetPoints(ctx, condition)
}

func (s *service) AddRouter(ctx context.Context, department string, router model.Router) (int, xerror.IError) {
	defer xtrace.Trace(ctx)(department, router)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return 0, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).AddRouter(ctx, router)
}

func (s *service) UpdateRouter(ctx context.Context, department string, router model.Router) xerror.IError {
	defer xtrace.Trace(ctx)(department, router)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).UpdateRouter(ctx, router)
}

func (s *service) DeleteRouters(ctx context.Context, department string, routers ...model.Identify) xerror.IError {
	defer xtrace.Trace(ctx)(department, routers)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).DeleteRouters(ctx, routers...)
}

func (s *service) GetRouters(ctx context.Context, department string, condition model.RouterCondition) ([]model.Router, model.PageInfo, xerror.IError) {
	defer xtrace.Trace(ctx)(department, condition)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return nil, model.PageInfo{}, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).GetRouters(ctx, condition)
}

func (s *service) AddPlan(ctx context.Context, department string, plan model.Plan) (int, xerror.IError) {
	defer xtrace.Trace(ctx)(department, plan)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return 0, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).AddPlan(ctx, plan)
}

func (s *service) UpdatePlan(ctx context.Context, department string, plan model.Plan) xerror.IError {
	defer xtrace.Trace(ctx)(department, plan)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).UpdatePlan(ctx, plan)
}

func (s *service) DeletePlans(ctx context.Context, department string, plans ...model.Identify) xerror.IError {
	defer xtrace.Trace(ctx)(department, plans)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).DeletePlans(ctx, plans...)
}

func (s *service) GetPlans(ctx context.Context, department string, condition model.PlanCondition) ([]model.Plan, model.PageInfo, xerror.IError) {
	defer xtrace.Trace(ctx)(department, condition)
	database, ok := s.IDatabase.Get(ctx, department)
	if !ok {
		s.ILogger.Error(ctx, "department %s database not found", department)
		return nil, model.PageInfo{}, xerror.Extend(xerror.ErrNotFound, "department "+department)
	}
	return database.GetDB(ctx).GetPlans(ctx, condition)
}
