package handlers

import (
	i "github.com/matheusrbarbosa/go-horse-challenge/domain/interfaces"
	m "github.com/matheusrbarbosa/go-horse-challenge/domain/models"
	"github.com/matheusrbarbosa/go-horse-challenge/infra/db/repositories"
)

type userHandler struct {
	repository i.UserRepository
}

func UserHandler() *userHandler {
	return &userHandler{
		repository: repositories.UserRepository(),
	}
}

func (h *userHandler) GetAll() ([]m.User, error) {
	users, err := h.repository.GetAll()

	if err != nil {
		return []m.User{}, err
	}

	return users, nil
}
