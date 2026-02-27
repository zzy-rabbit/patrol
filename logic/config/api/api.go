package api

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.logic.config"
)

type IPlugin interface {
	xplugin.IPlugin

	AddDepartment(ctx context.Context, department model.Department) xerror.IError
	UpdateDepartment(ctx context.Context, department model.Department) xerror.IError
	DeleteDepartments(ctx context.Context, departments ...model.Identify) xerror.IError
	GetDepartments(ctx context.Context, condition model.DepartmentCondition) ([]model.Department, model.PageInfo, xerror.IError)

	AddPoint(ctx context.Context, department string, point model.Point) (int, xerror.IError)
	UpdatePoint(ctx context.Context, department string, point model.Point) xerror.IError
	DeletePoints(ctx context.Context, department string, points ...model.Identify) xerror.IError
	GetPoints(ctx context.Context, department string, condition model.PointCondition) ([]model.Point, model.PageInfo, xerror.IError)

	AddRouter(ctx context.Context, department string, router model.Router) (int, xerror.IError)
	UpdateRouter(ctx context.Context, department string, router model.Router) xerror.IError
	DeleteRouters(ctx context.Context, department string, routers ...model.Identify) xerror.IError
	GetRouters(ctx context.Context, department string, condition model.RouterCondition) ([]model.Router, model.PageInfo, xerror.IError)

	AddPlan(ctx context.Context, department string, plan model.Plan) (int, xerror.IError)
	UpdatePlan(ctx context.Context, department string, plan model.Plan) xerror.IError
	DeletePlans(ctx context.Context, department string, plans ...model.Identify) xerror.IError
	GetPlans(ctx context.Context, department string, condition model.PlanCondition) ([]model.Plan, model.PageInfo, xerror.IError)
}
