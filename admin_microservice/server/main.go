package main

import (
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/database"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/grpcapi"
)

func main() {
	db := database.InitializeDB()
	grpcapi.Serve(db)
}
