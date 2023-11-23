package repository

import (
	"context"
	"database/sql"
)

type IWaitingListRepo interface {
	// CreateWaitingList creates a new waiting list in db given by waiting list model
	CreateWaitingList(ctx context.Context, wt WaitingListInputRepo) error
	// GetWaitingList returns a waiting list in db given by id
	GetWaitingList(ctx context.Context, id int) (WaitingListOutputRepo, error)
	// CheckWaitingListExists checks whether the specified waiting list exists by name
	IsWaitingListExists(ctx context.Context, classroomID int, userID string) (bool, error)
	// UpdateWaitingList updates the specified rs by id
	UpdateWaitingList(ctx context.Context, id int, wt WaitingListInputRepo) error
	// DeleteWaitingList deletes a rs in db given by id
	DeleteWaitingList(ctx context.Context, id int) error
	// GetWaitingLists returns a list of waiting lists in db with filter
	GetWaitingListsOfClassroom(ctx context.Context, classroomID int) ([]WaitingListOutputRepo, error)
	// CheckUserInWaitingListOfClassroom returns a boolean indicating whether user is in waiting list
	CheckUserInWaitingListOfClassroom(ctx context.Context, userID string) (bool, int, error)
	// GetWaitingList returns a list of waiting_lists in db
	GetWaitingLists(ctx context.Context) ([]WaitingListOutputRepo, error)
}

type WaitingListRepo struct {
	Database *sql.DB
}

func NewWaitingListRepo(db *sql.DB) IWaitingListRepo {
	return &WaitingListRepo{Database: db}
}
