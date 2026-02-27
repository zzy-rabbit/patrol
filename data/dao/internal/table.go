package internal

import (
	"time"
)

type Department struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement"`
	Identify    string `gorm:"column:identify;unique"`
	Name        string `gorm:"column:name"`
	Detail      string `gorm:"column:detail"`
	Description string `gorm:"column:description"`
}

func (*Department) TableName() string {
	return "table_department"
}

type Point struct {
	ID       int    `gorm:"column:id;primaryKey;autoIncrement"`
	Identify string `gorm:"column:identify;unique"`
	Name     string `gorm:"column:name"`
	Type     int    `gorm:"column:type"`
	Serial   string `gorm:"column:serial"`
}

func (*Point) TableName() string {
	return "table_point"
}

type Router struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement"`
	Identify  string   `gorm:"column:identify;unique"`
	Name      string   `gorm:"column:name"`
	Type      int      `gorm:"column:type"`
	Points    []string `gorm:"column:points;serializer:json"`
	Deviation int      `gorm:"column:deviation"`
}

func (*Router) TableName() string {
	return "table_router"
}

type Plan struct {
	ID       int       `gorm:"column:id;primaryKey;autoIncrement"`
	Identify string    `gorm:"column:identify;unique"`
	Name     string    `gorm:"column:name"`
	Type     int       `gorm:"column:type"`
	Router   string    `gorm:"column:router"`
	Util     time.Time `gorm:"column:util"`
	Start    time.Time `gorm:"column:start"`
	End      time.Time `gorm:"column:end"`
	Users    []string  `gorm:"column:users;serializer:json"`
}

func (*Plan) TableName() string {
	return "table_plan"
}
