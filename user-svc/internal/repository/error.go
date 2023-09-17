package repository

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUserExisted  = errors.New("user already exists")
)