package internal

import (
	"github.com/gofiber/fiber/v2"
	bpmodel "github.com/zzy-rabbit/bp/model"
	"github.com/zzy-rabbit/patrol/model"
	"github.com/zzy-rabbit/xtools/xerror"
)

func (s *service) AddPoint(ctx *fiber.Ctx) error {
	var point model.Point
	err := s.IHttp.ParseBodyParams(ctx, &point)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	_, xerr := s.IConfig.AddPoint(ctx.UserContext(), department, point)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) UpdatePoint(ctx *fiber.Ctx) error {
	var point model.Point
	err := s.IHttp.ParseBodyParams(ctx, &point)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	xerr := s.IConfig.UpdatePoint(ctx.UserContext(), department, point)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) DeletePoints(ctx *fiber.Ctx) error {
	var points []model.Identify
	err := s.IHttp.ParseBodyParams(ctx, &points)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	xerr := s.IConfig.DeletePoints(ctx.UserContext(), department, points...)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) GetPoints(ctx *fiber.Ctx) error {
	var condition model.PointCondition
	err := s.IHttp.ParseQueryParams(ctx, &condition)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	points, page, xerr := s.IConfig.GetPoints(ctx.UserContext(), department, condition)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
		Data: model.PaginatedData[model.Point]{
			PageInfo: page,
			List:     points,
		},
	})
}

func (s *service) AddRouter(ctx *fiber.Ctx) error {
	var router model.Router
	err := s.IHttp.ParseBodyParams(ctx, &router)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	_, xerr := s.IConfig.AddRouter(ctx.UserContext(), department, router)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) UpdateRouter(ctx *fiber.Ctx) error {
	var router model.Router
	err := s.IHttp.ParseBodyParams(ctx, &router)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	xerr := s.IConfig.UpdateRouter(ctx.UserContext(), department, router)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) DeleteRouters(ctx *fiber.Ctx) error {
	var routers []model.Identify
	err := s.IHttp.ParseBodyParams(ctx, &routers)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	xerr := s.IConfig.DeleteRouters(ctx.UserContext(), department, routers...)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) GetRouters(ctx *fiber.Ctx) error {
	var condition model.RouterCondition
	err := s.IHttp.ParseQueryParams(ctx, &condition)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	routers, page, xerr := s.IConfig.GetRouters(ctx.UserContext(), department, condition)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
		Data: model.PaginatedData[model.Router]{
			PageInfo: page,
			List:     routers,
		},
	})
}

func (s *service) AddPlan(ctx *fiber.Ctx) error {
	var plan model.Plan
	err := s.IHttp.ParseBodyParams(ctx, &plan)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	_, xerr := s.IConfig.AddPlan(ctx.UserContext(), department, plan)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) UpdatePlan(ctx *fiber.Ctx) error {
	var plan model.Plan
	err := s.IHttp.ParseBodyParams(ctx, &plan)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	xerr := s.IConfig.UpdatePlan(ctx.UserContext(), department, plan)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) DeletePlans(ctx *fiber.Ctx) error {
	var plans []model.Identify
	err := s.IHttp.ParseBodyParams(ctx, &plans)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	xerr := s.IConfig.DeletePlans(ctx.UserContext(), department, plans...)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
	})
}

func (s *service) GetPlans(ctx *fiber.Ctx) error {
	var condition model.PlanCondition
	err := s.IHttp.ParseQueryParams(ctx, &condition)
	if xerror.Error(err) {
		return err
	}
	department := ctx.Get("department")
	plans, page, xerr := s.IConfig.GetPlans(ctx.UserContext(), department, condition)
	if xerror.Error(xerr) {
		return ctx.JSON(&bpmodel.HttpResponse{
			IError: xerr,
		})
	}
	return ctx.JSON(&bpmodel.HttpResponse{
		IError: xerror.ErrSuccess,
		Data: model.PaginatedData[model.Plan]{
			PageInfo: page,
			List:     plans,
		},
	})
}
