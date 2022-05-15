package pdb

import (
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/models"
	"gorm.io/gorm"
)

func CreatePlan(db *gorm.DB, plan models.Plan) int32 {
	db.Create(&plan)
	return plan.Id
}

func GetAllPlan(db *gorm.DB) []models.Plan {
	var plans []models.Plan
	db.Find(&plans)
	return plans
}
