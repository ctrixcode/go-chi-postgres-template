package server

import (
	"net/http"

	"github.com/ctrixcode/go-chi-postgres/internal/database"
	"github.com/ctrixcode/go-chi-postgres/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
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
	exampleHandler := handlers.NewExampleHandler(exampleRepo)

	r.Mount("/examples", exampleHandler.RegisterRoutes())

	return r
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	// jsonResp, _ := json.Marshal(s.db.Health())
	// _, _ = w.Write(jsonResp)
	w.Write([]byte("health check"))
}
