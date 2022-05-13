package main

import (
	"context"
	"log"
	"net"

	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/database/pdb"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/database/sdb"
	"github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/models"
	pb "github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/protos"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type subscriptionServer struct {
	pb.UnimplementedSubscriptionServiceServer
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
	pb.RegisterSubscriptionServiceServer(s, &subscriptionServer{})
	if err := s.Serve(lst); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *subscriptionServer) GetPlan(ctx context.Context, in *pb.PlanRequest) (*pb.PlanInfo, error) {
	log.Printf("Received: %v", in)
	res := &pb.PlanInfo{}

	// get plan by passing planid
	var plan models.Plan
	plan = pdb.GetPlanByName(in.Name)
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
	var plan models.Plan
	plan = pdb.GetPlanById(in.Planid)

	subs := models.Subscription{
		Uid:          in.Userid,
		Pid:          in.Planid,
		PlanValidity: plan.Validity,
		ExpiryDate:   calculateExpiryDate(plan.Validity),
		IsActive:     true,
	}

	// add data to subscription database
	sid := sdb.AddSubscription(subs)
	res.Id = sid
	res.Userid = subs.Uid
	res.Planid = subs.Pid
	res.ExpiryDate = timestamppb.New(*subs.ExpiryDate)
	res.IsActive = subs.IsActive
	return res, nil
}

func (s *subscriptionServer) CancelPlan(ctx context.Context, in *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	log.Printf("Received: %v", in)
	res := &pb.SubscriptionResponse{}

	// find the subscription using the planid and userid and update isActive to false
	var subs models.Subscription
	subs = sdb.CancelSubscription(in.Planid, in.Userid)

	res.Id = subs.Id
	res.Userid = subs.Uid
	res.Planid = subs.Pid
	res.PlanValidity = subs.PlanValidity
	res.ExpiryDate = timestamppb.New(*subs.ExpiryDate)
	res.IsActive = subs.IsActive
	return res, nil
}

func (s *subscriptionServer) RenewPlan(ctx context.Context, in *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	log.Printf("Received: %v", in)
	res := &pb.SubscriptionResponse{}

	// find the plan passing usedid and planid and update isActive to true and change expiry date
	var subs models.Subscription
	subs = sdb.RenewSubscription(in.Planid, in.Userid)

	res.Id = subs.Id
	res.Userid = subs.Uid
	res.Planid = subs.Pid
	res.PlanValidity = subs.PlanValidity
	res.ExpiryDate = timestamppb.New(*subs.ExpiryDate)
	res.IsActive = subs.IsActive
	return res, nil
}

func (s *subscriptionServer) GetPlans() {}

func (s *subscriptionServer) SortPlans() {}

func (s *subscriptionServer) GetAllSubscriptions() {}

func (s *subscriptionServer) GetUserSubscriptions() {}

func (s *subscriptionServer) GetUserActiveSubscriptions() {}

func (s *subscriptionServer) GetUserPreviousSubscriptions() {}
