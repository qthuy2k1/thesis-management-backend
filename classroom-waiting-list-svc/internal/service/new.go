package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/internal/repository"
)

type IWaitingListSvc interface {
	// CreateWaitingList creates a new waiting list in db given by waiting list model
	CreateWaitingList(ctx context.Context, p WaitingListInputSvc) error
	// GetWaitingList returns a waiting list in db given by id
	GetWaitingList(ctx context.Context, id int) (WaitingListOutputSvc, error)
	// UpdateWaitingList updates the specified waiting lis by id
	UpdateWaitingList(ctx context.Context, id int, classroom WaitingListInputSvc) error
	// DeleteWaitingList deletes a classroom in db given by id
	DeleteWaitingList(ctx context.Context, id int) error
	// GetWaitingListsOfClassroom returns a list of waiting lists in a classroom
	GetWaitingListsOfClassroom(ctx context.Context, classroomID int) ([]WaitingListOutputSvc, error)
}

type WaitingListSvc struct {
	Repository repository.IWaitingListRepo
}

func NewWaitingListSvc(pRepo repository.IWaitingListRepo) IWaitingListSvc {
	return &WaitingListSvc{Repository: pRepo}
}
