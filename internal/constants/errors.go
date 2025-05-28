package constants

import "errors"

var (
	ErrInsertPost = errors.New("failed to insert post")

	ErrPostSlugTaken = errors.New("blog post slug is already registered")
)
