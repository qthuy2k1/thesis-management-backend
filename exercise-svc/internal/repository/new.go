package repository

import (
	"context"
	"database/sql"
	// model "github.com/qthuy2k1/thesis-management-backend/exercise-svc/internal/model"
)

type IExerciseRepo interface {
	// CreateExercise creates a new exercise in db given by exercise model
	CreateExercise(ctx context.Context, clr ExerciseInputRepo) error
	// GetExercise returns a exercise in db given by id
	GetExercise(ctx context.Context, id int) (ExerciseOutputRepo, error)
	// CheckExerciseExists checks whether the specified exercise exists by name
	IsExerciseExists(ctx context.Context, classroomID int, title string) (bool, error)
	// UpdateExercise updates the specified classroom by id
	UpdateExercise(ctx context.Context, id int, classroom ExerciseInputRepo) error
	// DeleteExercise deletes a classroom in db given by id
	DeleteExercise(ctx context.Context, id int) error
	// GetExercises returns a list of exercises in db with filter
	GetExercises(ctx context.Context, filter ExerciseFilterRepo) ([]ExerciseOutputRepo, int, error)
	// GetAllExercisesOfClassroom returns all posts of the specified classroom given by classroom id
	GetAllExercisesOfClassroom(ctx context.Context, filter ExerciseFilterRepo, classromID int) ([]ExerciseOutputRepo, int, error)
}

type ExerciseRepo struct {
	Database *sql.DB
}

func NewExerciseRepo(db *sql.DB) IExerciseRepo {
	return &ExerciseRepo{Database: db}
}
