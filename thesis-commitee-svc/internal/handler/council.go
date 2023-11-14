package handler

import (
	"context"

	commiteepb "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/internal/repository"
	"google.golang.org/grpc/status"
)

// CreateCouncil retrieves a commitee request from gRPC-gateway and calls to the Repository layer, then returns the response and status code.
func (h *CommiteeHdl) CreateCouncil(ctx context.Context, req *commiteepb.CreateCouncilRequest) (*commiteepb.CreateCouncilResponse, error) {
	r, err := validateAndConvertCouncil(req.Council)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	_, err = h.Repository.CreateCouncil(ctx, r)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &commiteepb.CreateCouncilResponse{
		Response: &commiteepb.CommonScheduleResponse{
			StatusCode: 201,
			Message:    "Created",
		},
	}

	return resp, nil
}

// GetCouncil returns a commitee in db given by id
func (h *CommiteeHdl) GetCouncil(ctx context.Context, req *commiteepb.GetCouncilRequest) (*commiteepb.GetCouncilResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	r, err := h.Repository.GetCouncil(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := commiteepb.CouncilResponse{
		Id:         int64(r.ID),
		LecturerID: r.LecturerID,
		ThesisID:   int64(r.ThesisID),
	}

	resp := &commiteepb.GetCouncilResponse{
		Council: &pResp,
	}
	return resp, nil
}

func (h *CommiteeHdl) GetCouncils(ctx context.Context, req *commiteepb.GetCouncilsRequest) (*commiteepb.GetCouncilsResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	rs, _, err := h.Repository.GetCouncils(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var rsResp []*commiteepb.CouncilResponse
	for _, r := range rs {
		rsResp = append(rsResp, &commiteepb.CouncilResponse{
			Id:         int64(r.ID),
			LecturerID: r.LecturerID,
			ThesisID:   int64(r.ThesisID),
		})
	}

	return &commiteepb.GetCouncilsResponse{
		Councils: rsResp,
	}, nil
}

func (h *CommiteeHdl) GetCouncilsByThesisID(ctx context.Context, req *commiteepb.GetCouncilsByThesisIDRequest) (*commiteepb.GetCouncilsByThesisIDResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	rs, _, err := h.Repository.GetCouncilsByThesisID(ctx, int(req.ThesisID))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var rsResp []*commiteepb.CouncilResponse
	for _, r := range rs {
		rsResp = append(rsResp, &commiteepb.CouncilResponse{
			Id:         int64(r.ID),
			LecturerID: r.LecturerID,
			ThesisID:   int64(r.ThesisID),
		})
	}

	return &commiteepb.GetCouncilsByThesisIDResponse{
		Councils: rsResp,
	}, nil
}

func validateAndConvertCouncil(pbCouncil *commiteepb.CouncilInput) (repository.CouncilInputRepo, error) {
	if err := pbCouncil.Validate(); err != nil {
		return repository.CouncilInputRepo{}, err
	}

	return repository.CouncilInputRepo{
		LecturerID: pbCouncil.LecturerID,
		ThesisID:   int(pbCouncil.ThesisID),
	}, nil
}
