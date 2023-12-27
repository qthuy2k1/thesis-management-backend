package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/reporting-stage"

	"google.golang.org/grpc/codes"
)

var (
	ErrReportingStageNotFound  = errors.New("reporting stage not found")
	ErrInvalidReportingStageID = errors.New("invalid reporting stage ID")
	ErrServerError             = errors.New("internal server error")
	ErrReportingStageExisted   = errors.New("reporting stage already exists")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrReportingStageNotFound:
		return codes.NotFound, ErrReportingStageNotFound
	case service.ErrReportingStageExisted:
		return codes.AlreadyExists, ErrReportingStageExisted
	case ErrInvalidReportingStageID:
		return codes.InvalidArgument, ErrInvalidReportingStageID
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
