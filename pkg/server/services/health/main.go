package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HealthService struct{}

func NewHealthService() *HealthService {
	return &HealthService{}
}

type HelloResponse struct {
	Msg  string `json:"msg"`
	Data int64  `json:"data"`
}

// @tags			Health
// @summary			Get Health check status
// @description		returns server time
// @accept			json
// @produce			json
// @success			200	{object}	HelloResponse
// @failure			500	{object}	types.ErrResponse
// @router			/ [get]
func (s *HealthService) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloResponse{Msg: "api ok", Data: time.Now().UnixNano()})
}
