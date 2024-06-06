package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"
)

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func main() {
	listenaddr := flag.String("listenAddr", ":6001", "")
	flag.Parse()

	http.HandleFunc("/invoice", makeAPIFunc(handleGetInvoice))
	logrus.Infof("Gateway HTTP server running on port %s", *listenaddr)
	log.Fatal(http.ListenAndServe(*listenaddr, nil))
}

func handleGetInvoice(w http.ResponseWriter, r *http.Request) error {
	return writeJSON(w, http.StatusOK, map[string]string{"invoice": "some invoice"})
}

func writeJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(v)
}

func makeAPIFunc(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}
