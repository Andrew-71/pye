package storage

// Storage is an interface for arbitrary storage
type Storage interface {
	Add(email, password string) error  // Add inserts a user into data
	ById(uuid string) (User, bool)     // ById retrieves a user by their UUID
	ByEmail(email string) (User, bool) // ByEmail retrieves a user by their email
	Taken(email string) bool           // Taken checks whether an email is taken
}

// Data stores active information for the app.
// It should be populated on startup
var Data Storage
