package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// RequestCounter
func RequestCounter() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "receiver_request_count",
			Help: "Total number of requests to Receiver",
		},
		[]string{"service", "method", "route", "status"},
	)
}

// DurationCounter
func DurationCounter() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "receiver_request_duration",
			Help:    "Duration of requests to Receiver",
			Buckets: []float64{0.1, 0.5, 1, 2, 5},
		},
		[]string{"service", "method", "route", "status"},
	)
}
