package handler

import (
	"context"
	"log"

	waitingListpb "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/internal/service"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateWaitingList retrieves a waitingList request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *WaitingListHdl) CreateWaitingList(ctx context.Context, req *waitingListpb.CreateWaitingListRequest) (*waitingListpb.CreateWaitingListResponse, error) {
	log.Println("calling insert waitingList...")
	p, err := validateAndConvertWaitingList(req.WaitingList)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.CreateWaitingList(ctx, p); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &waitingListpb.CreateWaitingListResponse{
		Response: &waitingListpb.CommonWaitingListResponse{
			StatusCode: 201,
			Message:    "Created",
		},
	}

	return resp, nil
}

// GetWaitingList returns a waitingList in db given by id
func (h *WaitingListHdl) GetWaitingList(ctx context.Context, req *waitingListpb.GetWaitingListRequest) (*waitingListpb.GetWaitingListResponse, error) {
	log.Println("calling get waitingList...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	wt, err := h.Service.GetWaitingList(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := waitingListpb.WaitingListResponse{
		Id:          int64(wt.ID),
		ClassroomID: int64(wt.ClassroomID),
		UserID:      wt.UserID,
		IsDefense:   wt.IsDefense,
		CreatedAt:   timestamppb.New(wt.CreatedAt),
	}

	resp := &waitingListpb.GetWaitingListResponse{
		Response: &waitingListpb.CommonWaitingListResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		WaitingList: &pResp,
	}
	return resp, nil
}

func (c *WaitingListHdl) UpdateWaitingList(ctx context.Context, req *waitingListpb.UpdateWaitingListRequest) (*waitingListpb.UpdateWaitingListResponse, error) {
	log.Println("calling update waitingList...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	wt, err := validateAndConvertWaitingList(req.WaitingList)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateWaitingList(ctx, int(req.GetId()), service.WaitingListInputSvc{
		ClassroomID: wt.ClassroomID,
		UserID:      wt.UserID,
		IsDefense:   wt.IsDefense,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &waitingListpb.UpdateWaitingListResponse{
		Response: &waitingListpb.CommonWaitingListResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *WaitingListHdl) DeleteWaitingList(ctx context.Context, req *waitingListpb.DeleteWaitingListRequest) (*waitingListpb.DeleteWaitingListResponse, error) {
	log.Println("calling delete waitingList...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteWaitingList(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &waitingListpb.DeleteWaitingListResponse{
		Response: &waitingListpb.CommonWaitingListResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *WaitingListHdl) GetWaitingListsOfClassroom(ctx context.Context, req *waitingListpb.GetWaitingListsRequest) (*waitingListpb.GetWaitingListsResponse, error) {
	log.Println("calling get all waitingLists...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	wts, err := h.Service.GetWaitingListsOfClassroom(ctx, int(req.GetClassroomID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var wtsResp []*waitingListpb.WaitingListResponse
	for _, wt := range wts {
		wtsResp = append(wtsResp, &waitingListpb.WaitingListResponse{
			Id:          int64(wt.ID),
			ClassroomID: int64(wt.ClassroomID),
			UserID:      wt.UserID,
			IsDefense:   wt.IsDefense,
			CreatedAt:   timestamppb.New(wt.CreatedAt),
		})
	}

	return &waitingListpb.GetWaitingListsResponse{
		Response: &waitingListpb.CommonWaitingListResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		WaitingLists: wtsResp,
	}, nil
}

func (h *WaitingListHdl) CheckUserInWaitingListOfClassroom(ctx context.Context, req *waitingListpb.CheckUserInWaitingListClassroomRequest) (*waitingListpb.CheckUserInWaitingListClassroomResponse, error) {
	isIn, classroomID, err := h.Service.CheckUserInWaitingListOfClassroom(ctx, req.GetUserID())
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &waitingListpb.CheckUserInWaitingListClassroomResponse{
		IsIn:        isIn,
		ClassroomID: int64(classroomID),
	}, nil
}

func validateAndConvertWaitingList(pbWaitingList *waitingListpb.WaitingListInput) (service.WaitingListInputSvc, error) {
	if err := pbWaitingList.Validate(); err != nil {
		return service.WaitingListInputSvc{}, err
	}

	return service.WaitingListInputSvc{
		ClassroomID: int(pbWaitingList.GetClassroomID()),
		UserID:      pbWaitingList.GetUserID(),
		IsDefense:   pbWaitingList.IsDefense,
	}, nil
}
