package internal

import (
	"time"
)

type Point struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement"`
	Identify   string `gorm:"column:identify;unique"`
	Department string `gorm:"column:department"`
	Name       string `gorm:"column:name"`
	Type       int    `gorm:"column:type"`
	Serial     string `gorm:"column:serial"`
}

func (*Point) TableName() string {
	return "table_point"
}

type Router struct {
	ID         int      `gorm:"column:id;primaryKey;autoIncrement"`
	Identify   string   `gorm:"column:identify;unique"`
	Department string   `gorm:"column:department"`
	Name       string   `gorm:"column:name"`
	Type       int      `gorm:"column:type"`
	Points     []string `gorm:"column:points;serializer:json"`
}

func (*Router) TableName() string {
	return "table_router"
}

type Plan struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement"`
	Identify   string    `gorm:"column:identify;unique"`
	Department string    `gorm:"column:department"`
	Name       string    `gorm:"column:name"`
	Type       int       `gorm:"column:type"`
	Router     string    `gorm:"column:router"`
	Util       time.Time `gorm:"column:util"`
	Start      time.Time `gorm:"column:start"`
	End        time.Time `gorm:"column:end"`
	Users      []string  `gorm:"column:users;serializer:json"`
}

func (*Plan) TableName() string {
	return "table_plan"
}
