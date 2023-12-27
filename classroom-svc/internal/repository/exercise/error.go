package repository

import "errors"

var (
	ErrExerciseNotFound = errors.New("exercise not found")
	ErrExerciseExisted  = errors.New("exercise already exists")
)
