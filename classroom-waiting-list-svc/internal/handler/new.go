package handler

import (
	waitingListpb "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/internal/service"
)

type WaitingListHdl struct {
	waitingListpb.UnimplementedWaitingListServiceServer
	Service service.IWaitingListSvc
}

// NewWaitingListHdl returns the Handler struct that contains the Service
func NewWaitingListHdl(svc service.IWaitingListSvc) *WaitingListHdl {
	return &WaitingListHdl{Service: svc}
}
