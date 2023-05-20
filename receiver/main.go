package main

import (
	"net/http"

	"go-prom-sender-receiver-measure/receiver/metrics"
	"go-prom-sender-receiver-measure/receiver/middleware"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var promMiddlewares *middleware.Middlewares

func init() {
	requestCount := metrics.RequestCounter()
	requestDuration := metrics.DurationCounter()
	promMiddlewares = middleware.NewMiddleware(requestCount, requestDuration)

	prometheus.MustRegister(requestCount)
	prometheus.MustRegister(requestDuration)
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "response sending from reveiver",
		})
	}, promMiddlewares.Measure)

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	if err := e.Start(":8002"); err != nil {
		panic(err)
	}
}
