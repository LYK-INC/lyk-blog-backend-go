package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func initRoutes(r *echo.Echo, s *Services, l *zerolog.Logger) {
	// Enable CORS middleware
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	r.GET("/", s.Health.HealthCheck)
	r.GET("docs/*", echoSwagger.WrapHandler)
	l.Info().Msgf("docs v1 at :: http://localhost:%d/docs/index.html \n", 8000)

	h := r.Group("/home")
	h.GET("/articles", s.HomePage.Articles)
	h.GET("/featured", s.HomePage.FeaturedBlog)
	h.GET("/press", s.HomePage.PressArticles)

	b := r.Group("/blog")
	b.GET("/:blog_slug", s.BlogPage.GetBlogBySlug)
	b.GET("/related/:blog_id", s.BlogPage.GetRelatedBlogs)

	a := r.Group("/admin")
	a.POST("/create-blog", s.AdminService.CreateBlogPage)
	a.POST("/create-press", s.AdminService.CreatePressPage)
	a.GET("/get-press", s.AdminService.GetAllPresses)
	a.GET("/get-blog/:limit/:offset", s.AdminService.GetBlogPage)
}
