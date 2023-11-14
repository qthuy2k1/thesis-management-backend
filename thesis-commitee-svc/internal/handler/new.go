package handler

import (
	commiteepb "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	repository "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/internal/repository"
)

type CommiteeHdl struct {
	commiteepb.UnimplementedCommiteeServiceServer
	commiteepb.UnimplementedScheduleServiceServer
	Repository repository.ICommiteeRepo
}

// NewCommiteeHdl returns the Handler struct that contains the Service
func NewCommiteeHdl(repo repository.ICommiteeRepo) *CommiteeHdl {
	return &CommiteeHdl{Repository: repo}
}
