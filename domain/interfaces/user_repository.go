package interfaces

import m "github.com/matheusrbarbosa/go-horse-challenge/domain/models"

type UserRepository interface {
	GetAll() ([]m.User, error)
	Search(term string) ([]m.User, error)
	GetById(id string) (m.User, error)
	Create(m.User) m.User
	Count() int64
}
