package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository"
)

type ClassroomInputSvc struct {
	ID            int
	Title         string
	Description   string
	Status        string
	LecturerID    int
	CodeClassroom string
	TopicTags     string
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// CreateClasroom creates a new classroom in db given by classroom model
func (s *ClassroomSvc) CreateClassroom(ctx context.Context, clr ClassroomInputSvc) error {
	clrRepo := repository.ClassroomInputRepo{
		Title:         clr.Title,
		Description:   clr.Description,
		Status:        clr.Status,
		LecturerID:    clr.LecturerID,
		CodeClassroom: clr.CodeClassroom,
		TopicTags:     clr.TopicTags,
		Quantity:      clr.Quantity,
	}

	if err := s.Repository.CreateClassroom(ctx, clrRepo); err != nil {
		if errors.Is(err, repository.ErrClassroomExisted) {
			return ErrClassroomExisted
		}
		return err
	}

	return nil
}

// GetClassroom returns a classroom in db given by id
func (s *ClassroomSvc) GetClassroom(ctx context.Context, id int) (ClassroomInputSvc, error) {
	clr, err := s.Repository.GetClassroom(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrClassroomNotFound) {
			return ClassroomInputSvc{}, ErrClassroomNotFound
		}
		return ClassroomInputSvc{}, err
	}

	return ClassroomInputSvc{
		ID:            clr.ID,
		Title:         clr.Title,
		Description:   clr.Description,
		Status:        clr.Status,
		LecturerID:    clr.LecturerID,
		CodeClassroom: clr.CodeClassroom,
		TopicTags:     clr.TopicTags,
		Quantity:      clr.Quantity,
		CreatedAt:     clr.CreatedAt,
		UpdatedAt:     clr.UpdatedAt,
	}, nil
}

// CheckClassroomExists checks if a classroom with given id exists in db
func (s *ClassroomSvc) CheckClassroomExists(ctx context.Context, id int) (bool, error) {
	if _, err := s.Repository.GetClassroom(ctx, id); err != nil {
		if errors.Is(err, repository.ErrClassroomNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// UpdateClassroom updates the specified classroom by id
func (s *ClassroomSvc) UpdateClassroom(ctx context.Context, id int, classroom ClassroomInputSvc) error {
	if err := s.Repository.UpdateClassroom(ctx, id, repository.ClassroomInputRepo{
		Title:         classroom.Title,
		Description:   classroom.Description,
		Status:        classroom.Status,
		LecturerID:    classroom.LecturerID,
		CodeClassroom: classroom.CodeClassroom,
		TopicTags:     classroom.TopicTags,
		Quantity:      classroom.Quantity,
	}); err != nil {
		if errors.Is(err, repository.ErrClassroomNotFound) {
			return ErrClassroomNotFound
		}
		return err
	}

	return nil
}

// DeleteClassroom deletes a classroom in db given by id
func (s *ClassroomSvc) DeleteClassroom(ctx context.Context, id int) error {
	if err := s.Repository.DeleteClassroom(ctx, id); err != nil {
		if errors.Is(err, repository.ErrClassroomNotFound) {
			return ErrClassroomNotFound
		}
		return err
	}

	return nil
}

type ClassroomOutputSvc struct {
	ID            int
	Title         string
	Description   string
	Status        string
	LecturerID    int
	CodeClassroom string
	TopicTags     string
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ClassroomFilterSvc struct {
	Limit       int
	Page        int
	TitleSearch string
	SortColumn  string
	SortOrder   string
}

// GetClassroom returns a list of classrooms in db with filter
func (s *ClassroomSvc) GetClassrooms(ctx context.Context, filter ClassroomFilterSvc) ([]ClassroomOutputSvc, int, error) {
	clrsRepo, count, err := s.Repository.GetClassrooms(ctx, repository.ClassroomFilterRepo{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   filter.SortOrder,
	})
	if err != nil {
		return nil, 0, err
	}

	var clrsSvc []ClassroomOutputSvc
	for _, c := range clrsRepo {
		clrsSvc = append(clrsSvc, ClassroomOutputSvc{
			ID:            c.ID,
			Title:         c.Title,
			Description:   c.Description,
			Status:        c.Status,
			LecturerID:    c.LecturerID,
			CodeClassroom: c.CodeClassroom,
			TopicTags:     c.TopicTags,
			Quantity:      c.Quantity,
			CreatedAt:     c.CreatedAt,
			UpdatedAt:     c.UpdatedAt,
		})
	}

	return clrsSvc, count, nil
}
