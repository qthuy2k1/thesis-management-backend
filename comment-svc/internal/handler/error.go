package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/comment-svc/internal/service"

	"google.golang.org/grpc/codes"
)

var (
	ErrCommentNotFound    = errors.New("comment not found")
	ErrInvalidCommentID   = errors.New("invalid comment ID")
	ErrServerError        = errors.New("internal server error")
	ErrCommentExisted     = errors.New("comment already exists")
	ErrInvalidLimit       = errors.New("invalid limit value")
	ErrInvalidPage        = errors.New("invalid page value")
	ErrInvalidTitleSearch = errors.New("invalid title search value")
	ErrInvalidSortColumn  = errors.New("invalid sort column value")
	ErrInvalidIsDesc      = errors.New("invalid isDesc value")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrCommentNotFound:
		return codes.NotFound, ErrCommentNotFound
	case service.ErrCommentExisted:
		return codes.AlreadyExists, ErrCommentExisted
	case ErrInvalidCommentID:
		return codes.InvalidArgument, ErrInvalidCommentID
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
