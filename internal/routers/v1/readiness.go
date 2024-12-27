package v1

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/handlers"
	"github.com/go-chi/chi"
)

// ReadinessRoutes sets up routes related to service readiness.
func ReadinessRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", handlers.Readiness)
	return router
}
