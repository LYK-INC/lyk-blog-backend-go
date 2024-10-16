-- name: CreatePress :one
INSERT INTO press 
    (publisher_name, 
    publisher_profile_img_link, 
    thumbnail_s3_path, 
    description, 
    title, 
    external_url, 
    category, 
    published_at)
VALUES ($1, $2, $3, $4, $5, $6, sqlc.arg(category)::TEXT[], $7)
RETURNING id;

-- name: GetPressById :one
SELECT * 
FROM press
WHERE id = $1 AND is_deleted = FALSE;


-- name: GetPressInCategory :many
SELECT * 
FROM press
WHERE $1 = ANY(category) AND is_deleted = FALSE
ORDER BY published_at DESC
LIMIT $2 OFFSET $3;


-- name: GetPresses :many
SELECT 
    p.id AS press_id,
    p.publisher_name,
    p.publisher_profile_img_link,
    p.thumbnail_s3_path AS press_thumbnail_url,
    p.description,
    p.title,
    p.external_url,
    p.category,
    p.is_featured,
    p.is_published,
    p.published_at AS press_published_at
FROM 
    press p
WHERE
    p.is_deleted = FALSE
ORDER BY 
    p.published_at DESC
LIMIT 
    $1 OFFSET $2;

-- name: GetAllPresses :many
SELECT 
    p.id AS press_id,
    p.publisher_name,
    p.publisher_profile_img_link,
    p.thumbnail_s3_path AS press_thumbnail_url,
    p.description,
    p.title,
    p.external_url,
    p.category,
    p.is_featured,
    p.is_published,
    p.published_at AS press_published_at
FROM 
    press p
WHERE
    p.is_deleted = FALSE
ORDER BY 
    p.published_at DESC;


-- name: SoftDeletePress :exec
UPDATE press
SET is_deleted = TRUE
WHERE id = $1;

-- name: PublishPress :exec
UPDATE press
SET is_published = $2
WHERE id = $1;

-- name: FeaturePress :exec
UPDATE press
SET is_featured = CASE
    WHEN id = $1 THEN true
    ELSE false
END;