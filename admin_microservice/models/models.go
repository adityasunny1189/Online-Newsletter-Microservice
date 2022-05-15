package models

import "time"

type Plan struct {
	Id       int32 `gorm:"primaryKey"`
	Name     string
	Price    float32
	Validity int32
}

type News struct {
	Id      int32 `gorm:"primaryKey"`
	Planid  int32
	Author  string
	Heading string
	Content string
}

type User struct {
	Id       int32 `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

type Subscription struct {
	Id           int32 `gorm:"primaryKey"`
	Uid          int32
	Pid          int32
	PlanValidity int32
	ExpiryDate   time.Time
	IsActive     bool
}
