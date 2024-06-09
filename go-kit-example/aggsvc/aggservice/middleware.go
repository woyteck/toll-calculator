package aggservice

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"github.com/woyteck/toll-calculator/types"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	log  log.Logger
	next Service
}

func newLoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next: next,
			log:  logger,
		}
	}
}

func (mw *loggingMiddleware) Aggregate(ctx context.Context, distance types.Distance) (err error) {
	defer func(start time.Time) {
		mw.log.Log("method", "Aggregate", "took", time.Since(start), "obu", distance.OBUID, "distance", distance.Value, "err", err)
	}(time.Now())
	err = mw.next.Aggregate(ctx, distance)

	return
}

func (mw *loggingMiddleware) Calculate(ctx context.Context, obuID int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		mw.log.Log("method", "Calculate", "took", time.Since(start), "obuID", obuID, "inv", inv, "err", err)
	}(time.Now())

	inv, err = mw.next.Calculate(ctx, obuID)
	return
}

type instrumentationMiddleware struct {
	next Service
}

func newInstrumentationMiddleware() Middleware {
	return func(next Service) Service {
		return &instrumentationMiddleware{
			next: next,
		}
	}
}

func (mw *instrumentationMiddleware) Aggregate(ctx context.Context, distance types.Distance) error {
	return mw.next.Aggregate(ctx, distance)
}

func (mw *instrumentationMiddleware) Calculate(ctx context.Context, obuID int) (*types.Invoice, error) {
	return mw.next.Calculate(ctx, obuID)
}
