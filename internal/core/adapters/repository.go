package adapters

import (
	"context"

	"github.com/naveenm4d/blog-svc/internal/core/entities"
)

type PostsRepository interface {
	GetPosts(ctx context.Context) ([]*entities.Post, error)
	CreatePost(ctx context.Context, post *entities.Post) error
}
