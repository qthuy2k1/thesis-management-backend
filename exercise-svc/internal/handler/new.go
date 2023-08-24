package handler

import (
	exercisepb "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/exercise-svc/internal/service"
)

type ExerciseHdl struct {
	exercisepb.UnimplementedExerciseServiceServer
	Service service.IExerciseSvc
}

// NewExerciseHdl returns the Handler struct that contains the Service
func NewExerciseHdl(svc service.IExerciseSvc) *ExerciseHdl {
	return &ExerciseHdl{Service: svc}
}
