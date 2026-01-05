package internal

import (
	"context"
	logApi "github.com/zzy-rabbit/bp/log/api"
	"github.com/zzy-rabbit/patrol/data/dao/api"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
	"gorm.io/gorm"
	"time"
)

type session struct {
	db     *gorm.DB
	tx     bool
	logger logApi.IPlugin
}

type transaction struct {
	*session
}

func (t *transaction) Commit(ctx context.Context) {
	t.db.Commit()
}

func (t *transaction) Rollback(ctx context.Context) {
	t.db.Rollback()
}

func (s *session) GetTransaction(ctx context.Context) api.ITransaction {
	return s.getTransaction(ctx)
}

func (s *session) getTransaction(ctx context.Context) *transaction {
	if s.tx {
		return &transaction{session: s}
	}
	return &transaction{session: &session{
		db:     s.db.Begin(),
		tx:     true,
		logger: s.logger,
	}}
}

func (s *session) AddPoint(ctx context.Context, point model.Point) (int, xerror.IError) {
	var table Point
	table.FromModel(point)
	err := s.db.Create(&table).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "add point %+v fail %v", point, err)
		return 0, transError(err)
	}
	return table.ID, nil
}

func (s *session) UpdatePoint(ctx context.Context, point model.Point) xerror.IError {
	var table Point
	table.FromModel(point)
	err := s.db.Updates(&table).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "update point %+v fail %v", point, err)
		return transError(err)
	}
	return nil
}

func (s *session) DeletePointsFromRouters(ctx context.Context, points []string) xerror.IError {
	pointMap := make(map[string]bool, len(points))
	for _, point := range points {
		pointMap[point] = true
	}

	// 查询涉及的department
	var departments []string
	err := s.db.Where("identify in ?", points).Distinct("department").Pluck("department", &departments).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "query departments by identify %+v fail %v", points, err)
		return transError(err)
	}

	// 获取涉及的routers
	routers, _, err := s.GetRouters(ctx, model.RouterCondition{Departments: departments})
	if xerror.Error(err) {
		s.logger.Error(ctx, "query routers by departments %+v fail %v", departments, err)
		return transError(err)
	}

	// 删除router中的point
	for i, router := range routers {
		tempPoints := make([]string, 0, len(router.Points))
		for _, point := range router.Points {
			if pointMap[point] {
				continue
			}
			tempPoints = append(tempPoints, point)
		}
		routers[i].Points = tempPoints
	}

	// 更新routers
	for _, router := range routers {
		err = s.UpdateRouter(ctx, router)
		if xerror.Error(err) {
			s.logger.Error(ctx, "update router %+v fail %v", router, err)
			return transError(err)
		}
	}
	return nil
}

func (s *session) DeletePoints(ctx context.Context, identifies ...model.Identify) xerror.IError {
	ids := make([]string, 0, len(identifies))
	for _, identify := range identifies {
		ids = append(ids, identify.ID)
	}

	tx := s.getTransaction(ctx)

	// 更新router
	xerr := tx.DeletePointsFromRouters(ctx, ids)
	if xerror.Error(xerr) {
		tx.Rollback(ctx)
		s.logger.Error(ctx, "delete points %+v from routers fail %v", ids, xerr)
		return xerr
	}

	// 删除point
	err := tx.db.Where("identify in ?", ids).Delete(&Point{}).Error
	if xerror.Error(err) {
		tx.Rollback(ctx)
		s.logger.Error(ctx, "delete points %+v fail %v", ids, err)
		return transError(err)
	}

	tx.Commit(ctx)
	return nil
}

