package v1

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/auth"
	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	"github.com/KennyMwendwaX/rss-scrapper/internal/handlers"
	"github.com/go-chi/chi"
)

// UserRoutes sets up routes related to users.
func UserRoutes(cfg *config.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", handlers.CreateUser(cfg))
	router.Get("/", auth.AuthMiddleware(cfg, handlers.GetUser))
	return router
}