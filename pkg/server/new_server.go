package server

import (
	"net/http"
	"strings"
	"time"

	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	adminService "github.com/LYK-INC/blog-backend-go/pkg/server/services/admin"
	"github.com/LYK-INC/blog-backend-go/pkg/server/services/blog"
	healthService "github.com/LYK-INC/blog-backend-go/pkg/server/services/health"
	"github.com/LYK-INC/blog-backend-go/pkg/server/services/homepage"
	"github.com/LYK-INC/blog-backend-go/utils/config"
	util_validator "github.com/LYK-INC/blog-backend-go/utils/validator"
	"github.com/go-playground/validator/v10"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

type Services struct {
	Health       *healthService.HealthService
	HomePage     *homepage.HomePageService
	BlogPage     *blog.BlogPageService
	AdminService *adminService.AdminPageService
}
type Server struct {
	config   config.Config
	router   *echo.Echo
	logger   *zerolog.Logger
	queries  *db.Queries
	services *Services
}

func (s *Server) GetServices() *Services {
	return s.services
}

func (s *Server) GetConfig() config.Config {
	return s.config
}

func (s *Server) GetRouter() *echo.Echo {
	return s.router
}

func (s *Server) GetLogger() *zerolog.Logger {
	return s.logger
}

func (s *Server) GetQueries() *db.Queries {
	return s.queries
}

type ServerParams struct {
	Config  config.Config
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func NewServer(c *ServerParams) (*Server, error) {
	router := echo.New()
	new_validator := validator.New()

	router.Validator = &util_validator.CustomValidator{Validator: new_validator}

	router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, request_id=${id}, remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency_nano=${latency}, bytes_in=${bytes_in}, bytes_out=${bytes_out}\n",
	}))

	router.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
	}))

	config := middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(100), Burst: 30, ExpiresIn: 1 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	router.Use(middleware.RateLimiterWithConfig(config))

	if c.Config.AppEnv != "local" {
		router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
			AllowCredentials: true,
			AllowHeaders:     []string{"application/json", "text/plain", "*/*"},
			Skipper: func(c echo.Context) bool {
				h := c.Request().Header.Get("User-Agent")
				user_agent := strings.Split(h, "/")

				return strings.EqualFold("PostmanRuntime", user_agent[0])

			},
		}))
	}

	services := initServices(initServicesParams{
		Config:  c.Config,
		Logger:  c.Logger,
		Queries: c.Queries,
	})

	initRoutes(router, services, c.Logger)

	return &Server{
		config:   c.Config,
		router:   router,
		logger:   c.Logger,
		queries:  c.Queries,
		services: services,
	}, nil
}
