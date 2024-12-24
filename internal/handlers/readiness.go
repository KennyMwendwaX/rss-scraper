package handlers

import "net/http"

func Readiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]bool{"ok": true})
}
