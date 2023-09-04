package users

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(r fiber.Router) {
	r.Get("/users", handleGet)
	r.Post("/users", handleCreate)
}
