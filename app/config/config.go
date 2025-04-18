package config

import (
	"os"
)

type Config struct {
	App      App
	Database Database
}

type App struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Path     string
	SSLMode  string
}

func LoadConfig() *Config {
	return &Config{
		App: App{
			Port: getEnv("APP_PORT", "5000"),
		},
		Database: Database{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "user"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "postgres"),
			Path:     getEnv("DB_PATH", "public"),
			SSLMode:  getEnv("SSL_MODE", "disable"),
		},
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
