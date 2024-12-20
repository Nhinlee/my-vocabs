-- name: ListVocabs :many
SELECT * FROM vocab ORDER BY created_at;

-- name: GetVocabByName :one
SELECT * FROM vocab WHERE word = $1;

-- name: CreateVocab :one
INSERT INTO vocab (
    vocab_id,
    word,
    image_urls
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: ReviewVocabs :many
SELECT * FROM vocab WHERE next_review <= NOW() ORDER BY next_review;

-- name: UpdateNextReviewByName :one
UPDATE vocab SET next_review = $2, reviewed_time = $3 WHERE word = $1 RETURNING *;