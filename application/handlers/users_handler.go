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

func (h *userHandler) Search(term string) ([]m.User, error) {
	users, err := h.repository.Search(term)

	return users, err
}

func (h *userHandler) GetById(id string) (m.User, error) {
	user, err := h.repository.GetById(id)
	return user, err
}

func (h *userHandler) Create(u m.User) (m.User, error) {
	u = h.repository.Create(u)
	return u, nil
}

func (h *userHandler) Count() int64 {
	total := h.repository.Count()
	return total
}
