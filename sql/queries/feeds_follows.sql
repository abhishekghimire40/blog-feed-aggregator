-- name: CreateFeedsFollow :one
INSERT INTO FeedFollows(id,feed_id,user_id,created_at,updated_at)
VALUES ($1,$2,$3,$4,$5)
RETURNING *;

-- name: DeleteFeedFollow :exec
DELETE from FeedFollows WHERE id=$1 and user_id=$2;

-- name: GetFeedFollows :many
SELECT * FROM FeedFollows WHERE user_id=$1;