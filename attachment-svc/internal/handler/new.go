package handler

import (
	attachmentpb "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/attachment-svc/internal/service"
)

type AttachmentHdl struct {
	attachmentpb.UnimplementedAttachmentServiceServer
	Service service.IAttachmentSvc
}

// NewAttachmentHdl returns the Handler struct that contains the Service
func NewAttachmentHdl(svc service.IAttachmentSvc) *AttachmentHdl {
	return &AttachmentHdl{Service: svc}
}
