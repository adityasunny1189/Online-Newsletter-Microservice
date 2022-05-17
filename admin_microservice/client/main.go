package main

import (
	"context"
	"log"

	pb "github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/protos"
	"google.golang.org/grpc"
)

const (
	Addr = "localhost:8000"
	body = `
	This is a test mail just testing my microservice
`
)

func main() {
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("erro connecting to port %v err %v", Addr, err)
	}
	defer conn.Close()
	client := pb.NewAdminServiceClient(conn)
	CreateNews(client, "Aditya Pathak", "Test Heading", body)
}

func CreateNews(client pb.AdminServiceClient, author, heading, content string) {
	ctx := context.Background()
	req := &pb.NewsInfo{
		Pid:     1,
		Author:  author,
		Heading: heading,
		Content: content,
	}
	res, err := client.CreateNews(ctx, req)
	if err != nil {
		log.Fatalf("error %v", err)
	}
	log.Printf("Response: %v", res)
}
