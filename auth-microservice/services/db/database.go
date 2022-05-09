package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/models"
)

const (
	dsn = "root:Adisunny123@tcp(127.0.0.1:3306)/nuclei?charset=utf8mb4&parseTime=True&loc=Local"
)

var db *gorm.DB

func InitializeDB() {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	db.AutoMigrate(&models.User{})
	log.Println("connected to database nuclei")
}
