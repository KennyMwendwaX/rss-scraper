package config

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
)

type APIConfig struct {
	DB *database.Queries
}
