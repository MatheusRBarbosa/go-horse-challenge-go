package requests

import (
	"strings"
	"time"

	m "github.com/matheusrbarbosa/go-horse-challenge/domain/models"
)

type CreateUserRequest struct {
	Name      string   `json:"name" validate:"required,lte=255"`
	Nickname  string   `json:"nickname" validate:"required,lte=255"`
	Birthdate string   `json:"birthDate" validate:"required,datetime=2006-01-02"`
	Stack     []string `json:"stack" validate:"required"`
}

func (r *CreateUserRequest) ParseToUser() m.User {
	birthDate, err := time.Parse("2006-01-02", r.Birthdate)
	if err != nil {
		panic(err)
	}

	return m.User{
		Name:      r.Name,
		Nickname:  r.Nickname,
		Birthdate: birthDate,
		Stack:     strings.Join(r.Stack[:], ","),
	}
}
