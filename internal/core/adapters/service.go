package adapters

import (
	"context"

	"github.com/naveenm4d/blog-svc/internal/core/entities"
)

type PostService interface {
	CreatePost(ctx context.Context, post *entities.Post) error
	GetPosts(ctx context.Context) ([]*entities.Post, error)
}
