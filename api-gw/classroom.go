package main

import (
	"context"
	"errors"
	"sort"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	attachmentSvcV1 "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	waitingListSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	reportingStageSvcV1 "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	submissionSvcV1 "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
	topicSvcV1 "github.com/qthuy2k1/thesis-management-backend/topic-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type classroomServiceGW struct {
	pb.UnimplementedClassroomServiceServer
	classroomClient      classroomSvcV1.ClassroomServiceClient
	postClient           postSvcV1.PostServiceClient
	exerciseClient       exerciseSvcV1.ExerciseServiceClient
	reportingStageClient reportingStageSvcV1.ReportingStageServiceClient
	userClient           userSvcV1.UserServiceClient
	topicClient          topicSvcV1.TopicServiceClient
	waitingListClient    waitingListSvcV1.WaitingListServiceClient
	attachmentClient     attachmentSvcV1.AttachmentServiceClient
	submissionClient     submissionSvcV1.SubmissionServiceClient
}

func NewClassroomsService(classroomClient classroomSvcV1.ClassroomServiceClient, postClient postSvcV1.PostServiceClient, exerciseClient exerciseSvcV1.ExerciseServiceClient, reportingStageClient reportingStageSvcV1.ReportingStageServiceClient, userClient userSvcV1.UserServiceClient, topicClient topicSvcV1.TopicServiceClient, waitingListClient waitingListSvcV1.WaitingListServiceClient, attachmentClient attachmentSvcV1.AttachmentServiceClient, submissionClient submissionSvcV1.SubmissionServiceClient) *classroomServiceGW {
	return &classroomServiceGW{
		classroomClient:      classroomClient,
		postClient:           postClient,
		exerciseClient:       exerciseClient,
		reportingStageClient: reportingStageClient,
		userClient:           userClient,
		topicClient:          topicClient,
		waitingListClient:    waitingListClient,
		attachmentClient:     attachmentClient,
		submissionClient:     submissionClient,
	}
}

