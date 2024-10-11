package main

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	uuid  uuid.UUID
	email string
	hash  []byte // bcrypt hash of password
}

func (u User) PasswordFits(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.hash, []byte(password))
	return err == nil
}

func NewUser(email, password string) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return User{}, err
	}
	return User{uuid.New(), email, hash}, nil
}

func CreateUser(User) {

}
