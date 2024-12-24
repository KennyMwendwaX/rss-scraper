package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/KennyMwendwaX/rss-scrapper/internal/auth"
	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
	"github.com/KennyMwendwaX/rss-scrapper/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	conn, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(err.Error())
	}

	apiCfg := &config.APIConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/readiness", handlers.Readiness)
	v1Router.Get("/error", handlers.Error)
	v1Router.Post("/users", handlers.CreateUser(apiCfg))
	v1Router.Get("/users", auth.AuthMiddleware(apiCfg, handlers.GetUser))
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server listening on port %s", portString)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
