package service

import (
	"context"
	"log"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type submissionServiceGW struct {
	pb.UnimplementedSubmissionServiceServer
	submissionClient classroomSvcV1.SubmissionServiceClient
	classroomClient  classroomSvcV1.ClassroomServiceClient
	exerciseClient   classroomSvcV1.ExerciseServiceClient
	attachmentClient classroomSvcV1.AttachmentServiceClient
	userClient       userSvcV1.UserServiceClient
}

func NewSubmissionsService(submissionClient classroomSvcV1.SubmissionServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, exerciseClient classroomSvcV1.ExerciseServiceClient, attachmentClient classroomSvcV1.AttachmentServiceClient, userClient userSvcV1.UserServiceClient) *submissionServiceGW {
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

	exerciseExists, err := u.exerciseClient.GetExercise(ctx, &classroomSvcV1.GetExerciseRequest{Id: req.GetSubmission().GetExerciseID()})
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

	userExists, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.Submission.AuthorID})
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

	res, err := u.submissionClient.CreateSubmission(ctx, &classroomSvcV1.CreateSubmissionRequest{
		Submission: &classroomSvcV1.SubmissionInput{
			UserID:     req.GetSubmission().GetAuthorID(),
			ExerciseID: req.GetSubmission().GetExerciseID(),
			Status:     req.GetSubmission().GetStatus(),
		},
	})
	if err != nil {
		return nil, err
	}

	var attIDList []int64
	for _, attReq := range req.GetSubmission().GetAttachments() {
		exerciseID := req.GetSubmission().GetExerciseID()

		attRes, err := u.attachmentClient.CreateAttachment(ctx, &classroomSvcV1.CreateAttachmentRequest{
			Attachment: &classroomSvcV1.AttachmentInput{
				FileURL:      attReq.FileURL,
				Status:       "",
				SubmissionID: &res.SubmissionID,
				ExerciseID:   &exerciseID,
				AuthorID:     req.GetSubmission().GetAuthorID(),
				Name:         attReq.Name,
				Type:         attReq.Type,
				Thumbnail:    attReq.Thumbnail,
				Size:         attReq.Size,
			},
		})
		if err != nil {
			// Rollback transaction
			for _, attID := range attIDList {
				_, err := u.attachmentClient.DeleteAttachment(ctx, &classroomSvcV1.DeleteAttachmentRequest{
					Id: attID,
				})
				if err != nil {
					return nil, err
				}
			}
			return nil, err
		}

		attIDList = append(attIDList, attRes.GetAttachmentRes().GetId())
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

	res, err := u.submissionClient.GetSubmission(ctx, &classroomSvcV1.GetSubmissionRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	attRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &classroomSvcV1.GetAttachmentsOfSubmissionRequest{
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
			Id:          res.GetSubmission().GetId(),
			AuthorID:    res.GetSubmission().GetUserID(),
			ExerciseID:  res.GetSubmission().GetExerciseID(),
			Status:      res.GetSubmission().GetStatus(),
			CreatedAt:   res.GetSubmission().CreatedAt,
			UpdatedAt:   res.Submission.UpdatedAt,
			Attachments: attachments,
		},
	}, nil
}

