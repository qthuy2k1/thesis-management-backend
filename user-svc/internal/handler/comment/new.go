package handler

import (
	commentpb "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service/comment"
)

type CommentHdl struct {
	commentpb.UnimplementedCommentServiceServer
	Service service.ICommentSvc
}

// NewCommentHdl returns the Handler struct that contains the Service
func NewCommentHdl(svc service.ICommentSvc) *CommentHdl {
	return &CommentHdl{Service: svc}
}
