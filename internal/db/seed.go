package db

import (
	m "app/internal/model"
	"app/internal/pkg/cypher"
	"log/slog"
)

var c = cypher.New(1, 64*1024, 32, 16, 4)

func GenerateData() {
	ok, err := AnyExists()
	if err != nil {
		slog.Error("Could not establish a connection to the database", "Error", err)
		return
	} 
	if ok {
		slog.Info("Database is already populated, skipping data generation")
		return 
	}

	users := []*m.User{
		newUser("Mike", "Judge", "mj42069@mail.com", "Pass123!"),
		newUser("Joe", "Dirt", "jd42069@mail.com", "Ass123!"),
	}

	count := 0
	for _, user := range users {
		if user == nil {
			continue;
		}
		err := Save(*user)
		if err != nil {
			slog.Error("Could not persist user in database", "Error", err)
			continue
		}
		count++
	}
	slog.Info("Generated user data", "Rows", count)
}

func newUser(firstname, lastname, email, password string) *m.User {
	pass := append([]byte{}, password[:]...)

	if passHash, err := c.Encrypt(pass, []byte{}); err == nil {
		return &m.User{
			Firstname: firstname, 
			Lastname: lastname, 
			Email: email,
			PasswordHash: passHash.Hash, 
			PasswordSalt: passHash.Salt, 
		}
	}

	return nil 
}
