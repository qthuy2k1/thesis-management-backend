package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/attachment-svc/internal/repository"
)

type AttachmentInputSvc struct {
	FileURL      string
	Status       string
	SubmissionID *int
	ExerciseID   *int
	PostID       *int
	AuthorID     string
	Name         string
	Type         string
	Thumbnail    string
	Size         int
}

// CreateClasroom creates a new attachment in db given by attachment model
func (s *AttachmentSvc) CreateAttachment(ctx context.Context, att AttachmentInputSvc) (AttachmentOutputSvc, error) {
	attRepo := repository.AttachmentInputRepo{
		FileURL:      att.FileURL,
		Status:       att.Status,
		SubmissionID: att.SubmissionID,
		ExerciseID:   att.ExerciseID,
		AuthorID:     att.AuthorID,
		PostID:       att.PostID,
		Name:         att.Name,
		Type:         att.Type,
		Thumbnail:    att.Thumbnail,
		Size:         att.Size,
	}
	attRes, err := s.Repository.CreateAttachment(ctx, attRepo)
	if err != nil {
		if errors.Is(err, repository.ErrAttachmentExisted) {
			return AttachmentOutputSvc{}, ErrAttachmentExisted
		}
		return AttachmentOutputSvc{}, err
	}

	return AttachmentOutputSvc{
		ID:           attRes.ID,
		FileURL:      attRes.FileURL,
		Status:       attRes.Status,
		SubmissionID: attRes.SubmissionID,
		ExerciseID:   attRes.ExerciseID,
		AuthorID:     attRes.AuthorID,
		PostID:       att.PostID,
		CreatedAt:    attRes.CreatedAt,
		Name:         attRes.Name,
		Type:         attRes.Type,
		Thumbnail:    attRes.Thumbnail,
		Size:         attRes.Size,
	}, nil
}

// GetAttachment returns a attachment in db given by id
func (s *AttachmentSvc) GetAttachment(ctx context.Context, id int) (AttachmentOutputSvc, error) {
	att, err := s.Repository.GetAttachment(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrAttachmentNotFound) {
			return AttachmentOutputSvc{}, ErrAttachmentNotFound
		}
		return AttachmentOutputSvc{}, err
	}

	return AttachmentOutputSvc{
		ID:           att.ID,
		FileURL:      att.FileURL,
		Status:       att.Status,
		SubmissionID: att.SubmissionID,
		ExerciseID:   att.ExerciseID,
		AuthorID:     att.AuthorID,
		PostID:       att.PostID,
		CreatedAt:    att.CreatedAt,
		Name:         att.Name,
		Type:         att.Type,
		Thumbnail:    att.Thumbnail,
		Size:         att.Size,
	}, nil
}

// UpdateAttachment updates the specified attachment by id
func (s *AttachmentSvc) UpdateAttachment(ctx context.Context, id int, attachment AttachmentInputSvc) error {
	if err := s.Repository.UpdateAttachment(ctx, id, repository.AttachmentInputRepo{
		FileURL:      attachment.FileURL,
		Status:       attachment.Status,
		SubmissionID: attachment.SubmissionID,
		ExerciseID:   attachment.ExerciseID,
		AuthorID:     attachment.AuthorID,
		PostID:       attachment.PostID,
		Type:         attachment.Type,
		Thumbnail:    attachment.Thumbnail,
		Size:         attachment.Size,
	}); err != nil {
		if errors.Is(err, repository.ErrAttachmentNotFound) {
			return ErrAttachmentNotFound
		}
		return err
	}

	return nil
}

// DeleteAttachment deletes a attachment in db given by id
func (s *AttachmentSvc) DeleteAttachment(ctx context.Context, id int) error {
	if err := s.Repository.DeleteAttachment(ctx, id); err != nil {
		if errors.Is(err, repository.ErrAttachmentNotFound) {
			return ErrAttachmentNotFound
		}
		return err
	}

	return nil
}

type AttachmentOutputSvc struct {
	ID           int
	FileURL      string
	Status       string
	SubmissionID *int
	ExerciseID   *int
	PostID       *int
	AuthorID     string
	CreatedAt    time.Time
	Name         string
	Type         string
	Thumbnail    string
	Size         int
}

// GetAttachment returns a list of attachments in db with filter
func (s *AttachmentSvc) GetAttachmentsOfExercise(ctx context.Context, exerciseID int) ([]AttachmentOutputSvc, error) {
	attsRepo, err := s.Repository.GetAttachmentsOfExercise(ctx, exerciseID)
	if err != nil {
		return nil, err
	}

	var attsSvc []AttachmentOutputSvc
	for _, c := range attsRepo {
		attsSvc = append(attsSvc, AttachmentOutputSvc{
			ID:           c.ID,
			FileURL:      c.FileURL,
			Status:       c.Status,
			SubmissionID: c.SubmissionID,
			ExerciseID:   c.ExerciseID,
			AuthorID:     c.AuthorID,
			PostID:       c.PostID,
			CreatedAt:    c.CreatedAt,
			Name:         c.Name,
			Type:         c.Type,
			Thumbnail:    c.Thumbnail,
			Size:         c.Size,
		})
	}

	return attsSvc, nil
}

// GetAttachment returns a list of attachments in db with filter
func (s *AttachmentSvc) GetAttachmentsOfSubmission(ctx context.Context, submissionID int) ([]AttachmentOutputSvc, error) {
	attsRepo, err := s.Repository.GetAttachmentsOfSubmission(ctx, submissionID)
	if err != nil {
		return nil, err
	}

	var attsSvc []AttachmentOutputSvc
	for _, c := range attsRepo {
		attsSvc = append(attsSvc, AttachmentOutputSvc{
			ID:           c.ID,
			FileURL:      c.FileURL,
			Status:       c.Status,
			SubmissionID: c.SubmissionID,
			ExerciseID:   c.ExerciseID,
			AuthorID:     c.AuthorID,
			PostID:       c.PostID,
			CreatedAt:    c.CreatedAt,
			Name:         c.Name,
			Type:         c.Type,
			Thumbnail:    c.Thumbnail,
			Size:         c.Size,
		})
	}

	return attsSvc, nil
}

// GetAttachment returns a list of attachments in db with filter
func (s *AttachmentSvc) GetAttachmentsOfPost(ctx context.Context, postID int) ([]AttachmentOutputSvc, error) {
	attsRepo, err := s.Repository.GetAttachmentsOfPost(ctx, postID)
	if err != nil {
		return nil, err
	}

	var attsSvc []AttachmentOutputSvc
	for _, c := range attsRepo {
		attsSvc = append(attsSvc, AttachmentOutputSvc{
			ID:           c.ID,
			FileURL:      c.FileURL,
			Status:       c.Status,
			SubmissionID: c.SubmissionID,
			ExerciseID:   c.ExerciseID,
			AuthorID:     c.AuthorID,
			PostID:       c.PostID,
			CreatedAt:    c.CreatedAt,
			Name:         c.Name,
			Type:         c.Type,
			Thumbnail:    c.Thumbnail,
			Size:         c.Size,
		})
	}

	return attsSvc, nil
}
