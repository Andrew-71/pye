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

func Register(w http.ResponseWriter, r *http.Request, data storage.Storage) {
	email, password, ok := r.BasicAuth()

	if ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		if !(validEmail(email) && validPass(password) && !data.EmailExists(email)) {
			slog.Debug("Outcome",
				"email", validEmail(email),
				"pass", validPass(password),
				"taken", !data.EmailExists(email))
			http.Error(w, "invalid auth credentials", http.StatusBadRequest)
			return
		}
		err := data.AddUser(email, password)
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

func Login(w http.ResponseWriter, r *http.Request, data storage.Storage) {
	email, password, ok := r.BasicAuth()

	if ok {
		email = strings.TrimSpace(email)
		password = strings.TrimSpace(password)
		user, ok := data.ByEmail(email)
		if !ok || !user.PasswordFits(password) {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
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