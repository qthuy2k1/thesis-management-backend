package handler

import (
	"context"
	"log"

	attachmentpb "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/attachment-svc/internal/service"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateAttachment retrieves a attachment request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *AttachmentHdl) CreateAttachment(ctx context.Context, req *attachmentpb.CreateAttachmentRequest) (*attachmentpb.CreateAttachmentResponse, error) {
	log.Println("calling insert attachment...")
	att, err := validateAndConvertAttachment(req.Attachment)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	attRes, err := h.Service.CreateAttachment(ctx, att)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	submissionID := int64(*att.SubmissionID)
	exerciseID := int64(*att.ExerciseID)
	postID := int64(*att.PostID)

	resp := &attachmentpb.CreateAttachmentResponse{
		Response: &attachmentpb.CommonAttachmentResponse{
			StatusCode: 201,
			Message:    "Created",
		},
		AttachmentRes: &attachmentpb.AttachmentResponse{
			Id:           int64(attRes.ID),
			FileURL:      attRes.FileURL,
			Status:       attRes.Status,
			SubmissionID: &submissionID,
			ExerciseID:   &exerciseID,
			AuthorID:     attRes.AuthorID,
			CreatedAt:    timestamppb.New(attRes.CreatedAt),
			PostID:       &postID,
			Name:         attRes.Name,
			Type:         attRes.Type,
			Thumbnail:    attRes.Thumbnail,
			Size:         int64(attRes.Size),
		},
	}

	return resp, nil
}

