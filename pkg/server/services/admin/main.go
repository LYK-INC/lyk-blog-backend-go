package admin

import (
	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/config"
	"github.com/rs/zerolog"
)

type AdminPageService struct {
	Config  config.Config
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func NewAdminPageService(params AdminPageService) *AdminPageService {
	return &AdminPageService{
		Config:  params.Config,
		Logger:  params.Logger,
		Queries: params.Queries,
	}
}
