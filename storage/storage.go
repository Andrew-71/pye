package storage

// Storage is an arbitrary storage interface
type Storage interface {
	AddUser(email, password string) error
	ById(uuid string) (User, bool)
	ByEmail(uuid string) (User, bool)
	EmailExists(email string) bool
}

// Data stores active information for the app
// It should be populated at app startup
var Data Storage
