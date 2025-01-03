package models

import (
	"github.com/KennyMwendwaX/rss-scraper/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type Feed struct {
	ID        pgtype.UUID      `json:"id"`
	Name      string           `json:"name"`
	Url       string           `json:"url"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	UserID    pgtype.UUID      `json:"user_id"`
}

func FromDatabaseFeed(databaseFeed database.Feed) Feed {
	return Feed{
		ID:        databaseFeed.ID,
		Name:      databaseFeed.Name,
		Url:       databaseFeed.Url,
		CreatedAt: databaseFeed.CreatedAt,
		UpdatedAt: databaseFeed.UpdatedAt,
		UserID:    databaseFeed.UserID,
	}
}

func FromDatabaseFeeds(databaseFeed []database.Feed) []Feed {
	feeds := []Feed{}

	for _, databaseFeed := range databaseFeed {
		feeds = append(feeds, FromDatabaseFeed(databaseFeed))
	}
	return feeds
}
