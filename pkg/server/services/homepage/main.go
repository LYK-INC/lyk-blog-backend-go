package homepage

import (
	"net/http"
	"strconv"
	"time"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/config"
	"github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type HomePageService struct {
	Config  config.Config
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func NewHomePageService(params HomePageService) *HomePageService {
	return &HomePageService{
		Config:  params.Config,
		Logger:  params.Logger,
		Queries: params.Queries,
	}
}

type ArticlesResponse struct {
	Msg  string `json:"msg"`
	Data []Blog `json:"data"`
}
type Blog struct {
	ID              int32       `json:"id"`
	AuthorID        int32       `json:"author_id"`
	Title           string      `json:"title"`
	Content         string      `json:"content"`
	TsvContent      interface{} `json:"tsv_content"`
	ThumbnailS3Path string      `json:"thumbnail_s3_path"`
	Category        []string    `json:"category"`
	CreatedAt       time.Time   `json:"created_at"`
}

func responseFmt(d []db.Blog) []Blog {
	// Create a slice of Blog to hold the formatted response
	formattedResponse := make([]Blog, len(d))

	// Loop through the db.Blog slice and map it to Blog
	for i, blog := range d {
		formattedResponse[i] = Blog{
			ID:              blog.ID,
			AuthorID:        blog.AuthorID,
			Title:           blog.Title,
			Content:         blog.Content,
			TsvContent:      blog.TsvContent,
			ThumbnailS3Path: blog.ThumbnailS3Path,
			Category:        blog.Category,
			CreatedAt:       blog.CreatedAt.Time,
		}
	}

	return formattedResponse
}

// @tags			Health
// @summary			Get Health check status
// @description		returns server time
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
