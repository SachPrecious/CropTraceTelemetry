package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppEnv         string
	AppPort        string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	MonolithAPIKey string
}

func LoadConfig() *Config {
	return &Config{
		AppEnv:         getEnv("APP_ENV", "production"),
		AppPort:        getEnv("API_PORT", "8000"),
		DBHost:         getEnv("DB_HOST", "postgres"),
		DBPort:         getEnv("DB_PORT", "5432"),
		DBUser:         getEnv("DB_USER", "croptrace"),
		DBPassword:     getEnv("DB_PASSWORD", "changeme"),
		DBName:         getEnv("DB_NAME", "croptrace"),
		MonolithAPIKey: getEnv("MONOLITH_API_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
