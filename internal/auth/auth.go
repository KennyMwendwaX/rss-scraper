package auth

import (
	"net/http"
	"strings"

	"errors"
)

// GetAPIKey extracts the API key from the request headers
func GetAPIKey(headers http.Header) (string, error) {
	header := headers.Get("Authorization")
	if header == "" {
		return "", errors.New("missing authorization header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("invalid authorization header")
	}
	if headerParts[0] != "API_KEY" {
		return "", errors.New("invalid authorization type")
	}

	return headerParts[1], nil
}
