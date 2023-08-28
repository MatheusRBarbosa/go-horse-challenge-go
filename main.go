package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/matheusrbarbosa/go-horse-challenge/api"
)

func main() {
	server := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api.ApiRouter(server)

	server.Listen(":8080")
}
