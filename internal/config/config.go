package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               int
	DatabaseURL        string
	JWTSecret          string
	Environment        string
	CorsAllowedOrigins string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Error loading .env file")
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 8080
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	return &Config{
		Port:               port,
		DatabaseURL:        databaseURL,
		JWTSecret:          os.Getenv("JWT_SECRET"),
		Environment:        os.Getenv("APP_ENV"),
		CorsAllowedOrigins: os.Getenv("CORS_ALLOWED_ORIGINS"),
	}
}
