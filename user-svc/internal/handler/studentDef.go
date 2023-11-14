package handler

import (
	"context"
	"log"

	studentDefpb "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service"
	"google.golang.org/grpc/status"
)

// CreateStudentDef retrieves a studentDef request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *UserHdl) CreateStudentDef(ctx context.Context, req *studentDefpb.CreateStudentDefRequest) (*studentDefpb.CreateStudentDefResponse, error) {
	log.Println("calling insert studentDef...")
	u, err := validateAndConvertStudentDef(req.StudentDef)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.CreateStudentDef(ctx, u); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &studentDefpb.CreateStudentDefResponse{
		Response: &studentDefpb.CommonUserResponse{
			StatusCode: 201,
			Message:    "Created",
		},
	}

	return resp, nil
}

// GetStudentDef returns a studentDef in db given by id
func (h *UserHdl) GetStudentDef(ctx context.Context, req *studentDefpb.GetStudentDefRequest) (*studentDefpb.GetStudentDefResponse, error) {
	log.Println("calling get studentDef...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	sd, err := h.Service.GetStudentDef(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	// user
	u, err := h.Service.GetUser(ctx, sd.UserID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	// instructor
	l, err := h.Service.GetUser(ctx, sd.UserID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	sdResp := studentDefpb.StudentDefResponse{
		Id: int64(sd.ID),
		User: &studentDefpb.UserResponse{
			Id:       u.ID,
			Class:    u.Class,
			Major:    u.Major,
			Phone:    u.Phone,
			PhotoSrc: u.PhotoSrc,
			Role:     u.Role,
			Name:     u.Name,
			Email:    u.Email,
		},
		Instructor: &studentDefpb.UserResponse{
			Id:       l.ID,
			Class:    l.Class,
			Major:    l.Major,
			Phone:    l.Phone,
			PhotoSrc: l.PhotoSrc,
			Role:     l.Role,
			Name:     l.Name,
			Email:    l.Email,
		},
		TimeSlotsID: int64(sd.TimeSlotsID),
	}

	resp := &studentDefpb.GetStudentDefResponse{
		Response: &studentDefpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		StudentDef: &sdResp,
	}

	return resp, nil
}

func (c *UserHdl) UpdateStudentDef(ctx context.Context, req *studentDefpb.UpdateStudentDefRequest) (*studentDefpb.UpdateStudentDefResponse, error) {
	log.Println("calling update studentDef...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	sd, err := validateAndConvertStudentDef(req.StudentDef)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateStudentDef(ctx, int(req.GetId()), service.StudentDefInputSvc{
		UserID:       sd.UserID,
		InstructorID: sd.InstructorID,
		TimeSlotsID:  sd.TimeSlotsID,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &studentDefpb.UpdateStudentDefResponse{
		Response: &studentDefpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *UserHdl) DeleteStudentDef(ctx context.Context, req *studentDefpb.DeleteStudentDefRequest) (*studentDefpb.DeleteStudentDefResponse, error) {
	log.Println("calling delete studentDef...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteStudentDef(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &studentDefpb.DeleteStudentDefResponse{
		Response: &studentDefpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *UserHdl) GetStudentDefs(ctx context.Context, req *studentDefpb.GetStudentDefsRequest) (*studentDefpb.GetStudentDefsResponse, error) {
	log.Println("calling get all studentDefs...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	sds, count, err := h.Service.GetStudentDefs(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var sdsResp []*studentDefpb.StudentDefResponse
	for _, sd := range sds {
		// user
		u, err := h.Service.GetUser(ctx, sd.UserID)
		if err != nil {
			code, err := convertCtrlError(err)
			return nil, status.Errorf(code, "err: %v", err)
		}

		// instructor
		l, err := h.Service.GetUser(ctx, sd.UserID)
		if err != nil {
			code, err := convertCtrlError(err)
			return nil, status.Errorf(code, "err: %v", err)
		}

		sdsResp = append(sdsResp, &studentDefpb.StudentDefResponse{
			Id: int64(sd.ID),
			User: &studentDefpb.UserResponse{
				Id:       u.ID,
				Class:    u.Class,
				Major:    u.Major,
				Phone:    u.Phone,
				PhotoSrc: u.PhotoSrc,
				Role:     u.Role,
				Name:     u.Name,
				Email:    u.Email,
			},
			Instructor: &studentDefpb.UserResponse{
				Id:       l.ID,
				Class:    l.Class,
				Major:    l.Major,
				Phone:    l.Phone,
				PhotoSrc: l.PhotoSrc,
				Role:     l.Role,
				Name:     l.Name,
				Email:    l.Email,
			},
			TimeSlotsID: int64(sd.TimeSlotsID),
		})
	}

	return &studentDefpb.GetStudentDefsResponse{
		Response: &studentDefpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		StudentDefs: sdsResp,
		TotalCount:  int64(count),
	}, nil
}

func (h *UserHdl) GetAllStudentDefsOfInstructor(ctx context.Context, req *studentDefpb.GetAllStudentDefsOfInstructorRequest) (*studentDefpb.GetAllStudentDefsOfInstructorResponse, error) {
	log.Println("calling get all studentDefs of a classroom...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	sds, count, err := h.Service.GetAllStudentDefsOfInstructor(ctx, req.GetInstructorID())
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var sdsResp []*studentDefpb.StudentDefResponse
	for _, sd := range sds {
		// user
		u, err := h.Service.GetUser(ctx, sd.UserID)
		if err != nil {
			code, err := convertCtrlError(err)
			return nil, status.Errorf(code, "err: %v", err)
		}

		// instructor
		l, err := h.Service.GetUser(ctx, sd.UserID)
		if err != nil {
			code, err := convertCtrlError(err)
			return nil, status.Errorf(code, "err: %v", err)
		}

		sdsResp = append(sdsResp, &studentDefpb.StudentDefResponse{
			Id: int64(sd.ID),
			User: &studentDefpb.UserResponse{
				Id:       u.ID,
				Class:    u.Class,
				Major:    u.Major,
				Phone:    u.Phone,
				PhotoSrc: u.PhotoSrc,
				Role:     u.Role,
				Name:     u.Name,
				Email:    u.Email,
			},
			Instructor: &studentDefpb.UserResponse{
				Id:       l.ID,
				Class:    l.Class,
				Major:    l.Major,
				Phone:    l.Phone,
				PhotoSrc: l.PhotoSrc,
				Role:     l.Role,
				Name:     l.Name,
				Email:    l.Email,
			},
			TimeSlotsID: int64(sd.TimeSlotsID),
		})
	}

	return &studentDefpb.GetAllStudentDefsOfInstructorResponse{
		Response: &studentDefpb.CommonUserResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		StudentDefs: sdsResp,
		TotalCount:  int64(count),
	}, nil
}

func validateAndConvertStudentDef(pbStudentDef *studentDefpb.StudentDefInput) (service.StudentDefInputSvc, error) {
	if err := pbStudentDef.Validate(); err != nil {
		return service.StudentDefInputSvc{}, err
	}

	return service.StudentDefInputSvc{
		UserID:       pbStudentDef.UserID,
		InstructorID: pbStudentDef.InstructorID,
		TimeSlotsID:  int(pbStudentDef.TimeSlotsID),
	}, nil
}

// GetStudentDef returns a studentDef in db given by id
func (h *UserHdl) GetStudentDefByTimeSlotsID(ctx context.Context, req *studentDefpb.GetStudentDefByTimeSlotsIDRequest) (*studentDefpb.GetStudentDefByTimeSlotsIDResponse, error) {
	log.Println("calling get studentDef...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	sd, err := h.Service.GetStudentDefByTimeSlotsID(ctx, int(req.GetTimeSlotsID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	// user
	u, err := h.Service.GetUser(ctx, sd.UserID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	// instructor
	l, err := h.Service.GetUser(ctx, sd.UserID)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	sdResp := studentDefpb.StudentDefResponse{
		Id: int64(sd.ID),
		User: &studentDefpb.UserResponse{
			Id:       u.ID,
			Class:    u.Class,
			Major:    u.Major,
			Phone:    u.Phone,
			PhotoSrc: u.PhotoSrc,
			Role:     u.Role,
			Name:     u.Name,
			Email:    u.Email,
		},
		Instructor: &studentDefpb.UserResponse{
			Id:       l.ID,
			Class:    l.Class,
			Major:    l.Major,
			Phone:    l.Phone,
			PhotoSrc: l.PhotoSrc,
			Role:     l.Role,
			Name:     l.Name,
			Email:    l.Email,
		},
		TimeSlotsID: int64(sd.TimeSlotsID),
	}

	resp := &studentDefpb.GetStudentDefByTimeSlotsIDResponse{
		StudentDef: &sdResp,
	}

	return resp, nil
}
