package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	model "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/model"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ClassroomInputRepo struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateClasroom creates a new classroom in db given by classroom model
func (r *ClassroomRepo) CreateClassroom(ctx context.Context, clr ClassroomInputRepo) error {
	clrModel := model.Classroom{
		Title:       clr.Title,
		Description: null.NewString(clr.Description, true),
		Status:      clr.Status,
	}

	// check classroom exists
	isExists, err := r.IsClassroomExists(ctx, clrModel.Title)
	if err != nil {
		return err
	}

	if isExists {
		return ErrClassroomExisted
	}

	if err := clrModel.Insert(ctx, r.Database, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// GetClassroom returns a classroom in db given by id
func (r *ClassroomRepo) GetClassroom(ctx context.Context, id int) (model.Classroom, error) {
	clr, err := model.FindClassroom(ctx, r.Database, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Classroom{}, ErrClassroomNotFound
		}
		return model.Classroom{}, err
	}

	return *clr, nil
}

// CheckClassroomExists checks whether the specified classroom exists by title (true == exist)
func (r *ClassroomRepo) IsClassroomExists(ctx context.Context, title string) (bool, error) {
	count, err := model.Classrooms(qm.Where("title LIKE ?", "%"+title+"%")).Count(ctx, r.Database)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
