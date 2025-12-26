package internal

func (s *service) registerRouter() {
	pointGroup := s.fiberApp.Group("/api/v1/point")
	pointGroup.Post("add", s.AddPoint)
	pointGroup.Post("update", s.UpdatePoint)
	pointGroup.Delete("", s.DeletePoints)
	pointGroup.Get("", s.GetPoints)

	routerGroup := s.fiberApp.Group("/api/v1/router")
	routerGroup.Post("add", s.AddRouter)
	routerGroup.Post("update", s.UpdateRouter)
	routerGroup.Delete("", s.DeleteRouters)
	routerGroup.Get("", s.GetRouters)

	planGroup := s.fiberApp.Group("/api/v1/plan")
	planGroup.Post("add", s.AddPlan)
	planGroup.Post("update", s.UpdatePlan)
	planGroup.Delete("", s.DeletePlans)
	planGroup.Get("", s.GetPlans)

	//checkPointGroup := s.fiberApp.Group("/api/v1/check-point")
	//checkPointGroup.Post("add", nil)
	//checkPointGroup.Post("update", nil)
	//checkPointGroup.Delete("", nil)
	//checkPointGroup.Get("", nil)
	//
	//scheduleGroup := s.fiberApp.Group("/api/v1/schedule")
	//scheduleGroup.Post("add", nil)
	//scheduleGroup.Post("update", nil)
	//scheduleGroup.Delete("", nil)
	//scheduleGroup.Get("", nil)
}
