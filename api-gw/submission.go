package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	attachmentSvcV1 "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	submissionSvcV1 "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type submissionServiceGW struct {
	pb.UnimplementedSubmissionServiceServer
	submissionClient submissionSvcV1.SubmissionServiceClient
	classroomClient  classroomSvcV1.ClassroomServiceClient
	exerciseClient   exerciseSvcV1.ExerciseServiceClient
	attachmentClient attachmentSvcV1.AttachmentServiceClient
	userClient       userSvcV1.UserServiceClient
}

func NewSubmissionsService(submissionClient submissionSvcV1.SubmissionServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, exerciseClient exerciseSvcV1.ExerciseServiceClient, attachmentClient attachmentSvcV1.AttachmentServiceClient, userClient userSvcV1.UserServiceClient) *submissionServiceGW {
	return &submissionServiceGW{
		submissionClient: submissionClient,
		classroomClient:  classroomClient,
		exerciseClient:   exerciseClient,
		attachmentClient: attachmentClient,
		userClient:       userClient,
	}
}

func (u *submissionServiceGW) CreateSubmission(ctx context.Context, req *pb.CreateSubmissionRequest) (*pb.CreateSubmissionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	exerciseExists, err := u.exerciseClient.GetExercise(ctx, &exerciseSvcV1.GetExerciseRequest{Id: req.GetSubmission().GetExerciseID()})
	if err != nil {
		return nil, err
	}

	if exerciseExists.GetResponse().StatusCode == 404 {
		return &pb.CreateSubmissionResponse{
			Response: &pb.CommonSubmissionResponse{
				StatusCode: 404,
				Message:    "Exercise does not exist",
			},
		}, nil
	}

	userExists, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.GetSubmission().GetUserID()})
	if err != nil {
		return nil, err
	}

	if userExists.GetResponse().StatusCode == 404 {
		return &pb.CreateSubmissionResponse{
			Response: &pb.CommonSubmissionResponse{
				StatusCode: 404,
				Message:    "User does not exist",
			},
		}, nil
	}

	var attResList []*attachmentSvcV1.AttachmentResponse
	var attIDList []int64
	for _, attReq := range req.GetSubmission().GetAttachments() {
		exerciseID := req.GetSubmission().GetExerciseID()

		attRes, err := u.attachmentClient.CreateAttachment(ctx, &attachmentSvcV1.CreateAttachmentRequest{
			Attachment: &attachmentSvcV1.AttachmentInput{
				FileURL:      attReq.FileURL,
				Status:       attReq.Status,
				SubmissionID: nil,
				ExerciseID:   &exerciseID,
				AuthorID:     req.GetSubmission().GetUserID(),
			},
		})
		if err != nil {
			return nil, err
		}

		attResList = append(attResList, attRes.GetAttachmentRes())
		attIDList = append(attIDList, attRes.GetAttachmentRes().GetId())
	}

	res, err := u.submissionClient.CreateSubmission(ctx, &submissionSvcV1.CreateSubmissionRequest{
		Submission: &submissionSvcV1.SubmissionInput{
			UserID:         req.GetSubmission().GetUserID(),
			ExerciseID:     req.GetSubmission().GetExerciseID(),
			SubmissionDate: req.GetSubmission().GetSubmissionDate(),
			Status:         req.GetSubmission().GetStatus(),
			AttachmentID:   attIDList,
		},
	})
	if err != nil {
		// Rollback transaction
		for _, attID := range attIDList {
			_, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
				Id: attID,
			})
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	// Update the submission ID for attachment
	for _, att := range attResList {
		submissionID := res.GetSubmissionID()

		if _, err = u.attachmentClient.UpdateAttachment(ctx, &attachmentSvcV1.UpdateAttachmentRequest{
			Id: att.Id,
			Attachment: &attachmentSvcV1.AttachmentInput{
				FileURL:      att.FileURL,
				Status:       att.Status,
				SubmissionID: &submissionID,
				ExerciseID:   att.ExerciseID,
				AuthorID:     att.AuthorID,
			},
		}); err != nil {
			return nil, err
		}
	}

	return &pb.CreateSubmissionResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *submissionServiceGW) GetSubmission(ctx context.Context, req *pb.GetSubmissionRequest) (*pb.GetSubmissionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.submissionClient.GetSubmission(ctx, &submissionSvcV1.GetSubmissionRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	attRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &attachmentSvcV1.GetAttachmentsOfSubmissionRequest{
		SubmissionID: res.GetSubmission().GetId(),
	})
	if err != nil {
		return nil, err
	}

	var attachments []*pb.AttachmentSubmissionResponse
	for _, a := range attRes.GetAttachments() {
		attachments = append(attachments, &pb.AttachmentSubmissionResponse{
			Id:      a.Id,
			FileURL: a.FileURL,
			Status:  a.Status,
		})
	}

	return &pb.GetSubmissionResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Submission: &pb.SubmissionResponse{
			Id:             res.GetSubmission().GetId(),
			UserID:         res.GetSubmission().GetUserID(),
			ExerciseID:     res.GetSubmission().GetExerciseID(),
			SubmissionDate: res.GetSubmission().GetSubmissionDate(),
			Status:         res.GetSubmission().GetStatus(),
			Attachments:    attachments,
		},
	}, nil
}

