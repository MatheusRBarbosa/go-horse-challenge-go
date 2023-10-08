package users

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(r fiber.Router) {
	r.Get("/users/:id", handleGet)
	r.Get("/count/users", handleCountUsers)
	r.Get("/users", handleGetAll)
	r.Post("/users", handleCreate)
}
