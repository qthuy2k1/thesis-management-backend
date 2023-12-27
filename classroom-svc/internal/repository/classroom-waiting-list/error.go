package repository

import "errors"

var (
	ErrWaitingListNotFound           = errors.New("waiting list not found")
	ErrWaitingListExisted            = errors.New("waiting list already exists")
	ErrWaitingListCreatedMoreThanTwo = errors.New("You don't send more than 2 requirements")
)
