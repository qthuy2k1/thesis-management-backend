package main

import (
	"context"
	"log"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
)

type exerciseServiceGW struct {
	pb.UnimplementedExerciseServiceServer
	exerciseClient  exerciseSvcV1.ExerciseServiceClient
	classroomClient classroomSvcV1.ClassroomServiceClient
}

func NewExercisesService(exerciseClient exerciseSvcV1.ExerciseServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient) *exerciseServiceGW {
	return &exerciseServiceGW{
		exerciseClient:  exerciseClient,
		classroomClient: classroomClient,
	}
}

func (u *exerciseServiceGW) CreateExercise(ctx context.Context, req *pb.CreateExerciseRequest) (*pb.CreateExerciseResponse, error) {
	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetExercise().ClassroomID})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.CreateExerciseResponse{
			Response: &pb.CommonExerciseResponse{
				StatusCode: 400,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	res, err := u.exerciseClient.CreateExercise(ctx, &exerciseSvcV1.CreateExerciseRequest{
		Exercise: &exerciseSvcV1.ExerciseInput{
			Title:       req.GetExercise().Title,
			Content:     req.GetExercise().Content,
			ClassroomID: req.GetExercise().ClassroomID,
			Deadline:    req.GetExercise().Deadline,
			Score:       req.GetExercise().Score,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateExerciseResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *exerciseServiceGW) GetExercise(ctx context.Context, req *pb.GetExerciseRequest) (*pb.GetExerciseResponse, error) {
	res, err := u.exerciseClient.GetExercise(ctx, &exerciseSvcV1.GetExerciseRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetExerciseResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Exercise: &pb.ExerciseResponse{
			Id:          res.GetExercise().Id,
			Title:       res.GetExercise().Title,
			Content:     res.GetExercise().Content,
			ClassroomID: res.GetExercise().ClassroomID,
			Deadline:    res.GetExercise().Deadline,
			Score:       res.GetExercise().Score,
			CreatedAt:   res.GetExercise().CreatedAt,
			UpdatedAt:   res.GetExercise().UpdatedAt,
		},
	}, nil
}

func (u *exerciseServiceGW) UpdateExercise(ctx context.Context, req *pb.UpdateExerciseRequest) (*pb.UpdateExerciseResponse, error) {
	log.Println(req)
	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetExercise().ClassroomID})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.UpdateExerciseResponse{
			Response: &pb.CommonExerciseResponse{
				StatusCode: 400,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	res, err := u.exerciseClient.UpdateExercise(ctx, &exerciseSvcV1.UpdateExerciseRequest{
		Id: req.GetId(),
		Exercise: &exerciseSvcV1.ExerciseInput{
			Title:       req.GetExercise().Title,
			Content:     req.GetExercise().Content,
			ClassroomID: req.GetExercise().ClassroomID,
			Deadline:    req.GetExercise().Deadline,
			Score:       req.GetExercise().Score,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateExerciseResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *exerciseServiceGW) DeleteExercise(ctx context.Context, req *pb.DeleteExerciseRequest) (*pb.DeleteExerciseResponse, error) {
	res, err := u.exerciseClient.DeleteExercise(ctx, &exerciseSvcV1.DeleteExerciseRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
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

	filter := &exerciseSvcV1.GetExercisesRequest{}

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
			"id":           "id",
			"title":        "title",
			"content":      "content",
			"classroom_id": "classroom_id",
			"deadline":     "deadline",
			"score":        "score",
			"created_at":   "created_at",
			"updated_at":   "updated_at",
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

	res, err := u.exerciseClient.GetExercises(ctx, &exerciseSvcV1.GetExercisesRequest{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   sortOrder,
	})
	if err != nil {
		return nil, err
	}

	var exercises []*pb.ExerciseResponse
	for _, e := range res.GetExercises() {
		exercises = append(exercises, &pb.ExerciseResponse{
			Id:          e.Id,
			Title:       e.Title,
			Content:     e.Content,
			ClassroomID: e.ClassroomID,
			Deadline:    e.Deadline,
			Score:       e.Score,
			CreatedAt:   e.CreatedAt,
			UpdatedAt:   e.UpdatedAt,
		})
	}

	return &pb.GetExercisesResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Exercises:  exercises,
	}, nil
}
