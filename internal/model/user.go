package model

type User struct {
	Firstname string
	Lastname string
	Email string
	PasswordHash []byte
	PasswordSalt []byte
}
