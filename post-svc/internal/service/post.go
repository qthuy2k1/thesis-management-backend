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

// UpdatePost updates the specified post by id
func (s *PostSvc) UpdatePost(ctx context.Context, id int, post PostInputSvc) error {
	if err := s.Repository.UpdatePost(ctx, id, repository.PostInputRepo{
		Title:       post.Title,
		Content:     post.Content,
		ClassroomID: post.ClassroomID,
	}); err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			return ErrPostNotFound
		}
		return err
	}

	return nil
}

// DeletePost deletes a post in db given by id
func (s *PostSvc) DeletePost(ctx context.Context, id int) error {
	if err := s.Repository.DeletePost(ctx, id); err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			return ErrPostNotFound
		}
		return err
	}

	return nil
}

type PostOutputSvc struct {
	ID          int
	Title       string
	Content     string
	ClassroomID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PostFilterSvc struct {
	Limit       int
	Page        int
	TitleSearch string
	SortColumn  string
	SortOrder   string
}

// GetPost returns a list of posts in db with filter
func (s *PostSvc) GetPosts(ctx context.Context, filter PostFilterSvc) ([]PostOutputSvc, int, error) {
	psRepo, count, err := s.Repository.GetPosts(ctx, repository.PostFilterRepo{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   filter.SortOrder,
	})
	if err != nil {
		return nil, 0, err
	}

	var psSvc []PostOutputSvc
	for _, p := range psRepo {
		psSvc = append(psSvc, PostOutputSvc{
			ID:          p.ID,
			Title:       p.Title,
			Content:     p.Content,
			ClassroomID: p.ClassroomID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return psSvc, count, nil
}
