package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

type Middlewares struct {
	RequestCounter  *prometheus.CounterVec
	RequestDuration *prometheus.HistogramVec
}

// NewMiddleware
func NewMiddleware(counter *prometheus.CounterVec, duration *prometheus.HistogramVec) *Middlewares {
	return &Middlewares{
		RequestCounter:  counter,
		RequestDuration: duration,
	}
}

// Measure
func (m *Middlewares) Measure(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		defer func() {
			duration := time.Since(start).Seconds()
			service := c.Request().Header.Get("X-From-Service")
			m.RequestDuration.With(
				prometheus.Labels{
					"service": service,
				},
			).Observe(duration)
		}()

		service := c.Request().Header.Get("X-From-Service")
		m.RequestCounter.With(
			prometheus.Labels{
				"service": service,
			},
		).Inc()

		return next(c)
	}
}
