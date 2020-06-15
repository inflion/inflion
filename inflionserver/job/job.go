package job

import (
	"context"
	spb "github.com/inflion/inflion/inflionserver/inflionserverpb"
	cpb "github.com/inflion/inflion/jobserver/jobserverpb"
	"google.golang.org/grpc"
	"log"
	"time"
)

type JobServer struct {
	endpoint string
}

func NewJobServer() JobServer {
	return JobServer{
		endpoint: "localhost:50052",
	}
}

type grpcConnection struct {
	conn     *grpc.ClientConn
	endpoint string
}

func (c *grpcConnection) connect() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	c.conn, err = grpc.DialContext(ctx, c.endpoint, grpc.WithInsecure(), grpc.WithBlock())
	return err
}

func (c *grpcConnection) close() error {
	return c.conn.Close()
}

func (j JobServer) Create(ctx context.Context, request *spb.CreateJobRequest) (*spb.CreateJobResponse, error) {
	log.Println("job create")
	c := grpcConnection{
		endpoint: j.endpoint,
	}
	err := c.connect()
	if err != nil {
		return nil, err
	}

	jc := cpb.NewJobClient(c.conn)

	_, err = jc.Create(ctx, &cpb.CreateJobRequest{
		Id:       request.Id,
		Project:  request.Project,
		FlowId:   request.FlowId,
		Schedule: request.Schedule,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &spb.CreateJobResponse{
		Id: "created",
	}, nil
}

func (j JobServer) Remove(ctx context.Context, request *spb.RemoveJobRequest) (*spb.RemoveJobResponse, error) {
	log.Println("remove job")
	c := grpcConnection{
		endpoint: j.endpoint,
	}
	err := c.connect()
	if err != nil {
		return nil, err
	}

	jc := cpb.NewJobClient(c.conn)

	r, err := jc.Remove(ctx, &cpb.RemoveJobRequest{
		Id:      request.Id,
		Project: request.Project,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &spb.RemoveJobResponse{
		Id: r.Id,
	}, nil
}
