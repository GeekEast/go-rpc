package main

import (
	"context"
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
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "James",
			LastName:  "Tan",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greeting RPC %v", err)
	}
	log.Printf("Response from Greet: %v", res.GetResult())
}
