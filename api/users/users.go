package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusrbarbosa/go-horse-challenge/application/handlers"
)

func handleGet(ctx *fiber.Ctx) error {
	handler := handlers.UserHandler()
	res, err := handler.GetAll()
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200).JSON(res)
	return nil
}
