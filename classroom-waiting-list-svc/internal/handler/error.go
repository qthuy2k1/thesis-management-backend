package handler

import (
	"errors"
	"log"

	service "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/internal/service"

	"google.golang.org/grpc/codes"
)

var (
	ErrWaitingListNotFound           = errors.New("waiting list not found")
	ErrInvalidWaitingListID          = errors.New("invalid waiting list ID")
	ErrServerError                   = errors.New("internal server error")
	ErrWaitingListExisted            = errors.New("waiting list already exists")
	ErrWaitingListCreatedMoreThanTwo = errors.New("You don't send more than 2 requirements")
)

// ConvertCtrlError compares the error return with the error in controller and returns the corresponding ErrorResponse
func convertCtrlError(err error) (codes.Code, error) {
	switch err {
	case service.ErrWaitingListNotFound:
		return codes.NotFound, ErrWaitingListNotFound
	case service.ErrWaitingListExisted:
		return codes.AlreadyExists, ErrWaitingListExisted
	case ErrInvalidWaitingListID:
		return codes.InvalidArgument, ErrInvalidWaitingListID
	case service.ErrWaitingListCreatedMoreThanTwo:
		return codes.InvalidArgument, ErrWaitingListCreatedMoreThanTwo
	default:
		log.Println("handler err: ", err)
		return codes.Internal, ErrServerError
	}
}
