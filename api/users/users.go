package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matheusrbarbosa/go-horse-challenge/api/requests"
	"github.com/matheusrbarbosa/go-horse-challenge/application/handlers"
)

func handleGetAll(ctx *fiber.Ctx) error {
	handler := handlers.UserHandler()
	res, err := handler.GetAll()
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200).JSON(res)
	return nil
}

func handleGet(ctx *fiber.Ctx) error {
	handler := handlers.UserHandler()
	id := ctx.Params("id")
	res, err := handler.GetById(id)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(200).JSON(res)
	return nil
}

func handleCreate(ctx *fiber.Ctx) error {
	req := new(requests.CreateUserRequest)

	err := requests.ValidateRequest(ctx, req)
	if err != nil {
		return nil
	}

	user := req.ParseToUser()
	handler := handlers.UserHandler()
	res, err := handler.Create(user)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	ctx.Status(201).JSON(res)
	return nil
}

func handleCountUsers(ctx *fiber.Ctx) error {
	handler := handlers.UserHandler()
	total := handler.Count()
	ctx.Status(200).JSON(total)
	return nil
}
