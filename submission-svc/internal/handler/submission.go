package handler

import (
	"context"
	"log"
	"time"

	submissionpb "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/submission-svc/internal/service"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc/status"
)

// CreateSubmission retrieves a submission request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *SubmissionHdl) CreateSubmission(ctx context.Context, req *submissionpb.CreateSubmissionRequest) (*submissionpb.CreateSubmissionResponse, error) {
	log.Println("calling insert submission...")
	e, err := validateAndConvertSubmission(req.Submission)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	id, err := h.Service.CreateSubmission(ctx, e)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &submissionpb.CreateSubmissionResponse{
		Response: &submissionpb.CommonSubmissionResponse{
			StatusCode: 201,
			Message:    "Created",
		},
		SubmissionID: id,
	}

	return resp, nil
}

func (c *SubmissionHdl) UpdateSubmission(ctx context.Context, req *submissionpb.UpdateSubmissionRequest) (*submissionpb.UpdateSubmissionResponse, error) {
	log.Println("calling update submission...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	s, err := validateAndConvertSubmission(req.Submission)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateSubmission(ctx, int(req.GetId()), service.SubmissionInputSvc{
		UserID:         s.UserID,
		ExerciseID:     s.ExerciseID,
		SubmissionDate: s.SubmissionDate,
		Status:         s.Status,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &submissionpb.UpdateSubmissionResponse{
		Response: &submissionpb.CommonSubmissionResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *SubmissionHdl) DeleteSubmission(ctx context.Context, req *submissionpb.DeleteSubmissionRequest) (*submissionpb.DeleteSubmissionResponse, error) {
	log.Println("calling delete submission...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteSubmission(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &submissionpb.DeleteSubmissionResponse{
		Response: &submissionpb.CommonSubmissionResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *SubmissionHdl) GetAllSubmissionsOfExercise(ctx context.Context, req *submissionpb.GetAllSubmissionsOfExerciseRequest) (*submissionpb.GetAllSubmissionsOfExerciseResponse, error) {
	log.Println("calling get all submissions of a classroom...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ss, count, err := h.Service.GetAllSubmissionsOfExercise(ctx, int(req.GetExerciseID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var ssResp []*submissionpb.SubmissionResponse
	for _, s := range ss {
		ssResp = append(ssResp, &submissionpb.SubmissionResponse{
			Id:         int64(s.ID),
			UserID:     s.UserID,
			ExerciseID: int64(s.ExerciseID),
			SubmissionDate: &datetime.DateTime{
				Year:    int32(s.SubmissionDate.Year()),
				Month:   int32(s.SubmissionDate.Minute()),
				Day:     int32(s.SubmissionDate.Day()),
				Hours:   int32(s.SubmissionDate.Hour()),
				Minutes: int32(s.SubmissionDate.Minute()),
				Seconds: int32(s.SubmissionDate.Second()),
			},
			Status: s.Status,
		})
	}

	return &submissionpb.GetAllSubmissionsOfExerciseResponse{
		Response: &submissionpb.CommonSubmissionResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Submissions: ssResp,
		TotalCount:  int64(count),
	}, nil
}

func validateAndConvertSubmission(pbSubmission *submissionpb.SubmissionInput) (service.SubmissionInputSvc, error) {
	if err := pbSubmission.Validate(); err != nil {
		return service.SubmissionInputSvc{}, err
	}

	submissionDate, err := time.Parse("year:2006 month:1 day:2 hours:15 minutes:4 seconds:5", pbSubmission.SubmissionDate.String())

	if err != nil {
		return service.SubmissionInputSvc{}, err
	}

	return service.SubmissionInputSvc{
		UserID:         pbSubmission.UserID,
		ExerciseID:     int(pbSubmission.ExerciseID),
		SubmissionDate: submissionDate,
		Status:         pbSubmission.Status,
	}, nil
}
