package users

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(r fiber.Router) {
	// r.Get("/users", handleGet) TODO: Buscar por ID
	// r.Get("/users", handleSearch) TODO: Buscar por termo
	// r.Get("/count/users", handleCountUsers) TODO: Contar users inseridos no banco
	r.Post("/users", handleCreate)
}
