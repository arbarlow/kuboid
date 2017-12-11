package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "request_count",
		Help: "Requests per service",
	},
		[]string{"source_service", "source_version", "destination_service", "destination_version", "response_code"},
	)

	connections = []struct {
		source      string
		destination string
		weighting   float64
	}{
		{"unknown", "nginx-ingress", 1},
		{"nginx-ingress", "admin", 0.4},
		{"nginx-ingress", "graphl-api", 0.95},
		{"nginx-ingress", "rest-api", 0.85},
		{"nginx-ingress", "webhooks", 0.4},
		{"admin", "auth", 0.1},
		{"admin", "orders", 0.6},
		{"admin", "stock", 0.4},
		{"admin", "products", 0.4},
		{"admin", "shipping", 0.5},
		{"admin", "accounts", 0.3},
		{"admin", "payments", 0.2},
		{"admin", "notifications", 0.1},
		{"graphl-api", "auth", 0.9},
		{"graphl-api", "cart", 0.8},
		{"graphl-api", "accounts", 0.3},
		{"graphl-api", "orders", 0.5},
		{"graphl-api", "shipping", 0.3},
		{"graphl-api", "products", 0.8},
		{"graphl-api", "stock", 0.6},
		{"graphl-api", "payments", 0.6},
		{"graphl-api", "notifications", 0.3},
		{"rest-api", "auth", 0.8},
		{"rest-api", "cart", 0.7},
		{"rest-api", "accounts", 0.32},
		{"rest-api", "orders", 0.4},
		{"rest-api", "shipping", 0.2},
		{"rest-api", "products", 0.6},
		{"rest-api", "stock", 0.1},
		{"rest-api", "payments", 0.4},
		{"rest-api", "notifications", 0.1},
		{"accounts", "notifications", 0.2},
		{"auth", "notifications", 0.2},
		{"carts", "stock", 0.4},
		{"carts", "shipping", 0.3},
		{"orders", "products", 0.6},
		{"orders", "stock", 0.5},
		{"orders", "notifications", 0.5},
		{"orders", "payments", 0.3},
		{"orders", "carts", 0.3},
		{"orders", "carts", 0.3},
		{"shipping", "notifications", 0.2},
		{"products", "notifications", 0.1},
		{"stock", "notifications", 0.1},
		{"notifications", "emails", 0.3},
		{"notifications", "sms", 0.1},
		{"webhooks", "notifications", 0.3},
	}
)

func main() {
	rg := prometheus.NewRegistry()
	rg.MustRegister(
		requestTotal,
	)

	go genFakeMetrics()

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(rg, promhttp.HandlerOpts{}))

	addr := ":9000"

	fmt.Printf("Http server is running on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

// add in some random errors
func genFakeMetrics() {
	for {
		for _, metric := range connections {
			rps := (10 * rand.Float64() * 20) * metric.weighting

			code := "200"
			if rand.Int31n(10) > 8 {
				code = "503"
			}

			requestTotal.WithLabelValues(
				metric.source,
				"unknown",
				metric.destination,
				"unknown",
				code,
			).Add(rps)
		}

		time.Sleep(1 * time.Second)
	}
}
