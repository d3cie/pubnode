package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseDSN string
	Dev         bool
	Port        string
	AssetPath   string
	AppVersion  string
}

var cfg Config

func Init() {
	godotenv.Load(".env")

	cfg = Config{
		DatabaseDSN: getEnv("DB_DSN", "./data/local.sqlite"),
		Dev:         getEnv("DEV", "true") == "true",
		Port:        getEnv("PORT", "3000"),
		AssetPath:   getEnv("ASSET_PATH", "/pub"),
		AppVersion:  "0.0.1",
	}
}

func Get() *Config {
	return &cfg
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
