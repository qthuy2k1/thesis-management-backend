package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service"

	"google.golang.org/grpc/codes"
)

var (
	ErrClassroomNotFound  = errors.New("classroom not found")
	ErrInvalidClassroomID = errors.New("invalid classroom ID")
	ErrServerError        = errors.New("internal server error")
	ErrClassroomExisted   = errors.New("classroom already exists")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrClassroomNotFound:
		return codes.NotFound, ErrClassroomNotFound
	case service.ErrClassroomExisted:
		return codes.AlreadyExists, ErrClassroomExisted
	case ErrInvalidClassroomID:
		return codes.InvalidArgument, ErrInvalidClassroomID
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
