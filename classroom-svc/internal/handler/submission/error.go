package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/submission"

	"google.golang.org/grpc/codes"
)

var (
	ErrSubmissionNotFound  = errors.New("submission not found")
	ErrInvalidSubmissionID = errors.New("invalid submission ID")
	ErrServerError         = errors.New("internal server error")
	ErrSubmissionExisted   = errors.New("submission already exists")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrSubmissionNotFound:
		return codes.NotFound, ErrSubmissionNotFound
	case service.ErrSubmissionExisted:
		return codes.AlreadyExists, ErrSubmissionExisted
	case ErrInvalidSubmissionID:
		return codes.InvalidArgument, ErrInvalidSubmissionID
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
