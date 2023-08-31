package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/exercise-svc/internal/repository"
)

type IExerciseSvc interface {
	// CreateExercise creates a new exercise in db given by exercise model
	CreateExercise(ctx context.Context, e ExerciseInputSvc) error
	// GetExercise returns a exercise in db given by id
	GetExercise(ctx context.Context, id int) (ExerciseOutputSvc, error)
	// UpdateExercise updates the specified classroom by id
	UpdateExercise(ctx context.Context, id int, classroom ExerciseInputSvc) error
	// DeleteExercise deletes a classroom in db given by id
	DeleteExercise(ctx context.Context, id int) error
	// GetExercises returns a list of exercises in db with filter
	GetExercises(ctx context.Context, filter ExerciseFilterSvc) ([]ExerciseOutputSvc, int, error)
	// GetAllExercisesOfClassroom returns a list of posts in db with filter
	GetAllExercisesOfClassroom(ctx context.Context, filter ExerciseFilterSvc, classroomID int) ([]ExerciseOutputSvc, int, error)
}

type ExerciseSvc struct {
	Repository repository.IExerciseRepo
}

func NewExerciseSvc(eRepo repository.IExerciseRepo) IExerciseSvc {
	return &ExerciseSvc{Repository: eRepo}
}
