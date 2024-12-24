package handlers

import "net/http"

func Error(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Something went wrong")
}
