package models

import (
	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID        pgtype.UUID      `json:"id"`
	Name      string           `json:"name"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func SerializeUser(databaseUser database.User) User {
	return User{
		ID:        databaseUser.ID,
		Name:      databaseUser.Name,
		CreatedAt: databaseUser.CreatedAt,
		UpdatedAt: databaseUser.UpdatedAt,
	}
}
