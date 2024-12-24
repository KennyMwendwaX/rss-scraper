package handlers

import (
	"net/http"

	"github.com/KennyMwendwaX/rss-scrapper/internal/utils"
)

func Error(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusBadRequest, "Something went wrong")
}
