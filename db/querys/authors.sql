-- name: CreateAuthor :one
INSERT INTO authors (name, password_hash, thumbnail_s3_path,role)
VALUES ($1, $2, $3,sqlc.arg(role)::TEXT[])
RETURNING id, name, password_hash;

-- name: UpdatePassword :exec
UPDATE authors
SET password_hash = $2
WHERE id = $1;

-- name: AddRole :exec
UPDATE authors
SET role = array_append(role, $2)
WHERE id = $1
  AND NOT ($2 = ANY(role));  -- Ensure role isn't already present

-- name: RemoveRole :exec
UPDATE authors
SET role = array_remove(role, sqlc.arg(role)::TEXT)
WHERE id = $1;

-- name: GetAuthors :many
SELECT * FROM authors;

-- name: GetAuthorByUsernameAndPassword :one
SELECT * FROM authors WHERE name = $1 AND password_hash = $2;