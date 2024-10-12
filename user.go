package main

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Uuid  uuid.UUID
	Email string
	Hash  []byte // bcrypt hash of password
}

func (u User) PasswordFits(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Hash, []byte(password))
	return err == nil
}

func NewUser(email, password string) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return User{}, err
	}
	return User{uuid.New(), email, hash}, nil
}

// TODO: Implement
func ByEmail(email string) (User, bool) {
	return UserByEmail(email)
}
