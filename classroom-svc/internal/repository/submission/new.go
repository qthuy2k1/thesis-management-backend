package repository

import (
	"context"
	"database/sql"
)

type ISubmissionRepo interface {
	// CreateSubmission creates a new submission in db given by submission model
	CreateSubmission(ctx context.Context, clr SubmissionInputRepo) (int64, error)
	// UpdateSubmission updates the specified submission by id
	UpdateSubmission(ctx context.Context, id int, submission SubmissionInputRepo) error
	// DeleteSubmission deletes a submission in db given by id
	DeleteSubmission(ctx context.Context, id int) error
	// GetAllSubmissionsOfExercise returns all submission of the specified exercise given by exercise id
	GetAllSubmissionsOfExercise(ctx context.Context, exerciseID int) ([]SubmissionOutputRepo, int, error)
	GetSubmissionOfUser(ctx context.Context, userID string, exerciseID int) ([]SubmissionOutputRepo, error)
	GetAllSubmissionFromUser(ctx context.Context, userID string) ([]SubmissionOutputRepo, error)
}

type SubmissionRepo struct {
	Database *sql.DB
}

func NewSubmissionRepo(db *sql.DB) ISubmissionRepo {
	return &SubmissionRepo{Database: db}
}
