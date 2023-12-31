package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/post"
)

type PostInputSvc struct {
	Title            string
	Content          string
	ClassroomID      int
	ReportingStageID int
	AuthorID         string
}

type PostOutputSvc struct {
	ID               int
	Title            string
	Content          string
	ClassroomID      int
	ReportingStageID int
	AuthorID         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// CreatePost creates a new post in db given by post model
func (s *PostSvc) CreatePost(ctx context.Context, p PostInputSvc) (PostOutputSvc, error) {
	pRepo := repository.PostInputRepo{
		Title:            p.Title,
		Content:          p.Content,
		ClassroomID:      p.ClassroomID,
		ReportingStageID: p.ReportingStageID,
		AuthorID:         p.AuthorID,
	}
	pRes, err := s.Repository.CreatePost(ctx, pRepo)
	if err != nil {
		if errors.Is(err, repository.ErrPostExisted) {
			return PostOutputSvc{}, ErrPostExisted
		}
		return PostOutputSvc{}, err
	}

	return PostOutputSvc{
		ID:               pRes.ID,
		Title:            pRes.Title,
		Content:          pRes.Content,
		ClassroomID:      pRes.ClassroomID,
		ReportingStageID: pRes.ReportingStageID,
		AuthorID:         pRes.AuthorID,
		CreatedAt:        pRes.CreatedAt,
		UpdatedAt:        pRes.UpdatedAt,
	}, nil
}

// GetPost returns a post in db given by id
func (s *PostSvc) GetPost(ctx context.Context, id int) (PostOutputSvc, error) {
	p, err := s.Repository.GetPost(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			return PostOutputSvc{}, ErrPostNotFound
		}
		return PostOutputSvc{}, err
	}

	return PostOutputSvc{
		ID:               p.ID,
		Title:            p.Title,
		Content:          p.Content,
		ClassroomID:      p.ClassroomID,
		ReportingStageID: p.ReportingStageID,
		AuthorID:         p.AuthorID,
		CreatedAt:        p.CreatedAt,
		UpdatedAt:        p.UpdatedAt,
	}, nil
}

// UpdatePost updates the specified post by id
func (s *PostSvc) UpdatePost(ctx context.Context, id int, post PostInputSvc) error {
	if err := s.Repository.UpdatePost(ctx, id, repository.PostInputRepo{
		Title:            post.Title,
		Content:          post.Content,
		ClassroomID:      post.ClassroomID,
		ReportingStageID: post.ReportingStageID,
		AuthorID:         post.AuthorID,
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

type PostFilterSvc struct {
	Limit       int
	Page        int
	TitleSearch string
	SortColumn  string
	SortOrder   string
}

// GetPosts returns a list of posts in db with filter
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
			ID:               p.ID,
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      p.ClassroomID,
			ReportingStageID: p.ReportingStageID,
			AuthorID:         p.AuthorID,
			CreatedAt:        p.CreatedAt,
			UpdatedAt:        p.UpdatedAt,
		})
	}

	return psSvc, count, nil
}

// GetAllPostsOfClassroom returns a list of posts in a classroom in db with filter
func (s *PostSvc) GetAllPostsOfClassroom(ctx context.Context, filter PostFilterSvc, classroomID int) ([]PostOutputSvc, int, error) {
	psRepo, count, err := s.Repository.GetAllPostsOfClassroom(ctx, repository.PostFilterRepo{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   filter.SortOrder,
	}, classroomID)
	if err != nil {
		return nil, 0, err
	}

	var psSvc []PostOutputSvc
	for _, p := range psRepo {
		psSvc = append(psSvc, PostOutputSvc{
			ID:               p.ID,
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      p.ClassroomID,
			ReportingStageID: p.ReportingStageID,
			AuthorID:         p.AuthorID,
			CreatedAt:        p.CreatedAt,
			UpdatedAt:        p.UpdatedAt,
		})
	}

	return psSvc, count, nil
}

// GetAllPostsInReportingStage returns all posts of the specified reporting stage given by reporting stage id
func (s *PostSvc) GetAllPostsInReportingStage(ctx context.Context, reportingStageID, classroomID int) ([]PostOutputSvc, int, error) {
	psRepo, count, err := s.Repository.GetAllPostsInReportingStage(ctx, reportingStageID, classroomID)
	if err != nil {
		return nil, 0, err
	}

	var psSvc []PostOutputSvc
	for _, p := range psRepo {
		psSvc = append(psSvc, PostOutputSvc{
			ID:               p.ID,
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      p.ClassroomID,
			ReportingStageID: p.ReportingStageID,
			AuthorID:         p.AuthorID,
			CreatedAt:        p.CreatedAt,
			UpdatedAt:        p.UpdatedAt,
		})
	}

	return psSvc, count, nil
}
