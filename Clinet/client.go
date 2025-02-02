package main

import (
	"context"
	"log"
	"time"

	pb "github.com/valdimir-makarov/Gprc-with-Go/hello" // Correct import path
	"google.golang.org/grpc"
)

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a client
	client := pb.NewGreeterClient(conn)

	// Prepare the request
	req := &pb.HelloRequest{Name: "World"}

	// Call the SayHello method
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	// Print the response
	log.Printf("Response: %s", res.GetMessage())
}
