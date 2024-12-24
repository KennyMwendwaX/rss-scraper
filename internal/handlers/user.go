package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KennyMwendwaX/rss-scrapper/internal/auth"
	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
	"github.com/KennyMwendwaX/rss-scrapper/internal/models"
	"github.com/KennyMwendwaX/rss-scrapper/internal/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type apiConfig struct {
	DB *database.Queries
}

func (apiConfig *apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	// Convert uuid.UUID to pgtype.UUID
	id := uuid.New()
	pgID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	// Convert time.Time to pgtype.Timestamp
	now := time.Now().UTC()
	pgTimestamp := pgtype.Timestamp{
		Time:  now,
		Valid: true,
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        pgID,
		Name:      params.Name,
		CreatedAt: pgTimestamp,
		UpdatedAt: pgTimestamp,
	})
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, models.SerializeUser(user))
}

func (apiConfig *apiConfig) GetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Authorization error: %v", err))
		return
	}

	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.SerializeUser(user))
}
