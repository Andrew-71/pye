package main

import (
	"database/sql"
	"errors"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `
  CREATE TABLE "users" (
	"uuid"	TEXT NOT NULL UNIQUE,
	"email"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	PRIMARY KEY("uuid")
  );`

var (
	DataFile         = "data.db"
	db       *sql.DB = LoadDb()
)

func LoadDb() *sql.DB {
	// I *think* we need some file, even if only empty
	if _, err := os.Stat(DataFile); errors.Is(err, os.ErrNotExist) {
		slog.Error("sqlite3 database file required", "file", DataFile)
		os.Exit(1)
	}
	db, err := sql.Open("sqlite3", DataFile)
	if err != nil {
		slog.Error("error opening database", "error", err)
		os.Exit(1)
	}
	if _, err := db.Exec(create); err != nil && err.Error() != "table \"users\" already exists" {
		slog.Info("error initialising database table", "error", err)
		os.Exit(1)
	}
	slog.Info("loaded database")
	return db
}

func addUser(user User) error {
	_, err := db.Exec("insert into users (uuid, email, password) values ($1, $2, $3)",
		user.Uuid.String(), user.Email, user.Hash)
	if err != nil {
		slog.Error("error adding user", "error", err, "user", user)
		return err
	}
	return nil
}

func byId(uuid string) (User, bool) {
	row := db.QueryRow("select * from users where uuid = $1", uuid)
	user := User{}
	err := row.Scan(&user.Uuid, &user.Email, &user.Hash)

	return user, err == nil
}

func byEmail(email string) (User, bool) {
	row := db.QueryRow("select * from users where email = $1", email)
	user := User{}
	err := row.Scan(&user.Uuid, &user.Email, &user.Hash)

	return user, err == nil
}

func emailExists(email string) bool {
	_, ok := byEmail(email)
	return ok
}
