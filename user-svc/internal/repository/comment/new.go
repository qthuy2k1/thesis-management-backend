package repository

import (
	"context"
	"database/sql"
)

type ICommentRepo interface {
	// CreateComment creates a new comment in db given by comment model
	CreateComment(ctx context.Context, comment CommentInputRepo) error
	// GetComment returns a comment in db given by id
	GetComment(ctx context.Context, id int) (CommentOutputRepo, error)
	// GetCommentAPost returns a list of comments in db
	GetCommentsOfAPost(ctx context.Context, postID int) ([]CommentOutputRepo, error)
	// GetCommentOfAExercise returns a list of comments in db
	GetCommentsOfAExercise(ctx context.Context, exerciseID int) ([]CommentOutputRepo, error)
}

type CommentRepo struct {
	Database *sql.DB
}

func NewCommentRepo(db *sql.DB) ICommentRepo {
	return &CommentRepo{Database: db}
}
