package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abhishekghimire40/blog-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

type FeedRequestData struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (cfg *apiConfig) createFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	requestData := FeedRequestData{}
	err := decoder.Decode(&requestData)
	if err != nil {
		responsdWithError(w, 404, "Invalid request")
		return
	}
	if err != nil {
		responsdWithError(w, 500, "internal server error")
		return
	}
	feedData := database.CreateFeedsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      requestData.Name,
		Url:       requestData.URL,
		UserID:    user.ID,
	}
	feed, err := cfg.DB.CreateFeeds(r.Context(), feedData)
	if err != nil {
		responsdWithError(w, 404, err.Error())
		return
	}
	respondWithJSON(w, 201, feed)
}
