package service

import (
	"context"
	"errors"
	"time"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/classroom-waiting-list"
)

type WaitingListInputSvc struct {
	ClassroomID int
	UserID      string
	IsDefense   bool
	Status      string
}

type WaitingListOutputSvc struct {
	ID          int
	ClassroomID int
	UserID      string
	IsDefense   bool
	Status      string
	CreatedAt   time.Time
}

// CreateWaitingList creates a new waiting list in db given by waiting list model
func (s *WaitingListSvc) CreateWaitingList(ctx context.Context, wt WaitingListInputSvc) error {
	pRepo := repository.WaitingListInputRepo{
		ClassroomID: wt.ClassroomID,
		UserID:      wt.UserID,
		IsDefense:   wt.IsDefense,
		Status:      wt.Status,
	}

	if err := s.Repository.CreateWaitingList(ctx, pRepo); err != nil {
		if errors.Is(err, repository.ErrWaitingListExisted) {
			return ErrWaitingListExisted
		}

		if errors.Is(err, repository.ErrWaitingListCreatedMoreThanTwo) {
			return ErrWaitingListCreatedMoreThanTwo
		}

		return err
	}

	return nil
}

// GetWaitingList returns a waiting list in db given by id
func (s *WaitingListSvc) GetWaitingList(ctx context.Context, id int) (WaitingListOutputSvc, error) {
	wt, err := s.Repository.GetWaitingList(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrWaitingListNotFound) {
			return WaitingListOutputSvc{}, ErrWaitingListNotFound
		}
		return WaitingListOutputSvc{}, err
	}

	return WaitingListOutputSvc{
		ID:          wt.ID,
		ClassroomID: wt.ClassroomID,
		UserID:      wt.UserID,
		IsDefense:   wt.IsDefense,
		Status:      wt.Status,
		CreatedAt:   wt.CreatedAt,
	}, nil
}

// UpdateWaitingList updates the specified waiting list by id
func (s *WaitingListSvc) UpdateWaitingList(ctx context.Context, id int, waitingList WaitingListInputSvc) error {
	if err := s.Repository.UpdateWaitingList(ctx, id, repository.WaitingListInputRepo{
		ClassroomID: waitingList.ClassroomID,
		UserID:      waitingList.UserID,
		IsDefense:   waitingList.IsDefense,
		Status:      waitingList.Status,
	}); err != nil {
		if errors.Is(err, repository.ErrWaitingListNotFound) {
			return ErrWaitingListNotFound
		}
		return err
	}

	return nil
}

// DeleteWaitingList deletes a waiting list in db given by id
func (s *WaitingListSvc) DeleteWaitingList(ctx context.Context, id int) error {
	if err := s.Repository.DeleteWaitingList(ctx, id); err != nil {
		if errors.Is(err, repository.ErrWaitingListNotFound) {
			return ErrWaitingListNotFound
		}
		return err
	}

	return nil
}

// GetWaitingListsOfClassroom returns a list of waiting lists in a classroom
func (s *WaitingListSvc) GetWaitingListsOfClassroom(ctx context.Context, classroomID int) ([]WaitingListOutputSvc, error) {
	wtsRepo, err := s.Repository.GetWaitingListsOfClassroom(ctx, classroomID)
	if err != nil {
		return nil, err
	}

	var wtsSvc []WaitingListOutputSvc
	for _, wt := range wtsRepo {
		wtsSvc = append(wtsSvc, WaitingListOutputSvc{
			ID:          wt.ID,
			ClassroomID: wt.ClassroomID,
			UserID:      wt.UserID,
			IsDefense:   wt.IsDefense,
			Status:      wt.Status,
			CreatedAt:   wt.CreatedAt,
		})
	}

	return wtsSvc, nil
}

// GetWaitingLists returns a list of waiting lists in a classroom
func (s *WaitingListSvc) GetWaitingLists(ctx context.Context) ([]WaitingListOutputSvc, error) {
	wtsRepo, err := s.Repository.GetWaitingLists(ctx)
	if err != nil {
		return nil, err
	}

	var wtsSvc []WaitingListOutputSvc
	for _, wt := range wtsRepo {
		wtsSvc = append(wtsSvc, WaitingListOutputSvc{
			ID:          wt.ID,
			ClassroomID: wt.ClassroomID,
			UserID:      wt.UserID,
			IsDefense:   wt.IsDefense,
			Status:      wt.Status,
			CreatedAt:   wt.CreatedAt,
		})
	}

	return wtsSvc, nil
}

// CheckUserInWaitingListOfClassroom returns a boolean indicating whether user is in waiting list
func (s *WaitingListSvc) CheckUserInWaitingListOfClassroom(ctx context.Context, userID string, classroomID int) (bool, int, error) {
	isIn, err := s.Repository.IsWaitingListExists(ctx, classroomID, userID)
	if err != nil {
		return false, 0, err
	}

	return isIn, classroomID, nil
}

// GetWaitingList returns a waiting_list in db given by id
func (s *WaitingListSvc) GetWaitingListByUser(ctx context.Context, userID string) (WaitingListOutputSvc, error) {
	wt, err := s.Repository.GetWaitingListByUser(ctx, userID)
	if err != nil {
		if errors.Is(err, repository.ErrWaitingListNotFound) {
			return WaitingListOutputSvc{}, ErrWaitingListNotFound
		}
		return WaitingListOutputSvc{}, err
	}

	return WaitingListOutputSvc{
		ID:          wt.ID,
		ClassroomID: wt.ClassroomID,
		UserID:      wt.UserID,
		IsDefense:   wt.IsDefense,
		Status:      wt.Status,
		CreatedAt:   wt.CreatedAt,
	}, nil
}