func (u *classroomServiceGW) CreateClassroom(ctx context.Context, req *pb.CreateClassroomRequest) (*pb.CreateClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	topicTags := ""
	if req.GetClassroom().TopicTags != nil {
		topicTags = req.GetClassroom().GetTopicTags()
	}

	lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.Classroom.LecturerID})
	if err != nil {
		return nil, err
	}

	if lecturerRes.Response.StatusCode == 404 {
		return nil, errors.New("user not found")
	}

	res, err := u.classroomClient.CreateClassroom(ctx, &classroomSvcV1.CreateClassroomRequest{
		Classroom: &classroomSvcV1.ClassroomInput{
			Title:           "",
			Description:     "",
			Status:          req.GetClassroom().GetStatus(),
			LecturerID:      req.GetClassroom().GetLecturerID(),
			ClassCourse:     req.GetClassroom().GetClassCourse(),
			TopicTags:       &topicTags,
			QuantityStudent: req.GetClassroom().GetQuantityStudent(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateClassroomResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
	}, nil
}

func (u *classroomServiceGW) GetClassroom(ctx context.Context, req *pb.GetClassroomRequest) (*pb.GetClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	filter := &postSvcV1.GetAllPostsOfClassroomRequest{
		Page:        1,
		Limit:       9999999,
		TitleSearch: "",
		SortColumn:  "created_at",
		SortOrder:   "asc",
	}

	titleSearch := strings.TrimSpace(req.TitleSearch)
	if titleSearch != "" {
		filter.TitleSearch = titleSearch
	}

	sortColumn := strings.TrimSpace(req.SortColumn)
	if sortColumn != "" {
		filter.SortColumn = sortColumn
	}

	// Get all posts of classroom
	resPost, err := u.postClient.GetAllPostsOfClassroom(ctx, &postSvcV1.GetAllPostsOfClassroomRequest{
		ClassroomID: res.GetClassroom().GetId(),
		Page:        filter.Page,
		Limit:       filter.Limit,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   filter.SortOrder,
	})
	if err != nil {
		return nil, err
	}

	var postsAndExercises []*pb.PostsAndExercisesOfClassroom
	for _, p := range resPost.GetPosts() {
		reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &reportingStageSvcV1.GetReportingStageRequest{
			Id: p.ReportingStageID,
		})
		if err != nil {
			return nil, err
		}

		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: p.AuthorID})
		if err != nil {
			return nil, err
		}

		if authorRes.Response.StatusCode == 404 {
			return nil, errors.New("user not found")
		}

		postsAndExercises = append(postsAndExercises, &pb.PostsAndExercisesOfClassroom{
			Id:          p.Id,
			Title:       p.Title,
			Description: p.Content,
			ClassroomID: p.ClassroomID,
			Category: &pb.ReportingStageClassroomResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Label:       reportingStageRes.ReportingStage.Label,
				Description: reportingStageRes.ReportingStage.Description,
				Value:       reportingStageRes.ReportingStage.Value,
			},
			Author: &pb.AuthorClassroomResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			Type:      "post",
		})
	}

	// Get all exercises of classroom
	resExercise, err := u.exerciseClient.GetAllExercisesOfClassroom(ctx, &exerciseSvcV1.GetAllExercisesOfClassroomRequest{
		ClassroomID: res.GetClassroom().GetId(),
		Page:        filter.Page,
		Limit:       filter.Limit,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   filter.SortOrder,
	})
	if err != nil {
		return nil, err
	}

	for _, p := range resExercise.GetExercises() {
		reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &reportingStageSvcV1.GetReportingStageRequest{
			Id: p.ReportingStageID,
		})
		if err != nil {
			return nil, err
		}

		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: p.AuthorID})
		if err != nil {
			return nil, err
		}

		if authorRes.Response.StatusCode == 404 {
			return nil, errors.New("user not found")
		}

		postsAndExercises = append(postsAndExercises, &pb.PostsAndExercisesOfClassroom{
			Id:          p.Id,
			Title:       p.Title,
			Description: p.Content,
			ClassroomID: p.ClassroomID,
			Deadline:    p.Deadline,
			Category: &pb.ReportingStageClassroomResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Label:       reportingStageRes.ReportingStage.Label,
				Description: reportingStageRes.ReportingStage.Description,
				Value:       reportingStageRes.ReportingStage.Value,
			},
			Author: &pb.AuthorClassroomResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			Type:      "exercise",
		})
	}

	// Sort the combined slice by CreatedAt field
	sort.Slice(postsAndExercises, func(i, j int) bool {
		return postsAndExercises[i].CreatedAt.AsTime().Before(postsAndExercises[j].CreatedAt.AsTime())
	})

	lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: res.GetClassroom().LecturerID,
	})
	if err != nil {
		return nil, err
	}

	userRes, err := u.userClient.GetAllMembersOfClassroom(ctx, &userSvcV1.GetAllMembersOfClassroomRequest{
		ClassroomID: res.GetClassroom().GetId(),
	})
	if err != nil {
		return nil, err
	}

	var userListID []string
	for _, user := range userRes.Members {
		userListID = append(userListID, user.MemberID)
	}

	topicRes, err := u.topicClient.GetAllTopicsOfListUser(ctx, &topicSvcV1.GetAllTopicsOfListUserRequest{
		UserID: userListID,
	})
	if err != nil {
		return nil, err
	}

	var topic []*pb.TopicClassroomResponse
	for _, t := range topicRes.GetTopic() {
		topic = append(topic, &pb.TopicClassroomResponse{
			Id:             t.Id,
			Title:          t.Title,
			TypeTopic:      t.TypeTopic,
			MemberQuantity: t.MemberQuantity,
			StudentID:      t.StudentID,
			MemberEmail:    t.MemberEmail,
			Description:    t.Description,
		})
	}

	topicTags := ""
	if res.GetClassroom().TopicTags != nil {
		topicTags = res.GetClassroom().GetTopicTags()
	}

	return &pb.GetClassroomResponse{
		Classroom: &pb.ClassroomResponse{
			Id:          res.GetClassroom().GetId(),
			Status:      res.GetClassroom().GetStatus(),
			ClassCourse: res.GetClassroom().GetClassCourse(),
			TopicTags:   &topicTags,
			Lecturer: &pb.AuthorClassroomResponse{
				Id:       lecturerRes.User.Id,
				Class:    lecturerRes.User.Class,
				Major:    lecturerRes.User.Major,
				Phone:    lecturerRes.User.Phone,
				PhotoSrc: lecturerRes.User.PhotoSrc,
				Role:     lecturerRes.User.Role,
				Name:     lecturerRes.User.Name,
				Email:    lecturerRes.User.Email,
			},
			QuantityStudent:   res.GetClassroom().GetQuantityStudent(),
			CreatedAt:         res.GetClassroom().GetCreatedAt(),
			UpdatedAt:         res.GetClassroom().GetUpdatedAt(),
			Topic:             topic,
			PostsAndExercises: postsAndExercises,
		},
	}, nil
}

