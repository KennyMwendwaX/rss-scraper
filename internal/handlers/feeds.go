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

func CreateFeed(cfg *config.APIConfig) func(http.ResponseWriter, *http.Request, database.User) {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		type parameters struct {
			Name string `json:"name"`
			Url  string `json:"url"`
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

		feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
			ID:        pgID,
			Name:      params.Name,
			Url:       params.Url,
			CreatedAt: pgTimestamp,
			UpdatedAt: pgTimestamp,
			UserID:    user.ID,
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, models.SerializeFeed(feed))
	}
}

func GetFeeds(cfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		feeds, err := cfg.DB.GetFeeds(r.Context())
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, models.SerializeFeeds(feeds))
	}
}
