package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/woyteck/toll-calculator/types"
)

type HTTPMetricsHandler struct {
	reqCounter prometheus.Counter
	reqLatency prometheus.Histogram
}

func newHTTPMetricsHandler(reqName string) *HTTPMetricsHandler {
	reqCounter := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: fmt.Sprintf("http_%s_%s", reqName, "request_counter"),
		Name:      "aggregator",
	})
	reqLatency := promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: fmt.Sprintf("http_%s_%s", reqName, "request_latency"),
		Name:      "aggregator",
		Buckets:   []float64{0.1, 0.5, 1},
	})

	return &HTTPMetricsHandler{
		reqCounter: reqCounter,
		reqLatency: reqLatency,
	}
}

func (h *HTTPMetricsHandler) instrument(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time) {
			h.reqLatency.Observe(time.Since(start).Seconds())
		}(time.Now())

		h.reqCounter.Inc()

		next(w, r)
	}
}

func handleGetInvoice(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "method not supported"})
			return
		}

		values, ok := r.URL.Query()["obu"]
		if !ok {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing obuid"})
			return
		}
		obuID, err := strconv.Atoi(values[0])
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid ouid"})
			return
		}

		invoice, err := svc.CalculateInvoice(obuID)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, invoice)
	}
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "method not supported"})
			return
		}

		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}
