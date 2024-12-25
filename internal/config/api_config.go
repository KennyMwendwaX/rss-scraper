package config

import (
	"log"

	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
)

type APIConfig struct {
	DB *database.Queries
}

func ApiCfg() *APIConfig {
	dBconn, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	return &APIConfig{
		DB: database.New(dBconn),
	}
}
