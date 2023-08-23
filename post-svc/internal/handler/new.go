package handler

import (
	postpb "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/post-svc/internal/service"
)

type PostHdl struct {
	postpb.UnimplementedPostServiceServer
	Service service.IPostSvc
}

// NewPostHdl returns the Handler struct that contains the Service
func NewPostHdl(svc service.IPostSvc) *PostHdl {
	return &PostHdl{Service: svc}
}
