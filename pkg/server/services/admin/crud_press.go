package admin

import (
	"fmt"
	"net/http"
	"strconv"

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
	presses, err := s.Queries.GetAllPresses(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve presses"})
	}

	fmt.Print(presses)

	return c.JSON(http.StatusOK, presses)
}

func (s *AdminPageService) UpdateFeaturePress(c echo.Context) error {
	idStr := c.Param("press_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id parameter"})
	}

	featureOk := s.Queries.FeaturePress(c.Request().Context(), int32(id))
	if featureOk != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update blog to Featured"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Blog updated to Featured"})
}

func (s *AdminPageService) DeletePress(c echo.Context) error {
	idStr := c.Param("press_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id parameter"})
	}

	deleteOk := s.Queries.SoftDeletePress(c.Request().Context(), int32(id))
	if deleteOk != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete blog"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Blog deleted"})
}

func (s *AdminPageService) PublishPress(c echo.Context) error {
	idStr := c.Param("press_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id parameter"})
	}

	publishBool := c.Param("publish")
	publish, err := strconv.ParseBool(publishBool)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid publish parameter"})
	}

	publishOk := s.Queries.PublishPress(c.Request().Context(), db.PublishPressParams{
		ID:          int32(id),
		IsPublished: publish,
	})
	if publishOk != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to publish blog"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Blog published"})
}
