#blog server

## apis

### blog

-- blog api /blog/<url_encoded_blog_title> ex: PostgreSQL Full-Text Search Guide -> PostgreSQL%2520Full-Text%2520Search%2520Guide

```sh
curl -X 'GET' \ 'http://localhost:8000/blog/PostgreSQL%2520Full-Text%2520Search%2520Guide' \ -H 'accept: application/json'

{
  "msg": "blog data",
  "data": {
    "blog_id": 1,
    "title": "PostgreSQL Full-Text Search Guide",
    "content": "In this blog post, we will explore full-text search capabilities of PostgreSQL, including indexing and querying.",
    "blog_thumbnail_url": "s3://path-to-thumbnail/thumbnail.jpg",
    "category": [
      "database",
      "postgresql",
      "search"
    ],
    "description": "A comprehensive guide to PostgreSQL full-text search",
    "blog_created_at": "2024-09-13T10:59:28.527593Z",
    "author_name": "John Doe",
    "read_time": 15,
    "author_profile_url": "s3://path-to-thumbnail/image.jpg"
  }
}
```

### homepage

-- articles api /home/articles

```sh
curl -X 'GET' \ 'http://localhost:8000/home/articles?limit=1&skip=0' \ -H 'accept: application/json'

{
  "msg": "blogs data",
  "data": [
    {
      "blog_id": 1,
      "title": "PostgreSQL Full-Text Search Guide",
      "blog_thumbnail_url": "s3://path-to-thumbnail/thumbnail.jpg",
      "category": [
        "database",
        "postgresql",
        "search"
      ],
      "description": "A comprehensive guide to PostgreSQL full-text search",
      "blog_created_at": "2024-09-13T10:59:28.527593Z",
      "author_name": "John Doe",
      "read_time": 15,
      "author_profile_url": "s3://path-to-thumbnail/image.jpg"
    }
  ]
}
```

-- featured api /home/featured

```sh
curl -X 'GET' \
  'http://localhost:8000/home/featured' \
  -H 'accept: application/json'

{
  "msg": "blogs data",
  "data": {
    "blog_id": 1,
    "title": "PostgreSQL Full-Text Search Guide",
    "blog_thumbnail_url": "s3://path-to-thumbnail/thumbnail.jpg",
    "category": [
      "database",
      "postgresql",
      "search"
    ],
    "description": "A comprehensive guide to PostgreSQL full-text search",
    "blog_created_at": "2024-09-13T10:59:28.527593Z",
    "author_name": "John Doe",
    "read_time": 15,
    "author_profile_url": "s3://path-to-thumbnail/image.jpg"
  }
}
```

### press

-- press articles api /home/press

```sh
curl -X 'GET' \
  'http://localhost:8000/home/press?limit=1&skip=0' \
  -H 'accept: application/json'

{
  "msg": "blogs data",
  "data": [
    {
      "press_id": 1,
      "publisher_name": "TechCrunch",
      "publisher_profile_img_link": "https://example.com/profile-img.jpg",
      "press_thumbnail_url": "s3://path-to-thumbnail/thumbnail.jpg",
      "description": "A comprehensive article on the latest tech trends.",
      "title": "Latest Tech Trends 2024",
      "external_url": "https://example.com/article-link",
      "category": [
        "technology",
        "trends",
        "2024"
      ],
      "press_published_at": "2024-09-15T10:00:00Z"
    }
  ]
}
```
