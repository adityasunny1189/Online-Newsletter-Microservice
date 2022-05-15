package udb

import (
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/models"
	"gorm.io/gorm"
)

func GetAllUser(db *gorm.DB) []models.User {
	var users []models.User
	db.Find(&users)
	return users
}
