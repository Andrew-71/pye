package config

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Config struct {
	Port       int    `json:"port"`
	KeyFile    string `json:"key-file"`
	SQLiteFile string `json:"sqlite-file"`
	LogToFile  bool   `json:"log-to-file"`
	LogFile    string `json:"log-file"`
}

var DefaultConfig = Config{
	Port:       7102,
	KeyFile:    "private.key",
	SQLiteFile: "data.db",
	LogToFile:  false,
	LogFile:    "pye.log",
}

var (
	DefaultLocation = "config.json"
	Cfg             Config
)

func Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	temp_config := DefaultConfig
	err = json.Unmarshal(data, &temp_config)
	if err != nil {
		return err
	}
	Cfg = temp_config
	slog.Debug("Loaded config", "file", filename, "config", Cfg)
	return nil
}

func MustLoad(filename string) {
	err := Load(filename)
	if err != nil {
		slog.Error("error initially loading config", "error", err)
		os.Exit(1)
	}
}