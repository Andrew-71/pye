package app

import (
	"log/slog"
	"net/http"
	"strconv"

	"git.a71.su/Andrew71/pye/internal/auth"
	"git.a71.su/Andrew71/pye/internal/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
)

var port int

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 0, "port to use")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start JWT service",
	Long:  `Start a simple authentication service`,
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	if port == 0 {
		port = config.Cfg.Port
	}

	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger, middleware.CleanPath, middleware.StripSlashes)

	r.Get("/pem", auth.ServePublicKey)
	r.Post("/register", auth.Register)
	r.Post("/login", auth.Login)

	// Note: likely temporary, possibly to be replaced by a fake "frontend"
	r.Get("/register", auth.Register)
	r.Get("/login", auth.Login)

	slog.Info("🪐 pye started", "port", port)
	http.ListenAndServe(":"+strconv.Itoa(port), r)
}
