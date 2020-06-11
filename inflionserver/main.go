//go:generate protoc -I inflionserverpb/ inflionserverpb/inflion.proto --go_out=plugins=grpc:inflionserverpb

package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/inflion/inflion/inflionserver/flow"
	pb "github.com/inflion/inflion/inflionserver/inflionserverpb"
	"github.com/inflion/inflion/inflionserver/rule"
	"github.com/inflion/inflion/internal/ops/flow/store"
	"github.com/inflion/inflion/internal/ops/rule/rulestore"
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

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_auth.UnaryServerInterceptor(ensureValidToken),
	)))

	pb.RegisterFlowServer(s, flow.DefaultFlowServer{Store: store.EtcdBackedFlowStore{}})
	pb.RegisterRuleServer(s, rule.DefaultRuleServer{Store: rulestore.EtcdStore{}})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
