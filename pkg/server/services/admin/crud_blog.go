package admin

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	// "net/url"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	// "github.com/LYK-INC/blog-backend-go/utils/types"
	"github.com/labstack/echo/v4"
)

func (s *AdminPageService) CreateBlogPage(c echo.Context) error {
	var req db.CreateBlogParams

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	fmt.Print(req)

	blog, err := s.Queries.CreateBlog(c.Request().Context(), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to create blog"})
	}

	return c.JSON(http.StatusOK, blog)
}

func (s *AdminPageService) GetBlogPage(c echo.Context) error {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid limit parameter"})
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid offset parameter"})
	}

	blog, err := s.Queries.GetBlogs(c.Request().Context(), db.GetBlogsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Blog not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve blog"})
	}

	fmt.Print(blog)

	return c.JSON(http.StatusOK, blog)
}

func (s *AdminPageService) GetAllAuthors(c echo.Context) error {
	authors, err := s.Queries.GetAuthors(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve authors"})
	}

	return c.JSON(http.StatusOK, authors)
}
