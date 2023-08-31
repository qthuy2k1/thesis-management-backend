package repository

import "errors"

var (
	ErrReportingStageNotFound = errors.New("reporting stage not found")
	ErrReportingStageExisted  = errors.New("reporting stage already exists")
)
