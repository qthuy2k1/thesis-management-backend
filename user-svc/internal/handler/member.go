package handler

import (
	"context"
	"errors"
	"log"

	memberpb "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateMember retrieves a member request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *UserHdl) CreateMember(ctx context.Context, req *memberpb.CreateMemberRequest) (*memberpb.CreateMemberResponse, error) {
	log.Println("calling insert member...")
	u, err := validateAndConvertMember(req.Member)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.CreateMember(ctx, u); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &memberpb.CreateMemberResponse{
		Response: &memberpb.CommonUserResponse{
			StatusCode: 201,
			Message:    "Created",
		},
	}

	return resp, nil
}

// GetMember returns a member in db given by id
func (h *UserHdl) GetMember(ctx context.Context, req *memberpb.GetMemberRequest) (*memberpb.GetMemberResponse, error) {
	log.Println("calling get member...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	u, err := h.Service.GetMember(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := memberpb.MemberResponse{
		Id:          int64(u.ID),
		ClassroomID: int64(u.ClassroomID),
		MemberID:    u.MemberID,
		Status:      u.Status,
		IsDefense:   u.IsDefense,
		CreatedAt:   timestamppb.New(u.CreatedAt),
	}

	resp := &memberpb.GetMemberResponse{
		Response: &memberpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Member: &pResp,
	}

	return resp, nil
}

func (c *UserHdl) UpdateMember(ctx context.Context, req *memberpb.UpdateMemberRequest) (*memberpb.UpdateMemberResponse, error) {
	log.Println("calling update member...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	u, err := validateAndConvertMember(req.Member)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateMember(ctx, int(req.GetId()), service.MemberInputSvc{
		ClassroomID: u.ClassroomID,
		MemberID:    u.MemberID,
		Status:      u.Status,
		IsDefense:   u.IsDefense,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &memberpb.UpdateMemberResponse{
		Response: &memberpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *UserHdl) DeleteMember(ctx context.Context, req *memberpb.DeleteMemberRequest) (*memberpb.DeleteMemberResponse, error) {
	log.Println("calling delete member...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteMember(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &memberpb.DeleteMemberResponse{
		Response: &memberpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *UserHdl) GetMembers(ctx context.Context, req *memberpb.GetMembersRequest) (*memberpb.GetMembersResponse, error) {
	log.Println("calling get all members...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ps, count, err := h.Service.GetMembers(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*memberpb.MemberResponse
	for _, u := range ps {
		psResp = append(psResp, &memberpb.MemberResponse{
			Id:          int64(u.ID),
			ClassroomID: int64(u.ClassroomID),
			MemberID:    u.MemberID,
			Status:      u.Status,
			IsDefense:   u.IsDefense,
			CreatedAt:   timestamppb.New(u.CreatedAt),
		})
	}

	return &memberpb.GetMembersResponse{
		Response: &memberpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Members:    psResp,
		TotalCount: int64(count),
	}, nil
}

func (h *UserHdl) GetAllMembersOfClassroom(ctx context.Context, req *memberpb.GetAllMembersOfClassroomRequest) (*memberpb.GetAllMembersOfClassroomResponse, error) {
	log.Println("calling get all members of a classroom...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ps, count, err := h.Service.GetAllMembersOfClassroom(ctx, int(req.GetClassroomID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*memberpb.MemberResponse
	for _, u := range ps {
		psResp = append(psResp, &memberpb.MemberResponse{
			Id:          int64(u.ID),
			ClassroomID: int64(u.ClassroomID),
			MemberID:    u.MemberID,
			Status:      u.Status,
			IsDefense:   u.IsDefense,
			CreatedAt:   timestamppb.New(u.CreatedAt),
		})
	}

	return &memberpb.GetAllMembersOfClassroomResponse{
		Response: &memberpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Members:    psResp,
		TotalCount: int64(count),
	}, nil
}

func (h *UserHdl) IsUserJoinedClassroom(ctx context.Context, req *memberpb.IsUserJoinedClassroomRequest) (*memberpb.IsUserJoinedClassroomResponse, error) {
	m, err := h.Service.IsUserJoinedClassroom(ctx, req.GetUserID())
	if err != nil {
		code, err := convertCtrlError(err)
		if errors.Is(err, ErrMemberNotFound) {
			return &memberpb.IsUserJoinedClassroomResponse{
				Response: &memberpb.CommonUserResponse{
					StatusCode: 404,
					Message:    "Member not found",
				},
			}, nil
		}

		return nil, status.Errorf(code, "err: %v", err)
	}

	return &memberpb.IsUserJoinedClassroomResponse{
		Response: &memberpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "",
		},
		Member: &memberpb.MemberResponse{
			Id:          int64(m.ID),
			ClassroomID: int64(m.ClassroomID),
			MemberID:    m.MemberID,
			Status:      m.Status,
			IsDefense:   m.IsDefense,
			CreatedAt:   timestamppb.New(m.CreatedAt),
		},
	}, nil
}

func validateAndConvertMember(pbMember *memberpb.MemberInput) (service.MemberInputSvc, error) {
	if err := pbMember.Validate(); err != nil {
		return service.MemberInputSvc{}, err
	}

	return service.MemberInputSvc{
		ClassroomID: int(pbMember.ClassroomID),
		MemberID:    pbMember.MemberID,
		Status:      pbMember.Status,
		IsDefense:   pbMember.IsDefense,
	}, nil
}
