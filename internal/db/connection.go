package db 

import (
	"app/internal/env"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func CanConnect() bool {
	conn, err := pgx.Connect(context.Background(), connectionString()) 
	if err != nil {
		slog.Error("Could not establish a connection to the database", "Error", err)
		return false 
	}

	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		slog.Error("Could not establish a connection to the database", "Error", err)
		return false
	}
	return true
}

func connect() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), connectionString()) 
}

func connectionString() string {
	host := env.Get("HOST_NAME")
	db := env.Get("DB_NAME")
	port := env.Get("DB_PORT")
	user := env.Get("DB_USER")
	password := env.Get("DB_USER_PW")

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s", 
		user, password, host, port, db)
}
