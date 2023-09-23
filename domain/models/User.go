package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Nickname  string    `json:"nickname"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
	Stack     string    `json:"stack"`
	Search    string    `json:"search"`
}
