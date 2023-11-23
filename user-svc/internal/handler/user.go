package handler

import (
	"context"

	userpb "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service"
	"google.golang.org/grpc/status"
)

// CreateUser retrieves a user request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *UserHdl) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	u, err := validateAndConvertUser(req.User)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.CreateUser(ctx, u); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &userpb.CreateUserResponse{
		Response: &userpb.CommonUserResponse{
			StatusCode: 201,
			Message:    "Created",
		},
	}

	return resp, nil
}

// GetUser returns a user in db given by id
func (h *UserHdl) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	u, err := h.Service.GetUser(ctx, req.GetId())
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := userpb.UserResponse{
		Id:             u.ID,
		Class:          u.Class,
		Major:          u.Major,
		Phone:          u.Phone,
		PhotoSrc:       u.PhotoSrc,
		Role:           u.Role,
		Name:           u.Name,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
	}

	resp := &userpb.GetUserResponse{
		Response: &userpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		User: &pResp,
	}

	return resp, nil
}

func (c *UserHdl) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	u, err := validateAndConvertUser(req.User)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateUser(ctx, req.GetId(), service.UserInputSvc{
		Class:          u.Class,
		Major:          u.Major,
		Phone:          u.Phone,
		PhotoSrc:       u.PhotoSrc,
		Role:           u.Role,
		Name:           u.Name,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &userpb.UpdateUserResponse{
		Response: &userpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *UserHdl) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteUser(ctx, req.GetId()); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &userpb.DeleteUserResponse{
		Response: &userpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *UserHdl) GetUsers(ctx context.Context, req *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ps, count, err := h.Service.GetUsers(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*userpb.UserResponse
	for _, u := range ps {
		psResp = append(psResp, &userpb.UserResponse{
			Id:             u.ID,
			Class:          u.Class,
			Major:          u.Major,
			Phone:          u.Phone,
			PhotoSrc:       u.PhotoSrc,
			Role:           u.Role,
			Name:           u.Name,
			Email:          u.Email,
			HashedPassword: u.HashedPassword,
		})
	}

	return &userpb.GetUsersResponse{
		Response: &userpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Users:      psResp,
		TotalCount: int64(count),
	}, nil
}

func (h *UserHdl) GetAllLecturers(ctx context.Context, req *userpb.GetAllLecturersRequest) (*userpb.GetAllLecturersResponse, error) {
	ps, count, err := h.Service.GetAllLecturers(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*userpb.UserResponse
	for _, u := range ps {
		psResp = append(psResp, &userpb.UserResponse{
			Id:             u.ID,
			Class:          u.Class,
			Major:          u.Major,
			Phone:          u.Phone,
			PhotoSrc:       u.PhotoSrc,
			Role:           u.Role,
			Name:           u.Name,
			Email:          u.Email,
			HashedPassword: u.HashedPassword,
		})
	}

	return &userpb.GetAllLecturersResponse{
		Response: &userpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Lecturers:  psResp,
		TotalCount: int64(count),
	}, nil
}

func (h *UserHdl) UnsubscribeClassroom(ctx context.Context, req *userpb.UnsubscribeClassroomRequest) (*userpb.UnsubscribeClassroomResponse, error) {
	if err := h.Service.UnsubscribeClassroom(ctx, req.MemberID, int(req.ClassroomID)); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &userpb.UnsubscribeClassroomResponse{
		Response: &userpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func validateAndConvertUser(pbUser *userpb.UserInput) (service.UserInputSvc, error) {
	if err := pbUser.Validate(); err != nil {
		return service.UserInputSvc{}, err
	}

	return service.UserInputSvc{
		ID:             pbUser.Id,
		Class:          pbUser.Class,
		Major:          pbUser.Major,
		Phone:          pbUser.Phone,
		PhotoSrc:       pbUser.PhotoSrc,
		Role:           pbUser.Role,
		Name:           pbUser.Name,
		Email:          pbUser.Email,
		HashedPassword: pbUser.HashedPassword,
	}, nil
}
