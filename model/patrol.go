package model

import (
	"encoding/json"
	"fmt"
	"github.com/zzy-rabbit/xtools/xerror"
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

func Paginate[T any](data []T, query *PageQuery) ([]T, PageInfo) {
	total := len(data)

	// 无分页条件 → 全部作为一页
	if query == nil || query.Num <= 0 || query.Size <= 0 {
		return data, PageInfo{
			Count: total,
			Total: total,
		}
	}

	start := (query.Num - 1) * query.Size
	if start >= total {
		// 页码超出范围
		return []T{}, PageInfo{
			Count: 0,
			Total: total,
		}
	}

	end := start + query.Size
	if end > total {
		end = total
	}

	result := data[start:end]

	return result, PageInfo{
		Count: len(result),
		Total: total,
	}
}

type PaginatedData[T any] struct {
	PageInfo
	List []T `json:"list"`
}

type Identify struct {
	ID string `json:"id"`
}

type Department struct {
	Identify
	Name        string `json:"name"`
	Detail      string `json:"detail"`
	Description string `json:"description"`
}

type DepartmentCondition struct {
	*PageQuery
	IDs  []string `json:"ids"`
	Name string   `json:"name"`
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
	IDs     []string    `json:"ids"`
	Types   []PointType `json:"types"`
	Name    string      `json:"name"`
	Serials []string    `json:"serials"`
}

type RouterType int

type Router struct {
	Identify
	Name      string     `json:"name"`
	Type      RouterType `json:"type"`
	Points    []string   `json:"points"`
	Deviation int        `json:"deviation"`
}

type RouterCondition struct {
	*PageQuery
	IDs   []string     `json:"ids"`
	Types []RouterType `json:"types"`
	Name  string       `json:"name"`
}

type PlanType int

const (
	PlanTypeAnyUser PlanType = iota + 1
	PlanTypeEachUser
)

type Plan struct {
	Identify
	Name   string    `json:"name"`
	Type   PlanType  `json:"type"`
	Router string    `json:"router"`
	Util   time.Time `json:"util"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Users  []string  `json:"users"`
}

func (p *Plan) UnmarshalJSON(data []byte) error {
	type planJSON struct {
		Identify
		Name   string   `json:"name"`
		Type   PlanType `json:"type"`
		Router string   `json:"router"`
		Util   string   `json:"util"`
		Start  string   `json:"start"`
		End    string   `json:"end"`
		Users  []string `json:"users"`
	}

	var aux planJSON
	if err := json.Unmarshal(data, &aux); xerror.Error(err) {
		return err
	}

	util, err := time.ParseInLocation("2006-01-02", aux.Util, time.Local)
	if xerror.Error(err) {
		return fmt.Errorf("util format error: %w", err)
	}

	start, err := time.ParseInLocation("15:04:05", aux.Start, time.Local)
	if xerror.Error(err) {
		return fmt.Errorf("start format error: %w", err)
	}

	end, err := time.ParseInLocation("15:04:05", aux.End, time.Local)
	if xerror.Error(err) {
		return fmt.Errorf("end format error: %w", err)
	}

	p.Identify = aux.Identify
	p.Name = aux.Name
	p.Type = aux.Type
	p.Router = aux.Router
	p.Util = util
	p.Start = start
	p.End = end
	p.Users = aux.Users

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
		Users      []string `json:"users"`
	}

	aux := planJSON{
		Identify: p.Identify,
		Name:     p.Name,
		Type:     p.Type,
		Router:   p.Router,
		Util:     p.Util.Format("2006-01-02"),
		Start:    p.Start.Format("15:04:05"),
		End:      p.End.Format("15:04:05"),
		Users:    p.Users,
	}
	return json.Marshal(aux)
}

type PlanCondition struct {
	*PageQuery
	IDs     []string   `json:"ids"`
	Types   []PlanType `json:"types"`
	Routers []string   `json:"routers"`
	Name    string     `json:"name"`
	Start   time.Time  `json:"start"`
	End     time.Time  `json:"end"`
}

func (p *PlanCondition) UnmarshalJSON(data []byte) error {
	type planJSON struct {
		*PageQuery
		IDs         []string   `json:"ids"`
		Departments []string   `json:"departments"`
		Types       []PlanType `json:"types"`
		Routers     []string   `json:"routers"`
		Name        string     `json:"name"`
		Start       string     `json:"start"`
		End         string     `json:"end"`
	}

	var aux planJSON
	if err := json.Unmarshal(data, &aux); xerror.Error(err) {
		return err
	}

	start, err := time.ParseInLocation("15:04:05", aux.Start, time.Local)
	if xerror.Error(err) {
		return fmt.Errorf("start format error: %w", err)
	}

	end, err := time.ParseInLocation("15:04:05", aux.End, time.Local)
	if xerror.Error(err) {
		return fmt.Errorf("end format error: %w", err)
	}

	p.PageQuery = aux.PageQuery
	p.IDs = aux.IDs
	p.Types = aux.Types
	p.Routers = aux.Routers
	p.Name = aux.Name
	p.Start = start
	p.End = end
	return nil
}

func (p *PlanCondition) MarshalJSON() ([]byte, error) {
	type planJSON struct {
		*PageQuery
		IDs     []string   `json:"ids"`
		Types   []PlanType `json:"types"`
		Routers []string   `json:"routers"`
		Name    string     `json:"name"`
		Start   string     `json:"start"`
		End     string     `json:"end"`
	}

	aux := planJSON{
		PageQuery: p.PageQuery,
		IDs:       p.IDs,
		Types:     p.Types,
		Routers:   p.Routers,
		Name:      p.Name,
		Start:     p.Start.Format("15:04:05"),
		End:       p.End.Format("15:04:05"),
	}
	return json.Marshal(aux)
}

type CheckPoint struct {
	User   string    `json:"user"`
	Serial string    `json:"serial"`
	Time   time.Time `json:"time"`
}

func (p *CheckPoint) UnmarshalJSON(data []byte) error {
	type checkPointJSON struct {
		User   string `json:"user"`
		Serial string `json:"serial"`
		Time   string `json:"time"`
	}
	var aux checkPointJSON
	if err := json.Unmarshal(data, &aux); xerror.Error(err) {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02 15:04:05", aux.Time, time.Local)
	if xerror.Error(err) {
		return err
	}
	p.User = aux.User
	p.Serial = aux.Serial
	p.Time = t
	return nil
}

func (p *CheckPoint) MarshalJSON() ([]byte, error) {
	type checkPointJSON struct {
		User   string `json:"user"`
		Serial string `json:"serial"`
		Time   string `json:"time"`
	}
	aux := checkPointJSON{
		User:   p.User,
		Serial: p.Serial,
		Time:   p.Time.Format("2006-01-02 15:04:05"),
	}
	return json.Marshal(aux)
}

type PointBak struct {
	Point
	ID int `json:"id"`
}

type RouterBak struct {
	Router
	ID int `json:"id"`
}

type PlanBak struct {
	Plan
	ID int `json:"id"`
}

type ExecutorParams struct {
	Department  string       `json:"department"`
	StartDate   time.Time    `json:"start_date"`
	Points      []Point      `json:"points"`
	Router      Router       `json:"router"`
	Plan        Plan         `json:"plan"`
	CheckPoints []CheckPoint `json:"check_points"`
}

type ExecuteStatus int

const (
	ExecuteStatusNotStart ExecuteStatus = iota // 未开始
	ExecuteStatusRunning                       // 巡更中
	ExecuteStatusWaiting                       // 等待结果
	ExecuteStatusAbnormal                      // 异常
	ExecuteStatusNormal                        // 正常
)

type ExecutorPoint struct {
	Point  string        `json:"point"`
	Status ExecuteStatus `json:"status"`
	Time   time.Time     `json:"time"`
}

type ExecuteResult struct {
	ID     int             `json:"id"`
	Status ExecuteStatus   `json:"status"`
	Start  time.Time       `json:"start"`
	End    time.Time       `json:"end"`
	Points []ExecutorPoint `json:"points"`
}
