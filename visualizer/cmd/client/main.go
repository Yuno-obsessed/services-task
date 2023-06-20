package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"services-task/pkg/servicespb"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	client := servicespb.NewVisualizerClient(conn)
	r := gin.Default()
	r.GET("/visualize", func(ctx *gin.Context) {

		var filters servicespb.Filters
		err := json.NewDecoder(ctx.Request.Body).Decode(&filters)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}

		payload, err := json.Marshal(&filters)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		provideReq, err := http.NewRequest("GET", "http://localhost:5002/fetch", bytes.NewBuffer(payload))
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}

		provideRes, err := http.DefaultClient.Do(provideReq)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		defer provideRes.Body.Close()
		var provideResponse servicespb.FetchResponse
		err = json.NewDecoder(provideRes.Body).Decode(&provideResponse)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		visualizeRequest := servicespb.Filters{}

		res, err := client.Visualize(ctx, &visualizeRequest)
		if err != nil {
			ctx.JSON(500, gin.H{"err:": err.Error()})
			return
		}
		ctx.JSON(200, res)
	})
	err = r.Run(":5003")
	if err != nil {
		log.Fatalf("%v", err)
	}
}