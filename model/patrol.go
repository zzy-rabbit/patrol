package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type PageQuery struct {
	Num  int `json:"num"`
	Size int `json:"size"`
}

type PageInfo struct {
	Count int `json:"count"`
	Total int `json:"total"`
}

type PaginatedData[T any] struct {
	PageInfo
	List []T `json:"list"`
}

type Identify struct {
	ID string `json:"id"`
}

type PointType int

const (
	PointTypeNFC PointType = iota + 1
	PointTypeQRCode
)

type Point struct {
	Identify
	Department string    `json:"department"`
	Name       string    `json:"name"`
	Type       PointType `json:"type"`   // 硬件设备类型
	Serial     string    `json:"serial"` // 硬件设备的唯一序列号
}

type PointCondition struct {
	*PageQuery
	ID         []string    `json:"ids"`
	Department []string    `json:"departments"`
	Type       []PointType `json:"types"`
	Name       string      `json:"name"`
	Serial     []string    `json:"serials"`
}

type RouterType int

type Router struct {
	Identify
	Department string     `json:"department"`
	Name       string     `json:"name"`
	Type       RouterType `json:"type"`
	Points     []string   `json:"points"`
}

type RouterCondition struct {
	*PageQuery
	ID         []string     `json:"id"`
	Department []string     `json:"department"`
	Type       []RouterType `json:"type"`
	Name       string       `json:"name"`
}

type PlanType int

type Plan struct {
	Identify
	Department string    `json:"department"`
	Name       string    `json:"name"`
	Type       PlanType  `json:"type"`
	Router     string    `json:"router"`
	Util       time.Time `json:"util"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
}

func (p *Plan) UnmarshalJSON(data []byte) error {
	type planJSON struct {
		Identify
		Department string   `json:"department"`
		Name       string   `json:"name"`
		Type       PlanType `json:"type"`
		Router     string   `json:"router"`
		Util       string   `json:"util"`
		Start      string   `json:"start"`
		End        string   `json:"end"`
	}

	var aux planJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	util, err := time.ParseInLocation("2006-01-02", aux.Util, time.Local)
	if err != nil {
		return fmt.Errorf("util format error: %w", err)
	}

	start, err := time.ParseInLocation("15:04:05", aux.Start, time.Local)
	if err != nil {
		return fmt.Errorf("start format error: %w", err)
	}

	end, err := time.ParseInLocation("15:04:05", aux.End, time.Local)
	if err != nil {
		return fmt.Errorf("end format error: %w", err)
	}

	p.Identify = aux.Identify
	p.Department = aux.Department
	p.Name = aux.Name
	p.Type = aux.Type
	p.Router = aux.Router
	p.Util = util
	p.Start = start
	p.End = end

	return nil
}

func (p *Plan) MarshalJSON() ([]byte, error) {
	type planJSON struct {
		Identify
		Department string   `json:"department"`
		Name       string   `json:"name"`
		Type       PlanType `json:"type"`
		Router     string   `json:"router"`
		Util       string   `json:"util"`
		Start      string   `json:"start"`
		End        string   `json:"end"`
	}

	aux := planJSON{
		Identify:   p.Identify,
		Department: p.Department,
		Name:       p.Name,
		Type:       p.Type,
		Router:     p.Router,
		Util:       p.Util.Format("2006-01-02"),
		Start:      p.Start.Format("15:04:05"),
		End:        p.End.Format("15:04:05"),
	}
	return json.Marshal(aux)
}

type PlanCondition struct {
	*PageQuery
	ID         []string   `json:"id"`
	Department []string   `json:"department"`
	Type       []PlanType `json:"type"`
	Name       string     `json:"name"`
	Start      string     `json:"start"`
	End        string     `json:"end"`
}
