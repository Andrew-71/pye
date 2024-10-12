package main

import (
	"log/slog"
	"net/http"
	"net/mail"
	"strings"
)

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func validPass(pass string) bool {
	// TODO: Obviously, we *might* want something more sophisticated here
	return len(pass) >= 8
}

func Register(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()

	if ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		if !(validEmail(email) && validPass(password) && !emailExists(email)) {
			slog.Debug("Outcome",
				"email", validEmail(email),
				"pass", validPass(password),
				"taken", !emailExists(email))
			http.Error(w, "invalid auth credentials", http.StatusBadRequest)
			return
		}
		user, err := NewUser(email, password)
		if err != nil {
			slog.Error("error creating a new user", "error", err)
			http.Error(w, "error creating a new user", http.StatusInternalServerError)
			return
		}
		err = addUser(user)
		if err != nil {
			slog.Error("error saving a new user", "error", err)
			http.Error(w, "error saving a new user", http.StatusInternalServerError)
			return
		}
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
		user, ok := byEmail(email)
		if !ok || !user.PasswordFits(password) {
			http.Error(w, "you did something wrong", http.StatusUnauthorized)
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
