package repository

import "errors"

var (
	ErrCommiteeNotFound = errors.New("commitee not found")
	ErrCommiteeExisted  = errors.New("commitee already exists")

	ErrCommiteeUserDetailNotFound = errors.New("commitee user detail not found")
	ErrCommiteeUserDetailExisted  = errors.New("commitee user detail already exists")

	ErrRoomNotFound = errors.New("room not found")
	ErrRoomExisted  = errors.New("room already exists")

	ErrCouncilNotFound = errors.New("council not found")
	ErrCouncilExisted  = errors.New("council already exists")

	ErrTimeSlotsNotFound = errors.New("time slots not found")
	ErrTimeSlotsExisted  = errors.New("time slots already exists")

	ErrScheduleNotFound = errors.New("schedule not found")
	ErrScheduleExisted  = errors.New("schedule already exists")

	ErrThesisNotFound = errors.New("thesis not found")
	ErrThesisExisted  = errors.New("thesis already exists")
)
