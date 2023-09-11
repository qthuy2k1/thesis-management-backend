package repository

import "errors"

var (
	ErrSubmissionNotFound = errors.New("submission not found")
	ErrSubmissionExisted  = errors.New("submission already exists")
)
