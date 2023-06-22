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
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		defer provideRes.Body.Close()
		var provideResponse servicespb.ProvideLogsResponse
		err = json.NewDecoder(provideRes.Body).Decode(&provideResponse)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		receiveRequest := servicespb.ReceiveLogsRequest{
			Logs: &provideResponse,
		}

		receive, err := client.Receive(ctx, &receiveRequest)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		if receive.Status != 200 {
			ctx.JSON(500, gin.H{"err:": "error occured while storing in db"})
			return
		}
		ctx.JSON(200, receive)
	})
	r.GET("/fetch", func(ctx *gin.Context) {
		var filters servicespb.Filters
		err := json.NewDecoder(ctx.Request.Body).Decode(&filters)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		fmt.Println(filters)

		m, err := client.Fetch(ctx, &filters)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}

		ctx.JSONP(200, m)
	})
	r.DELETE("/delete/:id", func(ctx *gin.Context) {
		status, err := client.Delete(ctx, &servicespb.DeleteRequest{Id: ctx.Param("id")})
		if err != nil {
			ctx.JSON(500, gin.H{"err": err.Error()})
			return
		}
		ctx.JSON(int(status.Status), nil)
	})
	err = r.Run(":5002")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
