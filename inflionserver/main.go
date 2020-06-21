//go:generate protoc -I inflionserverpb/ inflionserverpb/inflion.proto --go_out=plugins=grpc:inflionserverpb

package main

import (
	"github.com/inflion/inflion/inflionserver/event"
	"github.com/inflion/inflion/inflionserver/flow"
	pb "github.com/inflion/inflion/inflionserver/inflionserverpb"
	"github.com/inflion/inflion/inflionserver/job"
	"github.com/inflion/inflion/inflionserver/rule"
	"github.com/inflion/inflion/internal/ops/flow/store"
	rule2 "github.com/inflion/inflion/internal/ops/rule"
	"google.golang.org/grpc"
	"log"
	"net"
)

const bind = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", bind)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterFlowInfoServer(s, flow.DefaultFlowServer{Store: store.EtcdBackedFlowStore{}})
	pb.RegisterRuleServer(s, rule.DefaultRuleServer{Store: rule2.EtcdStore{}})
	pb.RegisterJobInfoServer(s, job.NewJobServer())
	pb.RegisterEventServer(s, event.DefaultEventServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
