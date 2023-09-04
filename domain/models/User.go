package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Nickname  string
	Name      string
	Birthdate time.Time
	Stack     string
	Search    string
}
