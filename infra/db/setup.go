package db

import (
	"time"

	"github.com/matheusrbarbosa/go-horse-challenge/crosscutting"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ctx *gorm.DB

func Ctx() *gorm.DB {
	return ctx
}

func ConnectDatabase() {
	connection := crosscutting.GetConnectionString()

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if err != nil {
		panic("Failed to connect to database!")
	}

	ctx = db
}
