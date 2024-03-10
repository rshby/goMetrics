package middleware

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		request := ctx.Request().Body()
		requestBody := map[string]any{}
		json.Unmarshal(request, &requestBody)

		ctx.Next()

		response := ctx.Response().Body()
		responseBody := map[string]any{}
		json.Unmarshal(response, &responseBody)

		log.WithFields(logrus.Fields{
			"url":      string(ctx.Request().URI().FullURI()),
			"request":  requestBody,
			"response": responseBody,
		}).Info("incmoing request")

		return nil
	}
}
