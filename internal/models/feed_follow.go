package models

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type FeedFollow struct {
	ID        pgtype.UUID      `json:"id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	UserID    pgtype.UUID      `json:"user_id"`
	FeedID    pgtype.UUID      `json:"feed_id"`
}

func FromDatabaseFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		UserID:    feedFollow.UserID,
		FeedID:    feedFollow.FeedID,
	}
}

func FromDatabaseFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedsFollows := []FeedFollow{}

	for _, dbFeedFollows := range dbFeedFollows {
		feedsFollows = append(feedsFollows, FromDatabaseFeedFollow(dbFeedFollows))
	}
	return feedsFollows
}
