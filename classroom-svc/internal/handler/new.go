package handler

import (
	classroompb "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service"
)

type ClassroomHdl struct {
	classroompb.UnimplementedClassroomServiceServer
	Service service.IClassroomSvc
}

// NewClassroomHdl returns the Handler struct that contains the Service
func NewClassroomHdl(svc service.IClassroomSvc) *ClassroomHdl {
	return &ClassroomHdl{Service: svc}
}
