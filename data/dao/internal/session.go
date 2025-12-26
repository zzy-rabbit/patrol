package internal

import (
	"context"
	"github.com/zzy-rabbit/patrol/data/dao/api"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
	"gorm.io/gorm"
)

type session struct {
	db *gorm.DB
	tx bool
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
		db: s.db,
		tx: true,
	}}
}

func (s *session) AddPoint(ctx context.Context, point model.Point) (int, xerror.IError) {
	var table Point
	table.FromModel(point)
	err := s.db.Create(&table).Error
	return table.ID, transError(err)
}

func (s *session) UpdatePoint(ctx context.Context, point model.Point) xerror.IError {
	var table Point
	table.FromModel(point)
	err := s.db.Updates(&table).Error
	return transError(err)
}

func (s *session) DeletePoints(ctx context.Context, identifies ...model.Identify) xerror.IError {
	ids := make([]string, 0, len(identifies))
	idMap := make(map[string]bool, len(identifies))
	for _, identify := range identifies {
		ids = append(ids, identify.ID)
		idMap[identify.ID] = true
	}

	tx := s.getTransaction(ctx)

	// 更新router
	xerr := tx.DeletePointsFromRouters(ctx, ids)
	if xerr != nil {
		tx.Rollback(ctx)
		return xerr
	}

	// 删除point
	err := tx.db.Where("identify in ?", ids).Delete(&Point{}).Error
	if err != nil {
		tx.Rollback(ctx)
		return transError(err)
	}

	tx.Commit(ctx)
	return nil
}

func (s *session) GetPoints(ctx context.Context, condition model.PointCondition) ([]model.Point, model.PageInfo, xerror.IError) {
	db := s.db
	if len(condition.ID) > 0 {
		db = db.Where("identify in ?", condition.ID)
	}
	if len(condition.Department) > 0 {
		db = db.Where("department in ?", condition.Department)
	}
	if len(condition.Type) > 0 {
		db = db.Where("type in ?", condition.Type)
	}
	if condition.Name != "" {
		db = db.Where("name like ?", "%"+condition.Name+"%")
	}
	if len(condition.Serial) > 0 {
		db = db.Where("serial in ?", condition.Serial)
	}
	total := int64(0)
	if condition.PageQuery != nil && condition.PageQuery.Num > 0 && condition.PageQuery.Size > 0 {
		err := db.Count(&total).Error
		if err != nil {
			return nil, model.PageInfo{}, transError(err)
		}
		offset := (condition.PageQuery.Num - 1) * condition.PageQuery.Size
		db = db.Offset(offset).Limit(condition.PageQuery.Size)
	}

	var tables []Point
	err := db.Find(&tables).Error
	if err != nil {
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
	return 0, nil
}

func (s *session) UpdateRouter(ctx context.Context, router model.Router) xerror.IError {
	return nil
}

func (s *session) DeleteRouters(ctx context.Context, routers ...model.Identify) xerror.IError {
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
	if err != nil {
		return transError(err)
	}

	// 获取涉及的routers
	routers, _, err := s.GetRouters(ctx, model.RouterCondition{Department: departments})
	if err != nil {
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
		if err != nil {
			return transError(err)
		}
	}
	return nil
}

func (s *session) GetRouters(ctx context.Context, condition model.RouterCondition) ([]model.Router, model.PageInfo, xerror.IError) {
	return nil, model.PageInfo{}, nil
}

func (s *session) AddPlan(ctx context.Context, plan model.Plan) (int, xerror.IError) {
	return 0, nil
}

func (s *session) UpdatePlan(ctx context.Context, plan model.Plan) xerror.IError {
	return nil
}

func (s *session) DeletePlans(ctx context.Context, plans ...model.Identify) xerror.IError {
	return nil
}

func (s *session) GetPlans(ctx context.Context, condition model.PlanCondition) ([]model.Plan, model.PageInfo, xerror.IError) {
	return nil, model.PageInfo{}, nil
}
