package main

import (
	"log"

	"github.com/woyteck/toll-calculator/aggregator/client"
)

const kafkaTopic = "obudata"
const aggregatorEndpointHTTP = "http://127.0.0.1:3000/aggregate"
const aggregatorEndpointGRPC = "127.0.0.1:3001"

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculdatorService()
	svc = NewLogMiddleware(svc)

	// httpClient := client.NewHTTPClient(aggregatorEndpointHTTP)
	grpcClient, err := client.NewGRPCClient(aggregatorEndpointGRPC)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
