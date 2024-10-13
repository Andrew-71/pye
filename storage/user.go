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

// Fits checks whether the password is correct
func (u User) Fits(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Hash, []byte(password))
	return err == nil
}

// New Creates a new User
func New(email, password string) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		slog.Error("error creating a new user", "error", err)
		return User{}, err
	}
	return User{uuid.New(), email, hash}, nil
}
