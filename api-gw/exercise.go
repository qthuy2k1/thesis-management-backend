package main

import (
	"context"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	commentSvcV1 "github.com/qthuy2k1/thesis-management-backend/comment-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	rpsSvcV1 "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
)

type exerciseServiceGW struct {
	pb.UnimplementedExerciseServiceServer
	exerciseClient       exerciseSvcV1.ExerciseServiceClient
	classroomClient      classroomSvcV1.ClassroomServiceClient
	reportingStageClient rpsSvcV1.ReportingStageServiceClient
	commentClient        commentSvcV1.CommentServiceClient
}

func NewExercisesService(exerciseClient exerciseSvcV1.ExerciseServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, reportStageClient rpsSvcV1.ReportingStageServiceClient, commentClient commentSvcV1.CommentServiceClient) *exerciseServiceGW {
	return &exerciseServiceGW{
		exerciseClient:       exerciseClient,
		classroomClient:      classroomClient,
		reportingStageClient: reportStageClient,
		commentClient:        commentClient,
	}
}

func (u *exerciseServiceGW) CreateExercise(ctx context.Context, req *pb.CreateExerciseRequest) (*pb.CreateExerciseResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
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

	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{Id: req.GetExercise().GetReportingStageID()})
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

	res, err := u.exerciseClient.CreateExercise(ctx, &exerciseSvcV1.CreateExerciseRequest{
		Exercise: &exerciseSvcV1.ExerciseInput{
			Title:            req.GetExercise().Title,
			Content:          req.GetExercise().Content,
			ClassroomID:      req.GetExercise().ClassroomID,
			Deadline:         req.GetExercise().Deadline,
			Score:            req.GetExercise().Score,
			ReportingStageID: req.GetExercise().ReportingStageID,
			AuthorID:         req.GetExercise().AuthorID,
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

	commentRes, err := u.commentClient.GetCommentsOfAExercise(ctx, &commentSvcV1.GetCommentsOfAExerciseRequest{
		ExerciseID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	var comments []*pb.CommentExerciseResponse
	for _, c := range commentRes.GetComments() {
		comments = append(comments, &pb.CommentExerciseResponse{
			Id:         c.Id,
			UserID:     c.UserID,
			ExerciseID: *c.ExerciseID,
			Content:    c.Content,
			CreatedAt:  c.CreatedAt,
		})
	}

	return &pb.GetExerciseResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Exercise: &pb.ExerciseResponse{
			Id:               res.GetExercise().Id,
			Title:            res.GetExercise().Title,
			Content:          res.GetExercise().Content,
			ClassroomID:      res.GetExercise().ClassroomID,
			Deadline:         res.GetExercise().Deadline,
			Score:            res.GetExercise().Score,
			ReportingStageID: res.GetExercise().ReportingStageID,
			AuthorID:         res.GetExercise().AuthorID,
			CreatedAt:        res.GetExercise().CreatedAt,
			UpdatedAt:        res.GetExercise().UpdatedAt,
		},
		Comments: comments,
	}, nil
}

func (u *exerciseServiceGW) UpdateExercise(ctx context.Context, req *pb.UpdateExerciseRequest) (*pb.UpdateExerciseResponse, error) {
	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{Id: req.GetExercise().GetReportingStageID()})
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
			Content:          req.GetExercise().Content,
			ClassroomID:      req.GetExercise().ClassroomID,
			Deadline:         req.GetExercise().Deadline,
			Score:            req.GetExercise().Score,
			ReportingStageID: req.GetExercise().ReportingStageID,
			AuthorID:         req.GetExercise().AuthorID,
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

	var limit int32 = 5
	var page int32 = 1
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
		exercises = append(exercises, &pb.ExerciseResponse{
			Id:               e.Id,
			Title:            e.Title,
			Content:          e.Content,
			ClassroomID:      e.ClassroomID,
			Deadline:         e.Deadline,
			Score:            e.Score,
			ReportingStageID: e.ReportingStageID,
			AuthorID:         e.AuthorID,
			CreatedAt:        e.CreatedAt,
			UpdatedAt:        e.UpdatedAt,
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

func (u *exerciseServiceGW) GetAllExercisesOfClassroom(ctx context.Context, req *pb.GetAllExercisesOfClassroomRequest) (*pb.GetAllExercisesOfClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var limit int32 = 5
	var page int32 = 1
	titleSearch := ""
	sortColumn := "id"
	sortOrder := "asc"
	var classroomID int32

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
		exercises = append(exercises, &pb.ExerciseResponse{
			Id:               p.Id,
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      p.ClassroomID,
			Deadline:         p.Deadline,
			Score:            p.Score,
			ReportingStageID: p.ReportingStageID,
			AuthorID:         p.AuthorID,
			CreatedAt:        p.CreatedAt,
			UpdatedAt:        p.UpdatedAt,
		})
	}

	return &pb.GetAllExercisesOfClassroomResponse{
		Response: &pb.CommonExerciseResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Exercises:  exercises,
	}, nil
}
