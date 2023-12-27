package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/post"

	"google.golang.org/grpc/codes"
)

var (
	ErrPostNotFound  = errors.New("post not found")
	ErrInvalidPostID = errors.New("invalid post ID")
	ErrServerError   = errors.New("internal server error")
	ErrPostExisted   = errors.New("post already exists")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrPostNotFound:
		return codes.NotFound, ErrPostNotFound
	case service.ErrPostExisted:
		return codes.AlreadyExists, ErrPostExisted
	case ErrInvalidPostID:
		return codes.InvalidArgument, ErrInvalidPostID
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
