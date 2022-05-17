package grpcapi

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/models"
	pb "github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/protos"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/services/db"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/utils"
	"google.golang.org/grpc"
)

const (
	port               = ":8080"
	InvalidPasswordErr = "incorrect password, please check your password"
	UserNotFoundErr    = "user not found"
	UserPresentErr     = "user with this email already present, try another email"
)

type userServer struct {
	pb.UnimplementedUserServer
}

func Serve() {
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
			Id:       userinfo.Id,
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
	res := &pb.UserInfo{}

	// Get user details from database
	userinfo := db.GetUser(in.Email)
	if userinfo.Email == "" {
		err := errors.New(UserNotFoundErr)
		return res, err
	}

	// If user found set the UserInfo object
	res.Id = userinfo.Id
	res.Name = userinfo.Name
	res.Email = userinfo.Email
	res.Password = userinfo.Password
	return res, nil
}

func (s *userServer) CreateUser(ctx context.Context, in *pb.UserInfo) (*pb.EmailInfo, error) {
	log.Printf("Received: %v", in)

	// Create the user model from the input data
	user := models.User{
		Name:     in.Name,
		Email:    in.Email,
		Password: utils.HashPassword(in.Password),
	}

	// Create a res object
	res := &pb.EmailInfo{}

	// Check for already present users
	u := db.GetUser(in.Email)
	if u.Name != "" {
		err := errors.New(UserPresentErr)
		return res, err
	}

	// Add user data to database
	db.AddUser(user)
	res.Email = in.Email
	return res, nil
}

func (s *userServer) LoginUser(ctx context.Context, in *pb.LoginInfo) (*pb.Status, error) {
	res := &pb.Status{
		Value: false,
	}

	// check for valid email and correct password
	user := db.GetUser(in.Email)
	if user.Email == "" {
		err := errors.New(UserNotFoundErr)
		return res, err
	} else if utils.CheckPasswordHash(in.Password, user.Password) {
		res.Value = true
	} else {
		err := errors.New(InvalidPasswordErr)
		return res, err
	}
	return res, nil
}
