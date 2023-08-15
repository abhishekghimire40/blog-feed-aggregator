package main

import (
	"net/http"
	"strconv"

	"github.com/abhishekghimire40/blog-feed-aggregator/internal/database"
)

func (cfg *apiConfig) GetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	var limit int32
	limitParam := r.URL.Query().Get("limit")
	if len(limitParam) == 0 {
		limit = 15
	} else {
		val, err := strconv.Atoi(limitParam)
		if err != nil {
			responsdWithError(w, 400, "Couldn't get posts with provided limit")
			return
		}
		limit = int32(val)
	}
	requestData := database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  limit,
	}
	posts, err := cfg.DB.GetPostsByUser(r.Context(), requestData)
	if err != nil {
		responsdWithError(w, 400, "Couldn't get posts")
		return
	}
	respondWithJSON(w, 200, databasePostToPosts(posts))
}
