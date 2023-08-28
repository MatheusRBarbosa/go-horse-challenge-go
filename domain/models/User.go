package models

import "time"

type User struct {
	ID        string
	Nickname  string
	Name      string
	birthDate time.Time
	Stack     string
	Search    string
}
