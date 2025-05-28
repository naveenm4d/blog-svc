package entities

import (
	"time"

	"github.com/naveenm4d/blog-svc/internal/constants"
)

type Post struct {
	PostID   string           `json:"postId" bson:"postId"`
	AuthorID string           `json:"authorId" bson:"authorId"`
	Slug     string           `json:"slug" bson:"slug"`
	Title    string           `json:"title" bson:"title"`
	Content  string           `json:"content" bson:"content"`
	Status   constants.Status `json:"status" bson:"status"`

	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updatedAt"`
}
