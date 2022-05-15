package ndb

import (
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/models"
	"gorm.io/gorm"
)

func CreateNews(db *gorm.DB, news models.News) int32 {
	db.Create(&news)
	return news.Id
}
