package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HTTPStatusCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "healthcheck_http_status_total",
			Help: "HTTP status codes returned from health check endpoints",
		},
		[]string{"domain", "status"},
	)
)

func Register() {
	prometheus.MustRegister(HTTPStatusCounter)
}