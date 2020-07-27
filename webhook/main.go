package main

import (
	"bytes"
	"context"
	"fmt"
	pb "github.com/inflion/inflion/inflionserver/byteevent/byteeventpb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net/http"
)

func main() {
	target := "localhost:5005"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal("byteEventClient connection error:", err)
	}

	defer func() {
		err = conn.Close()
		fmt.Printf("error:%#v \n", err)
	}()

	c := newGrpcClient(pb.NewByteEventClient(conn))

	http.HandleFunc("/webhook", c.handleWebhook)
	addr := ":8082"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("error:%#v \n", err)
	}
	log.Printf("listening on: %v", addr)
}

type webhookEventGrpcAdapter struct {
	byteEventClient pb.ByteEventClient
}

func newGrpcClient(client pb.ByteEventClient) *webhookEventGrpcAdapter {
	return &webhookEventGrpcAdapter{byteEventClient: client}
}

func (c webhookEventGrpcAdapter) handleWebhook(_ http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		fmt.Printf("error:%#v \n", err)
	}()

	b, err := toBytes(r.Body)
	if err != nil {
		fmt.Printf("error:%#v \n", err)
	}

	c.putByteEvent(b)
}

func toBytes(body io.Reader) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c webhookEventGrpcAdapter) putByteEvent(event []byte) {
	message := &pb.PutByteEventRequest{
		Event: event,
	}
	res, err := c.byteEventClient.Put(context.Background(), message)
	if err != nil {
		fmt.Printf("error:%#v \n", err)
	}
	fmt.Printf("result:%+v \n", res)
}
