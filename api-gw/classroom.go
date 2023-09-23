package main

import (
	"context"
	"sort"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	reportingStageSvcV1 "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type classroomServiceGW struct {
	pb.UnimplementedClassroomServiceServer
	classroomClient      classroomSvcV1.ClassroomServiceClient
	postClient           postSvcV1.PostServiceClient
	exerciseClient       exerciseSvcV1.ExerciseServiceClient
	reportingStageClient reportingStageSvcV1.ReportingStageServiceClient
	userClient           userSvcV1.UserServiceClient
}

func NewClassroomsService(classroomClient classroomSvcV1.ClassroomServiceClient, postClient postSvcV1.PostServiceClient, exerciseClient exerciseSvcV1.ExerciseServiceClient, reportingStageClient reportingStageSvcV1.ReportingStageServiceClient, userClient userSvcV1.UserServiceClient) *classroomServiceGW {
	return &classroomServiceGW{
		classroomClient:      classroomClient,
		postClient:           postClient,
		exerciseClient:       exerciseClient,
		reportingStageClient: reportingStageClient,
		userClient:           userClient,
	}
}

func (u *classroomServiceGW) CreateClassroom(ctx context.Context, req *pb.CreateClassroomRequest) (*pb.CreateClassroomResponse, error) {
	res, err := u.classroomClient.CreateClassroom(ctx, &classroomSvcV1.CreateClassroomRequest{
		Classroom: &classroomSvcV1.ClassroomInput{
			Title:         req.GetClassroom().GetTitle(),
			Description:   req.GetClassroom().GetDescription(),
			Status:        req.GetClassroom().GetStatus(),
			LecturerId:    req.GetClassroom().GetLecturerId(),
			CodeClassroom: req.GetClassroom().GetCodeClassroom(),
			TopicTags:     req.GetClassroom().GetTopicTags(),
			Quantity:      req.GetClassroom().GetQuantity(),
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

		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: p.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		postsAndExercises = append(postsAndExercises, &pb.PostsAndExercisesOfClassroom{
			Id:          p.Id,
			Title:       p.Title,
			Content:     p.Content,
			ClassroomID: p.ClassroomID,
			ReportingStage: &pb.ReportingStageClassroomResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Name:        reportingStageRes.ReportingStage.Name,
				Description: reportingStageRes.ReportingStage.Description,
			},
			Author: &pb.AuthorClassroomResponse{
				Id:          authorRes.User.Id,
				Class:       authorRes.User.Class,
				Major:       authorRes.User.Major,
				Phone:       authorRes.User.Phone,
				PhotoSrc:    authorRes.User.PhotoSrc,
				Role:        authorRes.User.Role,
				Name:        authorRes.User.Name,
				Email:       authorRes.User.Email,
				ClassroomID: &authorRes.User.ClassroomID,
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

		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: p.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		postsAndExercises = append(postsAndExercises, &pb.PostsAndExercisesOfClassroom{
			Id:          p.Id,
			Title:       p.Title,
			Content:     p.Content,
			ClassroomID: p.ClassroomID,
			Deadline:    p.Deadline,
			Score:       &p.Score,
			ReportingStage: &pb.ReportingStageClassroomResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Name:        reportingStageRes.ReportingStage.Name,
				Description: reportingStageRes.ReportingStage.Description,
			},
			Author: &pb.AuthorClassroomResponse{
				Id:          authorRes.User.Id,
				Class:       authorRes.User.Class,
				Major:       authorRes.User.Major,
				Phone:       authorRes.User.Phone,
				PhotoSrc:    authorRes.User.PhotoSrc,
				Role:        authorRes.User.Role,
				Name:        authorRes.User.Name,
				Email:       authorRes.User.Email,
				ClassroomID: &authorRes.User.ClassroomID,
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
		Id: res.GetClassroom().LecturerId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetClassroomResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
		Classroom: &pb.ClassroomResponse{
			Id:            res.GetClassroom().GetId(),
			Title:         res.GetClassroom().GetTitle(),
			Description:   res.GetClassroom().GetDescription(),
			Status:        res.GetClassroom().GetStatus(),
			CodeClassroom: res.GetClassroom().GetCodeClassroom(),
			TopicTags:     res.GetClassroom().GetTopicTags(),
			Lecturer: &pb.AuthorClassroomResponse{
				Id:          lecturerRes.User.Id,
				Class:       lecturerRes.User.Class,
				Major:       lecturerRes.User.Major,
				Phone:       lecturerRes.User.Phone,
				PhotoSrc:    lecturerRes.User.PhotoSrc,
				Role:        lecturerRes.User.Role,
				Name:        lecturerRes.User.Name,
				Email:       lecturerRes.User.Email,
				ClassroomID: &lecturerRes.User.ClassroomID,
			},
			Quantity:          res.GetClassroom().GetQuantity(),
			CreatedAt:         res.GetClassroom().GetCreatedAt(),
			UpdatedAt:         res.GetClassroom().GetUpdatedAt(),
			PostsAndExercises: postsAndExercises,
		},
	}, nil
}

func (u *classroomServiceGW) UpdateClassroom(ctx context.Context, req *pb.UpdateClassroomRequest) (*pb.UpdateClassroomResponse, error) {
	res, err := u.classroomClient.UpdateClassroom(ctx, &classroomSvcV1.UpdateClassroomRequest{
		Id: req.GetId(),
		Classroom: &classroomSvcV1.ClassroomInput{
			Title:         req.GetClassroom().GetTitle(),
			Description:   req.GetClassroom().GetDescription(),
			Status:        req.GetClassroom().GetStatus(),
			LecturerId:    req.GetClassroom().GetLecturerId(),
			CodeClassroom: req.GetClassroom().GetCodeClassroom(),
			TopicTags:     req.GetClassroom().GetTopicTags(),
			Quantity:      req.GetClassroom().GetQuantity(),
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
	res, err := u.classroomClient.DeleteClassroom(ctx, &classroomSvcV1.DeleteClassroomRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
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

			authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: p.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			postsAndExercises = append(postsAndExercises, &pb.PostsAndExercisesOfClassroom{
				Id:          p.Id,
				Title:       p.Title,
				Content:     p.Content,
				ClassroomID: p.ClassroomID,
				ReportingStage: &pb.ReportingStageClassroomResponse{
					Id:          reportingStageRes.ReportingStage.Id,
					Name:        reportingStageRes.ReportingStage.Name,
					Description: reportingStageRes.ReportingStage.Description,
				},
				Author: &pb.AuthorClassroomResponse{
					Id:          authorRes.User.Id,
					Class:       authorRes.User.Class,
					Major:       authorRes.User.Major,
					Phone:       authorRes.User.Phone,
					PhotoSrc:    authorRes.User.PhotoSrc,
					Role:        authorRes.User.Role,
					Name:        authorRes.User.Name,
					Email:       authorRes.User.Email,
					ClassroomID: &authorRes.User.ClassroomID,
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

			authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: p.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			postsAndExercises = append(postsAndExercises, &pb.PostsAndExercisesOfClassroom{
				Id:          p.Id,
				Title:       p.Title,
				Content:     p.Content,
				ClassroomID: p.ClassroomID,
				Deadline:    p.Deadline,
				Score:       &p.Score,
				ReportingStage: &pb.ReportingStageClassroomResponse{
					Id:          reportingStageRes.ReportingStage.Id,
					Name:        reportingStageRes.ReportingStage.Name,
					Description: reportingStageRes.ReportingStage.Description,
				},
				Author: &pb.AuthorClassroomResponse{
					Id:          authorRes.User.Id,
					Class:       authorRes.User.Class,
					Major:       authorRes.User.Major,
					Phone:       authorRes.User.Phone,
					PhotoSrc:    authorRes.User.PhotoSrc,
					Role:        authorRes.User.Role,
					Name:        authorRes.User.Name,
					Email:       authorRes.User.Email,
					ClassroomID: &authorRes.User.ClassroomID,
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
			Id: c.LecturerId,
		})
		if err != nil {
			return nil, err
		}
		classrooms = append(classrooms, &pb.ClassroomResponse{
			Id:          c.Id,
			Title:       c.Title,
			Description: c.Description,
			Status:      c.Status,
			Lecturer: &pb.AuthorClassroomResponse{
				Id:          lecturerRes.User.Id,
				Class:       lecturerRes.User.Class,
				Major:       lecturerRes.User.Major,
				Phone:       lecturerRes.User.Phone,
				PhotoSrc:    lecturerRes.User.PhotoSrc,
				Role:        lecturerRes.User.Role,
				Name:        lecturerRes.User.Name,
				Email:       lecturerRes.User.Email,
				ClassroomID: &lecturerRes.User.ClassroomID,
			},
			CodeClassroom:     c.CodeClassroom,
			TopicTags:         c.TopicTags,
			Quantity:          c.Quantity,
			CreatedAt:         c.CreatedAt,
			UpdatedAt:         c.UpdatedAt,
			PostsAndExercises: postsAndExercises,
		})
	}

	return &pb.GetClassroomsResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
		TotalCount: res.GetTotalCount(),
		Classrooms: classrooms,
	}, nil
}

func stringInMap(s string, m map[string]string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}
