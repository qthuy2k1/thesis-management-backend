package handler

import (
	"context"
	"log"
	"time"

	commiteepb "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/internal/repository"
	"google.golang.org/genproto/googleapis/type/datetime"
	"google.golang.org/grpc/status"
)

// CreateCommitee retrieves a commitee request from gRPC-gateway and calls to the Repository layer, then returns the response and status code.
func (h *CommiteeHdl) CreateCommitee(ctx context.Context, req *commiteepb.CreateCommiteeRequest) (*commiteepb.CreateCommiteeResponse, error) {
	log.Println("calling insert commitee...")
	p, err := validateAndConvertCommitee(req.Commitee)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	cOut, err := h.Repository.CreateCommitee(ctx, p)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &commiteepb.CreateCommiteeResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 201,
			Message:    "Created",
		},
		Commitee: &commiteepb.CommiteeResponse{
			Id: int64(cOut.ID),
			StartDate: &datetime.DateTime{
				Day:     int32(cOut.StartDate.Day()),
				Month:   int32(cOut.StartDate.Month()),
				Year:    int32(cOut.StartDate.Year()),
				Hours:   int32(cOut.StartDate.Hour()),
				Minutes: int32(cOut.StartDate.Minute()),
			},
			Shift: cOut.Period,
		},
	}

	return resp, nil
}

// GetCommitee returns a commitee in db given by id
func (h *CommiteeHdl) GetCommitee(ctx context.Context, req *commiteepb.GetCommiteeRequest) (*commiteepb.GetCommiteeResponse, error) {
	log.Println("calling get commitee...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	p, err := h.Repository.GetCommitee(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := commiteepb.CommiteeResponse{
		Id: int64(p.ID),
		StartDate: &datetime.DateTime{
			Day:     int32(p.StartDate.Day()),
			Month:   int32(p.StartDate.Month()),
			Year:    int32(p.StartDate.Year()),
			Hours:   int32(p.StartDate.Hour()),
			Minutes: int32(p.StartDate.Minute()),
		},
		Shift:  p.Period,
		RoomID: int64(p.RoomID),
	}

	resp := &commiteepb.GetCommiteeResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Commitee: &pResp,
	}
	return resp, nil
}

func (c *CommiteeHdl) UpdateCommitee(ctx context.Context, req *commiteepb.UpdateCommiteeRequest) (*commiteepb.UpdateCommiteeResponse, error) {
	log.Println("calling update commitee...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	p, err := validateAndConvertCommitee(req.Commitee)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Repository.UpdateCommitee(ctx, int(req.GetId()), repository.CommiteeInputRepo{
		StartDate: p.StartDate,
		Period:    p.Period,
		RoomID:    p.RoomID,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &commiteepb.UpdateCommiteeResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *CommiteeHdl) DeleteCommitee(ctx context.Context, req *commiteepb.DeleteCommiteeRequest) (*commiteepb.DeleteCommiteeResponse, error) {
	log.Println("calling delete commitee...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Repository.DeleteCommitee(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &commiteepb.DeleteCommiteeResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *CommiteeHdl) GetCommitees(ctx context.Context, req *commiteepb.GetCommiteesRequest) (*commiteepb.GetCommiteesResponse, error) {
	log.Println("calling get all commitees...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ps, count, err := h.Repository.GetCommitees(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*commiteepb.CommiteeResponse
	for _, p := range ps {
		psResp = append(psResp, &commiteepb.CommiteeResponse{
			Id: int64(p.ID),
			StartDate: &datetime.DateTime{
				Day:     int32(p.StartDate.Day()),
				Month:   int32(p.StartDate.Month()),
				Year:    int32(p.StartDate.Year()),
				Hours:   int32(p.StartDate.Hour()),
				Minutes: int32(p.StartDate.Minute()),
			},
			Shift:  p.Period,
			RoomID: int64(p.RoomID),
		})
	}

	return &commiteepb.GetCommiteesResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Commitees:  psResp,
		TotalCount: int64(count),
	}, nil
}

func validateAndConvertCommitee(pbCommitee *commiteepb.CommiteeInput) (repository.CommiteeInputRepo, error) {
	if err := pbCommitee.Validate(); err != nil {
		return repository.CommiteeInputRepo{}, err
	}

	startDate, err := time.Parse("year:2006 month:1 day:2 hours:15 minutes:4 seconds:5", pbCommitee.StartDate.String())

	if err != nil {
		return repository.CommiteeInputRepo{}, err
	}

	return repository.CommiteeInputRepo{
		StartDate: startDate,
		Period:    pbCommitee.Shift,
		RoomID:    int(pbCommitee.RoomID),
	}, nil
}
