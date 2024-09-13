package homepage

import "time"

type Blog struct {
	BlogID           int32     `json:"blog_id"`
	Title            string    `json:"title"`
	BlogThumbnailUrl string    `json:"blog_thumbnail_url"`
	Category         []string  `json:"category"`
	Description      string    `json:"description"`
	BlogCreatedAt    time.Time `json:"blog_created_at"`
	AuthorName       string    `json:"author_name"`
	ReadTime         int32     `json:"read_time"`
	AuthorProfileUrl string    `json:"author_profile_url"`
}
