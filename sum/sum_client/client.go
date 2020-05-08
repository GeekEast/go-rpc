package main

import (
	"context"
	"go-grpc/sum/sumpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	// request via connection
	c := sumpb.NewSumServiceClient(conn)
	doUnary(c)

}

func doUnary(c sumpb.SumServiceClient) {
	req := &sumpb.SumRequest{
		A: 32,
		B: 16,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greeting RPC %v", err)
	}
	log.Printf("Response from Greet: %v", res.GetSum())
}
