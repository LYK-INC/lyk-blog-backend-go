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

func (s *AdminPageService) UpdateFeatureBlog(c echo.Context) error {
	idStr := c.Param("blog_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id parameter"})
	}

	featureOk := s.Queries.FeatureBlog(c.Request().Context(), int32(id))
	if featureOk != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update blog to Featured"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Blog updated to Featured"})
}

func (s *AdminPageService) DeleteBlog(c echo.Context) error {
	idStr := c.Param("blog_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id parameter"})
	}

	deleteOk := s.Queries.SoftDeleteBlog(c.Request().Context(), int32(id))
	if deleteOk != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete blog"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Blog deleted"})
}

func (s *AdminPageService) PublishBlog(c echo.Context) error {
	idStr := c.Param("blog_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id parameter"})
	}

	publishBool := c.Param("publish")
	publish, err := strconv.ParseBool(publishBool)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid publish parameter"})
	}

	publishOk := s.Queries.PublishBlog(c.Request().Context(), db.PublishBlogParams{
		ID:          int32(id),
		IsPublished: publish,
	})
	if publishOk != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to publish blog"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Blog published"})
}

func (s *AdminPageService) GetBlogPageById(c echo.Context) error {
	idStr := c.Param("blog_id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid limit parameter"})
	}

	blog, err := s.Queries.GetBlogById(c.Request().Context(), int32(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, echo.Map{"error": "Blog not found"})
		}
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve blog"})
	}

	fmt.Print(blog)

	return c.JSON(http.StatusOK, blog)
}
