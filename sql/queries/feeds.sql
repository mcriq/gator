-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE url = $1;

-- name: DeleteFeeds :exec
DELETE FROM feeds CASCADE;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsAndUsername :many
SELECT feeds.id, feeds.created_at, feeds.updated_at, feeds.name, feeds.url, feeds.user_id, users.name as user_name
FROM feeds
JOIN users ON feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE feeds.url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET 
    last_fetched_at = $1,
    updated_at = $1
WHERE id = $2;

-- name: GetNextFeedToFetch :one
SELECT *
FROM
    feeds
ORDER BY
    last_fetched_at ASC NULLS FIRST
LIMIT 1;