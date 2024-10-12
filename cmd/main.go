package cmd

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"git.a71.su/Andrew71/pye/cmd/serve"
	"git.a71.su/Andrew71/pye/cmd/verify"
	"git.a71.su/Andrew71/pye/config"
)

func Run() {	

	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	serveConfig := serveCmd.String("config", "", "override config file")
	servePort := serveCmd.Int("port", 0, "override port")
	serveDb := serveCmd.String("db", "", "override sqlite database")

	verifyCmd := flag.NewFlagSet("verify", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println("expected 'serve' or 'verify' subcommands")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "serve":
		serveCmd.Parse(os.Args[2:])
		if *serveConfig != "" {
			err := config.LoadConfig(*serveConfig)
			if err != nil {
				slog.Error("error loading custom config", "error", err)
			}
		}
		if *servePort != 0 {
			config.Cfg.Port = *servePort
		}
		if *serveDb != "" {
			config.Cfg.SQLiteFile = *serveDb
		}
		serve.Serve()
	case "verify":
		verifyCmd.Parse(os.Args[2:])
		if len(os.Args) != 4 {
			fmt.Println("Usage: <jwt> <pem file>")
		}
		verify.Verify(os.Args[2], os.Args[3])
	default:
		fmt.Println("expected 'serve' or 'verify' subcommands")
		os.Exit(0)
	}
}
