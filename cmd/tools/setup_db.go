package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/ctrixcode/go-chi-postgres/internal/config"
	"github.com/ctrixcode/go-chi-postgres/pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	logger.Init()

	_ = config.LoadConfig()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSSLMode := os.Getenv("DB_SSLMODE")
	targetDBName := os.Getenv("DB_NAME")

	if targetDBName == "" {
		slog.Error("DB_NAME is not set")
		os.Exit(1)
	}

	// Connect to default 'postgres' database
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbSSLMode)

	db, err := sqlx.Connect("pgx", connString)
	if err != nil {
		slog.Error("Failed to connect to postgres database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Check if database exists
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)"
	err = db.Get(&exists, query, targetDBName)
	if err != nil {
		slog.Error("Failed to check if database exists", "error", err)
		os.Exit(1)
	}

	if exists {
		slog.Info("Database already exists", "db_name", targetDBName)
		return
	}

	// Create database
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", targetDBName))
	if err != nil {
		slog.Error("Failed to create database", "db_name", targetDBName, "error", err)
		os.Exit(1)
	}

	slog.Info("Database created successfully", "db_name", targetDBName)
}
