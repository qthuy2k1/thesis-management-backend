package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service"

	"google.golang.org/grpc/codes"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidUserID = errors.New("invalid user ID")
	ErrServerError   = errors.New("internal server error")
	ErrUserExisted   = errors.New("user already exists")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrUserNotFound:
		return codes.NotFound, ErrUserNotFound
	case service.ErrUserExisted:
		return codes.AlreadyExists, ErrUserExisted
	case ErrInvalidUserID:
		return codes.InvalidArgument, ErrInvalidUserID
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