func (s *session) GetPoints(ctx context.Context, condition model.PointCondition) ([]model.Point, model.PageInfo, xerror.IError) {
	db := s.db
	if len(condition.IDs) > 0 {
		db = db.Where("identify in ?", condition.IDs)
	}
	if len(condition.Departments) > 0 {
		db = db.Where("department in ?", condition.Departments)
	}
	if len(condition.Types) > 0 {
		db = db.Where("type in ?", condition.Types)
	}
	if condition.Name != "" {
		db = db.Where("name like ?", "%"+condition.Name+"%")
	}
	if len(condition.Serials) > 0 {
		db = db.Where("serial in ?", condition.Serials)
	}
	total := int64(0)
	if condition.PageQuery != nil && condition.PageQuery.Num > 0 && condition.PageQuery.Size > 0 {
		err := db.Count(&total).Error
		if xerror.Error(err) {
			s.logger.Error(ctx, "query points count by condition %+v fail %v", condition, err)
			return nil, model.PageInfo{}, transError(err)
		}
		offset := (condition.PageQuery.Num - 1) * condition.PageQuery.Size
		db = db.Offset(offset).Limit(condition.PageQuery.Size)
	}

	var tables []Point
	err := db.Find(&tables).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "query points by condition %+v fail %v", condition, err)
		return nil, model.PageInfo{}, transError(err)
	}
	points := make([]model.Point, 0, len(tables))
	for _, table := range tables {
		points = append(points, table.ToModel())
	}
	return points, model.PageInfo{
		Count: len(tables),
		Total: int(total),
	}, nil
}

func (s *session) AddRouter(ctx context.Context, router model.Router) (int, xerror.IError) {
	var table Router
	table.FromModel(router)
	err := s.db.Create(&table).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "add router %+v fail %v", router, err)
		return 0, transError(err)
	}
	return table.ID, nil
}

func (s *session) UpdateRouter(ctx context.Context, router model.Router) xerror.IError {
	var table Router
	table.FromModel(router)
	err := s.db.Updates(&table).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "update router %+v fail %v", router, err)
		return transError(err)
	}
	return nil
}

func (s *session) DeleteRoutersFromPlans(ctx context.Context, routers []string) xerror.IError {
	// 查询涉及的department
	var departments []string
	err := s.db.Where("identify in ?", routers).Distinct("department").Pluck("department", &departments).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "query departments by identify %+v fail %v", routers, err)
		return transError(err)
	}

	// 获取涉及的plans
	plans, _, err := s.GetPlans(ctx, model.PlanCondition{Departments: departments, Routers: routers})
	if xerror.Error(err) {
		s.logger.Error(ctx, "query plans by departments %+v fail %v", departments, err)
		return transError(err)
	}

	// 删除plan中的router
	for i := range plans {
		plans[i].Router = ""
	}

	// 更新plans
	for _, plan := range plans {
		err = s.UpdatePlan(ctx, plan)
		if xerror.Error(err) {
			s.logger.Error(ctx, "update plan %+v fail %v", plan, err)
			return transError(err)
		}
	}
	return nil
}

func (s *session) DeleteRouters(ctx context.Context, identifies ...model.Identify) xerror.IError {
	ids := make([]string, 0, len(identifies))
	for _, identify := range identifies {
		ids = append(ids, identify.ID)
	}

	tx := s.getTransaction(ctx)

	// 更新plan
	xerr := tx.DeleteRoutersFromPlans(ctx, ids)
	if xerror.Error(xerr) {
		tx.Rollback(ctx)
		s.logger.Error(ctx, "delete routers %+v from plans fail %v", ids, xerr)
		return xerr
	}

	// 删除router
	err := tx.db.Where("identify in ?", ids).Delete(&Router{}).Error
	if xerror.Error(err) {
		tx.Rollback(ctx)
		s.logger.Error(ctx, "delete routers %+v fail %v", ids, err)
		return transError(err)
	}

	tx.Commit(ctx)
	return nil
}

func (s *session) GetRouters(ctx context.Context, condition model.RouterCondition) ([]model.Router, model.PageInfo, xerror.IError) {
	db := s.db
	if len(condition.IDs) > 0 {
		db = db.Where("identify in ?", condition.IDs)
	}
	if len(condition.Departments) > 0 {
		db = db.Where("department in ?", condition.Departments)
	}
	if len(condition.Types) > 0 {
		db = db.Where("type in ?", condition.Types)
	}
	if condition.Name != "" {
		db = db.Where("name like ?", "%"+condition.Name+"%")
	}
	total := int64(0)
	if condition.PageQuery != nil && condition.PageQuery.Num > 0 && condition.PageQuery.Size > 0 {
		err := db.Count(&total).Error
		if xerror.Error(err) {
			s.logger.Error(ctx, "query routers count by condition %+v fail %v", condition, err)
			return nil, model.PageInfo{}, transError(err)
		}
		offset := (condition.PageQuery.Num - 1) * condition.PageQuery.Size
		db = db.Offset(offset).Limit(condition.PageQuery.Size)
	}

	var tables []Router
	err := db.Find(&tables).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "query routers by condition %+v fail %v", condition, err)
		return nil, model.PageInfo{}, transError(err)
	}
	routers := make([]model.Router, 0, len(tables))
	for _, table := range tables {
		routers = append(routers, table.ToModel())
	}
	return routers, model.PageInfo{
		Count: len(tables),
		Total: int(total),
	}, nil
}

