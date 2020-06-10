// Copyright 2020 The Inflion Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package clientv1

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

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
