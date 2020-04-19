// Code generated by sqlc. DO NOT EDIT.

package timescale

import (
	"context"
	"time"
)

type Querier interface {
	Average(ctx context.Context, instanceID string) (AverageRow, error)
	CountCpuUtilization(ctx context.Context, arg CountCpuUtilizationParams) (int64, error)
	GetCpuUtilization(ctx context.Context) ([]Metric, error)
	InsertCpuUtilization(ctx context.Context, arg InsertCpuUtilizationParams) (time.Time, error)
	UpsertCpuUtilization(ctx context.Context, arg UpsertCpuUtilizationParams) (time.Time, error)
}

var _ Querier = (*Queries)(nil)
