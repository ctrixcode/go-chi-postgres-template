package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ctrixcode/go-chi-postgres/internal/config"
	"github.com/ctrixcode/go-chi-postgres/internal/database"
)

type Server struct {
	port   int
	db     database.Service
	server *http.Server
	config *config.Config
}

func NewServer(cfg *config.Config, db database.Service) *Server {
	s := &Server{
		port:   cfg.Port,
		db:     db,
		config: cfg,
	}

	// Declare Server config
	s.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return s
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	return s.db.Close()
}

func (s *Server) GetHTTPServer() *http.Server {
	return s.server
}
