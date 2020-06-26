package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	pb "github.com/inflion/inflion/inflionserver/event/eventpb"
	"google.golang.org/grpc"
	"log"
	"os"
)

type LambdaResponse struct {
	Message string `json:"Answer:"`
}

type InflionEvent struct {
	Project string          `json:"project"`
	Body    json.RawMessage `json:"body"`
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, event events.CloudWatchEvent) (LambdaResponse, error) {
	conn, err := grpc.Dial(os.Getenv("INFLIONSERVER_ADDR"), grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	client := pb.NewEventClient(conn)

	eventJson, err := json.Marshal(event)
	if err != nil {
		return LambdaResponse{Message: "failed marshal json"}, err
	}

	ie := InflionEvent{
		Project: "sandbox",
		Body:    eventJson,
	}

	ieJson, err := json.Marshal(ie)
	if err != nil {
		return LambdaResponse{Message: "failed marshal inflion json"}, err
	}

	message := &pb.PutEventRequest{
		Event: string(ieJson),
	}

	res, err := client.Put(ctx, message)
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
