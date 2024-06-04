package main

import (
	"log"

	"github.com/woyteck/toll-calculator/aggregator/client"
)

const kafkaTopic = "obudata"
const aggregatorEndpoint = "http://localhost:3000/aggregate"

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc = NewCalculdatorService()
	svc = NewLogMiddleware(svc)

	// httpClient := client.NewHTTPClient(aggregatorEndpoint)
	grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, grpcClient)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
