package types

import "time"

type Press struct {
	PressID                 int32     `json:"press_id"`
	PublisherName           string    `json:"publisher_name"`
	PublisherProfileImgLink string    `json:"publisher_profile_img_link"`
	PressThumbnailUrl       string    `json:"press_thumbnail_url"`
	Description             string    `json:"description"`
	Title                   string    `json:"title"`
	ExternalUrl             string    `json:"external_url"`
	Category                []string  `json:"category"`
	PressPublishedAt        time.Time `json:"press_published_at"`
}
