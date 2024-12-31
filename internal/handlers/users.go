package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KennyMwendwaX/rss-scraper/internal/config"
	"github.com/KennyMwendwaX/rss-scraper/internal/database"
	"github.com/KennyMwendwaX/rss-scraper/internal/models"
	"github.com/KennyMwendwaX/rss-scraper/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func CreateUser(cfg *config.APIConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			Name string `json:"name"`
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

		user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
			ID:        pgID,
			Name:      params.Name,
			CreatedAt: pgTimestamp,
			UpdatedAt: pgTimestamp,
		})
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
			return
		}

		utils.RespondWithJSON(w, http.StatusCreated, models.FromDatabaseUser(user))
	}
}

func GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	utils.RespondWithJSON(w, http.StatusOK, models.FromDatabaseUser(user))
}
