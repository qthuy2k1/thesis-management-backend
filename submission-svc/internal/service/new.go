package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/submission-svc/internal/repository"
)

type ISubmissionSvc interface {
	// CreateSubmission creates a new submission in db given by submission model
	CreateSubmission(ctx context.Context, e SubmissionInputSvc) error
	// UpdateSubmission updates the specified submission by id
	UpdateSubmission(ctx context.Context, id int, classroom SubmissionInputSvc) error
	// DeleteSubmission deletes a submission in db given by id
	DeleteSubmission(ctx context.Context, id int) error
	// GetAllSubmissionsOfExercise returns a list of submissions in db
	GetAllSubmissionsOfExercise(ctx context.Context, classroomID int) ([]SubmissionOutputSvc, int, error)
}

type SubmissionSvc struct {
	Repository repository.ISubmissionRepo
}

func NewSubmissionSvc(eRepo repository.ISubmissionRepo) ISubmissionSvc {
	return &SubmissionSvc{Repository: eRepo}
}
