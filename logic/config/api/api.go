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

	AddPoint(ctx context.Context, point model.Point) (int, xerror.IError)
	UpdatePoint(ctx context.Context, point model.Point) xerror.IError
	DeletePoints(ctx context.Context, points ...model.Identify) xerror.IError
	GetPoints(ctx context.Context, condition model.PointCondition) ([]model.Point, model.PageInfo, xerror.IError)

	AddRouter(ctx context.Context, router model.Router) (int, xerror.IError)
	UpdateRouter(ctx context.Context, router model.Router) xerror.IError
	DeleteRouters(ctx context.Context, routers ...model.Identify) xerror.IError
	GetRouters(ctx context.Context, condition model.RouterCondition) ([]model.Router, model.PageInfo, xerror.IError)

	AddPlan(ctx context.Context, plan model.Plan) (int, xerror.IError)
	UpdatePlan(ctx context.Context, plan model.Plan) xerror.IError
	DeletePlans(ctx context.Context, plans ...model.Identify) xerror.IError
	GetPlans(ctx context.Context, condition model.PlanCondition) ([]model.Plan, model.PageInfo, xerror.IError)
}
