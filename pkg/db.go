package pkg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	uri string
}

func NewDB(user, password, name string) *DB {
	return &DB{
		uri: fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", user, password, name),
	}
}

func (db *DB) Start() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), db.uri)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (db *DB) Migrate() error {
    fmt.Println("implement migrate functionality")
    return nil
}
