package repositories

import (
	i "github.com/matheusrbarbosa/go-horse-challenge/domain/interfaces"
	m "github.com/matheusrbarbosa/go-horse-challenge/domain/models"
	"github.com/matheusrbarbosa/go-horse-challenge/infra/db"
	"gorm.io/gorm"
)

type userRepository struct {
	ctx *gorm.DB
}

func UserRepository() i.UserRepository {
	return &userRepository{
		ctx: db.Ctx(),
	}
}

func (r *userRepository) GetAll() ([]m.User, error) {
	persons := []m.User{}

	err := r.ctx.Find(&persons)
	if err != nil {
		return persons, err.Error
	}

	return persons, nil
}
