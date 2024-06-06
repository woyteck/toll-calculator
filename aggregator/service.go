package main

import (
	"github.com/sirupsen/logrus"
	"github.com/woyteck/toll-calculator/types"
)

const basePrice = 3.15

type Aggregator interface {
	AggregateDistance(types.Distance) error
	CalculateInvoice(int) (*types.Invoice, error)
}

type Storer interface {
	Insert(types.Distance) error
	Get(int) (float64, error)
}

type InvoiceAggregator struct {
	store Storer
}

func NewInvoiceAggregator(store Storer) Aggregator {
	return &InvoiceAggregator{
		store: store,
	}
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	logrus.WithFields(logrus.Fields{
		"obuid":    distance.OBUID,
		"distance": distance.Value,
		"unix":     distance.Unix,
	}).Info("aggregatinr distance")
	return i.store.Insert(distance)
}

func (i *InvoiceAggregator) CalculateInvoice(obuid int) (*types.Invoice, error) {
	dist, err := i.store.Get(obuid)
	if err != nil {
		return nil, err
	}

	invoice := &types.Invoice{
		OBUID:         obuid,
		TotalDistance: dist,
		TotalAmount:   basePrice * dist,
	}

	return invoice, nil
}
