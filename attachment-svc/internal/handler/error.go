package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/attachment-svc/internal/service"

	"google.golang.org/grpc/codes"
)

var (
	ErrAttachmentNotFound  = errors.New("attachment not found")
	ErrInvalidAttachmentID = errors.New("invalid attachment ID")
	ErrServerError         = errors.New("internal server error")
	ErrAttachmentExisted   = errors.New("attachment already exists")
	ErrInvalidLimit        = errors.New("invalid limit value")
	ErrInvalidPage         = errors.New("invalid page value")
	ErrInvalidTitleSearch  = errors.New("invalid title search value")
	ErrInvalidSortColumn   = errors.New("invalid sort column value")
	ErrInvalidIsDesc       = errors.New("invalid isDesc value")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrAttachmentNotFound:
		return codes.NotFound, ErrAttachmentNotFound
	case service.ErrAttachmentExisted:
		return codes.AlreadyExists, ErrAttachmentExisted
	case ErrInvalidAttachmentID:
		return codes.InvalidArgument, ErrInvalidAttachmentID
	case ErrInvalidLimit:
		return codes.InvalidArgument, ErrInvalidLimit
	case ErrInvalidPage:
		return codes.InvalidArgument, ErrInvalidPage
	case ErrInvalidTitleSearch:
		return codes.InvalidArgument, ErrInvalidTitleSearch
	case ErrInvalidSortColumn:
		return codes.InvalidArgument, ErrInvalidSortColumn
	case ErrInvalidIsDesc:
		return codes.InvalidArgument, ErrInvalidIsDesc
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
