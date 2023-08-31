package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/internal/repository"
)

type IReportingStageSvc interface {
	// CreateReportingStage creates a new reporting stage in db given by post model
	CreateReportingStage(ctx context.Context, p ReportingStageInputSvc) error
	// GetReportingStage returns a reporting stage in db given by id
	GetReportingStage(ctx context.Context, id int) (ReportingStageOutputSvc, error)
	// UpdateReportingStage updates the specified classroom by id
	UpdateReportingStage(ctx context.Context, id int, classroom ReportingStageInputSvc) error
	// DeleteReportingStage deletes a classroom in db given by id
	DeleteReportingStage(ctx context.Context, id int) error
	// GetReportingStages returns a list of reporting stages in db with filter
	GetReportingStages(ctx context.Context) ([]ReportingStageOutputSvc, error)
}

type ReportingStageSvc struct {
	Repository repository.IReportingStageRepo
}

func NewReportingStageSvc(pRepo repository.IReportingStageRepo) IReportingStageSvc {
	return &ReportingStageSvc{Repository: pRepo}
}
