package serve

import (
	"log/slog"
	"net/http"
	"strconv"

	"git.a71.su/Andrew71/pye/auth"
	"git.a71.su/Andrew71/pye/config"
	"git.a71.su/Andrew71/pye/storage"
	"git.a71.su/Andrew71/pye/storage/sqlite"
)

var data storage.Storage

func Serve() {
	data = sqlite.MustLoadSQLite(config.Cfg.SQLiteFile)
	
	router := http.NewServeMux()

	router.HandleFunc("GET /pem", auth.PublicKey)

	router.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) { auth.Register(w, r, data) })
	router.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) { auth.Login(w, r, data) })

	// Note: likely temporary, possibly to be replaced by a fake "frontend"
	router.HandleFunc("GET /register", func(w http.ResponseWriter, r *http.Request) { auth.Register(w, r, data) })
	router.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) { auth.Login(w, r, data) })

	slog.Info("🪐 pye started", "port", config.Cfg.Port)
	slog.Debug("debug mode active")
	http.ListenAndServe(":"+strconv.Itoa(config.Cfg.Port), router)
}
