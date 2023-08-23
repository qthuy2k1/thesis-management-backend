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

type ClassroomOutputRepo struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// UpdateClassroom updates the specified classroom by id
func (r *ClassroomRepo) UpdateClassroom(ctx context.Context, id int, classroom ClassroomInputRepo) error {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "UPDATE classrooms SET title=$2, description=$3, status=$4, updated_at=$5 WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the ID of the updated classroom
	result, err := stmt.ExecContext(ctx, id, classroom.Title, classroom.Description, classroom.Status, time.Now())
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrClassroomNotFound
	}

	return nil
}

// DeleteClassroom deletes a classroom in db given by id
func (r *ClassroomRepo) DeleteClassroom(ctx context.Context, id int) error {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "DELETE FROM classrooms WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the deleted classroom's details
	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrClassroomNotFound
	}

	return nil
}
