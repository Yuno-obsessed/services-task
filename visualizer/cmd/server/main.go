package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"services-task/pkg/servicespb"
	"services-task/visualizer/service"
)

func main() {
	s := grpc.NewServer()

	servicespb.RegisterVisualizerServer(s, service.NewVisualizerService())

	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50053))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
