package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ctrixcode/go-chi-postgres/internal/config"
	"github.com/ctrixcode/go-chi-postgres/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() *http.Server {
	cfg := config.LoadConfig()

	db := database.New(cfg.DatabaseURL)

	NewServer := &Server{
		port: cfg.Port,
		db:   db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
