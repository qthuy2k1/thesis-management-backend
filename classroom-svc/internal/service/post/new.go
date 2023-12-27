package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/post"
)

type IPostSvc interface {
	// CreatePost creates a new post in db given by post model
	CreatePost(ctx context.Context, p PostInputSvc) (PostOutputSvc, error)
	// GetPost returns a post in db given by id
	GetPost(ctx context.Context, id int) (PostOutputSvc, error)
	// UpdatePost updates the specified post by id
	UpdatePost(ctx context.Context, id int, classroom PostInputSvc) error
	// DeletePost deletes a post in db given by id
	DeletePost(ctx context.Context, id int) error
	// GetPosts returns a list of posts in db with filter
	GetPosts(ctx context.Context, filter PostFilterSvc) ([]PostOutputSvc, int, error)
	// GetAllPostsOfClassroom returns a list of posts in classroom with filte
	GetAllPostsOfClassroom(ctx context.Context, filter PostFilterSvc, classroomID int) ([]PostOutputSvc, int, error)
	// GetAllPostsInReportingStage returns all posts of the specified reporting stage given by reporting stage id
	GetAllPostsInReportingStage(ctx context.Context, reportingStageID, classroomID int) ([]PostOutputSvc, int, error)
}

type PostSvc struct {
	Repository repository.IPostRepo
}

func NewPostSvc(pRepo repository.IPostRepo) IPostSvc {
	return &PostSvc{Repository: pRepo}
}
