package handlers

import (
	"net/http"

	"github.com/KennyMwendwaX/rss-scraper/internal/utils"
)

func Readiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, map[string]bool{"ok": true})
}
