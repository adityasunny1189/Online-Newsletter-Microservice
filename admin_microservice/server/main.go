package main

import (
	"context"
	"log"
	"net"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/database"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/database/ndb"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/database/pdb"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/database/sdb"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/database/udb"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/models"
	pb "github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/admin_microservice/protos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

const (
	port = ":8000"
)

type adminServer struct {
	pb.UnimplementedAdminServiceServer
	db *gorm.DB
}

func (s *adminServer) Init() {
	s.db = database.InitializeDB()
}

func main() {
	lst, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error connecting at port %v error %v", port, err)
	}
	log.Printf("server up and running at port: %v", port)
	s := grpc.NewServer()
	srv := &adminServer{}
	srv.Init()
	pb.RegisterAdminServiceServer(s, srv)
	if err := s.Serve(lst); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}

func (s *adminServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.Status, error) {
	log.Printf("Received: %v", in)
	res := &pb.Status{}
	if in.Password == "admin1234" {
		res.Value = true
	}
	return res, nil
}

func (s *adminServer) CreatePlan(ctx context.Context, in *pb.PlanInfo) (*pb.ID, error) {
	log.Printf("Received: %v", in)
	res := &pb.ID{}
	plan := models.Plan{
		Name:     in.Name,
		Price:    in.Price,
		Validity: in.Validity,
	}
	res.Id = pdb.CreatePlan(s.db, plan)
	return res, nil
}

func (s *adminServer) CreateNews(ctx context.Context, in *pb.NewsInfo) (*pb.ID, error) {
	log.Printf("Received: %v", in)
	res := &pb.ID{}
	news := models.News{
		Planid:  in.Pid,
		Author:  in.Author,
		Heading: in.Heading,
		Content: in.Content,
	}
	res.Id = ndb.CreateNews(s.db, news)
	return res, nil
}

func (s *adminServer) GetPlans(in *pb.Empty, stream pb.AdminService_GetPlansServer) error {
	log.Printf("Received: %v", in)
	planlist := pdb.GetAllPlan(s.db)
	for _, planinfo := range planlist {
		plan := &pb.PlanInfo{
			Id:       planinfo.Id,
			Name:     planinfo.Name,
			Price:    planinfo.Price,
			Validity: planinfo.Validity,
		}
		if err := stream.Send(plan); err != nil {
			return err
		}
	}
	return nil
}

func (s *adminServer) GetUsers(in *pb.Empty, stream pb.AdminService_GetUsersServer) error {
	log.Printf("Received: %v", in)
	userlist := udb.GetAllUser(s.db)
	for _, userinfo := range userlist {
		user := &pb.UserInfo{
			Id:    userinfo.Id,
			Name:  userinfo.Name,
			Email: userinfo.Email,
		}
		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (s *adminServer) GetSubscriptions(in *pb.Empty, stream pb.AdminService_GetSubscriptionsServer) error {
	log.Printf("Received: %v", in)
	subslist := sdb.GetAllSubscription(s.db)
	for _, subsinfo := range subslist {
		subs := &pb.SubscriptionInfo{
			Id:           subsinfo.Id,
			Userid:       subsinfo.Uid,
			Planid:       subsinfo.Pid,
			PlanValidity: subsinfo.PlanValidity,
			ExpiryDate:   timestamppb.New(subsinfo.ExpiryDate),
			IsActive:     subsinfo.IsActive,
		}
		if err := stream.Send(subs); err != nil {
			return err
		}
	}
	return nil
}
