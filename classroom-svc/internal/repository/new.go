package repository

import (
	"context"
	"database/sql"

	model "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/model"
)

type IClassroomRepo interface {
	// CreateClassroom creates a new classroom in db given by classroom model
	CreateClassroom(ctx context.Context, clr ClassroomInputRepo) error
	// GetClassroom returns a classroom in db given by id
	GetClassroom(ctx context.Context, id int) (model.Classroom, error)
	// CheckClassroomExists checks whether the specified classroom exists by name
	IsClassroomExists(ctx context.Context, title string) (bool, error)
	// UpdateClassroom updates the specified classroom by id
	UpdateClassroom(ctx context.Context, id int, classroom ClassroomInputRepo) error
	// DeleteClassroom deletes a classroom in db given by id
	DeleteClassroom(ctx context.Context, id int) error
}

type ClassroomRepo struct {
	Database *sql.DB
}

func NewClassroomRepo(db *sql.DB) IClassroomRepo {
	return &ClassroomRepo{Database: db}
}
