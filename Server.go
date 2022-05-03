package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Nviswateja/CustomerTrackerService/service/protos"
)

type server struct {
	pb.UnimplementedCustomerServiceServer
}

func (c *server) AddCustomerDetails(ctx context.Context, in *pb.CustomerMessageRequest) (*pb.CustomerMessageReply, error) {
	fmt.Println("Customer details:", in.Name)
	return &pb.CustomerMessageReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error while connecting to tcp %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCustomerServiceServer(grpcServer, &server{})
	log.Printf("server listening at %v", conn.Addr())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
