package main

import (
	"context"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
)

type classroomServiceGW struct {
	pb.UnimplementedClassroomServiceServer
	classroomClient classroomSvcV1.ClassroomServiceClient
	postClient      postSvcV1.PostServiceClient
	exerciseClient  exerciseSvcV1.ExerciseServiceClient
}

func NewClassroomsService(classroomClient classroomSvcV1.ClassroomServiceClient, postClient postSvcV1.PostServiceClient, exerciseClient exerciseSvcV1.ExerciseServiceClient) *classroomServiceGW {
	return &classroomServiceGW{
		classroomClient: classroomClient,
		postClient:      postClient,
		exerciseClient:  exerciseClient,
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

	var posts []*pb.PostInClassroomResponse
	for _, p := range resPost.GetPosts() {
		posts = append(posts, &pb.PostInClassroomResponse{
			Id:               p.Id,
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      p.ClassroomID,
			ReportingStageID: p.ReportingStageID,
			AuthorID:         p.AuthorID,
			CreatedAt:        p.CreatedAt,
			UpdatedAt:        p.UpdatedAt,
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

	var exercises []*pb.ExerciseInClassroomResponse
	for _, p := range resExercise.GetExercises() {
		exercises = append(exercises, &pb.ExerciseInClassroomResponse{
			Id:               p.Id,
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      p.ClassroomID,
			ReportingStageID: p.ReportingStageID,
			AuthorID:         p.AuthorID,
			CreatedAt:        p.CreatedAt,
			UpdatedAt:        p.UpdatedAt,
		})
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
			LecturerId:    res.GetClassroom().GetLecturerId(),
			Quantity:      res.GetClassroom().GetQuantity(),
			CreatedAt:     res.GetClassroom().GetCreatedAt(),
			UpdatedAt:     res.GetClassroom().GetUpdatedAt(),

			Post:     posts,
			Exercise: exercises,
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
		classrooms = append(classrooms, &pb.ClassroomResponse{
			Id:            c.Id,
			Title:         c.Title,
			Description:   c.Description,
			Status:        c.Status,
			LecturerId:    c.LecturerId,
			CodeClassroom: c.CodeClassroom,
			TopicTags:     c.TopicTags,
			Quantity:      c.Quantity,
			CreatedAt:     c.CreatedAt,
			UpdatedAt:     c.UpdatedAt,
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
