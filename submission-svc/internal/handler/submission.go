package handler

import (
	"context"
	"log"

	submissionpb "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/submission-svc/internal/service"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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
		UserID:     s.UserID,
		ExerciseID: s.ExerciseID,
		Status:     s.Status,
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
			Status:     s.Status,
			CreatedAt:  timestamppb.New(s.CreatedAt),
			UpdatedAt:  timestamppb.New(s.UpdatedAt),
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

func (h *SubmissionHdl) GetSubmissionOfUser(ctx context.Context, req *submissionpb.GetSubmissionOfUserRequest) (*submissionpb.GetSubmissionOfUserResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ss, err := h.Service.GetSubmissionOfUser(ctx, req.UserID, int(req.ExerciseID))
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
			Status:     s.Status,
			CreatedAt:  timestamppb.New(s.CreatedAt),
			UpdatedAt:  timestamppb.New(s.UpdatedAt),
		})
	}

	return &submissionpb.GetSubmissionOfUserResponse{
		Response: &submissionpb.CommonSubmissionResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Submissions: ssResp,
	}, nil
}

func (h *SubmissionHdl) GetSubmissionFromUser(ctx context.Context, req *submissionpb.GetSubmissionFromUserRequest) (*submissionpb.GetSubmissionFromUserResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ss, err := h.Service.GetAllSubmissionFromUser(ctx, req.UserID)
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
			Status:     s.Status,
			CreatedAt:  timestamppb.New(s.CreatedAt),
			UpdatedAt:  timestamppb.New(s.UpdatedAt),
		})
	}

	return &submissionpb.GetSubmissionFromUserResponse{
		Response: &submissionpb.CommonSubmissionResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Submissions: ssResp,
	}, nil
}

func validateAndConvertSubmission(pbSubmission *submissionpb.SubmissionInput) (service.SubmissionInputSvc, error) {
	if err := pbSubmission.Validate(); err != nil {
		return service.SubmissionInputSvc{}, err
	}

	return service.SubmissionInputSvc{
		UserID:     pbSubmission.UserID,
		ExerciseID: int(pbSubmission.ExerciseID),
		Status:     pbSubmission.Status,
	}, nil
}