// GetAttachment returns a attachment in db given by id
func (h *AttachmentHdl) GetAttachment(ctx context.Context, req *attachmentpb.GetAttachmentRequest) (*attachmentpb.GetAttachmentResponse, error) {
	log.Println("calling get attachment...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	att, err := h.Service.GetAttachment(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	submissionID := int64(*att.SubmissionID)
	exerciseID := int64(*att.ExerciseID)
	postID := int64(*att.PostID)

	attResp := attachmentpb.AttachmentResponse{
		Id:           int64(att.ID),
		FileURL:      att.FileURL,
		Status:       att.Status,
		SubmissionID: &submissionID,
		ExerciseID:   &exerciseID,
		AuthorID:     att.AuthorID,
		CreatedAt:    timestamppb.New(att.CreatedAt),
		PostID:       &postID,
		Name:         att.Name,
		Type:         att.Type,
		Thumbnail:    att.Thumbnail,
		Size:         int64(att.Size),
	}

	resp := &attachmentpb.GetAttachmentResponse{
		Response: &attachmentpb.CommonAttachmentResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Attachment: &attResp,
	}
	return resp, nil
}

func (c *AttachmentHdl) UpdateAttachment(ctx context.Context, req *attachmentpb.UpdateAttachmentRequest) (*attachmentpb.UpdateAttachmentResponse, error) {
	log.Println("calling update attachment...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	att, err := validateAndConvertAttachment(req.Attachment)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateAttachment(ctx, int(req.GetId()), service.AttachmentInputSvc{
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
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &attachmentpb.UpdateAttachmentResponse{
		Response: &attachmentpb.CommonAttachmentResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *AttachmentHdl) DeleteAttachment(ctx context.Context, req *attachmentpb.DeleteAttachmentRequest) (*attachmentpb.DeleteAttachmentResponse, error) {
	log.Println("calling delete attachment...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteAttachment(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &attachmentpb.DeleteAttachmentResponse{
		Response: &attachmentpb.CommonAttachmentResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *AttachmentHdl) GetAttachmentsOfExercise(ctx context.Context, req *attachmentpb.GetAttachmentsOfExerciseRequest) (*attachmentpb.GetAttachmentsOfExerciseResponse, error) {
	log.Println("calling get all attachments...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	atts, err := h.Service.GetAttachmentsOfExercise(ctx, int(req.GetExerciseID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var attsResp []*attachmentpb.AttachmentResponse
	for _, c := range atts {
		submissionID := int64(*c.SubmissionID)
		exerciseID := int64(*c.ExerciseID)
		postID := int64(*c.PostID)

		attsResp = append(attsResp, &attachmentpb.AttachmentResponse{
			Id:           int64(c.ID),
			FileURL:      c.FileURL,
			Status:       c.Status,
			SubmissionID: &submissionID,
			ExerciseID:   &exerciseID,
			AuthorID:     c.AuthorID,
			CreatedAt:    timestamppb.New(c.CreatedAt),
			PostID:       &postID,
			Name:         c.Name,
			Type:         c.Type,
			Thumbnail:    c.Thumbnail,
			Size:         int64(c.Size),
		})
	}

	return &attachmentpb.GetAttachmentsOfExerciseResponse{
		Response: &attachmentpb.CommonAttachmentResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Attachments: attsResp,
	}, nil
}

func (h *AttachmentHdl) GetAttachmentsOfSubmission(ctx context.Context, req *attachmentpb.GetAttachmentsOfSubmissionRequest) (*attachmentpb.GetAttachmentsOfSubmissionResponse, error) {
	log.Println("calling get all attachments...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	atts, err := h.Service.GetAttachmentsOfSubmission(ctx, int(req.GetSubmissionID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var attsResp []*attachmentpb.AttachmentResponse
	for _, c := range atts {
		submissionID := int64(*c.SubmissionID)
		exerciseID := int64(*c.ExerciseID)
		postID := int64(*c.PostID)

		attsResp = append(attsResp, &attachmentpb.AttachmentResponse{
			Id:           int64(c.ID),
			FileURL:      c.FileURL,
			Status:       c.Status,
			SubmissionID: &submissionID,
			ExerciseID:   &exerciseID,
			AuthorID:     c.AuthorID,
			CreatedAt:    timestamppb.New(c.CreatedAt),
			PostID:       &postID,
			Name:         c.Name,
			Type:         c.Type,
			Thumbnail:    c.Thumbnail,
			Size:         int64(c.Size),
		})
	}

	return &attachmentpb.GetAttachmentsOfSubmissionResponse{
		Response: &attachmentpb.CommonAttachmentResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Attachments: attsResp,
	}, nil
}

func (h *AttachmentHdl) GetAttachmentsOfPost(ctx context.Context, req *attachmentpb.GetAttachmentsOfPostRequest) (*attachmentpb.GetAttachmentsOfPostResponse, error) {
	log.Println("calling get all attachments...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	atts, err := h.Service.GetAttachmentsOfPost(ctx, int(req.GetPostID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var attsResp []*attachmentpb.AttachmentResponse
	for _, c := range atts {
		submissionID := int64(*c.SubmissionID)
		exerciseID := int64(*c.ExerciseID)
		postID := int64(*c.PostID)

		attsResp = append(attsResp, &attachmentpb.AttachmentResponse{
			Id:           int64(c.ID),
			FileURL:      c.FileURL,
			Status:       c.Status,
			SubmissionID: &submissionID,
			ExerciseID:   &exerciseID,
			AuthorID:     c.AuthorID,
			PostID:       &postID,
			CreatedAt:    timestamppb.New(c.CreatedAt),
			Name:         c.Name,
			Type:         c.Type,
			Thumbnail:    c.Thumbnail,
			Size:         int64(c.Size),
		})
	}

	return &attachmentpb.GetAttachmentsOfPostResponse{
		Response: &attachmentpb.CommonAttachmentResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Attachments: attsResp,
	}, nil
}

func (h *AttachmentHdl) GetFinalFile(ctx context.Context, req *attachmentpb.GetFinalFileRequest) (*attachmentpb.GetFinalFileResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	att, err := h.Service.GetFinalFile(ctx, req.AuthorID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	submissionID := int64(*att.SubmissionID)
	exerciseID := int64(*att.ExerciseID)
	postID := int64(*att.PostID)

	attResp := attachmentpb.AttachmentResponse{
		Id:           int64(att.ID),
		FileURL:      att.FileURL,
		Status:       att.Status,
		SubmissionID: &submissionID,
		ExerciseID:   &exerciseID,
		AuthorID:     att.AuthorID,
		CreatedAt:    timestamppb.New(att.CreatedAt),
		PostID:       &postID,
		Name:         att.Name,
		Type:         att.Type,
		Thumbnail:    att.Thumbnail,
		Size:         int64(att.Size),
	}

	resp := &attachmentpb.GetFinalFileResponse{
		Attachment: &attResp,
	}
	return resp, nil
}

func validateAndConvertAttachment(pbAttachment *attachmentpb.AttachmentInput) (service.AttachmentInputSvc, error) {
	if err := pbAttachment.Validate(); err != nil {
		return service.AttachmentInputSvc{}, err
	}

	var submissionID int
	if pbAttachment.SubmissionID != nil {
		submissionID = int(*pbAttachment.SubmissionID)
	}

	var exerciseID int
	if pbAttachment.ExerciseID != nil {
		exerciseID = int(*pbAttachment.ExerciseID)
	}

	var postID int
	if pbAttachment.PostID != nil {
		postID = int(*pbAttachment.PostID)
	}

	return service.AttachmentInputSvc{
		FileURL:      pbAttachment.FileURL,
		Status:       pbAttachment.Status,
		SubmissionID: &submissionID,
		ExerciseID:   &exerciseID,
		PostID:       &postID,
		AuthorID:     pbAttachment.AuthorID,
		Name:         pbAttachment.Name,
		Type:         pbAttachment.Type,
		Thumbnail:    pbAttachment.Thumbnail,
		Size:         int(pbAttachment.Size),
	}, nil
}
