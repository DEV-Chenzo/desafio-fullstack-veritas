package config

import "os"

type Config struct {
	DatabaseURL string
	Port        string
	CORSOrigin  string
}

func Load() Config {
	return Config{
		DatabaseURL: env("DATABASE_URL", "postgres://kanban:kanban@127.0.0.1:5433/kanban?sslmode=disable"),
		Port:        env("PORT", "8080"),
		CORSOrigin:  env("CORS_ORIGIN", "http://localhost:5173"),
	}
}

func env(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
