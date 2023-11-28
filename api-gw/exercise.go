package main

import (
	"context"
	"log"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	attachmentSvcV1 "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	commentSvcV1 "github.com/qthuy2k1/thesis-management-backend/comment-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	rpsSvcV1 "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	scheduleSvcV1 "github.com/qthuy2k1/thesis-management-backend/schedule-svc/api/goclient/v1"
	submissionSvcV1 "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type exerciseServiceGW struct {
	pb.UnimplementedExerciseServiceServer
	exerciseClient       exerciseSvcV1.ExerciseServiceClient
	classroomClient      classroomSvcV1.ClassroomServiceClient
	reportingStageClient rpsSvcV1.ReportingStageServiceClient
	commentClient        commentSvcV1.CommentServiceClient
	userClient           userSvcV1.UserServiceClient
	submissionClient     submissionSvcV1.SubmissionServiceClient
	attachmentClient     attachmentSvcV1.AttachmentServiceClient
	scheduleClient       scheduleSvcV1.ScheduleServiceClient
}

func NewExercisesService(exerciseClient exerciseSvcV1.ExerciseServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, reportStageClient rpsSvcV1.ReportingStageServiceClient, commentClient commentSvcV1.CommentServiceClient, userClient userSvcV1.UserServiceClient, submissionClient submissionSvcV1.SubmissionServiceClient, attachmentClient attachmentSvcV1.AttachmentServiceClient, scheduleClient scheduleSvcV1.ScheduleServiceClient) *exerciseServiceGW {
	return &exerciseServiceGW{
		exerciseClient:       exerciseClient,
		classroomClient:      classroomClient,
		reportingStageClient: reportStageClient,
		commentClient:        commentClient,
		userClient:           userClient,
		submissionClient:     submissionClient,
		attachmentClient:     attachmentClient,
		scheduleClient:       scheduleClient,
	}
}

