package main

import (
	"context"
	pb "github.com/inflion/inflion/inflionserver/byteevent/byteeventpb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
)

func main() {
	ctx := context.Background()

	rawEvent, err := ioutil.ReadFile("event_trustedadvisor.json")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	_, err = pb.NewByteEventClient(conn).Put(ctx, &pb.PutByteEventRequest{
		Project: "sandbox",
		Event:   rawEvent,
	})
	if err != nil {
		log.Fatal(err)
	}
}
