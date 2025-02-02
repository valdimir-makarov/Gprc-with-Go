package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/valdimir-makarov/Gprc-with-Go/hello"
	"google.golang.org/grpc"
)

// Define the server struct
type server struct {
	pb.UnimplementedGreeterServer
}

// Implement the SayHello method
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.GetName() + "!"}, nil
}

func main() {
	// Start listening on a port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	grpcServer := grpc.NewServer()

	// Register the Greeter service
	pb.RegisterGreeterServer(grpcServer, &server{})

	// Start the server
	fmt.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
