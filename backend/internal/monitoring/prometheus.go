package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
	httpRequests        prometheus.Counter
	httpRequestDuration prometheus.Histogram
	memoryUsage         prometheus.Gauge
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		httpRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		}),
		httpRequestDuration: prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.DefBuckets, // Default latency buckets
		}),
		memoryUsage: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "memory_usage_bytes",
			Help: "Current memory usage in bytes.",
		}),
	}
	reg.MustRegister(m.httpRequests)
	reg.MustRegister(m.httpRequestDuration)
	reg.MustRegister(m.memoryUsage)
	return m
}
