package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/post-svc/internal/repository"
)

type IPostSvc interface {
	// CreatePost creates a new post in db given by post model
	CreatePost(ctx context.Context, p PostInputSvc) error
	// GetPost returns a post in db given by id
	GetPost(ctx context.Context, id int) (PostInputSvc, error)
	// UpdatePost updates the specified classroom by id
	UpdatePost(ctx context.Context, id int, classroom PostInputSvc) error
	// DeletePost deletes a classroom in db given by id
	DeletePost(ctx context.Context, id int) error
}

type PostSvc struct {
	Repository repository.IPostRepo
}

func NewPostSvc(pRepo repository.IPostRepo) IPostSvc {
	return &PostSvc{Repository: pRepo}
}
