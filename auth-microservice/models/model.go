package models

type User struct {
	Id       int32 `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}