func (u *classroomServiceGW) UpdateClassroom(ctx context.Context, req *pb.UpdateClassroomRequest) (*pb.UpdateClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	topicTags := ""
	if req.GetClassroom().TopicTags != nil {
		topicTags = req.GetClassroom().GetTopicTags()
	}

	res, err := u.classroomClient.UpdateClassroom(ctx, &classroomSvcV1.UpdateClassroomRequest{
		Id: req.GetId(),
		Classroom: &classroomSvcV1.ClassroomInput{
			Title:           "",
			Description:     "",
			Status:          req.GetClassroom().GetStatus(),
			LecturerID:      req.GetClassroom().GetLecturerID(),
			ClassCourse:     req.GetClassroom().GetClassCourse(),
			TopicTags:       &topicTags,
			QuantityStudent: req.GetClassroom().GetQuantityStudent(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateClassroomResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
	}, nil
}

func (u *classroomServiceGW) DeleteClassroom(ctx context.Context, req *pb.DeleteClassroomRequest) (*pb.DeleteClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.classroomClient.DeleteClassroom(ctx, &classroomSvcV1.DeleteClassroomRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	wltRes, err := u.waitingListClient.GetWaitingListsOfClassroom(ctx, &waitingListSvcV1.GetWaitingListsOfClassroomRequest{
		ClassroomID: req.Id,
	})
	if err != nil {
		return nil, err
	}

	// remove waiting list in the classroom was deleted
	for _, l := range wltRes.GetWaitingLists() {
		_, err := u.waitingListClient.DeleteWaitingList(ctx, &waitingListSvcV1.DeleteWaitingListRequest{
			Id: l.Id,
		})
		if err != nil {
			return nil, err
		}
	}

	exRes, err := u.exerciseClient.GetAllExercisesOfClassroom(ctx, &exerciseSvcV1.GetAllExercisesOfClassroomRequest{
		Page:        1,
		Limit:       99999,
		TitleSearch: "",
		SortColumn:  "id",
		SortOrder:   "asc",
		ClassroomID: req.Id,
	})
	if err != nil {
		return nil, err
	}

	for _, e := range exRes.GetExercises() {
		if _, err := u.exerciseClient.DeleteExercise(ctx, &exerciseSvcV1.DeleteExerciseRequest{
			Id: e.Id,
		}); err != nil {
			return nil, err
		}

		submissionRes, err := u.submissionClient.GetAllSubmissionsOfExercise(ctx, &submissionSvcV1.GetAllSubmissionsOfExerciseRequest{
			ExerciseID: e.Id,
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
			ExerciseID: e.Id,
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
	}

	postRes, err := u.postClient.GetAllPostsOfClassroom(ctx, &postSvcV1.GetAllPostsOfClassroomRequest{
		Page:        1,
		Limit:       99999,
		TitleSearch: "",
		SortColumn:  "id",
		SortOrder:   "asc",
		ClassroomID: req.Id,
	})
	if err != nil {
		return nil, err
	}

	for _, p := range postRes.GetPosts() {
		if _, err := u.postClient.DeletePost(ctx, &postSvcV1.DeletePostRequest{
			Id: p.Id,
		}); err != nil {
			return nil, err
		}

		attGetRes, err := u.attachmentClient.GetAttachmentsOfPost(ctx, &attachmentSvcV1.GetAttachmentsOfPostRequest{
			PostID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		for _, a := range attGetRes.Attachments {
			if _, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
				Id: a.Id,
			}); err != nil {
				return nil, err
			}
		}
	}

	return &pb.DeleteClassroomResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
	}, nil
}

func (u *classroomServiceGW) GetClassrooms(ctx context.Context, req *pb.GetClassroomsRequest) (*pb.GetClassroomsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	filter := &classroomSvcV1.GetClassroomsRequest{}

	if req.GetLimit() > 0 {
		filter.Limit = req.GetLimit()
	} else {
		filter.Limit = 5
	}

	if req.GetPage() > 0 {
		filter.Page = req.GetPage()
	} else {
		filter.Page = 1
	}

	titleSearchTrim := strings.TrimSpace(req.GetTitleSearch())
	if len(titleSearchTrim) > 0 {
		filter.TitleSearch = titleSearchTrim
	}

	sortColumnTrim := strings.TrimSpace(req.GetSortColumn())
	if len(sortColumnTrim) > 0 {
		columns := map[string]string{
			"id":             "id",
			"title":          "title",
			"description":    "description",
			"status":         "status",
			"lecturer_id":    "lecturer_id",
			"code_classroom": "code_classroom",
			"topic_tags":     "topic_tags",
			"quantity":       "quantity",
			"created_at":     "created_at",
			"updated_at":     "updated_at",
		}
		if stringInMap(sortColumnTrim, columns) {
			filter.SortColumn = sortColumnTrim
		} else {
			filter.SortColumn = "id"
		}
	} else {
		filter.SortColumn = "id"
	}

	sortOrder := "asc"
	if req.IsDesc {
		sortOrder = "desc"
	}

	res, err := u.classroomClient.GetClassrooms(ctx, &classroomSvcV1.GetClassroomsRequest{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   sortOrder,
	})
	if err != nil {
		return nil, err
	}

	var classrooms []*pb.ClassroomResponse
	for _, c := range res.GetClassrooms() {
		// Get all posts of classroom
		resPost, err := u.postClient.GetAllPostsOfClassroom(ctx, &postSvcV1.GetAllPostsOfClassroomRequest{
			ClassroomID: c.GetId(),
			Page:        filter.Page,
			Limit:       filter.Limit,
			TitleSearch: filter.TitleSearch,
			SortColumn:  filter.SortColumn,
			SortOrder:   filter.SortOrder,
		})
		if err != nil {
			return nil, err
		}

		var postsAndExercises []*pb.PostsAndExercisesOfClassroom
		for _, p := range resPost.GetPosts() {
			reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &reportingStageSvcV1.GetReportingStageRequest{
				Id: p.ReportingStageID,
			})
			if err != nil {
				return nil, err
			}

			authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: p.AuthorID})
			if err != nil {
				return nil, err
			}

			postsAndExercises = append(postsAndExercises, &pb.PostsAndExercisesOfClassroom{
				Id:          p.Id,
				Title:       p.Title,
				Description: p.Content,
				ClassroomID: p.ClassroomID,
				Category: &pb.ReportingStageClassroomResponse{
					Id:          reportingStageRes.ReportingStage.Id,
					Label:       reportingStageRes.ReportingStage.Label,
					Description: reportingStageRes.ReportingStage.Description,
					Value:       reportingStageRes.ReportingStage.Value,
				},
				Author: &pb.AuthorClassroomResponse{
					Id:       authorRes.User.Id,
					Class:    authorRes.User.Class,
					Major:    authorRes.User.Major,
					Phone:    authorRes.User.Phone,
					PhotoSrc: authorRes.User.PhotoSrc,
					Role:     authorRes.User.Role,
					Name:     authorRes.User.Name,
					Email:    authorRes.User.Email,
				},
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
				Type:      "post",
			})
		}

		// Get all exercises of classroom
		resExercise, err := u.exerciseClient.GetAllExercisesOfClassroom(ctx, &exerciseSvcV1.GetAllExercisesOfClassroomRequest{
			ClassroomID: c.GetId(),
			Page:        filter.Page,
			Limit:       filter.Limit,
			TitleSearch: filter.TitleSearch,
			SortColumn:  filter.SortColumn,
			SortOrder:   filter.SortOrder,
		})
		if err != nil {
			return nil, err
		}

		for _, p := range resExercise.GetExercises() {
			reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &reportingStageSvcV1.GetReportingStageRequest{
				Id: p.ReportingStageID,
			})
			if err != nil {
				return nil, err
			}

			authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: p.AuthorID})
			if err != nil {
				return nil, err
			}

			postsAndExercises = append(postsAndExercises, &pb.PostsAndExercisesOfClassroom{
				Id:          p.Id,
				Title:       p.Title,
				Description: p.Content,
				ClassroomID: p.ClassroomID,
				Deadline:    p.Deadline,
				Category: &pb.ReportingStageClassroomResponse{
					Id:          reportingStageRes.ReportingStage.Id,
					Label:       reportingStageRes.ReportingStage.Label,
					Description: reportingStageRes.ReportingStage.Description,
					Value:       reportingStageRes.ReportingStage.Value,
				},
				Author: &pb.AuthorClassroomResponse{
					Id:       authorRes.User.Id,
					Class:    authorRes.User.Class,
					Major:    authorRes.User.Major,
					Phone:    authorRes.User.Phone,
					PhotoSrc: authorRes.User.PhotoSrc,
					Role:     authorRes.User.Role,
					Name:     authorRes.User.Name,
					Email:    authorRes.User.Email,
				},
				CreatedAt: p.CreatedAt,
				UpdatedAt: p.UpdatedAt,
				Type:      "exercise",
			})
		}

		// Sort the combined slice by CreatedAt field
		sort.Slice(postsAndExercises, func(i, j int) bool {
			return postsAndExercises[i].CreatedAt.AsTime().Before(postsAndExercises[j].CreatedAt.AsTime())
		})

		lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: c.LecturerID})
		if err != nil {
			return nil, err
		}

		classrooms = append(classrooms, &pb.ClassroomResponse{
			Id:     c.Id,
			Status: c.Status,
			Lecturer: &pb.AuthorClassroomResponse{
				Id:       lecturerRes.User.Id,
				Class:    lecturerRes.User.Class,
				Major:    lecturerRes.User.Major,
				Phone:    lecturerRes.User.Phone,
				PhotoSrc: lecturerRes.User.PhotoSrc,
				Role:     lecturerRes.User.Role,
				Name:     lecturerRes.User.Name,
				Email:    lecturerRes.User.Email,
			},
			ClassCourse:       c.ClassCourse,
			TopicTags:         c.TopicTags,
			QuantityStudent:   c.QuantityStudent,
			CreatedAt:         c.CreatedAt,
			UpdatedAt:         c.UpdatedAt,
			PostsAndExercises: postsAndExercises,
		})
	}

	return &pb.GetClassroomsResponse{
		TotalCount: res.GetTotalCount(),
		Classrooms: classrooms,
	}, nil
}

