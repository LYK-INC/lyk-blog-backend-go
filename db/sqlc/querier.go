// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	AddRole(ctx context.Context, arg AddRoleParams) error
	CheckBlogId(ctx context.Context, id int32) (bool, error)
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (int32, error)
	CreateBlog(ctx context.Context, arg CreateBlogParams) (int32, error)
	CreatePress(ctx context.Context, arg CreatePressParams) (int32, error)
	FeatureBlog(ctx context.Context, id int32) error
	GetBlogById(ctx context.Context, id int32) (GetBlogByIdRow, error)
	GetBlogByTitleSlug(ctx context.Context, title string) (GetBlogByTitleSlugRow, error)
	GetBlogInCategory(ctx context.Context, arg GetBlogInCategoryParams) ([]Blog, error)
	GetBlogs(ctx context.Context, arg GetBlogsParams) ([]GetBlogsRow, error)
	GetFeaturedBlog(ctx context.Context) (GetFeaturedBlogRow, error)
	GetPressById(ctx context.Context, id int32) (Press, error)
	GetPressInCategory(ctx context.Context, arg GetPressInCategoryParams) ([]Press, error)
	GetPresses(ctx context.Context, arg GetPressesParams) ([]GetPressesRow, error)
	GetRealatedBlogsById(ctx context.Context, arg GetRealatedBlogsByIdParams) ([]GetRealatedBlogsByIdRow, error)
	// Ensure role isn't already present
	RemoveRole(ctx context.Context, arg RemoveRoleParams) error
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error
}

var _ Querier = (*Queries)(nil)
