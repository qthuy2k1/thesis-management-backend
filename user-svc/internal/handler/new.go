package handler

import (
	userpb "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service"
)

type UserHdl struct {
	userpb.UnimplementedUserServiceServer
	Service service.IUserSvc
}

// NewUserHdl returns the Handler struct that contains the Service
func NewUserHdl(svc service.IUserSvc) *UserHdl {
	return &UserHdl{Service: svc}
}
