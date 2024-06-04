package client

import (
	"context"

	"github.com/woyteck/toll-calculator/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
