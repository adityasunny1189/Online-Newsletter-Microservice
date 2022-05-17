package sdb

import (
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/models"
	"gorm.io/gorm"
)

func GetAllSubscription(db *gorm.DB) []models.Subscription {
	var subscriptions []models.Subscription
	db.Find(&subscriptions)
	return subscriptions
}

func GetSubscriptions(db *gorm.DB, pid int32) []models.Subscription {
	var subs []models.Subscription
	db.Find(&subs, "pid = ?", pid)
	return subs
}
