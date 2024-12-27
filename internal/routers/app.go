package routers

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	v1 "github.com/KennyMwendwaX/rss-scrapper/internal/routers/v1"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// AppRouter sets up the main application router.
func AppRouter() *chi.Mux {
	cfg := config.ApiCfg()

	router := chi.NewRouter()
	// CORS configuration
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Mount the versioned API router
	router.Mount("/v1", v1.Router(cfg))

	return router
}
