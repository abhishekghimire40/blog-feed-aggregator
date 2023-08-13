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

type FeedResponseData struct {
	Feed       database.Feed       `json:"feed"`
	FeedFollow database.Feedfollow `json:"feed_follow"`
}

func (cfg *apiConfig) createFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	decoder := json.NewDecoder(r.Body)
	requestData := FeedRequestData{}
	err := decoder.Decode(&requestData)
	if err != nil {
		responsdWithError(w, 404, "Invalid request")
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
	feedFollowData := database.CreateFeedsFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}
	feedFollow, err := cfg.DB.CreateFeedsFollow(r.Context(), feedFollowData)
	if err != nil {
		responsdWithError(w, 404, err.Error())
		return
	}

	respondWithJSON(w, 201, FeedResponseData{
		Feed:       feed,
		FeedFollow: feedFollow,
	})
}

func getAllFeeds(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		feeds, err := db.GetFeeds(r.Context())
		if err != nil {
			responsdWithError(w, 404, "Invalid request")
		}
		respondWithJSON(w, 200, feeds)
	}
}
