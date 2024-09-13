package homepage

import (
	"net/http"
	"strconv"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
)

type ArticlesResponse struct {
	Msg  string `json:"msg"`
	Data []Blog `json:"data"`
}

func responseFmt(d []db.GetBlogsRow) []Blog {
	// Create a slice of Blog to hold the formatted response
	formattedResponse := make([]Blog, len(d))

	// Loop through the db.GetBlogsRow slice and map it to Blog
	for i, b := range d {
		formattedResponse[i] = Blog{
			BlogID:           b.BlogID,
			Title:            b.Title,
			BlogThumbnailUrl: b.BlogThumbnailUrl,
			Category:         b.Category,
			Description:      b.Description,
			BlogCreatedAt:    b.BlogCreatedAt.Time,
			AuthorName:       b.AuthorName,
			AuthorProfileUrl: b.AuthorProfileUrl,
			ReadTime:         b.ReadTime,
		}
	}

	return formattedResponse
}

// @tags			Homepage
// @summary			Get homepage articles
// @description		returns array of articles
// @Param   		limit query	int	false	"int valid"	minimum(1)	maximum(100)
// @Param   		skip query	int	false	"int valid"	minimum(0)	maximum(100)
// @accept			json
// @produce			json
// @success			200	{object}	ArticlesResponse
// @failure			500	{object}	types.ErrResponse
// @router			/home/articles [get]
func (s *HomePageService) Articles(c echo.Context) error {
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64) // base 10, bit size 64
	skip, _ := strconv.ParseInt(c.QueryParam("skip"), 10, 64)   // base 10, bit size 64

	r, err := s.Queries.GetBlogs(c.Request().Context(), db.GetBlogsParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrResponse{Msg: "bad paramaters", Err: err.Error()})
	}

	return c.JSON(http.StatusOK, ArticlesResponse{Msg: "blogs data", Data: responseFmt(r)})
}
