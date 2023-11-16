package repository

import (
	"context"
	"database/sql"
)

type IAttachmentRepo interface {
	// CreateAttachment creates a new attachment in db given by attachment model
	CreateAttachment(ctx context.Context, clr AttachmentInputRepo) (AttachmentOutputRepo, error)
	// GetAttachment returns a attachment in db given by id
	GetAttachment(ctx context.Context, id int) (AttachmentOutputRepo, error)
	// UpdateAttachment updates the specified attachment by id
	UpdateAttachment(ctx context.Context, id int, attachment AttachmentInputRepo) error
	// DeleteAttachment deletes a attachment in db given by id
	DeleteAttachment(ctx context.Context, id int) error
	// GetAttachment returns a list of attachments of an exercise in db
	GetAttachmentsOfExercise(ctx context.Context, exerciseID int) ([]AttachmentOutputRepo, error)
	// GetAttachment returns a list of attachments of a submission in db
	GetAttachmentsOfSubmission(ctx context.Context, submissionID int) ([]AttachmentOutputRepo, error)
	// GetAttachment returns a list of attachments of a submission in db
	GetAttachmentsOfPost(ctx context.Context, postID int) ([]AttachmentOutputRepo, error)
}

type AttachmentRepo struct {
	Database *sql.DB
}

func NewAttachmentRepo(db *sql.DB) IAttachmentRepo {
	return &AttachmentRepo{Database: db}
}
