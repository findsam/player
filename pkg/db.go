package pkg

import (
	"fmt"

	"github.com/findsam/tbot/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	user, password, name string
}

func NewDB(user, password, name string) *DB {
	return &DB{	
		user: user,
		password: password,
		name: name,
	}
}

func (db *DB) Start() (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
        "localhost", db.user, db.password, db.name, 5432)

    conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = conn.AutoMigrate(&models.User{})
    if err != nil {
        panic("failed to migrate database: " + err.Error())
    }

    return conn, nil
}
