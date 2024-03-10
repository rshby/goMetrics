package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"strconv"
)

var GetVisitorCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_total_visitor",
		Help: "number of visitor",
	}, []string{"status_code", "status"})

var GetLatency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_request_get_duration",
		Help: "number of latency",
		//Buckets: prometheus.LinearBuckets(0.01, 0.05, 10),
	}, []string{"url"})

func MetricsMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Println("masuk middleware")

		var statusCode string
		timer := prometheus.NewTimer(prometheus.ObserverFunc(func(f float64) {
			GetLatency.WithLabelValues(string(ctx.Request().URI().FullURI())).Observe(f)
		}))

		defer func() {
			GetVisitorCounter.WithLabelValues(statusCode, "success").Inc()
			timer.ObserveDuration()
		}()

		ctx.Next()

		statusCode = strconv.Itoa(ctx.Response().StatusCode())

		log.Println("keluar middleware")
		return nil
	}
}
