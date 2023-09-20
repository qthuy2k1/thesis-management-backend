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
		Response: &classroompb.CommonClassroomResponse{
			StatusCode: 200,
			Message:    "OK",
		},
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
		Id:            int32(clr.ID),
		Title:         clr.Title,
		Description:   clr.Description,
		Status:        clr.Status,
		LecturerId:    clr.LecturerID,
		CodeClassroom: clr.CodeClassroom,
		TopicTags:     clr.TopicTags,
		Quantity:      int32(clr.Quantity),
		CreatedAt:     timestamppb.New(clr.CreatedAt),
		UpdatedAt:     timestamppb.New(clr.UpdatedAt),
	}

	resp := &classroompb.GetClassroomResponse{
		Response: &classroompb.CommonClassroomResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Classroom: &clrResp,
	}
	return resp, nil
}

func (h *ClassroomHdl) CheckClassroomExists(ctx context.Context, req *classroompb.CheckClassroomExistsRequest) (*classroompb.CheckClassroomExistsResponse, error) {
	log.Println("calling check classroom exists...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	exists, err := h.Service.CheckClassroomExists(ctx, int(req.GetClassroomID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &classroompb.CheckClassroomExistsResponse{
		Exists: exists,
	}, nil
}

func (c *ClassroomHdl) UpdateClassroom(ctx context.Context, req *classroompb.UpdateClassroomRequest) (*classroompb.UpdateClassroomResponse, error) {
	log.Println("calling update classroom...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	clr, err := validateAndConvertClassroom(req.Classroom)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateClassroom(ctx, int(req.GetId()), service.ClassroomInputSvc{
		Title:         clr.Title,
		Description:   clr.Description,
		Status:        clr.Status,
		LecturerID:    clr.LecturerID,
		CodeClassroom: clr.CodeClassroom,
		TopicTags:     clr.TopicTags,
		Quantity:      clr.Quantity,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &classroompb.UpdateClassroomResponse{
		Response: &classroompb.CommonClassroomResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *ClassroomHdl) DeleteClassroom(ctx context.Context, req *classroompb.DeleteClassroomRequest) (*classroompb.DeleteClassroomResponse, error) {
	log.Println("calling delete classroom...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteClassroom(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &classroompb.DeleteClassroomResponse{
		Response: &classroompb.CommonClassroomResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *ClassroomHdl) GetClassrooms(ctx context.Context, req *classroompb.GetClassroomsRequest) (*classroompb.GetClassroomsResponse, error) {
	log.Println("calling get all classrooms...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	filter := service.ClassroomFilterSvc{
		Limit:       int(req.GetLimit()),
		Page:        int(req.GetPage()),
		TitleSearch: req.GetTitleSearch(),
		SortColumn:  req.GetSortColumn(),
		SortOrder:   req.GetSortOrder(),
	}

	clrs, count, err := h.Service.GetClassrooms(ctx, filter)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var clrsResp []*classroompb.ClassroomResponse
	for _, c := range clrs {
		clrsResp = append(clrsResp, &classroompb.ClassroomResponse{
			Id:            int32(c.ID),
			Title:         c.Title,
			Description:   c.Description,
			Status:        c.Status,
			LecturerId:    c.LecturerID,
			CodeClassroom: c.CodeClassroom,
			TopicTags:     c.TopicTags,
			Quantity:      int32(c.Quantity),
			CreatedAt:     timestamppb.New(c.CreatedAt),
			UpdatedAt:     timestamppb.New(c.UpdatedAt),
		})
	}

	return &classroompb.GetClassroomsResponse{
		Response: &classroompb.CommonClassroomResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Classrooms: clrsResp,
		TotalCount: int32(count),
	}, nil
}

func validateAndConvertClassroom(pbClassroom *classroompb.ClassroomInput) (service.ClassroomInputSvc, error) {
	if err := pbClassroom.Validate(); err != nil {
		return service.ClassroomInputSvc{}, err
	}

	return service.ClassroomInputSvc{
		Title:         pbClassroom.Title,
		Description:   pbClassroom.Description,
		Status:        pbClassroom.Status,
		LecturerID:    pbClassroom.LecturerId,
		CodeClassroom: pbClassroom.CodeClassroom,
		TopicTags:     pbClassroom.TopicTags,
		Quantity:      int(pbClassroom.Quantity),
	}, nil
}
