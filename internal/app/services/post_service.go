package services

import (
	"context"

	"github.com/naveenm4d/blog-svc/internal/core/adapters"
	"github.com/naveenm4d/blog-svc/internal/core/entities"
)

var _ = adapters.PostService(&postService{})

type postService struct {
	postsRepository adapters.PostsRepository
}

func NewPostService(postsRepository adapters.PostsRepository) adapters.PostService {
	service := &postService{
		postsRepository: postsRepository,
	}

	return service
}

func (s *postService) CreatePost(
	ctx context.Context,
	post *entities.Post,
) error {
	return s.postsRepository.CreatePost(ctx, post)
}

func (s *postService) GetPosts(
	ctx context.Context,
) ([]*entities.Post, error) {
	return s.postsRepository.GetPosts(ctx)
}
