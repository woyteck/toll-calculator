package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/woyteck/toll-calculator/types"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("AggregateDistance")
	}(time.Now())

	err = m.next.AggregateDistance(distance)

	return
}

func (m *LogMiddleware) CalculateInvoice(obuid int) (invoice *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if invoice != nil {
			distance = invoice.TotalDistance
			amount = invoice.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took":        time.Since(start),
			"err":         err,
			"obuID":       obuid,
			"totalDist":   distance,
			"totalAmount": amount,
		}).Info("CalculateInvoice")
	}(time.Now())

	invoice, err = m.next.CalculateInvoice(obuid)

	return
}
