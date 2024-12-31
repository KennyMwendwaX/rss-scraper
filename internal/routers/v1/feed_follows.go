package v1

import (
	"github.com/KennyMwendwaX/rss-scraper/internal/auth"
	"github.com/KennyMwendwaX/rss-scraper/internal/config"
	"github.com/KennyMwendwaX/rss-scraper/internal/handlers"
	"github.com/go-chi/chi"
)

// FeedRoutes sets up routes related to feeds.
func FeedFollowRoutes(cfg *config.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", auth.AuthMiddleware(cfg, handlers.CreateFeedFollow(cfg)))
	router.Get("/", auth.AuthMiddleware(cfg, handlers.GetFeedFollows(cfg)))
	router.Delete("/{feedFollowID}", auth.AuthMiddleware(cfg, handlers.DeleteFeedFollow(cfg)))
	return router
}