func (u *classroomServiceGW) GetLecturerClassroom(ctx context.Context, req *pb.GetLecturerClassroomRequest) (*pb.GetLecturerClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	classroomRes, err := u.classroomClient.GetLecturerClassroom(ctx, &classroomSvcV1.GetLecturerClassroomRequest{
		LecturerID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetLecturerClassroomResponse{
		Classroom: &pb.ClassroomResponse{
			Id:     classroomRes.Classroom.Id,
			Status: classroomRes.Classroom.Status,
			Lecturer: &pb.AuthorClassroomResponse{
				Id:       lecturerRes.User.Id,
				Class:    lecturerRes.User.Class,
				Major:    lecturerRes.User.Major,
				Phone:    lecturerRes.User.Phone,
				PhotoSrc: lecturerRes.User.PhotoSrc,
				Role:     lecturerRes.User.Role,
				Name:     lecturerRes.User.Name,
				Email:    lecturerRes.User.Email,
			},
			ClassCourse:     classroomRes.Classroom.ClassCourse,
			TopicTags:       classroomRes.Classroom.TopicTags,
			QuantityStudent: classroomRes.Classroom.QuantityStudent,
			CreatedAt:       classroomRes.Classroom.CreatedAt,
			UpdatedAt:       classroomRes.Classroom.UpdatedAt,
		},
	}, nil
}

// func (u *classroomServiceGW) GetUserClassroom(ctx context.Context, req *pb.GetUserClassroomRequest) (*pb.GetUserClassroomResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	memberRes, err := u.userClient.GetUserMember(ctx, &userSvcV1.GetUserMemberRequest{
// 		UserID: req.UserID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	classroomRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
// 		Id: memberRes.GetMember().ClassroomID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
// 		Id: classroomRes.Classroom.LecturerID,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetUserClassroomResponse{
// 		Classroom: &pb.ClassroomResponse{
// 			Id:     classroomRes.Classroom.Id,
// 			Status: classroomRes.Classroom.Status,
// 			Lecturer: &pb.AuthorClassroomResponse{
// 				Id:       lecturerRes.User.Id,
// 				Class:    lecturerRes.User.Class,
// 				Major:    lecturerRes.User.Major,
// 				Phone:    lecturerRes.User.Phone,
// 				PhotoSrc: lecturerRes.User.PhotoSrc,
// 				Role:     lecturerRes.User.Role,
// 				Name:     lecturerRes.User.Name,
// 				Email:    lecturerRes.User.Email,
// 			},
// 			ClassCourse:     classroomRes.Classroom.ClassCourse,
// 			TopicTags:       classroomRes.Classroom.TopicTags,
// 			QuantityStudent: classroomRes.Classroom.QuantityStudent,
// 			CreatedAt:       classroomRes.Classroom.CreatedAt,
// 			UpdatedAt:       classroomRes.Classroom.UpdatedAt,
// 		},
// 	}, nil
// }

func stringInMap(s string, m map[string]string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}
