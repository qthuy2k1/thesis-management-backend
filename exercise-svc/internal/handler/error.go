package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/exercise-svc/internal/service"

	"google.golang.org/grpc/codes"
)

var (
	ErrExerciseNotFound  = errors.New("exercise not found")
	ErrInvalidExerciseID = errors.New("invalid exercise ID")
	ErrServerError       = errors.New("internal server error")
	ErrExerciseExisted   = errors.New("exercise already exists")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrExerciseNotFound:
		return codes.NotFound, ErrExerciseNotFound
	case service.ErrExerciseExisted:
		return codes.AlreadyExists, ErrExerciseExisted
	case ErrInvalidExerciseID:
		return codes.InvalidArgument, ErrInvalidExerciseID
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
