package blog

import (
	"net/http"
	"strconv"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
)

type RealtedBLogsResponse struct {
	Msg  string              `json:"msg"`
	Data []types.RealtedBlog `json:"data"`
}

func responseReated(d []db.GetRealatedBlogsByIdRow) []types.RealtedBlog {
	var relatedBlogs []types.RealtedBlog
	for _, v := range d {
		val := types.RealtedBlog{
			BlogID:           v.ID,
			Title:            v.Title,
			BlogThumbnailUrl: v.ThumbnailS3Path,
			Category:         v.Category,
			Description:      v.Description,
			BlogCreatedAt:    v.BlogCreatedAt.Time,
			AuthorName:       v.AuthorName,
			AuthorProfileUrl: v.AuthorProfileUrl,
		}

		relatedBlogs = append(relatedBlogs, val)
	}
	return relatedBlogs
}

// @tags			Blog
// @summary			Get Related Blog By Slug
// @description		returns Realted Blog Data
// @Param 			blog_id path string true "Blog Id"
// @Param   		limit query	int	false	"int valid"	minimum(1)	maximum(100)
// @Param   		skip query	int	false	"int valid"	minimum(0)	maximum(100)
// @accept			json
// @produce			json
// @success			200	{object}	RealtedBLogsResponse
// @failure			500	{object}	types.ErrResponse
// @router			/blog/related/{blog_id} [get]
func (s *BlogPageService) GetRelatedBlogs(c echo.Context) error {
	str_blog_id := c.Param("blog_id")
	blog_id, err := strconv.Atoi(str_blog_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid blog id",
		})
	}
	_, err = s.Queries.CheckBlogId(c.Request().Context(), int32(blog_id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Blog not found",
		})
	}
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64) // base 10, bit size 64
	skip, _ := strconv.ParseInt(c.QueryParam("skip"), 10, 64)   // base 10, bit size 64

	r, err := s.Queries.GetRealatedBlogsById(c.Request().Context(), db.GetRealatedBlogsByIdParams{
		Limit:  int32(limit),
		Offset: int32(skip),
		ID:     int32(blog_id),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrResponse{Msg: "bad paramaters", Err: err.Error()})
	}

	return c.JSON(http.StatusOK, RealtedBLogsResponse{Msg: "blog data", Data: responseReated(r)})
}
