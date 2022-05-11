package main

import (
	"context"
	"io"
	"log"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/models"
	pb "github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/protos"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connection error %v", err)
	}
	defer conn.Close()
	client := pb.NewUserClient(conn)

	// GetUsers(client)
	// GetUser(client, "adityapathak1189@gmail.com")

	// dummyUser := models.User{
	// 	Name:     "Amit Shahwal",
	// 	Email:    "amitshahwal24@gmail.com",
	// 	Password: "password98",
	// }
	// Signup(client, dummyUser)

	Login(client, "adityapathak1189@gmail.com", "1234pass")
}

func GetUsers(client pb.UserClient) {
	ctx := context.Background()
	req := &pb.Empty{}
	stream, err := client.GetUsers(ctx, req)
	if err != nil {
		log.Fatalf("error fetching users %v", err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error %v", err)
		}
		log.Printf("User Info: %v", row)
	}
}

func GetUser(client pb.UserClient, email string) {
	ctx := context.Background()
	req := &pb.EmailInfo{
		Email: email,
	}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("error getting user details %v", err)
	}
	log.Printf("User Info: %v", res)
}

func Signup(client pb.UserClient, user models.User) {
	ctx := context.Background()
	req := &pb.UserInfo{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
	res, err := client.CreateUser(ctx, req)
	if err != nil {
		log.Fatalf("error signing up user %v", err)
	}
	log.Printf("Response: %v", res)
}

func Login(client pb.UserClient, email, paswd string) {
	ctx := context.Background()
	req := &pb.LoginInfo{
		Email:    email,
		Password: paswd,
	}
	res, err := client.LoginUser(ctx, req)
	if err != nil {
		log.Fatalf("error logging user %v", err)
	}
	log.Printf("Response: %v", res)
}
