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
SELECT 
    b.id AS blog_id,
    b.title,
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
WHERE b.id =$1;

-- name: GetBlogByTitleSlug :one
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
WHERE b.title =$1;

-- name: FeatureBlog :exec
UPDATE blogs
SET is_featured = CASE
    WHEN id = $1 THEN true
    ELSE false
END;

-- name: GetFeaturedBlog :one
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
WHERE b.is_featured = TRUE;

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
    b.thumbnail_s3_path AS blog_thumbnail_url,
    b.category,
    b.description,
    b.read_time,
    b.created_at AS blog_created_at,
    a.name AS author_name,
    b.is_featured as is_featured,
    b.is_published as is_published,
    a.thumbnail_s3_path AS author_profile_url
FROM 
    blogs b
JOIN 
    authors a ON b.author_id = a.id
WHERE 
    b.is_deleted = false
ORDER BY 
    b.created_at DESC
LIMIT 
    $1 OFFSET $2;

-- name: GetAllBlogs :many
SELECT 
    b.id AS blog_id,
    b.title,
    b.thumbnail_s3_path AS blog_thumbnail_url,
    b.category,
    b.description,
    b.read_time,
    b.is_published,
    b.created_at AS blog_created_at
FROM
    blogs b
WHERE 
    b.is_deleted = false
ORDER BY 
    b.created_at DESC
LIMIT 
    $1 OFFSET $2;

-- name: GetRealatedBlogsById :many
SELECT  
b.id,
b.author_id,
b.category,
b.title,
b.thumbnail_s3_path,
b.description,
b.created_at AS blog_created_at,
a.name AS author_name,
a.thumbnail_s3_path AS author_profile_url
FROM blogs b
JOIN authors a ON b.author_id = a.id
WHERE b.id <> $3 -- Exclude the blog with the provided id
  AND b.category && (
    SELECT b2.category
    FROM blogs b2
    WHERE b2.id =$3
  )
LIMIT $1 
OFFSET $2;


-- name: CheckBlogId :one  
SELECT EXISTS (
    SELECT 1
    FROM blogs
    WHERE id = $1
);

-- name: SoftDeleteBlog :exec
UPDATE blogs
SET is_deleted = TRUE
WHERE id = $1;

-- name: PublishBlog :exec
UPDATE blogs
SET is_published = $2
WHERE id = $1;