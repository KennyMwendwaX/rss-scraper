package v1

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/auth"
	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	"github.com/KennyMwendwaX/rss-scrapper/internal/handlers"
	"github.com/go-chi/chi"
)

// FeedRoutes sets up routes related to feeds.
func FeedFollowRoutes(cfg *config.APIConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", auth.AuthMiddleware(cfg, handlers.CreateFeedFollow(cfg)))
	router.Get("/", auth.AuthMiddleware(cfg, handlers.GetFeedFollows(cfg)))
	return router
}
