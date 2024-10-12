package main

import (
	"log/slog"
	"net/http"
	"net/mail"
	"strings"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func ValidPass(pass string) bool {
	// TODO: Obviously, we *might* want something more sophisticated here
	return true
	//return len(pass) >= 8
}
func EmailTaken(email string) bool {
	// TODO: Implement properly
	return EmailExists(email)
}
func Register(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()

	if ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		if !(ValidEmail(email) && ValidPass(password) && !EmailTaken(email)) {
			// TODO: Provide descriptive error and check if 400 is best code?
			slog.Info("Outcome",
				"email", ValidEmail(email),
				"pass", ValidPass(password),
				"taken", !EmailTaken(email))
			http.Error(w, "Invalid auth credentials", http.StatusBadRequest)
			return
		}
		user, err := NewUser(email, password)
		if err != nil {
			slog.Error("error creating a new user", "error", err)
		}
		slog.Info("user", "user", user)
		AddUser(user)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created"))
		return
	}

	// No email and password was provided
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "This API requires authorization", http.StatusUnauthorized)
}

func Login(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()

	if ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		user, ok := ByEmail(email)
		if !ok || !user.PasswordFits(password) {
			http.Error(w, "You did something wrong", http.StatusUnauthorized)
			return
		}

		s, err := CreateJWT(user)
		if err != nil {
			http.Error(w, "error creating jwt", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(s))
		return
	}

	// No email and password was provided
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "This API requires authorization", http.StatusUnauthorized)
}
