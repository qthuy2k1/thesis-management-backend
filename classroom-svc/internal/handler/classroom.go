package handler

import (
	"context"
	"log"

	classroompb "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateClassroom retrieves a classroom request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *ClassroomHdl) CreateClassroom(ctx context.Context, req *classroompb.CreateClassroomRequest) (*classroompb.CreateClassroomResponse, error) {
	log.Println("calling insert classroom...")
	clr, err := validateAndConvertClassroom(req.Classroom)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.CreateClassroom(ctx, clr); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &classroompb.CreateClassroomResponse{
		StatusCode: 201,
		Message:    "Created",
	}

	return resp, nil
}

// GetClassroom returns a classroom in db given by id
func (h *ClassroomHdl) GetClassroom(ctx context.Context, req *classroompb.GetClassroomRequest) (*classroompb.GetClassroomResponse, error) {
	log.Println("calling get classroom...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	clr, err := h.Service.GetClassroom(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	clrResp := classroompb.ClassroomResponse{
		Id:          int32(clr.ID),
		Title:       clr.Title,
		Description: clr.Description,
		Status:      clr.Status,
		CreatedAt:   timestamppb.New(clr.CreatedAt),
		UpdatedAt:   timestamppb.New(clr.UpdatedAt),
	}

	resp := &classroompb.GetClassroomResponse{
		StatusCode: 200,
		Message:    "OK",
		Classroom:  &clrResp,
	}
	return resp, nil
}

func validateAndConvertClassroom(pbClassroom *classroompb.ClassroomInput) (service.ClassroomInputSvc, error) {
	if err := pbClassroom.Validate(); err != nil {
		return service.ClassroomInputSvc{}, err
	}

	return service.ClassroomInputSvc{
		Title:       pbClassroom.Title,
		Description: pbClassroom.Description,
		Status:      pbClassroom.Status,
	}, nil
}
