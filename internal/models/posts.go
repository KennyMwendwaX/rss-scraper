package models

import (
	"github.com/KennyMwendwaX/rss-scraper/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type Post struct {
	ID          pgtype.UUID      `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	PublishedAt pgtype.Timestamp `json:"published_at"`
	Url         string           `json:"url"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
	FeedID      pgtype.UUID      `json:"feed_id"`
}

func FromDatabasePost(databasePost database.Post) Post {
	return Post{
		ID:          databasePost.ID,
		Title:       databasePost.Title,
		Description: databasePost.Description,
		PublishedAt: databasePost.PublishedAt,
		Url:         databasePost.Url,
		CreatedAt:   databasePost.CreatedAt,
		UpdatedAt:   databasePost.UpdatedAt,
		FeedID:      databasePost.FeedID,
	}
}

func FromDatabasePosts(databasePost []database.Post) []Post {
	posts := []Post{}

	for _, databasePost := range databasePost {
		posts = append(posts, FromDatabasePost(databasePost))
	}
	return posts
}
