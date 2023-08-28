package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusrbarbosa/go-horse-challenge/api/users"
	"github.com/matheusrbarbosa/go-horse-challenge/api/utils"
)

func ApiRouter(server *fiber.App) {
	router := server.Group("/api")

	utils.RegisterUtilsRouter(router)
	users.RegisterUserRoutes(router)
}
