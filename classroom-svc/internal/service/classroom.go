package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository"
)

type ClassroomInputSvc struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateClasroom creates a new classroom in db given by classroom model
func (s *ClassroomSvc) CreateClassroom(ctx context.Context, clr ClassroomInputSvc) error {
	clrRepo := repository.ClassroomInputRepo{
		Title:       clr.Title,
		Description: clr.Description,
		Status:      clr.Status,
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
		ID:          clr.ID,
		Title:       clr.Title,
		Description: clr.Description.String,
		Status:      clr.Status,
		CreatedAt:   clr.CreatedAt,
		UpdatedAt:   clr.UpdatedAt,
	}, nil
}
