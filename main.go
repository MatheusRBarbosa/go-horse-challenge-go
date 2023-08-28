package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/matheusrbarbosa/go-horse-challenge/api"
	"github.com/matheusrbarbosa/go-horse-challenge/infra/db"
)

func main() {
	server := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	db.ConnectDatabase()
	api.ApiRouter(server)

	server.Listen(":8080")
}
