package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/attachment-svc/internal/repository"
)

type IAttachmentSvc interface {
	// CreateClasroom creates a new attachment in db given by attachment model
	CreateAttachment(ctx context.Context, clr AttachmentInputSvc) error
	// GetAttachment returns a attachment in db given by id
	GetAttachment(ctx context.Context, id int) (AttachmentOutputSvc, error)
	// UpdateAttachment updates the specified attachment by id
	UpdateAttachment(ctx context.Context, id int, attachment AttachmentInputSvc) error
	// DeleteAttachment deletes a attachment in db given by id
	DeleteAttachment(ctx context.Context, id int) error
	// GetAttachmentOfExercise returns a list of attachments of a exercise in db
	GetAttachmentsOfExercise(ctx context.Context, exerciseID int) ([]AttachmentOutputSvc, error)
}

type AttachmentSvc struct {
	Repository repository.IAttachmentRepo
}

func NewAttachmentSvc(clrRepo repository.IAttachmentRepo) IAttachmentSvc {
	return &AttachmentSvc{Repository: clrRepo}
}
