package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/comment-svc/internal/repository"
)

type CommentInputSvc struct {
	UserID     string
	PostID     *int
	ExerciseID *int
	Content    string
}

// CreateClasroom creates a new comment in db given by comment model
func (s *CommentSvc) CreateComment(ctx context.Context, comment CommentInputSvc) error {
	commentRepo := repository.CommentInputRepo{
		UserID:     comment.UserID,
		PostID:     comment.PostID,
		ExerciseID: comment.ExerciseID,
		Content:    comment.Content,
	}

	if err := s.Repository.CreateComment(ctx, commentRepo); err != nil {
		if errors.Is(err, repository.ErrCommentExisted) {
			return ErrCommentExisted
		}
		return err
	}

	return nil
}

// GetComment returns a comment in db given by id
func (s *CommentSvc) GetComment(ctx context.Context, id int) (CommentOutputSvc, error) {
	comment, err := s.Repository.GetComment(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrCommentNotFound) {
			return CommentOutputSvc{}, ErrCommentNotFound
		}
		return CommentOutputSvc{}, err
	}

	return CommentOutputSvc{
		ID:         comment.ID,
		UserID:     comment.UserID,
		PostID:     comment.PostID,
		ExerciseID: comment.ExerciseID,
		Content:    comment.Content,
		CreatedAt:  comment.CreatedAt,
	}, nil
}

type CommentOutputSvc struct {
	ID         int
	UserID     string
	PostID     *int
	ExerciseID *int
	Content    string
	CreatedAt  time.Time
}

// GetComment returns a list of comments in db with filter
func (s *CommentSvc) GetCommentsOfAPost(ctx context.Context, postID int) ([]CommentOutputSvc, error) {
	commentsRepo, err := s.Repository.GetCommentsOfAPost(ctx, postID)
	if err != nil {
		return nil, err
	}

	var commentsSvc []CommentOutputSvc
	for _, c := range commentsRepo {
		commentsSvc = append(commentsSvc, CommentOutputSvc{
			ID:         c.ID,
			UserID:     c.UserID,
			PostID:     c.PostID,
			ExerciseID: c.ExerciseID,
			Content:    c.Content,
			CreatedAt:  c.CreatedAt,
		})
	}

	return commentsSvc, nil
}

// GetComment returns a list of comments in db with filter
func (s *CommentSvc) GetCommentsOfAExercise(ctx context.Context, exerciseID int) ([]CommentOutputSvc, error) {
	commentsRepo, err := s.Repository.GetCommentsOfAExercise(ctx, exerciseID)
	if err != nil {
		return nil, err
	}

	var commentsSvc []CommentOutputSvc
	for _, c := range commentsRepo {
		commentsSvc = append(commentsSvc, CommentOutputSvc{
			ID:         c.ID,
			UserID:     c.UserID,
			PostID:     c.PostID,
			ExerciseID: c.ExerciseID,
			Content:    c.Content,
			CreatedAt:  c.CreatedAt,
		})
	}

	return commentsSvc, nil
}
