package handlers

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/naveenm4d/blog-svc/internal/constants"
	"github.com/naveenm4d/blog-svc/internal/core/adapters"
	"github.com/naveenm4d/blog-svc/proto"
)

var _ = adapters.Handler(&server{})

type server struct {
	proto.BlogSvcServer

	postsService adapters.PostService
}

func NewHandler(
	postsService adapters.PostService,
) adapters.Handler {
	server := &server{
		postsService: postsService,
	}

	return server
}

func (s *server) GetPosts(
	ctx context.Context,
	request *proto.GetPostsRequest,
) (*proto.GetPostsResponse, error) {
	posts, err := s.postsService.GetPosts(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := make([]*proto.Post, 0)
	for _, post := range posts {
		response = append(response, proto.ComposePostEntityToProto(post))
	}

	return &proto.GetPostsResponse{
		Posts: response,
	}, nil
}

func (s *server) CreatePost(
	ctx context.Context,
	request *proto.CreatePostRequest,
) (*proto.CreatePostResponse, error) {
	newPost := request.GetCreatePostEntity()

	if err := s.postsService.CreatePost(ctx, newPost); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &proto.CreatePostResponse{Success: constants.ResponseSuccess}, nil
}
