package repositories

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"github.com/naveenm4d/blog-svc/internal/constants"
	"github.com/naveenm4d/blog-svc/internal/core/adapters"
	"github.com/naveenm4d/blog-svc/internal/core/entities"
)

var _ = adapters.PostsRepository(&postsRepository{})

type postsRepository struct {
	collection *mongo.Collection
}

func NewPostsRepository(
	collection *mongo.Collection,
) adapters.PostsRepository {
	repo := &postsRepository{
		collection: collection,
	}

	return repo
}

func (r *postsRepository) GetPosts(ctx context.Context) ([]*entities.Post, error) {
	var posts []*entities.Post

	filter := bson.M{
		"status": bson.M{"$ne": constants.StatusPublished},
	}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post entities.Post

		err := cursor.Decode(&post)
		if err != nil {
			continue
		}

		posts = append(posts, &post)
	}

	return posts, nil
}

func (r *postsRepository) CreatePost(
	ctx context.Context,
	post *entities.Post,
) error {
	filter := bson.M{
		"slug": post.Slug,
	}

	err := r.collection.FindOne(ctx, filter).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		_, err = r.collection.InsertOne(ctx, post)
		if err != nil {
			return constants.ErrInsertPost
		}

		return nil
	}

	if err != nil {
		return constants.ErrInsertPost
	}

	return constants.ErrPostSlugTaken
}
