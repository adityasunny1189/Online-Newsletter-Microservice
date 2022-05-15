package main

import (
	"context"
	"log"
	"net"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/database"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/database/pdb"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/database/sdb"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/models"
	pb "github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/protos"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/services"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type subscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
	db *gorm.DB
}

func (s *subscriptionServer) Init() {
	s.db = database.InitializeDB()
}

const (
	port = ":8080"
)

func main() {
	lst, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen at %v error: %v", port, err)
	}
	log.Printf("server running at port: %v", port)
	s := grpc.NewServer()
	srv := &subscriptionServer{}
	srv.Init()
	pb.RegisterSubscriptionServiceServer(s, srv)
	if err := s.Serve(lst); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *subscriptionServer) GetPlan(ctx context.Context, in *pb.PlanRequest) (*pb.PlanInfo, error) {
	log.Printf("Received: %v", in)
	res := &pb.PlanInfo{}

	// get plan by passing planid
	plan := pdb.GetPlanByName(s.db, in.Name)
	res.Id = plan.Id
	res.Name = plan.Name
	res.Price = plan.Price
	res.Validity = plan.Validity
	return res, nil
}

func (s *subscriptionServer) SubscribePlan(ctx context.Context, in *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	log.Printf("Received: %v", in)
	res := &pb.SubscriptionResponse{}

	// get plan by passing planid
	plan := pdb.GetPlanById(s.db, in.Planid)

	subs := models.Subscription{
		Uid:          in.Userid,
		Pid:          in.Planid,
		PlanValidity: plan.Validity,
		ExpiryDate:   services.CalculateExpiryDate(plan.Validity),
		IsActive:     true,
	}

	// add data to subscription database
	sid := sdb.AddSubscription(s.db, subs)
	res.Id = sid
	res.Userid = subs.Uid
	res.Planid = subs.Pid
	res.ExpiryDate = timestamppb.New(subs.ExpiryDate)
	res.IsActive = subs.IsActive
	return res, nil
}

func (s *subscriptionServer) CancelPlan(ctx context.Context, in *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	log.Printf("Received: %v", in)
	res := &pb.SubscriptionResponse{}

	// find the subscription using the planid and userid and update isActive to false
	subs := sdb.CancelSubscription(s.db, in.Planid, in.Userid)

	res.Id = subs.Id
	res.Userid = subs.Uid
	res.Planid = subs.Pid
	res.PlanValidity = subs.PlanValidity
	res.ExpiryDate = timestamppb.New(subs.ExpiryDate)
	res.IsActive = subs.IsActive
	return res, nil
}

func (s *subscriptionServer) RenewPlan(ctx context.Context, in *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	log.Printf("Received: %v", in)
	res := &pb.SubscriptionResponse{}

	// find the plan passing usedid and planid and update isActive to true and change expiry date
	subs := sdb.RenewSubscription(s.db, in.Planid, in.Userid)

	res.Id = subs.Id
	res.Userid = subs.Uid
	res.Planid = subs.Pid
	res.PlanValidity = subs.PlanValidity
	res.ExpiryDate = timestamppb.New(subs.ExpiryDate)
	res.IsActive = subs.IsActive
	return res, nil
}

func (s *subscriptionServer) GetPlans(in *pb.Empty, stream pb.SubscriptionService_GetPlansServer) error {
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

func (s *subscriptionServer) SortPlans(in *pb.SortRequest, stream pb.SubscriptionService_SortPlansServer) error {
	log.Printf("Received: %v", in)
	planlist := pdb.GetSortedPlan(s.db, in.Value)
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

func (s *subscriptionServer) GetAllSubscriptions(in *pb.Empty, stream pb.SubscriptionService_GetAllSubscriptionsServer) error {
	log.Printf("Received: %v", in)
	subscriptionlist := sdb.GetAllSubscriptions(s.db)
	for _, subsinfo := range subscriptionlist {
		subs := &pb.SubscriptionResponse{
			Id:           subsinfo.Id,
			Planid:       subsinfo.Pid,
			Userid:       subsinfo.Uid,
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

func (s *subscriptionServer) GetUserSubscriptions(in *pb.ID, stream pb.SubscriptionService_GetUserSubscriptionsServer) error {
	log.Printf("Received: %v", in)
	subscriptionlist := sdb.GetUserSubscriptions(s.db, in.Value)
	for _, subsinfo := range subscriptionlist {
		subs := &pb.SubscriptionResponse{
			Id:           subsinfo.Id,
			Planid:       subsinfo.Pid,
			Userid:       subsinfo.Uid,
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

func (s *subscriptionServer) GetUserActiveSubscriptions(in *pb.ID, stream pb.SubscriptionService_GetUserActiveSubscriptionsServer) error {
	log.Printf("Received: %v", in)
	subscriptionlist := sdb.GetUserActiveSubscriptions(s.db, in.Value)
	for _, subsinfo := range subscriptionlist {
		subs := &pb.SubscriptionResponse{
			Id:           subsinfo.Id,
			Planid:       subsinfo.Pid,
			Userid:       subsinfo.Uid,
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

func (s *subscriptionServer) GetUserPreviousSubscriptions(in *pb.ID, stream pb.SubscriptionService_GetUserPreviousSubscriptionsServer) error {
	log.Printf("Received: %v", in)
	subscriptionlist := sdb.GetUserPreviousSubscriptions(s.db, in.Value)
	for _, subsinfo := range subscriptionlist {
		subs := &pb.SubscriptionResponse{
			Id:           subsinfo.Id,
			Planid:       subsinfo.Pid,
			Userid:       subsinfo.Uid,
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
