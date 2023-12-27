package service

import "errors"

var (
	ErrCommentNotFound = errors.New("comment not found")
	ErrCommentExisted  = errors.New("comment already exists")
)
