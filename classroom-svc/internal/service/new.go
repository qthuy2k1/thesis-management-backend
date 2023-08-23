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
}

type ClassroomSvc struct {
	Repository repository.IClassroomRepo
}

func NewClassroomSvc(clrRepo repository.IClassroomRepo) IClassroomSvc {
	return &ClassroomSvc{Repository: clrRepo}
}
