package cmd

import (
	"fmt"
	"os"

	"git.a71.su/Andrew71/pye/auth"
	"git.a71.su/Andrew71/pye/config"
	"git.a71.su/Andrew71/pye/logging"
	"git.a71.su/Andrew71/pye/storage"
	"git.a71.su/Andrew71/pye/storage/sqlite"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pye",
	Short: "Pye is a simple JWT system",
	Long:  `A bare-bones authentication system built by Andrew71 as an assignment`,
}

var (
	cfgFile   string
	cfgDb     string
	debugMode *bool
)

func initConfig() {
	logging.LogInit(*debugMode)
	config.MustLoadConfig(cfgFile)
	if cfgDb != "" {
		config.Cfg.SQLiteFile = cfgDb
	}

	auth.MustLoadKey()
	storage.Data = sqlite.MustLoadSQLite(config.Cfg.SQLiteFile)
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
