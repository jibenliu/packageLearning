-- name: GetAuthor :one
SELECT *
FROM authors
WHERE id = ?
LIMIT 1;

-- name: GetAuthorArticles :many
SELECT *
FROM author_articles
WHERE author_id = ?
ORDER BY created_at DESC;

-- name: ListAuthors :many
SELECT *
FROM authors
ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (name, nick_name)
VALUES (?, ?);

-- name: DeleteAuthor :exec
DELETE
FROM authors
WHERE id = ?;