func (u *exerciseServiceGW) CreateExercise(ctx context.Context, req *pb.CreateExerciseRequest) (*pb.CreateExerciseResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	log.Println("reqqqqqqqqqqqqqqqq", req.GetExercise())
	log.Println("======================================================")

	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetExercise().ClassroomID})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.CreateExerciseResponse{
			Response: &pb.CommonExerciseResponse{
				StatusCode: 404,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{Id: req.GetExercise().GetCategoryID()})
	if err != nil {
		return nil, err
	}

	if rpsRes.GetResponse().GetStatusCode() == 404 {
		return &pb.CreateExerciseResponse{
			Response: &pb.CommonExerciseResponse{
				StatusCode: 404,
				Message:    "Reporting stage does not exist",
			},
		}, nil
	}

	_, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.Exercise.AuthorID,
	})
	if err != nil {
		return nil, err
	}

	res, err := u.exerciseClient.CreateExercise(ctx, &exerciseSvcV1.CreateExerciseRequest{
		Exercise: &exerciseSvcV1.ExerciseInput{
			Title:            req.GetExercise().Title,
			Content:          req.GetExercise().Description,
			ClassroomID:      req.GetExercise().ClassroomID,
			Deadline:         req.GetExercise().Deadline,
			ReportingStageID: req.GetExercise().CategoryID,
			AuthorID:         req.GetExercise().AuthorID,
		},
	})
	if err != nil {
		return nil, err
	}

	var attCreated []int64
	if len(req.Exercise.GetAttachments()) > 0 {
		for _, att := range req.Exercise.Attachments {
			attRes, err := u.attachmentClient.CreateAttachment(ctx, &attachmentSvcV1.CreateAttachmentRequest{
				Attachment: &attachmentSvcV1.AttachmentInput{
					FileURL:    att.FileURL,
					ExerciseID: &res.ExerciseID,
					AuthorID:   req.Exercise.AuthorID,
					Name:       att.Name,
					Status:     "",
					Type:       att.Type,
					Thumbnail:  att.Thumbnail,
					Size:       att.Size,
				},
			})
			if err != nil {
				if len(attCreated) > 0 {
					for _, aErr := range attCreated {
						if _, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
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

	return &pb.CreateExerciseResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *exerciseServiceGW) GetExercise(ctx context.Context, req *pb.GetExerciseRequest) (*pb.GetExerciseResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.exerciseClient.GetExercise(ctx, &exerciseSvcV1.GetExerciseRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	commentRes, err := u.commentClient.GetCommentsOfAExercise(ctx, &commentSvcV1.GetCommentsOfAExerciseRequest{
		ExerciseID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	var comments []*pb.CommentExerciseResponse
	if len(commentRes.GetComments()) > 0 {
		for _, c := range commentRes.GetComments() {
			userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: c.UserID,
			})
			if err != nil {
				return nil, err
			}

			comments = append(comments, &pb.CommentExerciseResponse{
				Id: c.Id,
				User: &pb.AuthorExerciseResponse{
					Id:       userRes.User.Id,
					Class:    userRes.User.Class,
					Major:    userRes.User.Major,
					Phone:    userRes.User.Phone,
					PhotoSrc: userRes.User.PhotoSrc,
					Role:     userRes.User.Role,
					Name:     userRes.User.Name,
					Email:    userRes.User.Email,
				},
				ExerciseID: *c.ExerciseID,
				Content:    c.Content,
				CreatedAt:  c.CreatedAt,
			})
		}
	}

	reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{
		Id: res.Exercise.ReportingStageID,
	})
	if err != nil {
		return nil, err
	}

	authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: res.Exercise.AuthorID,
	})
	if err != nil {
		return nil, err
	}

	attachment, err := u.attachmentClient.GetAttachmentsOfExercise(ctx, &attachmentSvcV1.GetAttachmentsOfExerciseRequest{
		ExerciseID: res.Exercise.Id,
	})
	if err != nil {
		return nil, err
	}

	var attachments []*pb.AttachmentExerciseResponse
	for _, a := range attachment.GetAttachments() {
		author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: a.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, &pb.AttachmentExerciseResponse{
			Id:      a.Id,
			FileURL: a.FileURL,
			Status:  a.Status,
			Author: &pb.AuthorExerciseResponse{
				Id:       author.User.Id,
				Class:    author.User.Class,
				Major:    author.User.Major,
				Phone:    author.User.Phone,
				PhotoSrc: author.User.PhotoSrc,
				Role:     author.User.Role,
				Name:     author.User.Name,
				Email:    author.User.Email,
			},
			CreatedAt: a.CreatedAt,
			Size:      a.Size,
			MimeType:  a.Type,
			Thumbnail: a.Thumbnail,
			FileName:  a.Name,
		})
	}

	return &pb.GetExerciseResponse{
		Exercise: &pb.ExerciseResponse{
			Id:          res.GetExercise().Id,
			Title:       res.GetExercise().Title,
			Description: res.GetExercise().Content,
			ClassroomID: res.GetExercise().ClassroomID,
			Deadline:    res.GetExercise().Deadline,
			Category: &pb.ReportingStageExerciseResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Label:       reportingStageRes.ReportingStage.Label,
				Description: reportingStageRes.ReportingStage.Description,
				Value:       reportingStageRes.ReportingStage.Value,
			},
			Author: &pb.AuthorExerciseResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt:   res.GetExercise().CreatedAt,
			UpdatedAt:   res.GetExercise().UpdatedAt,
			Attachments: attachments,
		},
		Comments: comments,
	}, nil
}

