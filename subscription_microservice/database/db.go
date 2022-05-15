package database

import (
	"log"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:Adisunny123@tcp(127.0.0.1:3306)/nuclei?charset=utf8mb4&parseTime=True&loc=Local"
)

var db *gorm.DB

func InitializeDB() *gorm.DB {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	log.Printf("connected to database")
	db.AutoMigrate(&models.Plan{})
	db.AutoMigrate(&models.Subscription{})
	return db
}
