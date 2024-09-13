-- name: CreateBlog :one
INSERT INTO blogs 
    (author_id, 
    title, 
    content, 
    tsv_content, 
    thumbnail_s3_path, 
    "description",
    read_time,
    category)
VALUES ($1, $2, $3, to_tsvector($3), $4, $5, $6, sqlc.arg(category)::TEXT[])
RETURNING id;


-- name: GetBlogById :one
SELECT * 
FROM blogs
WHERE id = $1;

-- name: FeatureBlog :exec
UPDATE blogs
SET is_featured = true
WHERE id =$1;

-- name: GetBlogInCategory :many
SELECT * 
FROM blogs
WHERE $1 = ANY(category)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: GetBlogs :many
SELECT 
    b.id AS blog_id,
    b.title,
    b.content,
    b.thumbnail_s3_path AS blog_thumbnail_url,
    b.category,
    b.description,
    b.read_time,
    b.created_at AS blog_created_at,
    a.name AS author_name,
    a.thumbnail_s3_path AS author_profile_url
FROM 
    blogs b
JOIN 
    authors a ON b.author_id = a.id
ORDER BY 
    b.created_at DESC
LIMIT 
    $1 OFFSET $2;
