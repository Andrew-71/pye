package storage

type Storage interface {
	AddUser(email, password string) error
	ById(uuid string) (User, bool)
	ByEmail(uuid string) (User, bool)
	EmailExists(email string) bool
}
