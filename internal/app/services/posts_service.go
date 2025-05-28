package services

import (
	"context"

	"github.com/naveenm4d/blog-svc/internal/core/adapters"
	"github.com/naveenm4d/blog-svc/internal/core/entities"
)

var _ = adapters.PostsService(&postsService{})

type postsService struct {
	postsRepository adapters.PostsRepository
}

func NewPostsService(postsRepository adapters.PostsRepository) adapters.PostsService {
	service := &postsService{
		postsRepository: postsRepository,
	}

	return service
}

func (s *postsService) CreatePost(
	ctx context.Context,
	post *entities.Post,
) error {
	return s.postsRepository.CreatePost(ctx, post)
}

func (s *postsService) GetPosts(
	ctx context.Context,
) ([]*entities.Post, error) {
	return s.postsRepository.GetPosts(ctx)
}
