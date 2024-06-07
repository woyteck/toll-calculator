package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/woyteck/toll-calculator/types"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	httpListenAddr := os.Getenv("AGG_HTTP_ENDPOINT")
	grpcListenAddr := os.Getenv("AGG_GPRC_ENDPOINT")
	store := makeStore()

	svc := NewInvoiceAggregator(store)
	svc = NewMetricsMiddleware(svc)
	svc = NewLogMiddleware(svc)

	go func() {
		log.Fatal(makeGRPCTransport(grpcListenAddr, svc))
	}()

	log.Fatal(makeHTTPTransport(httpListenAddr, svc))
}

func makeGRPCTransport(listenAddr string, svc Aggregator) error {
	fmt.Println("GRPC transport running on port", listenAddr)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	defer ln.Close()

	server := grpc.NewServer([]grpc.ServerOption{}...)
	types.RegisterAggregatorServer(server, NewGRPCAggregatorServer(svc))

	return server.Serve(ln)
}

func makeHTTPTransport(listenAddr string, svc Aggregator) error {
	aggMetricHandler := newHTTPMetricsHandler("aggregate")
	aggregateHandler := makeHTTPHandlerFunc(aggMetricHandler.instrument(handleAggregate(svc)))
	http.HandleFunc("/aggregate", aggregateHandler)

	invMetricHandler := newHTTPMetricsHandler("invoice")
	invoiceHandler := makeHTTPHandlerFunc(invMetricHandler.instrument(handleGetInvoice(svc)))
	http.HandleFunc("/invoice", invoiceHandler)

	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("HTTP transport running on port", listenAddr)
	return http.ListenAndServe(listenAddr, nil)
}

func makeStore() Storer {
	t := os.Getenv("AGG_STORE_TYPE")
	switch t {
	case "memory":
		return NewMemoryStore()
	default:
		log.Fatalf("invalid store %s", t)
		return nil
	}
}

func writeJSON(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(rw).Encode(v)
}
