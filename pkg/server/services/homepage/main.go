package homepage

import (
	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/config"
	"github.com/rs/zerolog"
)

type HomePageService struct {
	Config  config.Config
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func NewHomePageService(params HomePageService) *HomePageService {
	return &HomePageService{
		Config:  params.Config,
		Logger:  params.Logger,
		Queries: params.Queries,
	}
}
