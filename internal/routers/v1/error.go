package v1

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/handlers"
	"github.com/go-chi/chi"
)

// ErrorRoutes sets up routes for testing error handling.
func ErrorRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", handlers.Error)
	return router
}