func (u *submissionServiceGW) UpdateSubmission(ctx context.Context, req *pb.UpdateSubmissionRequest) (*pb.UpdateSubmissionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	log.Println(req)
	exists, err := u.exerciseClient.GetExercise(ctx, &classroomSvcV1.GetExerciseRequest{Id: req.GetSubmission().GetExerciseID()})
	if err != nil {
		return nil, err
	}

	if exists.GetResponse().StatusCode == 404 {
		return &pb.UpdateSubmissionResponse{
			Response: &pb.CommonSubmissionResponse{
				StatusCode: 404,
				Message:    "Exercise does not exist",
			},
		}, nil
	}

	res, err := u.submissionClient.UpdateSubmission(ctx, &classroomSvcV1.UpdateSubmissionRequest{
		Id: req.GetId(),
		Submission: &classroomSvcV1.SubmissionInput{
			UserID:     req.GetSubmission().GetAuthorID(),
			ExerciseID: req.GetSubmission().GetExerciseID(),
			Status:     req.GetSubmission().GetStatus(),
		},
	})
	if err != nil {
		return nil, err
	}

	attGetRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &classroomSvcV1.GetAttachmentsOfSubmissionRequest{
		SubmissionID: req.Id,
	})
	if err != nil {
		return nil, err
	}

	// delete old attachments
	for _, a := range attGetRes.Attachments {
		if _, err := u.attachmentClient.DeleteAttachment(ctx, &classroomSvcV1.DeleteAttachmentRequest{
			Id: a.Id,
		}); err != nil {
			return nil, err
		}
	}

	// create new attachments
	var attCreated []int64
	if len(req.Submission.GetAttachments()) > 0 {
		for _, att := range req.Submission.Attachments {
			attRes, err := u.attachmentClient.CreateAttachment(ctx, &classroomSvcV1.CreateAttachmentRequest{
				Attachment: &classroomSvcV1.AttachmentInput{
					FileURL:      att.FileURL,
					ExerciseID:   &req.Submission.ExerciseID,
					AuthorID:     req.Submission.AuthorID,
					Name:         att.Name,
					Status:       "",
					Type:         att.Type,
					Thumbnail:    att.Thumbnail,
					Size:         att.Size,
					SubmissionID: &req.Id,
				},
			})
			if err != nil {
				if len(attCreated) > 0 {
					for _, aErr := range attCreated {
						if _, err := u.attachmentClient.DeleteAttachment(ctx, &classroomSvcV1.DeleteAttachmentRequest{
							Id: aErr,
						}); err != nil {
							return nil, err
						}
					}
				}
				return nil, err
			}

			attCreated = append(attCreated, attRes.AttachmentRes.Id)
		}
	}

	return &pb.UpdateSubmissionResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *submissionServiceGW) DeleteSubmission(ctx context.Context, req *pb.DeleteSubmissionRequest) (*pb.DeleteSubmissionResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.submissionClient.DeleteSubmission(ctx, &classroomSvcV1.DeleteSubmissionRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	attGetRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &classroomSvcV1.GetAttachmentsOfSubmissionRequest{
		SubmissionID: req.Id,
	})
	if err != nil {
		return nil, err
	}

	for _, a := range attGetRes.Attachments {
		if _, err := u.attachmentClient.DeleteAttachment(ctx, &classroomSvcV1.DeleteAttachmentRequest{
			Id: a.Id,
		}); err != nil {
			return nil, err
		}
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

	res, err := u.submissionClient.GetAllSubmissionsOfExercise(ctx, &classroomSvcV1.GetAllSubmissionsOfExerciseRequest{
		ExerciseID: req.GetExerciseID(),
	})
	if err != nil {
		return nil, err
	}

	var submissions []*pb.SubmissionResponse
	for _, p := range res.GetSubmissions() {
		attRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &classroomSvcV1.GetAttachmentsOfSubmissionRequest{
			SubmissionID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentSubmissionResponse
		for _, a := range attRes.GetAttachments() {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentSubmissionResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorSubmissionResponse{
					Id:       author.GetUser().GetId(),
					Class:    author.GetUser().Class,
					Major:    author.GetUser().Major,
					Phone:    author.GetUser().Phone,
					PhotoSrc: author.GetUser().GetPhotoSrc(),
					Role:     author.GetUser().GetRole(),
					Name:     author.GetUser().GetName(),
					Email:    author.GetUser().GetEmail(),
					// HashedPassword: author.GetUser().HashedPassword,
				},
				CreatedAt: a.CreatedAt,
				Size:      a.Size,
				MimeType:  a.Type,
				Thumbnail: a.Thumbnail,
				FileName:  a.Name,
			})
		}

		submissions = append(submissions, &pb.SubmissionResponse{
			Id:          p.Id,
			AuthorID:    p.UserID,
			ExerciseID:  p.ExerciseID,
			Status:      p.Status,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Attachments: attachments,
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
func (u *submissionServiceGW) GetSubmissionOfUser(ctx context.Context, req *pb.GetSubmissionOfUserRequest) (*pb.GetSubmissionOfUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.submissionClient.GetSubmissionOfUser(ctx, &classroomSvcV1.GetSubmissionOfUserRequest{
		UserID:     req.UserID,
		ExerciseID: req.ExerciseID,
	})
	if err != nil {
		return nil, err
	}

	var submissions []*pb.SubmissionResponse
	for _, p := range res.GetSubmissions() {
		attRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &classroomSvcV1.GetAttachmentsOfSubmissionRequest{
			SubmissionID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentSubmissionResponse
		for _, a := range attRes.GetAttachments() {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentSubmissionResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorSubmissionResponse{
					Id:       author.GetUser().GetId(),
					Class:    author.GetUser().Class,
					Major:    author.GetUser().Major,
					Phone:    author.GetUser().Phone,
					PhotoSrc: author.GetUser().GetPhotoSrc(),
					Role:     author.GetUser().GetRole(),
					Name:     author.GetUser().GetName(),
					Email:    author.GetUser().GetEmail(),
					// HashedPassword: author.GetUser().HashedPassword,
				},
				CreatedAt: a.CreatedAt,
				Size:      a.Size,
				MimeType:  a.Type,
				Thumbnail: a.Thumbnail,
				FileName:  a.Name,
			})
		}

		submissions = append(submissions, &pb.SubmissionResponse{
			Id:          p.Id,
			AuthorID:    p.UserID,
			ExerciseID:  p.ExerciseID,
			Status:      p.Status,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Attachments: attachments,
		})
	}

	return &pb.GetSubmissionOfUserResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Submissions: submissions,
	}, nil
}

func (u *submissionServiceGW) GetSubmissionFromUser(ctx context.Context, req *pb.GetSubmissionFromUserRequest) (*pb.GetSubmissionFromUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.submissionClient.GetSubmissionFromUser(ctx, &classroomSvcV1.GetSubmissionFromUserRequest{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	var submissions []*pb.SubmissionResponse
	for _, p := range res.GetSubmissions() {
		attRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &classroomSvcV1.GetAttachmentsOfSubmissionRequest{
			SubmissionID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentSubmissionResponse
		for _, a := range attRes.GetAttachments() {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentSubmissionResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorSubmissionResponse{
					Id:       author.GetUser().GetId(),
					Class:    author.GetUser().Class,
					Major:    author.GetUser().Major,
					Phone:    author.GetUser().Phone,
					PhotoSrc: author.GetUser().GetPhotoSrc(),
					Role:     author.GetUser().GetRole(),
					Name:     author.GetUser().GetName(),
					Email:    author.GetUser().GetEmail(),
					// HashedPassword: author.GetUser().HashedPassword,
				},
				CreatedAt: a.CreatedAt,
				Size:      a.Size,
				MimeType:  a.Type,
				Thumbnail: a.Thumbnail,
				FileName:  a.Name,
			})
		}

		submissions = append(submissions, &pb.SubmissionResponse{
			Id:          p.Id,
			AuthorID:    p.UserID,
			ExerciseID:  p.ExerciseID,
			Status:      p.Status,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Attachments: attachments,
		})
	}

	return &pb.GetSubmissionFromUserResponse{
		Response: &pb.CommonSubmissionResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Submissions: submissions,
	}, nil
}
