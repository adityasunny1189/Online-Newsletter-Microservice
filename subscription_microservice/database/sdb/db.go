package sdb

import (
	"log"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:Adisunny123@tcp(127.0.0.1:3306)/nuclei?charset=utf8mb4&parseTime=True&loc=Local"
)

type SDB struct {
	db *gorm.DB
}

var sdb SDB

func GetAllSubscriptions() []models.Subscription {

}

func AddSubscription(subs models.Subscription) int32 {

}

func CancelSubscription(pid, uid int32) models.Subscription {

}

func RenewSubscription(pid, uid int32) models.Subscription {

}

func InitializeDB() {
	var err error
	sdb.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	sdb.db.AutoMigrate(&models.Subscription{})
	log.Printf("connected to database")
}
