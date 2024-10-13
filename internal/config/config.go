package config

import (
	"encoding/json"
	"log/slog"
	"os"
)

// Config stores app configuration
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

// Load parses a JSON-formatted configuration file
func load(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	conf := DefaultConfig
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return Config{}, err
	}
	slog.Debug("Loaded config", "file", filename, "config", conf)
	return conf, nil
}

// MustLoad handles initial configuration loading
func MustLoad(filename string) {
	conf, err := load(filename)
	if err != nil {
		slog.Error("error initially loading config", "error", err)
		os.Exit(1)
	}
	Cfg = conf
}
