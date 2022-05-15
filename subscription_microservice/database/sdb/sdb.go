package sdb

import (
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/models"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/services"
	"gorm.io/gorm"
)

func GetAllSubscriptions(db *gorm.DB) []models.Subscription {
	var subs []models.Subscription
	db.Find(&subs)
	return subs
}

func AddSubscription(db *gorm.DB, subs models.Subscription) int32 {
	db.Create(&subs)
	return subs.Id
}

func CancelSubscription(db *gorm.DB, pid, uid int32) models.Subscription {
	var subs models.Subscription
	db.First(&subs, "uid = ? AND pid = ?", uid, pid)
	subs.IsActive = false
	db.Save(&subs)
	return subs
}

func RenewSubscription(db *gorm.DB, pid, uid int32) models.Subscription {
	var subs models.Subscription
	db.First(&subs, "uid = ? AND pid = ?", uid, pid)
	subs.ExpiryDate = services.CalculateExpiryDate(subs.PlanValidity)
	subs.IsActive = true
	db.Save(&subs)
	return subs
}

func GetUserSubscriptions(db *gorm.DB, id int32) []models.Subscription {
	var subs []models.Subscription
	db.Find(&subs, "uid = ?", id)
	return subs
}

func GetUserActiveSubscriptions(db *gorm.DB, id int32) []models.Subscription {
	var subs []models.Subscription
	db.Find(&subs, "uid = ? AND is_active = ?", id, true)
	return subs
}

func GetUserPreviousSubscriptions(db *gorm.DB, id int32) []models.Subscription {
	var subs []models.Subscription
	db.Find(&subs, "uid = ? AND is_active = ?", id, false)
	return subs
}
