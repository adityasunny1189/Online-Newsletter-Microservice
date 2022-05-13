package models

import "time"

type Plan struct {
	Id       int32 `gorm:"primaryKey"`
	Name     string
	Price    float32
	Validity int32
}

type Subscription struct {
	Id           int32 `gorm:"primaryKey"`
	Uid          int32
	Pid          int32
	PlanValidity int32
	ExpiryDate   *time.Time
	IsActive     bool
}
