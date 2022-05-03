package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb1 "github.com/Nviswateja/CustomerTrackerService/service/protos/Version1Protos"
	pb2 "github.com/Nviswateja/CustomerTrackerService/service/protos/Version2Protos"
)

type server struct {
	pb1.UnimplementedCustomerServiceServer
}

type serverV2 struct {
	pb2.UnimplementedCustomerServiceServer
}

var customerData = make([]string, 2)
var customerDetails = make([]*pb2.Customer, 2)

func (c *server) AddCustomerDetails(ctx context.Context, in *pb1.CustomerMessageRequest) (*pb1.CustomerMessageReply, error) {
	fmt.Println("Customer details:", in.Name)
	customerData = append(customerData, in.Name)
	customerDetails = append(customerDetails, &pb2.Customer{Name: in.Name, PhoneNumber: "123"})
	return &pb1.CustomerMessageReply{Message: "Hello " + in.GetName()}, nil
}

func (c *server) GetCustomers(ctx context.Context, in *pb1.GetCustomerMessageRequest) (*pb1.CustomerDetailsReply, error) {
	return &pb1.CustomerDetailsReply{Customers: customerData}, nil
}

func (c *serverV2) GetCustomerWithName(ctx context.Context, in *pb2.GetCustomerMessageRequest) (*pb2.CustomerDetailsReply, error) {
	return &pb2.CustomerDetailsReply{Customers: customerDetails}, nil
}

func main() {
	conn, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error while connecting to tcp %v", err)
	}

	grpcServer := grpc.NewServer()

	pb1.RegisterCustomerServiceServer(grpcServer, &server{})
	pb2.RegisterCustomerServiceServer(grpcServer, &serverV2{})
	log.Printf("server listening at %v", conn.Addr())
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
