package handler

import (
	"context"
	"log"
	"time"

	exercisepb "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/exercise-svc/internal/service"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateExercise retrieves a exercise request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *ExerciseHdl) CreateExercise(ctx context.Context, req *exercisepb.CreateExerciseRequest) (*exercisepb.CreateExerciseResponse, error) {
	log.Println("calling insert exercise...")
	e, err := validateAndConvertExercise(req.Exercise)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.CreateExercise(ctx, e); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &exercisepb.CreateExerciseResponse{
		Response: &exercisepb.CommonExerciseResponse{
			StatusCode: 201,
			Message:    "Created",
		},
	}

	return resp, nil
}

// GetExercise returns a exercise in db given by id
func (h *ExerciseHdl) GetExercise(ctx context.Context, req *exercisepb.GetExerciseRequest) (*exercisepb.GetExerciseResponse, error) {
	log.Println("calling get exercise...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	e, err := h.Service.GetExercise(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := exercisepb.ExerciseResponse{
		Id:          int32(e.ID),
		Title:       e.Title,
		Content:     e.Content,
		ClassroomID: int32(e.ClassroomID),
		Deadline: &datetime.DateTime{
			Day:     int32(e.Deadline.Day()),
			Month:   int32(e.Deadline.Month()),
			Year:    int32(e.Deadline.Year()),
			Hours:   int32(e.Deadline.Hour()),
			Minutes: int32(e.Deadline.Minute()),
			Seconds: int32(e.Deadline.Second()),
		},
		Score:            int32(e.Score),
		ReportingStageID: int32(e.ReportingStageID),
		AuthorID:         int32(e.AuthorID),
		CreatedAt:        timestamppb.New(e.CreatedAt),
		UpdatedAt:        timestamppb.New(e.UpdatedAt),
	}

	resp := &exercisepb.GetExerciseResponse{
		Response: &exercisepb.CommonExerciseResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Exercise: &pResp,
	}
	return resp, nil
}

func (c *ExerciseHdl) UpdateExercise(ctx context.Context, req *exercisepb.UpdateExerciseRequest) (*exercisepb.UpdateExerciseResponse, error) {
	log.Println("calling update exercise...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	e, err := validateAndConvertExercise(req.Exercise)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateExercise(ctx, int(req.GetId()), service.ExerciseInputSvc{
		Title:            e.Title,
		Content:          e.Content,
		ClassroomID:      e.ClassroomID,
		Deadline:         e.Deadline,
		Score:            e.Score,
		ReportingStageID: e.ReportingStageID,
		AuthorID:         e.AuthorID,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &exercisepb.UpdateExerciseResponse{
		Response: &exercisepb.CommonExerciseResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *ExerciseHdl) DeleteExercise(ctx context.Context, req *exercisepb.DeleteExerciseRequest) (*exercisepb.DeleteExerciseResponse, error) {
	log.Println("calling delete exercise...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteExercise(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &exercisepb.DeleteExerciseResponse{
		Response: &exercisepb.CommonExerciseResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *ExerciseHdl) GetExercises(ctx context.Context, req *exercisepb.GetExercisesRequest) (*exercisepb.GetExercisesResponse, error) {
	log.Println("calling get all exercises...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	filter := service.ExerciseFilterSvc{
		Limit:       int(req.GetLimit()),
		Page:        int(req.GetPage()),
		TitleSearch: req.GetTitleSearch(),
		SortColumn:  req.GetSortColumn(),
		SortOrder:   req.GetSortOrder(),
	}

	ps, count, err := h.Service.GetExercises(ctx, filter)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*exercisepb.ExerciseResponse
	for _, e := range ps {
		psResp = append(psResp, &exercisepb.ExerciseResponse{
			Id:          int32(e.ID),
			Title:       e.Title,
			Content:     e.Content,
			ClassroomID: int32(e.ClassroomID),
			Deadline: &datetime.DateTime{
				Day:     int32(e.Deadline.Day()),
				Month:   int32(e.Deadline.Month()),
				Year:    int32(e.Deadline.Year()),
				Hours:   int32(e.Deadline.Hour()),
				Minutes: int32(e.Deadline.Minute()),
				Seconds: int32(e.Deadline.Second()),
			},
			Score:            int32(e.Score),
			ReportingStageID: int32(e.ReportingStageID),
			AuthorID:         int32(e.AuthorID),
			CreatedAt:        timestamppb.New(e.CreatedAt),
			UpdatedAt:        timestamppb.New(e.UpdatedAt),
		})
	}

	return &exercisepb.GetExercisesResponse{
		Response: &exercisepb.CommonExerciseResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Exercises:  psResp,
		TotalCount: int32(count),
	}, nil
}

func (h *ExerciseHdl) GetAllExercisesOfClassroom(ctx context.Context, req *exercisepb.GetAllExercisesOfClassroomRequest) (*exercisepb.GetAllExercisesOfClassroomResponse, error) {
	log.Println("calling get all exercises of a classroom...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	filter := service.ExerciseFilterSvc{
		Limit:       int(req.GetLimit()),
		Page:        int(req.GetPage()),
		TitleSearch: req.GetTitleSearch(),
		SortColumn:  req.GetSortColumn(),
		SortOrder:   req.GetSortOrder(),
	}

	es, count, err := h.Service.GetAllExercisesOfClassroom(ctx, filter, int(req.GetClassroomID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var esResp []*exercisepb.ExerciseResponse
	for _, e := range es {
		esResp = append(esResp, &exercisepb.ExerciseResponse{
			Id:          int32(e.ID),
			Title:       e.Title,
			Content:     e.Content,
			ClassroomID: int32(e.ClassroomID),
			Deadline: &datetime.DateTime{
				Day:     int32(e.Deadline.Day()),
				Month:   int32(e.Deadline.Month()),
				Year:    int32(e.Deadline.Year()),
				Hours:   int32(e.Deadline.Hour()),
				Minutes: int32(e.Deadline.Minute()),
				Seconds: int32(e.Deadline.Second()),
			},
			Score:            int32(e.Score),
			ReportingStageID: int32(e.ReportingStageID),
			AuthorID:         int32(e.AuthorID),
			CreatedAt:        timestamppb.New(e.CreatedAt),
			UpdatedAt:        timestamppb.New(e.UpdatedAt),
		})
	}

	return &exercisepb.GetAllExercisesOfClassroomResponse{
		Response: &exercisepb.CommonExerciseResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Exercises:  esResp,
		TotalCount: int32(count),
	}, nil
}

func validateAndConvertExercise(pbExercise *exercisepb.ExerciseInput) (service.ExerciseInputSvc, error) {
	if err := pbExercise.Validate(); err != nil {
		return service.ExerciseInputSvc{}, err
	}

	deadline, err := time.Parse("year:2006 month:1 day:2 hours:15 minutes:4 seconds:5", pbExercise.Deadline.String())

	log.Println(deadline)
	if err != nil {
		return service.ExerciseInputSvc{}, err
	}

	return service.ExerciseInputSvc{
		Title:            pbExercise.Title,
		Content:          pbExercise.Content,
		ClassroomID:      int(pbExercise.ClassroomID),
		Deadline:         deadline,
		Score:            int(pbExercise.Score),
		ReportingStageID: int(pbExercise.ReportingStageID),
		AuthorID:         int(pbExercise.AuthorID),
	}, nil
}
