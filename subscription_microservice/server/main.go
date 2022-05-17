package main

import (
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/database"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/grpcapi"
)

func main() {
	db := database.InitializeDB()
	grpcapi.Serve(db)
}
