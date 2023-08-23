package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/post-svc/internal/repository"
)

type PostInputSvc struct {
	ID          int
	Title       string
	Content     string
	ClassroomID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreatePost creates a new post in db given by post model
func (s *PostSvc) CreatePost(ctx context.Context, p PostInputSvc) error {
	pRepo := repository.PostInputRepo{
		Title:       p.Title,
		Content:     p.Content,
		ClassroomID: p.ClassroomID,
	}

	if err := s.Repository.CreatePost(ctx, pRepo); err != nil {
		if errors.Is(err, repository.ErrPostExisted) {
			return ErrPostExisted
		}
		return err
	}

	return nil
}

// GetPost returns a post in db given by id
func (s *PostSvc) GetPost(ctx context.Context, id int) (PostInputSvc, error) {
	p, err := s.Repository.GetPost(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			return PostInputSvc{}, ErrPostNotFound
		}
		return PostInputSvc{}, err
	}

	return PostInputSvc{
		ID:          p.ID,
		Title:       p.Title,
		Content:     p.Content,
		ClassroomID: p.ClassroomID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}, nil
}
