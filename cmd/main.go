package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	"github.com/KennyMwendwaX/rss-scrapper/internal/routers"
	"github.com/KennyMwendwaX/rss-scrapper/internal/utils"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	cfg := config.ApiCfg()

	go utils.StartScraping(cfg, 10, time.Minute)

	router := routers.AppRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server listening on port %s", portString)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
