package handler

import (
	commentpb "github.com/qthuy2k1/thesis-management-backend/comment-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/comment-svc/internal/service"
)

type CommentHdl struct {
	commentpb.UnimplementedCommentServiceServer
	Service service.ICommentSvc
}

// NewCommentHdl returns the Handler struct that contains the Service
func NewCommentHdl(svc service.ICommentSvc) *CommentHdl {
	return &CommentHdl{Service: svc}
}
