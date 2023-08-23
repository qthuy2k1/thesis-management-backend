package repository

import "errors"

var (
	ErrClassroomNotFound = errors.New("classroom not found")
	ErrClassroomExisted  = errors.New("classroom already exists")
)
