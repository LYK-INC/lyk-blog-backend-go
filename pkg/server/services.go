package server

import (
	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	adminService "github.com/LYK-INC/blog-backend-go/pkg/server/services/admin"
	"github.com/LYK-INC/blog-backend-go/pkg/server/services/blog"
	healthService "github.com/LYK-INC/blog-backend-go/pkg/server/services/health"
	"github.com/LYK-INC/blog-backend-go/pkg/server/services/homepage"
	"github.com/LYK-INC/blog-backend-go/utils/config"
	"github.com/rs/zerolog"
)

type initServicesParams struct {
	Config  config.Config
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func initServices(p initServicesParams) *Services {
	health_service := healthService.NewHealthService()
	admin_service := adminService.NewAdminPageService(adminService.AdminPageService{
		Config:  p.Config,
		Logger:  p.Logger,
		Queries: p.Queries,
	})

	homepage_service := homepage.NewHomePageService(homepage.HomePageService{
		Config:  p.Config,
		Logger:  p.Logger,
		Queries: p.Queries,
	})

	blogpage_service := blog.NewBlogPageService(blog.BlogPageService{
		Config:  p.Config,
		Logger:  p.Logger,
		Queries: p.Queries,
	})
	return &Services{
		Health:       health_service,
		HomePage:     homepage_service,
		BlogPage:     blogpage_service,
		AdminService: admin_service,
	}
}
