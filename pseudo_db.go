package main

import (
	"encoding/json"
	"log/slog"
	"os"
)

// SQLite seems to hate my Mac.
// And I'd rather deal with something easily tinker-able in PoC stage
// So.................
// JSON.
//
// TODO: Kill this, preferably with fire.

func ReadUsers() []User {
	data, err := os.ReadFile("./data.json")
	if err != nil {
		slog.Error("error reading file", "error", err)
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		slog.Error("error unmarshalling data", "error", err)
	}
	return users
}

func AddUser(user User) {
	users := ReadUsers()
	users = append(users, user)
	// slog.Info("users", "users", users)
	data, err := json.Marshal(users)
	if err != nil {
		slog.Error("error marshalling", "error", err)
		return
	}

	f, err := os.OpenFile("./data.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		slog.Error("error opening/creating file data.json")
		return
	}
	if _, err := f.Write(data); err != nil {
		slog.Error("error writing to file data.json", "error", err)
		return
	}
}

func EmailExists(email string) bool {
	users := ReadUsers()
	for i := 0; i < len(users); i++ {
		if users[i].Email == email {
			return true
		}
	}
	return false
}

func UserByEmail(email string) (User, bool) {
	users := ReadUsers()
	for i := 0; i < len(users); i++ {
		if users[i].Email == email {
			return users[i], true
		}
	}
	return User{}, false
}
