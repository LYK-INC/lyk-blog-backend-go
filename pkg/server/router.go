package server

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func initRoutes(router *echo.Echo, s *Services, l *zerolog.Logger) {
	router.GET("/", s.Health.HealthCheck)

	router.GET("docs/*", echoSwagger.WrapHandler)

	router.GET("/home/articles", s.HomePage.Articles)

	l.Info().Msgf("docs v1 at :: http://localhost:%d/docs/index.html \n", 8000)

}
