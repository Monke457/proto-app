package db 

import (
	"app/internal/model"
	"context"
)

func Save(user model.User) error {
	conn, err := connect() 
	if err != nil {
		return err 
	}
	defer conn.Close(context.Background())

	_, err = conn.Exec(
		context.Background(), 
		"INSERT INTO users(firstname, lastname, email, password_hash, password_salt) VALUES($1, $2, $3, $4, $5);", 
		user.Firstname, user.Lastname, user.Email, user.PasswordHash, user.PasswordSalt)
	
	return err 
}

func AnyExists() (bool, error) {
	conn, err := connect()
	if err != nil {
		return false, err
	}
	rows, _ := conn.Query(context.Background(), "SELECT id FROM users LIMIT 1;") 
	return rows.Next(), nil 
}

