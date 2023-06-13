package main

import (
	"log"
	"services-task/pkg/servicespb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := servicespb.NewProviderClient(conn)
	r := gin.Default()
	r.GET("/provide", func(ctx *gin.Context) {
		res, err := client.Provide(ctx, &servicespb.SymbolsRequest{})
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
