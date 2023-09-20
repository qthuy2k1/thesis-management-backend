package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/comment-svc/internal/repository"
)

type ICommentSvc interface {
	// CreateClasroom creates a new comment in db given by comment model
	CreateComment(ctx context.Context, comment CommentInputSvc) error
	// GetComment returns a comment in db given by id
	GetComment(ctx context.Context, id int) (CommentOutputSvc, error)
	// GetCommentOfAPost returns a list of comments in db
	GetCommentsOfAPost(ctx context.Context, postID int) ([]CommentOutputSvc, error)
	// GetComment returns a list of comments in db
	GetCommentsOfAExercise(ctx context.Context, exerciseID int) ([]CommentOutputSvc, error)
}

type CommentSvc struct {
	Repository repository.ICommentRepo
}

func NewCommentSvc(commentRepo repository.ICommentRepo) ICommentSvc {
	return &CommentSvc{Repository: commentRepo}
}
