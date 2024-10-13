package app

import (
	"fmt"
	"os"

	"git.a71.su/Andrew71/pye/internal/auth"
	"git.a71.su/Andrew71/pye/internal/config"
	"git.a71.su/Andrew71/pye/internal/logging"
	"git.a71.su/Andrew71/pye/internal/storage"
	"git.a71.su/Andrew71/pye/internal/storage/sqlite"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pye",
	Short: "Pye is a simple JWT system",
	Long:  `A bare-bones authentication system with RS256`,
}

var (
	cfgFile   string
	cfgDb     string
	debugMode *bool
)

func initConfig() {
	logging.Load(*debugMode)
	config.MustLoad(cfgFile)
	if cfgDb != "" {
		config.Cfg.SQLiteFile = cfgDb
	}

	auth.MustLoadKey()
	storage.Data = sqlite.MustLoad(config.Cfg.SQLiteFile)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.json", "config file")
	rootCmd.PersistentFlags().StringVar(&cfgDb, "db", "", "database to use")
	debugMode = rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
