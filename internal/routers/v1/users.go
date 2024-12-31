package v1

import (
	"github.com/KennyMwendwaX/rss-scraper/internal/auth"
	"github.com/KennyMwendwaX/rss-scraper/internal/config"
	"github.com/KennyMwendwaX/rss-scraper/internal/handlers"
	"github.com/go-chi/chi"
)

// UserRoutes sets up routes related to users.
func UserRoutes(cfg *config.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", handlers.CreateUser(cfg))
	router.Get("/", auth.AuthMiddleware(cfg, handlers.GetUser))
	router.Get("/posts", auth.AuthMiddleware(cfg, handlers.GetUserPosts(cfg)))
	return router
}
