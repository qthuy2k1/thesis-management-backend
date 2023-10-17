package handler

import (
	"context"
	"log"

	commiteepb "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/internal/repository"
	"google.golang.org/grpc/status"
)

// CreateCommiteeUserDetail retrieves a commitee request from gRPC-gateway and calls to the Repository layer, then returns the response and status code.
func (h *CommiteeHdl) CreateCommiteeUserDetail(ctx context.Context, req *commiteepb.CreateCommiteeUserDetailRequest) (*commiteepb.CreateCommiteeUserDetailResponse, error) {
	log.Println("calling insert commitee...")
	p, err := validateAndConvertCommiteeUserDetail(req.CommiteeUserDetail)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	cDetailOut, err := h.Repository.CreateCommiteeUserDetail(ctx, p)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &commiteepb.CreateCommiteeUserDetailResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 201,
			Message:    "Created",
		},
		CommiteeUserDetail: &commiteepb.CommiteeUserDetail{
			CommiteeID: int64(cDetailOut.CommiteeID),
			LecturerID: cDetailOut.LecturerID,
			StudentID:  cDetailOut.StudentID,
		},
	}

	return resp, nil
}

// GetCommiteeUserDetail returns a commitee in db given by id
// func (h *CommiteeHdl) GetCommiteeUserDetail(ctx context.Context, req *commiteepb.GetCommiteeUserDetailRequest) (*commiteepb.GetCommiteeUserDetailResponse, error) {
// 	log.Println("calling get commitee...")
// 	if err := req.Validate(); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}
// 	p, err := h.Repository.GetCommiteeUserDetail(ctx, int(req.GetId()))
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	pResp := commiteepb.CommiteeUserDetail{
// 		CommiteeID: int64(p.CommiteeID),
// 		LecturerID: p.LecturerID,
// 		StudentID:  p.StudentID,
// 	}

// 	resp := &commiteepb.GetCommiteeUserDetailResponse{
// 		Response: &commiteepb.CommonCommiteeResponse{
// 			StatusCode: 200,
// 			Message:    "OK",
// 		},
// 		CommiteeUserDetail: &pResp,
// 	}
// 	return resp, nil
// }

// func (c *CommiteeHdl) UpdateCommiteeUserDetail(ctx context.Context, req *commiteepb.UpdateCommiteeUserDetailRequest) (*commiteepb.UpdateCommiteeUserDetailResponse, error) {
// 	log.Println("calling update commitee...")
// 	if err := req.Validate(); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	p, err := validateAndConvertCommiteeUserDetail(req.CommiteeUserDetail)
// 	if err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	if err := c.Repository.UpdateCommiteeUserDetail(ctx, repository.CommiteeUserDetailInputRepo{
// 		CommiteeID: p.CommiteeID,
// 		LecturerID: p.LecturerID,
// 		StudentID:  p.StudentID,
// 	}); err != nil {
// 		code, err := convertCtrlError(err)
// 		return nil, status.Errorf(code, "err: %v", err)
// 	}

// 	return &commiteepb.UpdateCommiteeUserDetailResponse{
// 		Response: &commiteepb.CommonCommiteeResponse{
// 			StatusCode: 200,
// 			Message:    "Success",
// 		},
// 	}, nil
// }

func (h *CommiteeHdl) DeleteCommiteeUserDetail(ctx context.Context, req *commiteepb.DeleteCommiteeUserDetailRequest) (*commiteepb.DeleteCommiteeUserDetailResponse, error) {
	log.Println("calling delete commitee...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Repository.DeleteCommiteeUserDetail(ctx, int(req.GetCommiteeID()), req.GetLecturerID(), req.GetStudentID()); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &commiteepb.DeleteCommiteeUserDetailResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *CommiteeHdl) GetCommiteeUserDetails(ctx context.Context, req *commiteepb.GetCommiteeUserDetailsRequest) (*commiteepb.GetCommiteeUserDetailsResponse, error) {
	log.Println("calling get all commitees...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ps, count, err := h.Repository.GetCommiteeUserDetails(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*commiteepb.CommiteeUserDetail
	for _, p := range ps {
		psResp = append(psResp, &commiteepb.CommiteeUserDetail{
			CommiteeID: int64(p.CommiteeID),
			LecturerID: p.LecturerID,
			StudentID:  p.StudentID,
		})
	}

	return &commiteepb.GetCommiteeUserDetailsResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		CommiteeUserDetails: psResp,
		TotalCount:          int64(count),
	}, nil
}

func (h *CommiteeHdl) GetAllCommiteeUserDetailsFromCommitee(ctx context.Context, req *commiteepb.GetAllCommiteeUserDetailsFromCommiteeRequest) (*commiteepb.GetAllCommiteeUserDetailsFromCommiteeResponse, error) {
	log.Println("calling get all commitees...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ps, err := h.Repository.GetAllCommiteeUserDetailsFromCommitee(ctx, int(req.GetCommiteeID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*commiteepb.CommiteeUserDetail
	for _, p := range ps {
		psResp = append(psResp, &commiteepb.CommiteeUserDetail{
			CommiteeID: int64(p.CommiteeID),
			LecturerID: p.LecturerID,
			StudentID:  p.StudentID,
		})
	}

	return &commiteepb.GetAllCommiteeUserDetailsFromCommiteeResponse{
		Response: &commiteepb.CommonCommiteeResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		CommiteeUserDetails: psResp,
	}, nil
}

func validateAndConvertCommiteeUserDetail(pbCommiteeUserDetail *commiteepb.CommiteeUserDetail) (repository.CommiteeUserDetailInputRepo, error) {
	if err := pbCommiteeUserDetail.Validate(); err != nil {
		return repository.CommiteeUserDetailInputRepo{}, err
	}

	return repository.CommiteeUserDetailInputRepo{
		CommiteeID: int(pbCommiteeUserDetail.CommiteeID),
		LecturerID: pbCommiteeUserDetail.LecturerID,
		StudentID:  pbCommiteeUserDetail.StudentID,
	}, nil
}
