package main

import (
	"context"
	"fmt"
	"go-grpc/sum/sumpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

// declare server
type server struct {
}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	a := req.GetA()
	b := req.GetB()
	sum := a + b
	return &sumpb.SumResponse{
		Sum: sum,
	}, nil
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

	sumpb.RegisterSumServiceServer(s, &server{})

	// serve gRPC at listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
