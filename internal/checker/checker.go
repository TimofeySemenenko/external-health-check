package checker

import (
	"log"
	"net/http"
	"time"

	"external-health-check/internal/metrics"
)

func Start(domains []string, interval time.Duration) {
	go func() {
		for {
			for _, domain := range domains {
				go check(domain)
			}
			time.Sleep(interval)
		}
	}()
}

func check(domain string) {
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(domain)
	if err != nil {
		log.Printf("[ERROR] %s - %v\n", domain, err)
		metrics.HTTPStatusCounter.WithLabelValues(domain, "error").Inc()
		return
	}
	defer resp.Body.Close()

	statusCode := resp.Status
	log.Printf("[INFO] %s - %s\n", domain, statusCode)
	metrics.HTTPStatusCounter.WithLabelValues(domain, statusCode).Inc()
}