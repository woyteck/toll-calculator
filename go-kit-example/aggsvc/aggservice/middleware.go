package aggservice

import (
	"context"

	"github.com/woyteck/toll-calculator/types"
)

type Middleware func(Service) Service

type loggingMiddleware struct {
	next Service
}

func newLoggingMiddleware() Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next: next,
		}
	}
}

func (mw *loggingMiddleware) Aggregate(_ context.Context, distance types.Distance) error {
	return nil
}

func (mw *loggingMiddleware) Calculate(_ context.Context, obuID int) (*types.Invoice, error) {
	return nil, nil
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

func (mw *instrumentationMiddleware) Aggregate(_ context.Context, distance types.Distance) error {
	return nil
}

func (mw *instrumentationMiddleware) Calculate(_ context.Context, obuID int) (*types.Invoice, error) {
	return nil, nil
}
