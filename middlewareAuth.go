package main

import (
	"errors"
	"net/http"
	"strings"
)

func getAPIKey(r *http.Request) (string, error) {
	apiKey := r.Header.Get("Authorization")
	if len(apiKey) == 0 {
		return "", errors.New("api key not provided")
	}
	apiKey = strings.TrimSpace(strings.TrimPrefix(apiKey, "ApiKey"))
	return apiKey, nil
}

func (cfg *apiConfig) middlewareAuth(nextHandler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := getAPIKey(r)
		if err != nil {
			responsdWithError(w, 401, "Unauthorized")
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responsdWithError(w, 401, "Unauthorized")
			return
		}
		nextHandler(w, r, user)
	}
}
