package routers

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/auth"
	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	"github.com/KennyMwendwaX/rss-scrapper/internal/handlers"
	"github.com/go-chi/chi"
)

func AppRouter() *chi.Mux {
	apiCfg := config.ApiCfg()

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", handlers.Readiness)
	v1Router.Get("/error", handlers.Error)
	v1Router.Post("/users", handlers.CreateUser(apiCfg))
	v1Router.Get("/users", auth.AuthMiddleware(apiCfg, handlers.GetUser))
	v1Router.Post("/feeds", auth.AuthMiddleware(apiCfg, handlers.CreateFeed(apiCfg)))

	return v1Router
}