func (u *exerciseServiceGW) UpdateExercise(ctx context.Context, req *pb.UpdateExerciseRequest) (*pb.UpdateExerciseResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{Id: req.GetExercise().GetCategoryID()})
	if err != nil {
		return nil, err
	}

	if rpsRes.GetResponse().GetStatusCode() == 404 {
		return &pb.UpdateExerciseResponse{
			Response: &pb.CommonExerciseResponse{
				StatusCode: 404,
				Message:    "Reporting stage does not exist",
			},
		}, nil
	}

	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetExercise().ClassroomID})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.UpdateExerciseResponse{
			Response: &pb.CommonExerciseResponse{
				StatusCode: 404,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	res, err := u.exerciseClient.UpdateExercise(ctx, &exerciseSvcV1.UpdateExerciseRequest{
		Id: req.GetId(),
		Exercise: &exerciseSvcV1.ExerciseInput{
			Title:            req.GetExercise().Title,
			Content:          req.GetExercise().Description,
			ClassroomID:      req.GetExercise().ClassroomID,
			Deadline:         req.GetExercise().Deadline,
			ReportingStageID: req.GetExercise().CategoryID,
			AuthorID:         req.GetExercise().AuthorID,
		},
	})
	if err != nil {
		return nil, err
	}

	attGetRes, err := u.attachmentClient.GetAttachmentsOfExercise(ctx, &attachmentSvcV1.GetAttachmentsOfExerciseRequest{
		ExerciseID: req.Id,
	})
	if err != nil {
		return nil, err
	}

	// delete old attachments
	for _, a := range attGetRes.Attachments {
		if _, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
			Id: a.Id,
		}); err != nil {
			return nil, err
		}
	}

	// create new attachments
	var attCreated []int64
	if len(req.Exercise.GetAttachments()) > 0 {
		for _, att := range req.Exercise.GetAttachments() {
			attRes, err := u.attachmentClient.CreateAttachment(ctx, &attachmentSvcV1.CreateAttachmentRequest{
				Attachment: &attachmentSvcV1.AttachmentInput{
					FileURL:    att.FileURL,
					ExerciseID: &req.Id,
					AuthorID:   req.Exercise.AuthorID,
					Name:       att.Name,
					Status:     "",
					Type:       att.Type,
					Thumbnail:  att.Thumbnail,
					Size:       att.Size,
				},
			})
			if err != nil {
				if len(attCreated) > 0 {
					for _, aErr := range attCreated {
						if _, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
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

	return &pb.UpdateExerciseResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *exerciseServiceGW) DeleteExercise(ctx context.Context, req *pb.DeleteExerciseRequest) (*pb.DeleteExerciseResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.exerciseClient.DeleteExercise(ctx, &exerciseSvcV1.DeleteExerciseRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	submissionRes, err := u.submissionClient.GetAllSubmissionsOfExercise(ctx, &submissionSvcV1.GetAllSubmissionsOfExerciseRequest{
		ExerciseID: req.Id,
	})
	if err != nil {
		return nil, err
	}

	for _, s := range submissionRes.GetSubmissions() {
		if _, err := u.submissionClient.DeleteSubmission(ctx, &submissionSvcV1.DeleteSubmissionRequest{
			Id: s.Id,
		}); err != nil {
			return nil, err
		}

		attSubRes, err := u.attachmentClient.GetAttachmentsOfSubmission(ctx, &attachmentSvcV1.GetAttachmentsOfSubmissionRequest{
			SubmissionID: s.Id,
		})
		if err != nil {
			return nil, err
		}

		for _, a := range attSubRes.GetAttachments() {
			if _, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
				Id: a.Id,
			}); err != nil {
				return nil, err
			}
		}
	}

	attGetRes, err := u.attachmentClient.GetAttachmentsOfExercise(ctx, &attachmentSvcV1.GetAttachmentsOfExerciseRequest{
		ExerciseID: req.Id,
	})
	if err != nil {
		return nil, err
	}

	for _, a := range attGetRes.GetAttachments() {
		if _, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
			Id: a.Id,
		}); err != nil {
			return nil, err
		}
	}

	return &pb.DeleteExerciseResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *exerciseServiceGW) GetExercises(ctx context.Context, req *pb.GetExercisesRequest) (*pb.GetExercisesResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var limit int64 = 5
	var page int64 = 1
	titleSearch := ""
	sortColumn := "id"
	sortOrder := "asc"

	if req.GetLimit() > 0 {
		limit = req.GetLimit()
	}

	if req.GetPage() > 0 {
		page = req.GetPage()
	}

	titleSearchTrim := strings.TrimSpace(req.GetTitleSearch())
	if len(titleSearchTrim) > 0 {
		titleSearch = titleSearchTrim
	}

	sortColumnTrim := strings.TrimSpace(req.GetSortColumn())
	if len(sortColumnTrim) > 0 {
		columns := map[string]string{
			"id":                 "id",
			"title":              "title",
			"content":            "content",
			"classroom_id":       "classroom_id",
			"deadline":           "deadline",
			"score":              "score",
			"reporting_stage_id": "reporting_stage_id",
			"author_id":          "author_id",
			"created_at":         "created_at",
			"updated_at":         "updated_at",
		}
		if stringInMap(sortColumnTrim, columns) {
			sortColumn = sortColumnTrim
		}
	}

	if req.IsDesc {
		sortOrder = "desc"
	}

	res, err := u.exerciseClient.GetExercises(ctx, &exerciseSvcV1.GetExercisesRequest{
		Limit:       limit,
		Page:        page,
		TitleSearch: titleSearch,
		SortColumn:  sortColumn,
		SortOrder:   sortOrder,
	})
	if err != nil {
		return nil, err
	}

	var exercises []*pb.ExerciseResponse
	for _, e := range res.GetExercises() {
		reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{
			Id: e.ReportingStageID,
		})
		if err != nil {
			return nil, err
		}

		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: e.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		attachment, err := u.attachmentClient.GetAttachmentsOfExercise(ctx, &attachmentSvcV1.GetAttachmentsOfExerciseRequest{
			ExerciseID: e.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentExerciseResponse
		for _, a := range attachment.Attachments {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentExerciseResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorExerciseResponse{
					Id:       author.User.Id,
					Class:    author.User.Class,
					Major:    author.User.Major,
					Phone:    author.User.Phone,
					PhotoSrc: author.User.PhotoSrc,
					Role:     author.User.Role,
					Name:     author.User.Name,
					Email:    author.User.Email,
				},
				CreatedAt: a.CreatedAt,
				Size:      a.Size,
				MimeType:  a.Type,
				Thumbnail: a.Thumbnail,
				FileName:  a.Name,
			})
		}

		exercises = append(exercises, &pb.ExerciseResponse{
			Id:          e.Id,
			Title:       e.Title,
			Description: e.Content,
			ClassroomID: e.ClassroomID,
			Deadline:    e.Deadline,
			Category: &pb.ReportingStageExerciseResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Label:       reportingStageRes.ReportingStage.Label,
				Description: reportingStageRes.ReportingStage.Description,
				Value:       reportingStageRes.ReportingStage.Value,
			},
			Author: &pb.AuthorExerciseResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt:   e.CreatedAt,
			UpdatedAt:   e.UpdatedAt,
			Attachments: attachments,
		})
	}

	return &pb.GetExercisesResponse{
		TotalCount: res.GetTotalCount(),
		Exercises:  exercises,
	}, nil
}

