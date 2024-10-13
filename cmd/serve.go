package cmd

import (
	"log/slog"
	"net/http"
	"strconv"

	"git.a71.su/Andrew71/pye/auth"
	"git.a71.su/Andrew71/pye/config"
	"github.com/spf13/cobra"
)

var port int

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", config.Cfg.Port, "port to use")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start JWT service",
	Long:  `Start a simple authentication service`,
	Run:   serveAuth,
}

func serveAuth(cmd *cobra.Command, args []string) {
	router := http.NewServeMux()

	router.HandleFunc("GET /pem", auth.PublicKey)

	router.HandleFunc("POST /register", auth.Register)
	router.HandleFunc("POST /login", auth.Login)

	// Note: likely temporary, possibly to be replaced by a fake "frontend"
	router.HandleFunc("GET /register", auth.Register)
	router.HandleFunc("GET /login", auth.Login)

	slog.Info("🪐 pye started", "port", port)
	slog.Debug("debug mode active")
	http.ListenAndServe(":"+strconv.Itoa(port), router)
}
