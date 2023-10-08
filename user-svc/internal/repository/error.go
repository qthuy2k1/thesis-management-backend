package repository

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrUserExisted    = errors.New("user already exists")
	ErrMemberNotFound = errors.New("member not found")
	ErrMemberExisted  = errors.New("member already exists")
)
