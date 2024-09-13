-- name: CreateBlog :one
INSERT INTO blogs (author_id, title, content, tsv_content, thumbnail_s3_path, category)
VALUES ($1, $2, $3, to_tsvector($3), $4, sqlc.arg(category)::TEXT[])
RETURNING id;


-- name: GetBlogById :one
SELECT * 
FROM blogs
WHERE id = $1;


-- name: GetBlogInCategory :many
SELECT * 
FROM blogs
WHERE $1 = ANY(category)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: GetBlogs :many
SELECT * 
FROM blogs
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;