func (u *exerciseServiceGW) GetAllExercisesOfClassroom(ctx context.Context, req *pb.GetAllExercisesOfClassroomRequest) (*pb.GetAllExercisesOfClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var limit int64 = 5
	var page int64 = 1
	titleSearch := ""
	sortColumn := "id"
	sortOrder := "asc"
	var classroomID int64

	if req.GetLimit() > 0 {
		limit = req.GetLimit()
	}

	if req.GetPage() > 0 {
		page = req.GetPage()
	}

	titleSearchTrim := strings.TrimSpace(req.GetTitleSearch())
	if len(titleSearchTrim) > 0 {
		titleSearch = titleSearchTrim
	}

	sortColumnTrim := strings.TrimSpace(req.GetSortColumn())
	if len(sortColumnTrim) > 0 {
		columns := map[string]string{
			"id":                 "id",
			"title":              "title",
			"content":            "content",
			"classroom_id":       "classroom_id",
			"deadline":           "deadline",
			"score":              "score",
			"reporting_stage_id": "reporting_stage_id",
			"author_id":          "author_id",
			"created_at":         "created_at",
			"updated_at":         "updated_at",
		}
		if stringInMap(sortColumnTrim, columns) {
			sortColumn = sortColumnTrim
		}
	}

	if req.IsDesc {
		sortOrder = "desc"
	}

	if req.GetClassroomID() > 0 {
		classroomID = req.GetClassroomID()
	}

	res, err := u.exerciseClient.GetAllExercisesOfClassroom(ctx, &exerciseSvcV1.GetAllExercisesOfClassroomRequest{
		Limit:       limit,
		Page:        page,
		TitleSearch: titleSearch,
		SortColumn:  sortColumn,
		SortOrder:   sortOrder,
		ClassroomID: classroomID,
	})
	if err != nil {
		return nil, err
	}

	var exercises []*pb.ExerciseResponse
	for _, p := range res.GetExercises() {
		reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{
			Id: p.ReportingStageID,
		})
		if err != nil {
			return nil, err
		}

		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: p.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		attachment, err := u.attachmentClient.GetAttachmentsOfExercise(ctx, &attachmentSvcV1.GetAttachmentsOfExerciseRequest{
			ExerciseID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentExerciseResponse
		for _, a := range attachment.Attachments {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentExerciseResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorExerciseResponse{
					Id:       author.User.Id,
					Class:    author.User.Class,
					Major:    author.User.Major,
					Phone:    author.User.Phone,
					PhotoSrc: author.User.PhotoSrc,
					Role:     author.User.Role,
					Name:     author.User.Name,
					Email:    author.User.Email,
				},
				CreatedAt: a.CreatedAt,
				Size:      a.Size,
				MimeType:  a.Type,
				Thumbnail: a.Thumbnail,
				FileName:  a.Name,
			})
		}

		exercises = append(exercises, &pb.ExerciseResponse{
			Id:          p.Id,
			Title:       p.Title,
			Description: p.Content,
			ClassroomID: p.ClassroomID,
			Deadline:    p.Deadline,
			Category: &pb.ReportingStageExerciseResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Label:       reportingStageRes.ReportingStage.Label,
				Description: reportingStageRes.ReportingStage.Description,
				Value:       reportingStageRes.ReportingStage.Value,
			},
			Author: &pb.AuthorExerciseResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Attachments: attachments,
		})
	}

	return &pb.GetAllExercisesOfClassroomResponse{
		TotalCount: res.GetTotalCount(),
		Exercises:  exercises,
	}, nil
}

