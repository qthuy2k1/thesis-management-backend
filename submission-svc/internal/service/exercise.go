package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/submission-svc/internal/repository"
)

type SubmissionInputSvc struct {
	UserID         string
	ExerciseID     int
	SubmissionDate time.Time
	Status         string
}

// CreateSubmission creates a new submission in db given by exercise model
func (s *SubmissionSvc) CreateSubmission(ctx context.Context, submission SubmissionInputSvc) error {
	sRepo := repository.SubmissionInputRepo{
		UserID:         submission.UserID,
		ExerciseID:     submission.ExerciseID,
		SubmissionDate: submission.SubmissionDate,
		Status:         submission.Status,
	}

	if err := s.Repository.CreateSubmission(ctx, sRepo); err != nil {
		if errors.Is(err, repository.ErrSubmissionExisted) {
			return ErrSubmissionExisted
		}
		return err
	}

	return nil
}

// UpdateSubmission updates the specified submission by id
func (s *SubmissionSvc) UpdateSubmission(ctx context.Context, id int, submission SubmissionInputSvc) error {
	if err := s.Repository.UpdateSubmission(ctx, id, repository.SubmissionInputRepo{
		UserID:         submission.UserID,
		ExerciseID:     submission.ExerciseID,
		SubmissionDate: submission.SubmissionDate,
		Status:         submission.Status,
	}); err != nil {
		if errors.Is(err, repository.ErrSubmissionNotFound) {
			return ErrSubmissionNotFound
		}
		return err
	}

	return nil
}

// DeleteSubmission deletes a submission in db given by id
func (s *SubmissionSvc) DeleteSubmission(ctx context.Context, id int) error {
	if err := s.Repository.DeleteSubmission(ctx, id); err != nil {
		if errors.Is(err, repository.ErrSubmissionNotFound) {
			return ErrSubmissionNotFound
		}
		return err
	}

	return nil
}

type SubmissionOutputSvc struct {
	ID             int
	UserID         string
	ExerciseID     int
	SubmissionDate time.Time
	Status         string
}

// GetAllSubmissionsOfExercise returns a list of submissions in a exercise in db
func (s *SubmissionSvc) GetAllSubmissionsOfExercise(ctx context.Context, exerciseID int) ([]SubmissionOutputSvc, int, error) {
	sRepo, count, err := s.Repository.GetAllSubmissionsOfExercise(ctx, exerciseID)
	if err != nil {
		return nil, 0, err
	}

	var ssSvc []SubmissionOutputSvc
	for _, submission := range sRepo {
		ssSvc = append(ssSvc, SubmissionOutputSvc{
			ID:             submission.ID,
			UserID:         submission.UserID,
			ExerciseID:     submission.ExerciseID,
			SubmissionDate: submission.SubmissionDate,
			Status:         submission.Status,
		})
	}

	return ssSvc, count, nil
}
