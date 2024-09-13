package homepage

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HomePageService struct{}

func NewHomePageService() *HomePageService {

	return &HomePageService{}
}

// @tags			Health
// @summary			Get Health check status
// @description		returns server time
// @Param   		int	limit	int	false	"int valid"	minimum(1)	maximum(100)
// @Param   		int	skip	int	false	"int valid"	minimum(0)	maximum(100)
// @accept			json
// @produce			json
// @success			200	{object}	int64
// @failure			500	{object}	utils.ErrorResponse
// @router			/home/articles [get]
func (s *HomePageService) Articles(c echo.Context) error {
	return c.JSON(http.StatusOK, time.Now().UnixNano())
}
