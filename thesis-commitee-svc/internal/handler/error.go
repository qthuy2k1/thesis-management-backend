package handler

import (
	"errors"
	"log"

	repository "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/internal/repository"

	"google.golang.org/grpc/codes"
)

var (
	ErrCommiteeNotFound  = errors.New("commitee not found")
	ErrInvalidCommiteeID = errors.New("invalid commitee ID")
	ErrServerError       = errors.New("internal server error")
	ErrCommiteeExisted   = errors.New("commitee already exists")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case repository.ErrCommiteeNotFound:
		return codes.NotFound, ErrCommiteeNotFound
	case repository.ErrCommiteeExisted:
		return codes.AlreadyExists, ErrCommiteeExisted
	case ErrInvalidCommiteeID:
		return codes.InvalidArgument, ErrInvalidCommiteeID
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
