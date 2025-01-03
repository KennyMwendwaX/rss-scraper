package v1

import (
	"github.com/KennyMwendwaX/rss-scraper/internal/auth"
	"github.com/KennyMwendwaX/rss-scraper/internal/config"
	"github.com/KennyMwendwaX/rss-scraper/internal/handlers"
	"github.com/go-chi/chi"
)

// FeedRoutes sets up routes related to feeds.
func FeedRoutes(cfg *config.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", handlers.GetFeeds(cfg))
	router.Post("/", auth.AuthMiddleware(cfg, handlers.CreateFeed(cfg)))
	return router
}
