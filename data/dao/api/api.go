package api

import (
	"context"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
	"github.com/zzy-rabbit/xtools/xplugin"
)

const (
	PluginName = "patrol.data.dao"
)

type Config struct {
	Driver string `json:"driver"`
	Sqlite Sqlite `json:"sqlite"`
}

type Sqlite struct {
	File     string `json:"file"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ISession interface {
	GetTransaction(ctx context.Context) ITransaction

	SetDepartment(ctx context.Context, department model.Department) xerror.IError
	GetDepartment(ctx context.Context) (model.Department, xerror.IError)

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

type ITransaction interface {
	ISession
	Commit(ctx context.Context)
	Rollback(ctx context.Context)
}

type IDatabase interface {
	GetDB(ctx context.Context) ISession
	GetTransaction(ctx context.Context) ITransaction
	Close(ctx context.Context) error
}

type IPlugin interface {
	xplugin.IPlugin
	OpenDatabase(ctx context.Context, config Config) (IDatabase, xerror.IError)
}
