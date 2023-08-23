package repository

import "errors"

var (
	ErrPostNotFound = errors.New("post not found")
	ErrPostExisted  = errors.New("post already exists")
)
