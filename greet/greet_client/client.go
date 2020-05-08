package main

import (
	"fmt"
	"go-grpc/greet/greetpb"
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
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("created client %f", c)
}
