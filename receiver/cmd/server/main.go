package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"services-task/pkg/servicespb"
	"services-task/receiver/service"
)

func main() {
	err := godotenv.Load()
	s := grpc.NewServer()

	servicespb.RegisterReceiverServer(s, service.NewReceiverService())

	fmt.Println("gRPC server running ...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50052))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
