package server

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func initRoutes(r *echo.Echo, s *Services, l *zerolog.Logger) {
	r.GET("/", s.Health.HealthCheck)
	r.GET("docs/*", echoSwagger.WrapHandler)
	l.Info().Msgf("docs v1 at :: http://localhost:%d/docs/index.html \n", 8000)

	h := r.Group("/home")
	h.GET("/articles", s.HomePage.Articles)
	h.GET("/featured", s.HomePage.FeaturedBlog)
	h.GET("/press", s.HomePage.PressArticles)

	b := r.Group("/blog")
	b.GET("/:blog_slug", s.BlogPage.BlogBySlug)

}
