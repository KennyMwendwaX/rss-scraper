package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KennyMwendwaX/rss-scrapper/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5XX level error:", message)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errorResponse{Error: message})
}
