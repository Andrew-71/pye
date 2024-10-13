package storage

import (
	"errors"

	"git.a71.su/Andrew71/pye/internal/models/user"
)

var ErrExist = errors.New("user already exists")

// Storage is an interface for arbitrary storage
type Storage interface {
	Add(email, password string) error       // Add inserts a user into data
	ById(uuid string) (user.User, bool)     // ById retrieves a user by their UUID
	ByEmail(email string) (user.User, bool) // ByEmail retrieves a user by their email
	Taken(email string) bool                // Taken checks whether an email is taken
}

// Data stores active information for the app.
// It should be populated on startup
var Data Storage
