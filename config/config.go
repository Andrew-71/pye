package config

type Config struct {
	Port       int    `json:"port"`
	KeyFile    string `json:"key-file"`
	SQLiteFile string `json:"sqlite-file"`
}

var DefaultConfig = Config{
	Port: 7102,
	KeyFile: "private.key",
	SQLiteFile: "data.db",
}

var Cfg = MustLoadConfig()

// TODO: Implement
func MustLoadConfig() Config {
	return DefaultConfig
}