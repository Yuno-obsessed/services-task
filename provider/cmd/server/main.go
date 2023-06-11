package main

import (
	"fmt"
	"log"
	"net"
	"services-task/pkg/servicespb"
	"services-task/provider/service"

	"google.golang.org/grpc"
)

func main() {

	s := grpc.NewServer()

	servicespb.RegisterProviderServer(s, service.NewProviderServer())

	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
