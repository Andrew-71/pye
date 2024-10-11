package main

import (
	"encoding/json"
	"log"
	"os"
)

// So SQLite seems to hate my Mac.
// And I'd rather deal with something easily tinker-able in PoC stage
// So.................
// JSON.
//
// TODO: Kill this, preferably with fire.

func ReadUsers() []User {
	data, err := os.ReadFile("./data.json")
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatal(err)
	}
	return users
}

func AddUser(user User) {
	users := ReadUsers()
	users = append(users, user)
	data, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("./data.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func EmailExists(email string) bool {
	users := ReadUsers()
	for i := 0; i < len(users); i++ {
		if users[i].email == email {
			return true
		}
	}
	return false
}

func UserByEmail(email string) (User, bool) {
	users := ReadUsers()
	for i := 0; i < len(users); i++ {
		if users[i].email == email {
			return users[i], true
		}
	}
	return User{}, false
}