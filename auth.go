package main

import (
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
func TakenEmail(email string) bool {
	// TODO: Implement
	return false
}
func Register(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()

	if !ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		if !(ValidEmail(email) || ValidPass(password) || TakenEmail(email)) {
			// TODO: Provide descriptive error and check if 400 is best code?
			http.Error(w, "Invalid auth credentials", http.StatusBadRequest)
		}
	}

	// No email and password was provided
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "This API requires authorization", http.StatusUnauthorized)
}
