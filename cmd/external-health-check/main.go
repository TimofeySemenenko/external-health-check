package main

import (
	"log"
	"net/http"
	"time"

	"external-health-check/config"
	"external-health-check/internal/checker"
	"external-health-check/internal/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.Load("configs/config.yaml")
	metrics.Register()

	checker.Start(cfg.Domains, 5*time.Second)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Starting metrics server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}