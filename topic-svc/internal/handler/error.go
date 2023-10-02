package handler

import (
	"errors"
	"log"

	repository "github.com/qthuy2k1/thesis-management-backend/topic-svc/internal/repository"

	"google.golang.org/grpc/codes"
)

var (
	ErrTopicNotFound      = errors.New("topic not found")
	ErrInvalidTopicID     = errors.New("invalid topic ID")
	ErrServerError        = errors.New("internal server error")
	ErrTopicExisted       = errors.New("topic already exists")
	ErrInvalidLimit       = errors.New("invalid limit value")
	ErrInvalidPage        = errors.New("invalid page value")
	ErrInvalidTitleSearch = errors.New("invalid title search value")
	ErrInvalidSortColumn  = errors.New("invalid sort column value")
	ErrInvalidIsDesc      = errors.New("invalid isDesc value")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case repository.ErrTopicNotFound:
		return codes.NotFound, ErrTopicNotFound
	case repository.ErrTopicExisted:
		return codes.AlreadyExists, ErrTopicExisted
	case ErrInvalidTopicID:
		return codes.InvalidArgument, ErrInvalidTopicID
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
