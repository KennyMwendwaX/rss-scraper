package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KennyMwendwaX/rss-scrapper/database"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing json")
		return
	}

	// Convert uuid.UUID to pgtype.UUID
	id := uuid.New()
	pgID := pgtype.UUID{
		Bytes: id, // Assign the UUID bytes
		Valid: true,
	}

	// Convert time.Time to pgtype.Timestamp
	now := time.Now().UTC()
	pgTimestamp := pgtype.Timestamp{
		Time:  now,  // Assign the time value
		Valid: true, // Set to valid
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        pgID,
		Name:      params.Name,
		CreatedAt: pgTimestamp,
		UpdatedAt: pgTimestamp,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
