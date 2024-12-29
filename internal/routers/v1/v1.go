package v1

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	"github.com/go-chi/chi"
)

// Router sets up the version 1 API routes.
func Router(cfg *config.APIConfig) *chi.Mux {
	router := chi.NewRouter()

	// Add route groups
	router.Mount("/readiness", ReadinessRoutes())
	router.Mount("/error", ErrorRoutes())
	router.Mount("/users", UserRoutes(cfg))
	router.Mount("/feeds", FeedRoutes(cfg))
	router.Mount("/feed-follows", FeedFollowRoutes(cfg))

	return router
}
