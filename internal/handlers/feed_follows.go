package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KennyMwendwaX/rss-scrapper/internal/config"
	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
	"github.com/KennyMwendwaX/rss-scrapper/internal/models"
	"github.com/KennyMwendwaX/rss-scrapper/internal/utils"
	"github.com/go-chi/chi"
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

func GetFeedFollows(cfg *config.APIConfig) func(http.ResponseWriter, *http.Request, database.User) {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {

		feedFollows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error could not get feed follows")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, models.FromDatabaseFeedFollows(feedFollows))
	}
}

func DeleteFeedFollow(cfg *config.APIConfig) func(http.ResponseWriter, *http.Request, database.User) {
	return func(w http.ResponseWriter, r *http.Request, user database.User) {
		feedFollowIDStr := chi.URLParam(r, "feedFollowID")
		feedFollowID, err := uuid.Parse(feedFollowIDStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Error could not parse feed follow id")
			return
		}

		err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
			ID: pgtype.UUID{
				Bytes: feedFollowID,
				Valid: true,
			},
			UserID: user.ID,
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Could not delete feed follow")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, struct{}{})
	}
}
