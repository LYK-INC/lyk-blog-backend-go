package blog

import (
	db "github.com/LYK-INC/blog-backend-go/db/sqlc"
	"github.com/LYK-INC/blog-backend-go/utils/config"
	"github.com/rs/zerolog"
)

type BlogPageService struct {
	Config  config.Config
	Logger  *zerolog.Logger
	Queries *db.Queries
}

func NewBlogPageService(params BlogPageService) *BlogPageService {
	return &BlogPageService{
		Config:  params.Config,
		Logger:  params.Logger,
		Queries: params.Queries,
	}
}
