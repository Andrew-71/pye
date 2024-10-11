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
	return len(pass) >= 8 // TODO: Obviously, we *might* want something more sophisticated here
}
func EmailTaken(email string) bool {
	// TODO: Implement properly
	return EmailExists(email)
}
func Register(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()

	if !ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		if !(ValidEmail(email) || ValidPass(password) || EmailTaken(email)) {
			// TODO: Provide descriptive error and check if 400 is best code?
			http.Error(w, "Invalid auth credentials", http.StatusBadRequest)
			return
		}
		user, err := NewUser(email, password)
		if err != nil {
			slog.Error("Error creating a new user", "error", err)
		}
		AddUser(user)
	}

	// No email and password was provided
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "This API requires authorization", http.StatusUnauthorized)
}

func Login(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()

	if !ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		user, ok := ByEmail(email)
		if !ok || !user.PasswordFits(password) {
			http.Error(w, "You did something wrong", http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusOK)
		
	}

	// No email and password was provided
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "This API requires authorization", http.StatusUnauthorized)
}