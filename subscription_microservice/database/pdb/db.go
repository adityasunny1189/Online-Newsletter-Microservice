package pdb

import (
	"log"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dsn = "root:Adisunny123@tcp(127.0.0.1:3306)/nuclei?charset=utf8mb4&parseTime=True&loc=Local"
)

type PDB struct {
	db *gorm.DB
}

var pdb PDB

func GetPlanById(id int32) models.Plan {
	var plan models.Plan
	pdb.db.First(&plan, "id = ?", id)
	return plan
}

func GetPlanByName(name string) models.Plan {
	var plan models.Plan
	pdb.db.First(&plan, "name = ?", name)
	return plan
}

func InitializeDB() {
	var err error
	pdb.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	pdb.db.AutoMigrate(&models.Plan{})
	log.Printf("connected to database")
}