func (s *session) AddPlan(ctx context.Context, plan model.Plan) (int, xerror.IError) {
	var table Plan
	table.FromModel(plan)
	err := s.db.Create(&table).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "add plan %+v fail %v", plan, err)
		return 0, transError(err)
	}
	return table.ID, nil
}

func (s *session) UpdatePlan(ctx context.Context, plan model.Plan) xerror.IError {
	var table Plan
	table.FromModel(plan)
	err := s.db.Updates(&table).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "update plan %+v fail %v", plan, err)
		return transError(err)
	}
	return nil
}

func (s *session) DeletePlans(ctx context.Context, identifies ...model.Identify) xerror.IError {
	ids := make([]string, 0, len(identifies))
	for _, identify := range identifies {
		ids = append(ids, identify.ID)
	}
	// 删除plan
	err := s.db.Where("identify in ?", ids).Delete(&Router{}).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "delete plans %+v fail %v", ids, err)
		return transError(err)
	}
	return nil
}

func (s *session) GetPlans(ctx context.Context, condition model.PlanCondition) ([]model.Plan, model.PageInfo, xerror.IError) {
	db := s.db
	if len(condition.IDs) > 0 {
		db = db.Where("identify in ?", condition.IDs)
	}
	if len(condition.Departments) > 0 {
		db = db.Where("department in ?", condition.Departments)
	}
	if len(condition.Routers) > 0 {
		db = db.Where("router in ?", condition.Routers)
	}
	if len(condition.Types) > 0 {
		db = db.Where("type in ?", condition.Types)
	}
	if condition.Name != "" {
		db = db.Where("name like ?", "%"+condition.Name+"%")
	}
	if condition.Start != (time.Time{}) {
		db = db.Where("start >= ?", condition.Start)
	}
	if condition.End != (time.Time{}) {
		db = db.Where("end <= ?", condition.End)
	}

	total := int64(0)
	if condition.PageQuery != nil && condition.PageQuery.Num > 0 && condition.PageQuery.Size > 0 {
		err := db.Count(&total).Error
		if xerror.Error(err) {
			s.logger.Error(ctx, "query plans count by condition %+v fail %v", condition, err)
			return nil, model.PageInfo{}, transError(err)
		}
		offset := (condition.PageQuery.Num - 1) * condition.PageQuery.Size
		db = db.Offset(offset).Limit(condition.PageQuery.Size)
	}

	var tables []Plan
	err := db.Find(&tables).Error
	if xerror.Error(err) {
		s.logger.Error(ctx, "query plans by condition %+v fail %v", condition, err)
		return nil, model.PageInfo{}, transError(err)
	}
	plans := make([]model.Plan, 0, len(tables))
	for _, table := range tables {
		plans = append(plans, table.ToModel())
	}
	return plans, model.PageInfo{
		Count: len(tables),
		Total: int(total),
	}, nil
}

func (s *session) DeleteUsersFromPlans(ctx context.Context, users []string) xerror.IError {
	// 获取涉及的plans
	plans, _, err := s.GetPlans(ctx, model.PlanCondition{})
	if xerror.Error(err) {
		s.logger.Error(ctx, "query all plans %+v fail %v", err)
		return transError(err)
	}

	userMap := make(map[string]bool, len(users))
	for _, user := range users {
		userMap[user] = true
	}

	// 删除plan中的router
	for i, plan := range plans {
		for _, user := range plan.Users {
			if userMap[user] {
				plans[i].Users = append(plan.Users[:i], plan.Users[i+1:]...)
			}
		}
	}

	// 更新plans
	for _, plan := range plans {
		err = s.UpdatePlan(ctx, plan)
		if xerror.Error(err) {
			s.logger.Error(ctx, "update plan %+v fail %v", plan, err)
			return transError(err)
		}
	}
	return nil
}
