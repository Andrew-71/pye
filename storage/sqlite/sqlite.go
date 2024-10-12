package sqlite

import (
	"database/sql"
	"errors"
	"log/slog"
	"os"

	"git.a71.su/Andrew71/pye/storage"
	_ "github.com/mattn/go-sqlite3"
)

const create string = `
  CREATE TABLE "users" (
	"uuid"	TEXT NOT NULL UNIQUE,
	"email"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	PRIMARY KEY("uuid")
  );`

type SQLiteStorage struct {
	db *sql.DB
}

func (s SQLiteStorage) AddUser(email, password string) error {
	user, err := storage.NewUser(email, password)
	if err != nil {
		return err
	}
	_, err = s.db.Exec("insert into users (uuid, email, password) values ($1, $2, $3)",
		user.Uuid.String(), user.Email, user.Hash)
	if err != nil {
		slog.Error("error adding user to database", "error", err, "user", user)
		return err
	}
	return nil
}

func (s SQLiteStorage) ById(uuid string) (storage.User, bool) {
	row := s.db.QueryRow("select * from users where uuid = $1", uuid)
	user := storage.User{}
	err := row.Scan(&user.Uuid, &user.Email, &user.Hash)

	return user, err == nil
}

func (s SQLiteStorage) ByEmail(email string) (storage.User, bool) {
	row := s.db.QueryRow("select * from users where email = $1", email)
	user := storage.User{}
	err := row.Scan(&user.Uuid, &user.Email, &user.Hash)

	return user, err == nil
}

func (s SQLiteStorage) EmailExists(email string) bool {
	_, ok := s.ByEmail(email)
	return ok
}

func MustLoadSQLite(dataFile string) SQLiteStorage {
	// I *think* we need some file, even if only empty
	if _, err := os.Stat(dataFile); errors.Is(err, os.ErrNotExist) {
		slog.Error("sqlite3 database file required", "file", dataFile)
		os.Exit(1)
	}
	db, err := sql.Open("sqlite3", dataFile)
	if err != nil {
		slog.Error("error opening database", "error", err)
		os.Exit(1)
	}

	// TODO: Apparently "prepare" works here
	if _, err := db.Exec(create); err != nil && err.Error() != "table \"users\" already exists" {
		slog.Info("error initialising database table", "error", err)
		os.Exit(1)
	}
	slog.Info("loaded database", "file", dataFile)
	return SQLiteStorage{db}
}