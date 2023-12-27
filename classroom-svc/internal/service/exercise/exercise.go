package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/exercise"
)

type ExerciseInputSvc struct {
	Title            string
	Content          string
	ClassroomID      int
	Deadline         time.Time
	ReportingStageID int
	AuthorID         string
}

// CreateExercise creates a new exercise in db given by exercise model
func (s *ExerciseSvc) CreateExercise(ctx context.Context, e ExerciseInputSvc) (int64, error) {
	eRepo := repository.ExerciseInputRepo{
		Title:            e.Title,
		Content:          e.Content,
		ClassroomID:      e.ClassroomID,
		Deadline:         e.Deadline,
		ReportingStageID: e.ReportingStageID,
		AuthorID:         e.AuthorID,
	}

	id, err := s.Repository.CreateExercise(ctx, eRepo)
	if err != nil {
		if errors.Is(err, repository.ErrExerciseExisted) {
			return 0, ErrExerciseExisted
		}
		return 0, err
	}

	return id, nil
}

// GetExercise returns a exercise in db given by id
func (s *ExerciseSvc) GetExercise(ctx context.Context, id int) (ExerciseOutputSvc, error) {
	e, err := s.Repository.GetExercise(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrExerciseNotFound) {
			return ExerciseOutputSvc{}, ErrExerciseNotFound
		}
		return ExerciseOutputSvc{}, err
	}

	return ExerciseOutputSvc{
		ID:               e.ID,
		Title:            e.Title,
		Content:          e.Content,
		ClassroomID:      e.ClassroomID,
		Deadline:         e.Deadline,
		ReportingStageID: e.ReportingStageID,
		AuthorID:         e.AuthorID,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
	}, nil
}

// UpdateExercise updates the specified exercise by id
func (s *ExerciseSvc) UpdateExercise(ctx context.Context, id int, exercise ExerciseInputSvc) error {
	if err := s.Repository.UpdateExercise(ctx, id, repository.ExerciseInputRepo{
		Title:            exercise.Title,
		Content:          exercise.Content,
		ClassroomID:      exercise.ClassroomID,
		Deadline:         exercise.Deadline,
		ReportingStageID: exercise.ReportingStageID,
		AuthorID:         exercise.AuthorID,
	}); err != nil {
		if errors.Is(err, repository.ErrExerciseNotFound) {
			return ErrExerciseNotFound
		}
		return err
	}

	return nil
}

// DeleteExercise deletes a exercise in db given by id
func (s *ExerciseSvc) DeleteExercise(ctx context.Context, id int) error {
	if err := s.Repository.DeleteExercise(ctx, id); err != nil {
		if errors.Is(err, repository.ErrExerciseNotFound) {
			return ErrExerciseNotFound
		}
		return err
	}

	return nil
}

type ExerciseOutputSvc struct {
	ID               int
	Title            string
	Content          string
	ClassroomID      int
	Deadline         time.Time
	ReportingStageID int
	AuthorID         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type ExerciseFilterSvc struct {
	Limit       int
	Page        int
	TitleSearch string
	SortColumn  string
	SortOrder   string
}

// GetExercise returns a list of exercises in db with filter
func (s *ExerciseSvc) GetExercises(ctx context.Context, filter ExerciseFilterSvc) ([]ExerciseOutputSvc, int, error) {
	esRepo, count, err := s.Repository.GetExercises(ctx, repository.ExerciseFilterRepo{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   filter.SortOrder,
	})
	if err != nil {
		return nil, 0, err
	}

	var psSvc []ExerciseOutputSvc
	for _, e := range esRepo {
		psSvc = append(psSvc, ExerciseOutputSvc{
			ID:               e.ID,
			Title:            e.Title,
			Content:          e.Content,
			ClassroomID:      e.ClassroomID,
			Deadline:         e.Deadline,
			ReportingStageID: e.ReportingStageID,
			AuthorID:         e.AuthorID,
			CreatedAt:        e.CreatedAt,
			UpdatedAt:        e.UpdatedAt,
		})
	}

	return psSvc, count, nil
}

// GetAllExercisesOfClassroom returns a list of exercises in a classroom in db with filter
func (s *ExerciseSvc) GetAllExercisesOfClassroom(ctx context.Context, filter ExerciseFilterSvc, classroomID int) ([]ExerciseOutputSvc, int, error) {
	psRepo, count, err := s.Repository.GetAllExercisesOfClassroom(ctx, repository.ExerciseFilterRepo{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   filter.SortOrder,
	}, classroomID)
	if err != nil {
		return nil, 0, err
	}

	var esSvc []ExerciseOutputSvc
	for _, e := range psRepo {
		esSvc = append(esSvc, ExerciseOutputSvc{
			ID:               e.ID,
			Title:            e.Title,
			Content:          e.Content,
			ClassroomID:      e.ClassroomID,
			Deadline:         e.Deadline,
			ReportingStageID: e.ReportingStageID,
			AuthorID:         e.AuthorID,
			CreatedAt:        e.CreatedAt,
			UpdatedAt:        e.UpdatedAt,
		})
	}

	return esSvc, count, nil
}

// GetAllExercisesInReportingStage returns all posts of the specified reporting stage given by reporting stage id
func (s *ExerciseSvc) GetAllExercisesInReportingStage(ctx context.Context, reportingStageID, classroomID int) ([]ExerciseOutputSvc, int, error) {
	psRepo, count, err := s.Repository.GetAllExercisesInReportingStage(ctx, reportingStageID, classroomID)
	if err != nil {
		return nil, 0, err
	}

	var esSvc []ExerciseOutputSvc
	for _, e := range psRepo {
		esSvc = append(esSvc, ExerciseOutputSvc{
			ID:               e.ID,
			Title:            e.Title,
			Content:          e.Content,
			ClassroomID:      e.ClassroomID,
			Deadline:         e.Deadline,
			ReportingStageID: e.ReportingStageID,
			AuthorID:         e.AuthorID,
			CreatedAt:        e.CreatedAt,
			UpdatedAt:        e.UpdatedAt,
		})
	}

	return esSvc, count, nil
}
