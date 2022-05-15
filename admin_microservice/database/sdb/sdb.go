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
