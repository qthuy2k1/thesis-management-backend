package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	attachmentSvcV1 "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
	submissionSvcV1 "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type attachmentServiceGW struct {
	pb.UnimplementedAttachmentServiceServer
	attachmentClient attachmentSvcV1.AttachmentServiceClient
	userClient       userSvcV1.UserServiceClient
	submissionClient submissionSvcV1.SubmissionServiceClient
}

func NewAttachmentsService(attachmentClient attachmentSvcV1.AttachmentServiceClient, userClient userSvcV1.UserServiceClient, submissionClient submissionSvcV1.SubmissionServiceClient) *attachmentServiceGW {
	return &attachmentServiceGW{
		attachmentClient: attachmentClient,
		userClient:       userClient,
		submissionClient: submissionClient,
	}
}

func (u *attachmentServiceGW) CreateAttachment(ctx context.Context, req *pb.CreateAttachmentRequest) (*pb.CreateAttachmentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.attachmentClient.CreateAttachment(ctx, &attachmentSvcV1.CreateAttachmentRequest{
		Attachment: &attachmentSvcV1.AttachmentInput{
			FileURL:      req.GetAttachment().GetFileURL(),
			Status:       req.GetAttachment().GetStatus(),
			SubmissionID: req.GetAttachment().SubmissionID,
			ExerciseID:   req.GetAttachment().ExerciseID,
			PostID:       req.GetAttachment().PostID,
			AuthorID:     req.GetAttachment().GetAuthorID(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateAttachmentResponse{
		Response: &pb.CommonAttachmentResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
	}, nil
}

func (u *attachmentServiceGW) GetAttachment(ctx context.Context, req *pb.GetAttachmentRequest) (*pb.GetAttachmentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.attachmentClient.GetAttachment(ctx, &attachmentSvcV1.GetAttachmentRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	submissionRes, err := u.submissionClient.GetSubmission(ctx, &submissionSvcV1.GetSubmissionRequest{
		Id: res.GetAttachment().GetSubmissionID(),
	})
	if err != nil {
		return nil, err
	}

	authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: res.Attachment.AuthorID})
	if err != nil {
		return nil, err
	}

	return &pb.GetAttachmentResponse{
		Attachment: &pb.AttachmentResponse{
			Id:      res.GetAttachment().GetId(),
			FileURL: res.GetAttachment().GetFileURL(),
			Status:  res.GetAttachment().GetStatus(),
			Submission: &pb.SubmissionAttachmentResponse{
				Id:     submissionRes.GetSubmission().GetId(),
				Status: submissionRes.GetSubmission().GetStatus(),
			},
			Author: &pb.AuthorAttachmentResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt: res.GetAttachment().GetCreatedAt(),
		},
	}, nil
}

func (u *attachmentServiceGW) DeleteAttachment(ctx context.Context, req *pb.DeleteAttachmentRequest) (*pb.DeleteAttachmentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAttachmentResponse{
		Response: &pb.CommonAttachmentResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
	}, nil
}

func (u *attachmentServiceGW) GetAttachmentsOfExercise(ctx context.Context, req *pb.GetAttachmentsOfExerciseRequest) (*pb.GetAttachmentsOfExerciseResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.attachmentClient.GetAttachmentsOfExercise(ctx, &attachmentSvcV1.GetAttachmentsOfExerciseRequest{
		ExerciseID: req.GetExerciseID(),
	})
	if err != nil {
		return nil, err
	}

	var attachments []*pb.AttachmentResponse
	for _, c := range res.GetAttachments() {
		attachments = append(attachments, &pb.AttachmentResponse{
			Id:        c.Id,
			FileURL:   c.FileURL,
			Status:    c.Status,
			CreatedAt: c.CreatedAt,
		})
	}

	return &pb.GetAttachmentsOfExerciseResponse{
		Attachments: attachments,
	}, nil
}

// func (u *attachmentServiceGW) GetAttachmentsOfPost(ctx context.Context, req *pb.GetAttachmentsOfPostRequest) (*pb.GetAttachmentsOfExerciseResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	res, err := u.attachmentClient.GetAttachmentsOfPost(ctx, &attachmentSvcV1.GetAttachmentsOfPostRequest{
// 		PostID: req.GetPostID(),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var attachments []*pb.AttachmentResponse
// 	for _, c := range res.GetAttachments() {
// 		attachments = append(attachments, &pb.AttachmentResponse{
// 			Id:        c.Id,
// 			FileURL:   c.FileURL,
// 			Status:    c.Status,
// 			CreatedAt: c.CreatedAt,
// 		})
// 	}

// 	return &pb.GetAttachmentsOfExerciseResponse{
// 		Attachments: attachments,
// 	}, nil
// }
