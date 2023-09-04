package requests

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func getValidator() *validator.Validate {
	return validate
}

type apiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

func ValidateRequest(ctx *fiber.Ctx, req interface{}) error {
	if err := ctx.BodyParser(req); err != nil {
		return err
	}

	if err := getValidator().Struct(req); err != nil {
		ctx.Status(fiber.ErrBadRequest.Code)
		ctx.JSON(fiber.Map{"errors": parseError(err)})
		return err
	}

	return nil
}

func parseError(err error) []apiError {
	var validationErros validator.ValidationErrors
	var timeError *time.ParseError

	if errors.As(err, &validationErros) {
		out := make([]apiError, len(validationErros))
		for i, fe := range validationErros {
			out[i] = apiError{fe.Field(), parseErrorMessage(fe)}
		}
		return out
	} else if errors.As(err, &timeError) {
		out := make([]apiError, 1)
		out[0] = apiError{timeError.Value, fmt.Sprintf("Should have %s format", timeError.Layout)}

		return out
	}

	return nil
}

func parseErrorMessage(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "Este campo é obrigatório"
	case "email":
		return "E-mail inválido"
	case "gt":
		return fmt.Sprintf("Este campo deve ser maior que %s", fieldError.Param())
	case "gte":
		return fmt.Sprintf("Este campo deve ser maior ou igual a %s", fieldError.Param())
	case "lt":
		return fmt.Sprintf("Este campo deve ser menor que %s", fieldError.Param())
	case "lte":
		return fmt.Sprintf("Este campo deve ser menor ou igual a %s", fieldError.Param())
	case "exists":
		return fmt.Sprintf("%s não foi encontrado na base", fieldError.Value())
	case "unique":
		return fmt.Sprintf("Este %s já foi usado.", fieldError.Field())
	case "datetime":
		return fmt.Sprintf("Este campo deve ter o formato igual a yyyy-mm-dd")
	}

	return fieldError.Error() // default error
}