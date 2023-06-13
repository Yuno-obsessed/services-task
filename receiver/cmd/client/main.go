package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"services-task/pkg/servicespb"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := servicespb.NewReceiverClient(conn)
	r := gin.Default()
	r.GET("/receive", func(ctx *gin.Context) {
		provideRes, err := http.Get("http://localhost:5001/provide")
		if err != nil {
			return
		}
		var request servicespb.SymbolsResponse
		err = json.NewDecoder(provideRes.Body).Decode(&request)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}
		res, err := client.Receive(ctx, &request)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		ctx.JSON(200, res)
	})
	err = r.Run(":5002")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
