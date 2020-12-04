package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	pb "github.com/inflion/inflion/inflionserver/byteevent/byteeventpb"
	"google.golang.org/grpc"
	"log"
	"os"
)

type LambdaResponse struct {
	Message string `json:"Answer:"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, event events.CloudWatchEvent) (LambdaResponse, error) {
	eventJson, err := json.Marshal(event)
	if err != nil {
		return LambdaResponse{Message: "failed marshal json"}, err
	}

	eventJsonForLog, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("EVENT: %s", eventJsonForLog)

	conn, err := grpc.Dial(os.Getenv("INFLIONSERVER_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	res, err := pb.NewByteEventClient(conn).Put(ctx, &pb.PutByteEventRequest{
		Project: "sandbox",
		Event:   eventJson,
	})
	if err != nil {
		return LambdaResponse{
			fmt.Sprintf("error::%#v \n", err),
		}, err
	}

	fmt.Printf("result:%+v \n", res)
	return LambdaResponse{
		"ok",
	}, nil
}
