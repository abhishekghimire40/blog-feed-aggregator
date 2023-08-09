package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abhishekghimire40/blog-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

type CreateUserRequestBody struct {
	Name string `json:"name"`
}

func createUser(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		requestData := CreateUserRequestBody{}
		err := decoder.Decode(&requestData)
		if err != nil {
			responsdWithError(w, 400, "Provide username to create user")
			return
		}
		userData := database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      requestData.Name,
		}
		user, err := db.CreateUser(r.Context(), userData)
		if err != nil {
			responsdWithError(w, 400, "User cannot be created")
			return
		}
		respondWithJSON(w, 201, user)

	}
}

// get user by apiKey
func (cfg *apiConfig) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, user)
}
