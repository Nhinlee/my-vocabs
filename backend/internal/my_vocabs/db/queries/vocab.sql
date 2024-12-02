-- name: ListVocabs :many
SELECT * FROM vocab ORDER BY created_at;

-- name: CreateVocab :one
INSERT INTO vocab (
    vocab_id,
    word,
    image_urls
) VALUES (
    $1, $2, $3
) RETURNING *;