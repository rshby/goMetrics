package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"goMetrics/logging"
	"goMetrics/middleware"
	"log"
	"net/http"
)

func init() {
	prometheus.MustRegister(middleware.GetVisitorCounter, middleware.GetLatency)
}

func main() {
	logConsole := logging.GenerateConsoleLogging()

	app := fiber.New()

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	v1 := app.Group("/v1")
	v1.Use(middleware.MetricsMiddleware())
	v1.Use(middleware.LoggerMiddleware(logConsole))

	v1.Get("/test", func(ctx *fiber.Ctx) error {
		statusCode := http.StatusOK
		ctx.Status(statusCode)
		return ctx.JSON(&map[string]any{
			"status_code": statusCode,
			"status":      "ok",
			"message":     "API ready to use",
		})
	})

	if err := app.Listen(":5005"); err != nil {
		log.Fatalf("cant run application : %v\n", err)
	} else {
		log.Println("success run application")
	}
}