func (u *exerciseServiceGW) GetAllExercisesInReportingStage(ctx context.Context, req *pb.GetAllExercisesInReportingStageRequest) (*pb.GetAllExercisesInReportingStageResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	classRes, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{
		ClassroomID: req.GetClassroomID(),
	})
	if err != nil {
		return nil, err
	}

	if !classRes.GetExists() {
		return &pb.GetAllExercisesInReportingStageResponse{
			Response: &pb.CommonExerciseResponse{
				StatusCode: 404,
				Message:    "classroom does not exist",
			},
		}, nil
	}

	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{
		Id: req.GetCategoryID(),
	})
	if err != nil {
		return nil, err
	}

	if rpsRes.GetResponse().StatusCode == 404 {
		return &pb.GetAllExercisesInReportingStageResponse{
			Response: &pb.CommonExerciseResponse{
				StatusCode: rpsRes.Response.StatusCode,
				Message:    rpsRes.Response.Message,
			},
		}, nil
	}

	res, err := u.exerciseClient.GetAllExercisesInReportingStage(ctx, &exerciseSvcV1.GetAllExercisesInReportingStageRequest{
		ClassroomID:      req.GetClassroomID(),
		ReportingStageID: req.GetCategoryID(),
	})
	if err != nil {
		return nil, err
	}

	var exercises []*pb.ExerciseResponse
	for _, p := range res.GetExercises() {
		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: p.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		attachment, err := u.attachmentClient.GetAttachmentsOfExercise(ctx, &attachmentSvcV1.GetAttachmentsOfExerciseRequest{
			ExerciseID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentExerciseResponse
		for _, a := range attachment.Attachments {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentExerciseResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorExerciseResponse{
					Id:       author.User.Id,
					Class:    author.User.Class,
					Major:    author.User.Major,
					Phone:    author.User.Phone,
					PhotoSrc: author.User.PhotoSrc,
					Role:     author.User.Role,
					Name:     author.User.Name,
					Email:    author.User.Email,
				},
				CreatedAt: a.CreatedAt,
				Size:      a.Size,
				MimeType:  a.Type,
				Thumbnail: a.Thumbnail,
				FileName:  a.Name,
			})
		}

		exercises = append(exercises, &pb.ExerciseResponse{
			Id:          p.Id,
			Title:       p.Title,
			Description: p.Content,
			ClassroomID: p.ClassroomID,
			Deadline:    p.Deadline,
			Category: &pb.ReportingStageExerciseResponse{
				Id:          rpsRes.ReportingStage.Id,
				Label:       rpsRes.ReportingStage.Label,
				Description: rpsRes.ReportingStage.Description,
				Value:       rpsRes.ReportingStage.Value,
			},
			Author: &pb.AuthorExerciseResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Attachments: attachments,
		})
	}

	return &pb.GetAllExercisesInReportingStageResponse{
		TotalCount: res.GetTotalCount(),
		Exercises:  exercises,
	}, nil
}