// func (u *submissionServiceGW) UpdateSubmission(ctx context.Context, req *pb.UpdateSubmissionRequest) (*pb.UpdateSubmissionResponse, error) {
// 	exists, err := u.exerciseClient.GetExercise(ctx, &exerciseSvcV1.GetExerciseRequest{Id: req.GetSubmission().GetExerciseID()})
// 	if err != nil {
// 		return nil, err
// 	}

// 	if exists.GetResponse().StatusCode == 404 {
// 		return &pb.UpdateSubmissionResponse{
// 			Response: &pb.CommonSubmissionResponse{
// 				StatusCode: 404,
// 				Message:    "Exercise does not exist",
// 			},
// 		}, nil
// 	}

// 	res, err := u.submissionClient.UpdateSubmission(ctx, &submissionSvcV1.UpdateSubmissionRequest{
// 		Id: req.GetId(),
// 		Submission: &submissionSvcV1.SubmissionInput{
// 			UserID:         req.GetSubmission().GetUserID(),
// 			ExerciseID:     req.GetSubmission().GetExerciseID(),
// 			SubmissionDate: req.GetSubmission().GetSubmissionDate(),
// 			Status:         req.GetSubmission().GetStatus(),
// 			AttachmentID:   req.GetSubmission().GetAttachmentID(),
// 		},
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.UpdateSubmissionResponse{
// 		Response: &pb.CommonSubmissionResponse{
// 			StatusCode: res.GetResponse().StatusCode,
// 			Message:    res.GetResponse().Message,
// 		},
// 	}, nil
// }

func (u *submissionServiceGW) DeleteSubmission(ctx context.Context, req *pb.DeleteSubmissionRequest) (*pb.DeleteSubmissionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.submissionClient.DeleteSubmission(ctx, &submissionSvcV1.DeleteSubmissionRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteSubmissionResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *submissionServiceGW) GetAllSubmissionsOfExercise(ctx context.Context, req *pb.GetAllSubmissionsOfExerciseRequest) (*pb.GetAllSubmissionsOfExerciseResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.submissionClient.GetAllSubmissionsOfExercise(ctx, &submissionSvcV1.GetAllSubmissionsOfExerciseRequest{
		ExerciseID: req.GetExerciseID(),
	})
	if err != nil {
		return nil, err
	}

	var submissions []*pb.SubmissionResponse
	for _, p := range res.GetSubmissions() {
		attRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &attachmentSvcV1.GetAttachmentsOfSubmissionRequest{
			SubmissionID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentSubmissionResponse
		for _, a := range attRes.GetAttachments() {
			attachments = append(attachments, &pb.AttachmentSubmissionResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
			})
		}

		submissions = append(submissions, &pb.SubmissionResponse{
			Id:             p.Id,
			UserID:         p.UserID,
			ExerciseID:     p.ExerciseID,
			SubmissionDate: p.SubmissionDate,
			Status:         p.Status,
			Attachments:    attachments,
		})
	}

	return &pb.GetAllSubmissionsOfExerciseResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount:  res.GetTotalCount(),
		Submissions: submissions,
	}, nil
}
