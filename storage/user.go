package storage

import (
	"log/slog"

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
		slog.Error("error creating a new user", "error", err)
		return User{}, err
	}
	return User{uuid.New(), email, hash}, nil
}
