package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/abhishekghimire40/blog-feed-aggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type FeedFollowRequestData struct {
	Feed_id string `json:"feed_id"`
}

func (cfg *apiConfig) createFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	// get data from request body
	decoder := json.NewDecoder(r.Body)
	requestData := FeedFollowRequestData{}
	err := decoder.Decode(&requestData)
	if err != nil {
		responsdWithError(w, 404, "Invalid request!")
		return
	}
	// parse feed_id from request body into uuid
	feed_id, err := uuid.Parse(requestData.Feed_id)
	if err != nil {
		responsdWithError(w, 404, "Invalid request!")
		return
	}
	// check if a feed is available with provided feed_id
	feed, err := cfg.DB.GetSingleFeed(r.Context(), feed_id)
	if err != nil {
		respondWithJSON(w, 400, "No feed with the given feed id!")
		return
	}
	// create feed follow
	feedFollowData := database.CreateFeedsFollowParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	feedFollow, err := cfg.DB.CreateFeedsFollow(r.Context(), feedFollowData)
	if err != nil {
		responsdWithError(w, 404, err.Error())
		return
	}
	respondWithJSON(w, 201, feedFollow)
}

func (cfg *apiConfig) deleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID, err := uuid.Parse(chi.URLParam(r, "feedFollowID"))
	if err != nil {
		responsdWithError(w, 404, "Invalid request!")
		return
	}
	deleteFeedParams := database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	}
	err = cfg.DB.DeleteFeedFollow(r.Context(), deleteFeedParams)
	if err != nil {
		responsdWithError(w, 400, "Couldn't unfollow the feed please try again")
		return
	}
	respondWithJSON(w, 204, nil)
}

func (cfg *apiConfig) getFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responsdWithError(w, 404, "Couldn't get any feeds for the user")
		return
	}
	respondWithJSON(w, 200, feedFollows)

}
