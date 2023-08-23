package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository"
)

type IClassroomSvc interface {
	// CreateClasroom creates a new classroom in db given by classroom model
	CreateClassroom(ctx context.Context, clr ClassroomInputSvc) error
	// GetClassroom returns a classroom in db given by id
	GetClassroom(ctx context.Context, id int) (ClassroomInputSvc, error)
	// CheckClassroomExists checks if a classroom with given id exists in db
	CheckClassroomExists(ctx context.Context, id int) (bool, error)
	// UpdateClassroom updates the specified classroom by id
	UpdateClassroom(ctx context.Context, id int, classroom ClassroomInputSvc) error
	// DeleteClassroom deletes a classroom in db given by id
	DeleteClassroom(ctx context.Context, id int) error
}

type ClassroomSvc struct {
	Repository repository.IClassroomRepo
}

func NewClassroomSvc(clrRepo repository.IClassroomRepo) IClassroomSvc {
	return &ClassroomSvc{Repository: clrRepo}
}
