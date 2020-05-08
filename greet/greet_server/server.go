package main

import (
	"context"
	"fmt"
	"go-grpc/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// declare server
type server struct {
}

// implement GreetServiceServer interface
func (s server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetResponse, error) {
	firstName := req.Greeting.GetFirstName()
	lastName := req.Greeting.GetLastName()
	res := &greetpb.GreetResponse{
		Result: "hello " + firstName + " " + lastName,
	}
	return res, nil
}

func main() {
	fmt.Println("Server running at 50051")
	// create listener
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create grpc server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	// serve gRPC at listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
