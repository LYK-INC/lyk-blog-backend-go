package blog

import (
	"net/http"
	"net/url"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
)

type BlogBySlug struct {
	Msg  string                `json:"msg"`
	Data types.BlogWithContent `json:"data"`
}

func responseFeatured(d db.GetBlogByTitleSlugRow) types.BlogWithContent {
	return types.BlogWithContent{
		BlogID:           d.BlogID,
		Title:            d.Title,
		BlogThumbnailUrl: d.BlogThumbnailUrl,
		Category:         d.Category,
		Description:      d.Description,
		BlogCreatedAt:    d.BlogCreatedAt.Time,
		AuthorName:       d.AuthorName,
		ReadTime:         d.ReadTime,
		AuthorProfileUrl: d.AuthorProfileUrl,
		Content:          d.Content,
	}
}

// @tags			Blog
// @summary			Get Blog By Slug
// @description		returns Blog Data
// @Param 			blog_slug path string true "Blog Slug"
// @accept			json
// @produce			json
// @success			200	{object}	BlogBySlug
// @failure			500	{object}	types.ErrResponse
// @router			/blog/{blog_slug} [get]
func (s *BlogPageService) BlogBySlug(c echo.Context) error {
	encloded_slug := c.Param("blog_slug")
	// Decode the URL-encoded path parameter
	decoded_slug, err := url.QueryUnescape(encloded_slug)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to decode URL parameter",
		})
	}
	r, err := s.Queries.GetBlogByTitleSlug(c.Request().Context(), decoded_slug)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrResponse{Msg: "bad paramaters", Err: err.Error()})
	}

	return c.JSON(http.StatusOK, BlogBySlug{Msg: "blog data", Data: responseFeatured(r)})
}
