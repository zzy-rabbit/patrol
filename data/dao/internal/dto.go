package internal

import "github.com/zzy-rabbit/patrol/model"

func (d *Department) FromModel(m model.Department) {
	d.Identify = m.Identify.ID
	d.Name = m.Name
	d.Detail = m.Detail
	d.Description = m.Description
}

func (p *Department) ToModel() model.Department {
	return model.Department{
		Identify: model.Identify{
			ID: p.Identify,
		},
		Name:        p.Name,
		Detail:      p.Detail,
		Description: p.Description,
	}
}

func (p *Point) FromModel(m model.Point) {
	p.Identify = m.Identify.ID
	p.Department = m.Department
	p.Name = m.Name
	p.Type = int(m.Type)
	p.Serial = m.Serial
}

func (p *Point) ToModel() model.Point {
	return model.Point{
		Identify: model.Identify{
			ID: p.Identify,
		},
		Department: p.Department,
		Name:       p.Name,
		Type:       model.PointType(p.Type),
		Serial:     p.Serial,
	}
}

func (r *Router) FromModel(m model.Router) {
	r.Identify = m.Identify.ID
	r.Department = m.Department
	r.Name = m.Name
	r.Type = int(m.Type)
	r.Points = m.Points
}

func (r *Router) ToModel() model.Router {
	return model.Router{
		Identify: model.Identify{
			ID: r.Identify,
		},
		Department: r.Department,
		Name:       r.Name,
		Type:       model.RouterType(r.Type),
		Points:     r.Points,
	}
}

func (p *Plan) FromModel(m model.Plan) {
	users := make([]string, len(p.Users))
	copy(users, p.Users)

	p.Identify = m.Identify.ID
	p.Department = m.Department
	p.Name = m.Name
	p.Type = int(m.Type)
	p.Router = m.Router
	p.Start = m.Start
	p.End = m.End
	p.Util = m.Util
	p.Users = users
}

func (p *Plan) ToModel() model.Plan {
	users := make([]string, len(p.Users))
	copy(users, p.Users)
	return model.Plan{
		Identify: model.Identify{
			ID: p.Identify,
		},
		Department: p.Department,
		Name:       p.Name,
		Type:       model.PlanType(p.Type),
		Router:     p.Router,
		Start:      p.Start,
		End:        p.End,
		Util:       p.Util,
		Users:      users,
	}
}
