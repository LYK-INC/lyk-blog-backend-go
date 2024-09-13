package homepage

import (
	"net/http"
	"strconv"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
)

type PressResponse struct {
	Msg  string        `json:"msg"`
	Data []types.Press `json:"data"`
}

func responsePressFmt(d []db.GetPressesRow) []types.Press {
	// Create a slice of Blog to hold the formatted response
	formattedResponse := make([]types.Press, len(d))

	// Loop through the db.GetPressesRow slice and map it to Blog
	for i, b := range d {
		formattedResponse[i] = types.Press{
			PressID:                 b.PressID,
			PublisherName:           b.PublisherName,
			PublisherProfileImgLink: b.PublisherProfileImgLink,
			PressThumbnailUrl:       b.PressThumbnailUrl,
			Description:             b.Description,
			Title:                   b.Title,
			ExternalUrl:             b.ExternalUrl,
			Category:                b.Category,
			PressPublishedAt:        b.PressPublishedAt.Time,
		}
	}

	return formattedResponse
}

// @tags			Homepage
// @summary			Get types.Press articles
// @description		returns array of articles
// @Param   		limit query	int	false	"int valid"	minimum(1)	maximum(100)
// @Param   		skip query	int	false	"int valid"	minimum(0)	maximum(100)
// @accept			json
// @produce			json
// @success			200	{object}	PressResponse
// @failure			500	{object}	types.ErrResponse
// @router			/home/press [get]
func (s *HomePageService) PressArticles(c echo.Context) error {
	limit, _ := strconv.ParseInt(c.QueryParam("limit"), 10, 64) // base 10, bit size 64
	skip, _ := strconv.ParseInt(c.QueryParam("skip"), 10, 64)   // base 10, bit size 64

	r, err := s.Queries.GetPresses(c.Request().Context(), db.GetPressesParams{
		Limit:  int32(limit),
		Offset: int32(skip),
	})
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrResponse{Msg: "bad paramaters", Err: err.Error()})
	}

	return c.JSON(http.StatusOK, PressResponse{Msg: "blogs data", Data: responsePressFmt(r)})
}
