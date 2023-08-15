package main

import (
	"time"

	"github.com/abhishekghimire40/blog-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

type PostResponse struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(post database.Post) PostResponse {
	var description *string
	if post.Description.Valid {
		description = &post.Description.String
	}

	return PostResponse{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: description,
		PublishedAt: post.PublishedAt,
		FeedID:      post.FeedID,
	}
}

func databasePostToPosts(dbPosts []database.Post) []PostResponse {
	posts := []PostResponse{}
	for _, post := range dbPosts {
		posts = append(posts, databasePostToPost(post))
	}
	return posts
}
