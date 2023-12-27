package repository

import "errors"

var (
	ErrAttachmentNotFound = errors.New("attachment not found")
	ErrAttachmentExisted  = errors.New("attachment already exists")
)
