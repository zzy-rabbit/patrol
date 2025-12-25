package model

type PointType int

const (
	PointTypeNFC PointType = iota + 1
	PointTypeQRCode
)

type Point struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Type   PointType `json:"type"`   // 硬件设备类型
	Serial string    `json:"serial"` // 硬件设备的唯一序列号
}

type RouterType int

type Router struct {
	ID     string     `json:"id"`
	Name   string     `json:"name"`
	Type   RouterType `json:"type"`
	Points []string   `json:"points"`
}

type PlanType int

type Plan struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Type   PlanType `json:"type"`
	Router string   `json:"router"`
	Util   string   `json:"util"`
	Start  string   `json:"start"`
	End    string   `json:"end"`
}
