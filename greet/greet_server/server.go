package main

import (
	"go-grpc/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
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
