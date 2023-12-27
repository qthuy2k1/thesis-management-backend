package handler

import (
	submissionpb "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/submission"
)

type SubmissionHdl struct {
	submissionpb.UnimplementedSubmissionServiceServer
	Service service.ISubmissionSvc
}

// NewSubmissionHdl returns the Handler struct that contains the Service
func NewSubmissionHdl(svc service.ISubmissionSvc) *SubmissionHdl {
	return &SubmissionHdl{Service: svc}
}
