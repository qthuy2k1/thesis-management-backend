package handler

// import (
// 	"context"
// 	"log"

// 	"github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/internal/repository"
// 	"google.golang.org/grpc/status"
// )

// // CreateTimeSlots retrieves a commitee request from gRPC-gateway and calls to the Repository layer, then returns the response and status code.
// func (h *CommiteeHdl) CreateTimeSlots(ctx context.Context, req *commiteepb.CreateTimeSlotsRequest) (*commiteepb.CreateTimeSlotsResponse, error) {
// 	log.Println("calling insert commitee...")
// 	p, err := validateAndConvertTimeSlots(req.TimeSlots)
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	cDetailOut, err := h.Repository.CreateTimeSlots(ctx, p)
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	resp := &commiteepb.CreateTimeSlotsResponse{
// 		Response: &commiteepb.CommonCommiteeResponse{
// 			StatusCode: 201,
// 			Message:    "Created",
// 		},
// 		TimeSlots: &commiteepb.TimeSlots{
// 			CommiteeID:   int64(cDetailOut.CommiteeID),
// 			InstructorID: cDetailOut.InstructorID,
// 			StudentID:    cDetailOut.StudentID,
// 		},
// 	}

// 	return resp, nil
// }

// // GetTimeSlots returns a commitee in db given by id
// func (h *CommiteeHdl) GetTimeSlots(ctx context.Context, req *commiteepb.GetTimeSlotsRequest) (*commiteepb.GetTimeSlotsResponse, error) {
// 	log.Println("calling get commitee...")
// 	if err := req.Validate(); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}
// 	p, err := h.Repository.GetTimeSlots(ctx, int(req.GetId()))
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	pResp := commiteepb.TimeSlots{
// 		CommiteeID:   int64(p.CommiteeID),
// 		InstructorID: p.InstructorID,
// 		StudentID:    p.StudentID,
// 	}

// 	resp := &commiteepb.GetTimeSlotsResponse{
// 		Response: &commiteepb.CommonCommiteeResponse{
// 			StatusCode: 200,
// 			Message:    "OK",
// 		},
// 		TimeSlots: &pResp,
// 	}
// 	return resp, nil
// }

// func (c *CommiteeHdl) UpdateTimeSlots(ctx context.Context, req *commiteepb.UpdateTimeSlotsRequest) (*commiteepb.UpdateTimeSlotsResponse, error) {
// 	log.Println("calling update commitee...")
// 	if err := req.Validate(); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	p, err := validateAndConvertTimeSlots(req.TimeSlots)
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	if err := c.Repository.UpdateTimeSlots(ctx, repository.TimeSlotsInputRepo{
// 		CommiteeID:   p.CommiteeID,
// 		InstructorID: p.InstructorID,
// 		StudentID:    p.StudentID,
// 	}); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	return &commiteepb.UpdateTimeSlotsResponse{
// 		Response: &commiteepb.CommonCommiteeResponse{
// 			StatusCode: 200,
// 			Message:    "Success",
// 		},
// 	}, nil
// }

// func (h *CommiteeHdl) DeleteTimeSlots(ctx context.Context, req *commiteepb.DeleteTimeSlotsRequest) (*commiteepb.DeleteTimeSlotsResponse, error) {
// 	log.Println("calling delete commitee...")
// 	if err := req.Validate(); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	if err := h.Repository.DeleteTimeSlots(ctx, int(req.GetCommiteeID()), req.GetInstructorID(), req.GetStudentID()); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	return &commiteepb.DeleteTimeSlotsResponse{
// 		Response: &commiteepb.CommonCommiteeResponse{
// 			StatusCode: 200,
// 			Message:    "Success",
// 		},
// 	}, nil
// }

// func (h *CommiteeHdl) GetTimeSlotss(ctx context.Context, req *commiteepb.GetTimeSlotssRequest) (*commiteepb.GetTimeSlotssResponse, error) {
// 	log.Println("calling get all commitees...")
// 	if err := req.Validate(); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	ps, count, err := h.Repository.GetTimeSlotss(ctx)
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	var psResp []*commiteepb.TimeSlots
// 	for _, p := range ps {
// 		psResp = append(psResp, &commiteepb.TimeSlots{
// 			CommiteeID:   int64(p.CommiteeID),
// 			InstructorID: p.InstructorID,
// 			StudentID:    p.StudentID,
// 		})
// 	}

// 	return &commiteepb.GetTimeSlotssResponse{
// 		Response: &commiteepb.CommonCommiteeResponse{
// 			StatusCode: 200,
// 			Message:    "Success",
// 		},
// 		TimeSlotss: psResp,
// 		TotalCount: int64(count),
// 	}, nil
// }

// func (h *CommiteeHdl) GetAllTimeSlotssFromCommitee(ctx context.Context, req *commiteepb.GetAllTimeSlotssFromCommiteeRequest) (*commiteepb.GetAllTimeSlotssFromCommiteeResponse, error) {
// 	log.Println("calling get all commitees...")
// 	if err := req.Validate(); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	ps, err := h.Repository.GetAllTimeSlotssFromCommitee(ctx, int(req.GetCommiteeID()))
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	var psResp []*commiteepb.TimeSlots
// 	for _, p := range ps {
// 		psResp = append(psResp, &commiteepb.TimeSlots{
// 			CommiteeID:   int64(p.CommiteeID),
// 			InstructorID: p.InstructorID,
// 			StudentID:    p.StudentID,
// 		})
// 	}

// 	return &commiteepb.GetAllTimeSlotssFromCommiteeResponse{
// 		Response: &commiteepb.CommonCommiteeResponse{
// 			StatusCode: 200,
// 			Message:    "Success",
// 		},
// 		TimeSlotss: psResp,
// 	}, nil
// }

// func validateAndConvertTimeSlots(pbTimeSlots *commiteepb.TimeSlots) (repository.TimeSlotsInputRepo, error) {
// 	if err := pbTimeSlots.Validate(); err != nil {
// 		return repository.TimeSlotsInputRepo{}, err
// 	}

// 	return repository.TimeSlotsInputRepo{
// 		CommiteeID:   int(pbTimeSlots.CommiteeID),
// 		InstructorID: pbTimeSlots.InstructorID,
// 		StudentID:    pbTimeSlots.StudentID,
// 	}, nil
// }
