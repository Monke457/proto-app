package database

import (
	"app/internal/env"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Open() error {
	var conn, err = newConnection()
	if err != nil {
		return err
	}

	db, err := pgx.Connect(context.Background(), *conn) 
	if err != nil {
		return err
	}
	defer db.Close(context.Background())

	err = db.Ping(context.Background())
	return err 
}

func newConnection() (*string, error) {
	host, err := env.Get("HOST_NAME")
	if err != nil {
		return nil, err
	}
	db, err := env.Get("DB_NAME")
	if err != nil {
		return nil, err
	}
	port, err := env.Get("DB_PORT")
	if err != nil {
		return nil, err
	}
	user, err := env.Get("DB_USER")
	if err != nil {
		return nil, err
	}
	password, err := env.Get("DB_USER_PW")
	if err != nil {
		return nil, err
	}

	conn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", 
		user, password, host, port, db)

	return &conn, nil
}
