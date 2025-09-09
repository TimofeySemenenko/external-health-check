# External Health Check

A lightweight Go service for performing external health checks (HTTP endpoints) and exposing results as **Prometheus metrics**.

---

## Features

- Periodic HTTP checks for configured domains.
- Exposes results as Prometheus-compatible metrics.
- Built-in Go runtime metrics (`go_memstats`, `go_goroutines`, etc.).
- Simple to extend with new domains.

---

## Exposed Metrics

Example `/metrics` output:

- **Healthcheck metrics**
    - `healthcheck_http_status_total{domain, status}` – count of responses by HTTP status code.

- **Go runtime metrics**
    - `go_goroutines` – number of goroutines.
    - `go_gc_duration_seconds` – GC pause duration.
    - `go_memstats_alloc_bytes` – current memory usage.

- **Prometheus scrape metrics**
    - `promhttp_metric_handler_requests_total`
    - `promhttp_metric_handler_requests_in_flight`

Example:
```
text
healthcheck_http_status_total{domain="https://example.com",status="200 OK"} 15
healthcheck_http_status_total{domain="https://google.com",status="200 OK"} 15
```


Getting Started
Prerequisites

Go 1.21+

Git

(optional) Docker

Clone

```
git clone https://github.com/<your-org>/external-health-check.git
cd external-health-check
```

Run locally
```
go run cmd/external-health-check/main.go
```

By default, the service exposes metrics at:
```
http://localhost:8080/metrics
```

Build binary
```
go build -o external-health-check ./cmd/external-health-check
```

Run
```./external-health-check```

Docker
Build image
```docker build -t external-health-check .```

Run container
```
docker run --rm -p 8080:8080 --env-file .env external-health-check
```

Configuration

Domains for health checks can be configured via environment variables.

Example of configuration in configs/config.yaml:
```
domains:
  - https://google.com
  - https://example.com
```


Prometheus Integration
```
scrape_configs:
  - job_name: 'external-health-check'
    static_configs:
      - targets: ['external-health-check:8080']
```

Roadmap

Config file for domains

Alertmanager integration examples

Helm chart for Kubernetes deployment
