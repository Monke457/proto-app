package database

import (
	"app/internal/env"
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type User struct {
	firstname string
	lastname string
	email string
	password string
}

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

func Save(user User) bool {
	conn, err := pgx.Connect(context.Background(), connectionString())
	if err != nil {
		slog.Error("Could not establish a connection to the database", "Error", err)
		return false
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(
		context.Background(), 
		"INSERT INTO users(firstname, lastname, email, password) VALUES($1, $2, $2, $3)", 
		user.firstname, user.lastname, user.email, user.password)
	
	if err != nil {
		slog.Error("Could not insert user", "Error", err)
		return false
	}
	return true
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
