package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/submission"
)

type SubmissionInputSvc struct {
	UserID     string
	ExerciseID int
	Status     string
}

// CreateSubmission creates a new submission in db given by exercise model
func (s *SubmissionSvc) CreateSubmission(ctx context.Context, submission SubmissionInputSvc) (int64, error) {
	sRepo := repository.SubmissionInputRepo{
		UserID:     submission.UserID,
		ExerciseID: submission.ExerciseID,
		Status:     submission.Status,
	}

	id, err := s.Repository.CreateSubmission(ctx, sRepo)
	if err != nil {
		if errors.Is(err, repository.ErrSubmissionExisted) {
			return 0, ErrSubmissionExisted
		}
		return 0, err
	}

	return id, nil
}

// UpdateSubmission updates the specified submission by id
func (s *SubmissionSvc) UpdateSubmission(ctx context.Context, id int, submission SubmissionInputSvc) error {
	if err := s.Repository.UpdateSubmission(ctx, id, repository.SubmissionInputRepo{
		UserID:     submission.UserID,
		ExerciseID: submission.ExerciseID,
		Status:     submission.Status,
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
	ID         int
	UserID     string
	ExerciseID int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
			ID:         submission.ID,
			UserID:     submission.UserID,
			ExerciseID: submission.ExerciseID,
			Status:     submission.Status,
			CreatedAt:  submission.CreatedAt,
			UpdatedAt:  submission.UpdatedAt,
		})
	}

	return ssSvc, count, nil
}

func (s *SubmissionSvc) GetSubmissionOfUser(ctx context.Context, userID string, exerciseID int) ([]SubmissionOutputSvc, error) {
	submission, err := s.Repository.GetSubmissionOfUser(ctx, userID, exerciseID)
	if err != nil {
		return nil, err
	}

	var ssSvc []SubmissionOutputSvc
	for _, submission := range submission {
		ssSvc = append(ssSvc, SubmissionOutputSvc{
			ID:         submission.ID,
			UserID:     submission.UserID,
			ExerciseID: submission.ExerciseID,
			Status:     submission.Status,
			CreatedAt:  submission.CreatedAt,
			UpdatedAt:  submission.UpdatedAt,
		})
	}

	return ssSvc, nil
}

func (s *SubmissionSvc) GetAllSubmissionFromUser(ctx context.Context, userID string) ([]SubmissionOutputSvc, error) {
	submission, err := s.Repository.GetAllSubmissionFromUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	var ssSvc []SubmissionOutputSvc
	for _, submission := range submission {
		ssSvc = append(ssSvc, SubmissionOutputSvc{
			ID:         submission.ID,
			UserID:     submission.UserID,
			ExerciseID: submission.ExerciseID,
			Status:     submission.Status,
			CreatedAt:  submission.CreatedAt,
			UpdatedAt:  submission.UpdatedAt,
		})
	}

	return ssSvc, nil
}
