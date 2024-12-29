package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
	"github.com/KennyMwendwaX/rss-scrapper/internal/models"
	"github.com/KennyMwendwaX/rss-scrapper/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateFeedFollow(cfg *config.APIConfig) func(http.ResponseWriter, *http.Request, database.User) {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		type parameters struct {
			FeedID pgtype.UUID `json:"feed_id"`
		}

		decoder := json.NewDecoder(r.Body)

		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error parsing json")
			return
		}

		id := uuid.New()
		pgID := pgtype.UUID{
			Bytes: id,
			Valid: true,
		}

		now := time.Now().UTC()
		pgTimestamp := pgtype.Timestamp{
			Time:  now,
			Valid: true,
		}

		feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
			ID:        pgID,
			CreatedAt: pgTimestamp,
			UpdatedAt: pgTimestamp,
			UserID:    user.ID,
			FeedID:    params.FeedID,
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating feed follow")
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, models.FromDatabaseFeedFollow(feedFollow))
	}
}
