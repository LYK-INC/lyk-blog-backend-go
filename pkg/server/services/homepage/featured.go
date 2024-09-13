package homepage

import (
	"net/http"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
)

type FeaturedResponse struct {
	Msg  string `json:"msg"`
	Data Blog   `json:"data"`
}

func responseFeatured(d db.GetFeaturedBlogRow) Blog {
	return Blog{
		BlogID:           d.BlogID,
		Title:            d.Title,
		BlogThumbnailUrl: d.BlogThumbnailUrl,
		Category:         d.Category,
		Description:      d.Description,
		BlogCreatedAt:    d.BlogCreatedAt.Time,
		AuthorName:       d.AuthorName,
		ReadTime:         d.ReadTime,
		AuthorProfileUrl: d.AuthorProfileUrl,
	}
}

// @tags			Homepage
// @summary			Get featured article
// @description		returns featured article
// @accept			json
// @produce			json
// @success			200	{object}	FeaturedResponse
// @failure			500	{object}	types.ErrResponse
// @router			/home/featured [get]
func (s *HomePageService) FeaturedBlog(c echo.Context) error {
	r, err := s.Queries.GetFeaturedBlog(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrResponse{Msg: "bad paramaters", Err: err.Error()})
	}

	return c.JSON(http.StatusOK, FeaturedResponse{Msg: "blogs data", Data: responseFeatured(r)})
}
