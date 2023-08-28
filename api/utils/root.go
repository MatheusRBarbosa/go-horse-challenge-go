package utils

import "github.com/gofiber/fiber/v2"

func RegisterUtilsRouter(r fiber.Router) {
	r.Get("/health", handleHealthCheck)
}
