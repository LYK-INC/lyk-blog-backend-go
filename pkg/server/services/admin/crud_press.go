package admin

import (
	"net/http"
	// "net/url"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	// "github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
)

func (s *AdminPageService) CreatePressPage(c echo.Context) error {
	var req db.CreatePressParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	blog, err := s.Queries.CreatePress(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create blog"})
	}

	return c.JSON(http.StatusOK, blog)
}

func (s *AdminPageService) GetAllPresses(c echo.Context) error {
	presses, err := s.Queries.GetPresses(c.Request().Context(), db.GetPressesParams{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve presses"})
	}

	return c.JSON(http.StatusOK, presses)
}

func (s *AdminPageService) CreateNewAuthor(c echo.Context) error {
	var req db.CreateAuthorParams
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	blog, err := s.Queries.CreateAuthor(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create author"})
	}

	return c.JSON(http.StatusOK, blog)
}
