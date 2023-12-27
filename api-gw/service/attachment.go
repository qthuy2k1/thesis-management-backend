package service

import (
	"context"
	"log"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type attachmentServiceGW struct {
	pb.UnimplementedAttachmentServiceServer
	classroomClient classroomSvcV1.AttachmentServiceClient
	userClient      userSvcV1.UserServiceClient
}

func NewAttachmentsService(classroomClient classroomSvcV1.AttachmentServiceClient, userClient userSvcV1.UserServiceClient, submissionClient classroomSvcV1.SubmissionServiceClient) *attachmentServiceGW {
	return &attachmentServiceGW{
		classroomClient: classroomClient,
		userClient:      userClient,
	}
}

func (u *attachmentServiceGW) CreateAttachment(ctx context.Context, req *pb.CreateAttachmentRequest) (*pb.CreateAttachmentResponse, error) {
	log.Println(req)
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.classroomClient.CreateAttachment(ctx, &classroomSvcV1.CreateAttachmentRequest{
		Attachment: &classroomSvcV1.AttachmentInput{
			FileURL:   req.GetAttachment().GetFileURL(),
			Status:    req.GetAttachment().GetStatus(),
			AuthorID:  req.GetAttachment().GetAuthorID(),
			Name:      req.GetAttachment().GetName(),
			Type:      req.GetAttachment().GetType(),
			Thumbnail: req.GetAttachment().GetThumbnail(),
			Size:      req.GetAttachment().GetSize(),
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

	res, err := u.classroomClient.GetAttachment(ctx, &classroomSvcV1.GetAttachmentRequest{Id: req.GetId()})
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

	res, err := u.classroomClient.DeleteAttachment(ctx, &classroomSvcV1.DeleteAttachmentRequest{
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

	res, err := u.classroomClient.GetAttachmentsOfExercise(ctx, &classroomSvcV1.GetAttachmentsOfExerciseRequest{
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

// 	res, err := u.classroomClient.GetAttachmentsOfPost(ctx, &classroomSvcV1.GetAttachmentsOfPostRequest{
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
