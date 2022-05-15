package pdb

import (
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/models"
	"gorm.io/gorm"
)

func GetPlanById(db *gorm.DB, id int32) models.Plan {
	var plan models.Plan
	db.First(&plan, "id = ?", id)
	return plan
}

func GetPlanByName(db *gorm.DB, name string) models.Plan {
	var plan models.Plan
	db.First(&plan, "name = ?", name)
	return plan
}

func GetAllPlan(db *gorm.DB) []models.Plan {
	var plans []models.Plan
	db.Find(&plans)
	return plans
}

func GetSortedPlan(db *gorm.DB, sortType string) []models.Plan {
	var sortedPlans []models.Plan
	db.Order(sortType).Find(&sortedPlans)
	return sortedPlans
}
