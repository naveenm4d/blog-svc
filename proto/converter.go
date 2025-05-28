package proto

import (
	"github.com/google/uuid"
	"github.com/naveenm4d/blog-svc/internal/constants"
	"github.com/naveenm4d/blog-svc/internal/core/entities"
)

func (r *CreatePostRequest) GetCreatePostEntity() *entities.Post {
	return &entities.Post{
		PostID:   uuid.NewString(),
		AuthorID: r.GetAuthorId(),
		Slug:     r.GetSlug(),
		Title:    r.GetTitle(),
		Content:  r.GetContent(),
		Status:   composeProtoToStatus(r.GetStatus()),
	}
}

func ComposePostEntityToProto(post *entities.Post) *Post {
	return &Post{
		Id:       post.PostID,
		AuthorId: post.AuthorID,
		Slug:     post.Slug,
		Title:    post.Title,
		Content:  post.Content,
		Status:   composeStatusToProto(post.Status),
	}
}

func composeProtoToStatus(status PostStatus) constants.Status {
	switch status {
	case PostStatus_Drafted:
		return constants.StatusDrafted
	case PostStatus_Published:
		return constants.StatusPublished
	case PostStatus_Archived:
		return constants.StatusArchived
	}

	return constants.StatusUndefined
}

func composeStatusToProto(status constants.Status) PostStatus {
	switch status {
	case constants.StatusDrafted:
		return PostStatus_Drafted
	case constants.StatusPublished:
		return PostStatus_Published
	case constants.StatusArchived:
		return PostStatus_Archived
	}

	return PostStatus_Undefined
}
