-- +goose Up
ALTER TABLE FeedFollows ADD CONSTRAINT uq_feed_user UNIQUE(feed_id, user_id);

-- +goose Down
ALTER TABLE FeedFollows DROP CONSTRAINT uq_feed_user;