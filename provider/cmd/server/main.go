package main

import (
	"fmt"
	"github.com/Yuno-obsessed/services-task/pkg/servicespb"
	"github.com/Yuno-obsessed/services-task/provider/service"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	go initClient()

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

func initClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := servicespb.NewProviderClient(conn)
	r := gin.Default()
	r.GET("/provide", func(ctx *gin.Context) {
		res, err := client.Provide(ctx, &servicespb.ProvideLogsRequest{})
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		ctx.JSON(200, res)
	})
	err = r.Run(":5001")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
