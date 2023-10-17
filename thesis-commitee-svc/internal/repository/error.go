package repository

import "errors"

var (
	ErrCommiteeNotFound = errors.New("commitee not found")
	ErrCommiteeExisted  = errors.New("commitee already exists")

	ErrCommiteeUserDetailNotFound = errors.New("commitee user detail not found")
	ErrCommiteeUserDetailExisted  = errors.New("commitee user detail already exists")
)
