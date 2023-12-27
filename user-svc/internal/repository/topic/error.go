package repository

import "errors"

var (
	ErrTopicNotFound = errors.New("topic not found")
	ErrTopicExisted  = errors.New("topic already exists")
)
