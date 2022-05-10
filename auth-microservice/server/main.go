package main

import (
	"context"
	"log"
	"net"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/models"
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
	userslist := db.GetUsers()
	for _, userinfo := range userslist {
		user := &pb.UserInfo{
			Name:     userinfo.Name,
			Email:    userinfo.Email,
			Password: userinfo.Password,
		}
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (s *userServer) GetUser(ctx context.Context, in *pb.EmailInfo) (*pb.UserInfo, error) {
	log.Printf("Received: %v", in)
	userinfo := db.GetUser(in.Email)
	res := &pb.UserInfo{
		Name:     userinfo.Name,
		Email:    userinfo.Email,
		Password: userinfo.Password,
	}
	return res, nil
}

func (s *userServer) CreateUser(ctx context.Context, in *pb.UserInfo) (*pb.EmailInfo, error) {
	log.Printf("Received: %v", in)
	user := models.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
	db.AddUser(user)
	res := &pb.EmailInfo{
		Email: user.Email,
	}
	return res, nil
}
