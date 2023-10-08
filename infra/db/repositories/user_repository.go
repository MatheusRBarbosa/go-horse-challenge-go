package repositories

import (
	"github.com/google/uuid"
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

func (r *userRepository) GetById(id string) (m.User, error) {
	user := m.User{}
	uuid := uuid.MustParse(id)
	err := r.ctx.Where(&m.User{ID: uuid}).First(&user).Error

	return user, err
}

func (r *userRepository) Create(u m.User) m.User {
	result := r.ctx.Create(&u)
	if result.Error != nil {
		panic(result.Error.Error())
	}

	return u
}

func (r *userRepository) Count() int64 {
	var total int64
	r.ctx.Model(&m.User{}).Count(&total)

	return total
}
