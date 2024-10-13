package auth

import (
	"log/slog"
	"net/http"
	"net/mail"
	"strings"

	"git.a71.su/Andrew71/pye/storage"
)

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func validPass(pass string) bool {
	// TODO: Obviously, we *might* want something more sophisticated here
	return len(pass) >= 8
}

// Register creates a new user with credentials provided through Basic Auth
func Register(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()

	if ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		if !(validEmail(email) && validPass(password) && !storage.Data.Taken(email)) {
			slog.Debug("outcome",
				"email", validEmail(email),
				"pass", validPass(password),
				"taken", !storage.Data.Taken(email))
			http.Error(w, "invalid auth credentials", http.StatusBadRequest)
			return
		}
		err := storage.Data.Add(email, password)
		if err != nil {
			slog.Error("error adding a new user", "error", err)
			http.Error(w, "error adding a new user", http.StatusInternalServerError)
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

// Login returns JWT for a registered user through Basic Auth
func Login(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()

	if ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		user, ok := storage.Data.ByEmail(email)
		if !ok || !user.Fits(password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
			http.Error(w, "you did something wrong", http.StatusUnauthorized)
			return
		}

		token, err := Create(user)
		if err != nil {
			http.Error(w, "error creating jwt", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(token))
		return
	}

	// No email and password was provided
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "This API requires authorization", http.StatusUnauthorized)
}
