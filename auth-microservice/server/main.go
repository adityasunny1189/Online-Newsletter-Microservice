package main

import (
	"context"
	"log"
	"net"

	pb "github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/protos"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/services/db"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type userServer struct {
	pb.UnimplementedUserServer
}

func main() {
	db.InitializeDB()
	lst, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listed: %v", err)
	}
	log.Printf("server running at port %v\n", port)
	s := grpc.NewServer()
	pb.RegisterUserServer(s, &userServer{})
	if err := s.Serve(lst); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *userServer) GetUsers(in *pb.Empty, stream pb.User_GetUsersServer) error {
	log.Printf("Received: %v", in)

	return nil
}

func (s *userServer) GetUser(ctx context.Context, in *pb.ID) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in)
	res := &pb.UserInfo{
		Id:       1,
		Name:     "Aditya",
		Email:    "adi@gmail.com",
		Password: "1234",
	}
	return res, nil
}

func (s *userServer) CreateUser(ctx context.Context, in *pb.UserInfo) (*pb.ID, error) {
	log.Printf("Received: %v", in)
	res := &pb.ID{
		Val: 1,
	}
	return res, nil
}
