package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"git.a71.su/Andrew71/pye/internal/models/user"
	"git.a71.su/Andrew71/pye/internal/storage"
	sqlite "github.com/mattn/go-sqlite3"
)

const create string = `
  CREATE TABLE "users" (
	"uuid"	TEXT NOT NULL UNIQUE,
	"email"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	PRIMARY KEY("uuid")
  );`

// SQLiteStorage is a Storage implementation with SQLite3
type SQLiteStorage struct {
	db *sql.DB
}

func (s SQLiteStorage) Add(email, password string) error {
	user, err := user.New(email, password)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("insert into users (uuid, email, password) values ($1, $2, $3)",
		user.Uuid.String(), user.Email, user.Hash)
	if err != nil {
		e, ok := err.(sqlite.Error)
		if ok && errors.Is(e.ExtendedCode, sqlite.ErrConstraintUnique) {
			// Return a standard error if the user already exists
			slog.Info("can't add user because email already taken", "user", user)
			return fmt.Errorf("%w (%s)", storage.ErrExist, user.Email)
		}
		slog.Error("error adding user to database", "error", err, "user", user)
		return err
	}
	return nil
}

func (s SQLiteStorage) ById(uuid string) (user.User, bool) {
	row := s.db.QueryRow("select * from users where uuid = $1", uuid)
	user := user.User{}
	err := row.Scan(&user.Uuid, &user.Email, &user.Hash)

	return user, err == nil
}

func (s SQLiteStorage) ByEmail(email string) (user.User, bool) {
	row := s.db.QueryRow("select * from users where email = $1", email)
	user := user.User{}
	err := row.Scan(&user.Uuid, &user.Email, &user.Hash)

	return user, err == nil
}

func (s SQLiteStorage) Taken(email string) bool {
	_, ok := s.ByEmail(email)
	return ok
}

func MustLoad(dataFile string) SQLiteStorage {
	if _, err := os.Stat(dataFile); errors.Is(err, os.ErrNotExist) {
		os.Create(dataFile)
		slog.Debug("created sqlite3 database file", "file", dataFile)
	}
	db, err := sql.Open("sqlite3", dataFile)
	if err != nil {
		slog.Error("error opening sqlite3 database", "error", err)
		os.Exit(1)
	}

	statement, err := db.Prepare(create)
	if err != nil {
		if err.Error() != "table \"users\" already exists" {
			slog.Error("error initialising sqlite3 database table", "error", err)
			os.Exit(1)
		}
	} else {
		statement.Exec()
	}

	slog.Debug("loaded database", "file", dataFile)
	return SQLiteStorage{db}
}
