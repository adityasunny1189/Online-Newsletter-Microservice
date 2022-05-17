package main

import (
	"context"
	"io"
	"log"

	pb "github.com/adityasunny1189/gRPC-GORM-Auth-Microservice/subscription-microservice/protos"
	"google.golang.org/grpc"
)

const (
	Addr = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error connecting %v", err)
	}
	defer conn.Close()
	client := pb.NewSubscriptionServiceClient(conn)

	// GetAllPlan(client)
	// GetPlan(client, "weekly")
	// GetSortedPlans(client, "validity")
	// GetSortedPlans(client, "price")
	SubscribePlan(client, 1, 1)
	SubscribePlan(client, 2, 1)
	SubscribePlan(client, 3, 1)
	SubscribePlan(client, 4, 1)
	SubscribePlan(client, 5, 1)
	SubscribePlan(client, 6, 1)
	// CancelPlan(client, 1, 3)
	// RenewPlan(client, 1, 3)
	// GetAllSubscriptions(client)
	// GetUserSubscriptions(client, 1)
	// GetUserActiveSubscriptions(client, 1)
	// GetUserPreviousSubscriptions(client, 1)
}

func GetAllPlan(client pb.SubscriptionServiceClient) {
	ctx := context.Background()
	req := &pb.Empty{}
	stream, err := client.GetPlans(ctx, req)
	if err != nil {
		log.Fatalf("error fetching plans: %v", err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error %v", err)
		}
		log.Printf("Plan Info: %v", row)
	}
}

func GetPlan(client pb.SubscriptionServiceClient, name string) {
	ctx := context.Background()
	req := &pb.PlanRequest{
		Name: name,
	}
	res, err := client.GetPlan(ctx, req)
	if err != nil {
		log.Println("erro getting plan %v", err)
	}
	log.Printf("Response: %v", res)
}

func GetSortedPlans(client pb.SubscriptionServiceClient, value string) {
	ctx := context.Background()
	req := &pb.SortRequest{
		Value: value,
	}
	stream, err := client.SortPlans(ctx, req)
	if err != nil {
		log.Fatalf("error fetching plans %v", err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error %v", err)
		}
		log.Printf("Plan Info: %v", row)
	}
}

func SubscribePlan(client pb.SubscriptionServiceClient, uid, pid int32) {
	ctx := context.Background()
	req := &pb.SubscriptionRequest{
		Userid: uid,
		Planid: pid,
	}
	res, err := client.SubscribePlan(ctx, req)
	if err != nil {
		log.Fatalf("error subscribing plan %v", err)
	}
	log.Printf("Response: %v", res)
}

func CancelPlan(client pb.SubscriptionServiceClient, uid, pid int32) {
	ctx := context.Background()
	req := &pb.SubscriptionRequest{
		Userid: uid,
		Planid: pid,
	}
	res, err := client.CancelPlan(ctx, req)
	if err != nil {
		log.Fatalf("error canceling plan %v", err)
	}
	log.Printf("Response: %v", res)
}

func RenewPlan(client pb.SubscriptionServiceClient, uid, pid int32) {
	ctx := context.Background()
	req := &pb.SubscriptionRequest{
		Userid: uid,
		Planid: pid,
	}
	res, err := client.RenewPlan(ctx, req)
	if err != nil {
		log.Fatalf("error renewing plan %v", err)
	}
	log.Printf("Response: %v", res)
}

func GetAllSubscriptions(client pb.SubscriptionServiceClient) {
	ctx := context.Background()
	req := &pb.Empty{}
	stream, err := client.GetAllSubscriptions(ctx, req)
	if err != nil {
		log.Fatalf("error getting subscription %v", err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error %v", err)
		}
		log.Printf("Subscription Info: %v", row)
	}
}

func GetUserSubscriptions(client pb.SubscriptionServiceClient, uid int32) {
	ctx := context.Background()
	req := &pb.ID{
		Value: uid,
	}
	stream, err := client.GetUserSubscriptions(ctx, req)
	if err != nil {
		log.Fatalf("error getting subscription %v", err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error %v", err)
		}
		log.Printf("Subscription Info: %v", row)
	}
}

func GetUserActiveSubscriptions(client pb.SubscriptionServiceClient, uid int32) {
	ctx := context.Background()
	req := &pb.ID{
		Value: uid,
	}
	stream, err := client.GetUserSubscriptions(ctx, req)
	if err != nil {
		log.Fatalf("error getting subscription %v", err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error %v", err)
		}
		log.Printf("Subscription Info: %v", row)
	}
}

func GetUserPreviousSubscriptions(client pb.SubscriptionServiceClient, uid int32) {
	ctx := context.Background()
	req := &pb.ID{
		Value: uid,
	}
	stream, err := client.GetUserSubscriptions(ctx, req)
	if err != nil {
		log.Fatalf("error getting subscription %v", err)
	}
	for {
		row, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error %v", err)
		}
		log.Printf("Subscription Info: %v", row)
	}
}
