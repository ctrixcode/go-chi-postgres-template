package server

import (
	"net/http"
	"strings"

	"github.com/ctrixcode/go-chi-postgres/internal/database"
	"github.com/ctrixcode/go-chi-postgres/internal/handlers"
	"github.com/ctrixcode/go-chi-postgres/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	allowedOrigins := []string{"https://*", "http://*"}
	if s.config.Environment == "production" {
		allowedOrigins = strings.Split(s.config.CorsAllowedOrigins, ",")
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", handlers.HelloWorldHandler)
	r.Get("/health", s.healthHandler)

	// Example Routes
	exampleRepo := database.NewExampleRepository(s.db.GetDB())
	exampleService := services.NewExampleService(exampleRepo)
	exampleHandler := handlers.NewExampleHandler(exampleService)

	r.Mount("/examples", exampleHandler.RegisterRoutes())

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	// jsonResp, _ := json.Marshal(s.db.Health())
	// _, _ = w.Write(jsonResp)
	w.Write([]byte("health check"))
}
