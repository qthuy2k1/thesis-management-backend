package repository

import (
	"context"
	"database/sql"
)

type IReportingStageRepo interface {
	// CreateReportingStage creates a new reporting stage in db given by reporting stage model
	CreateReportingStage(ctx context.Context, clr ReportingStageInputRepo) error
	// GetReportingStage returns a reporting stage in db given by id
	GetReportingStage(ctx context.Context, id int) (ReportingStageOutputRepo, error)
	// CheckReportingStageExists checks whether the specified reporting stage exists by name
	IsReportingStageExists(ctx context.Context, title string) (bool, error)
	// UpdateReportingStage updates the specified rs by id
	UpdateReportingStage(ctx context.Context, id int, classroom ReportingStageInputRepo) error
	// DeleteReportingStage deletes a rs in db given by id
	DeleteReportingStage(ctx context.Context, id int) error
	// GetReportingStages returns a list of reporting stages in db with filter
	GetReportingStages(ctx context.Context) ([]ReportingStageOutputRepo, error)
}

type ReportingStageRepo struct {
	Database *sql.DB
}

func NewReportingStageRepo(db *sql.DB) IReportingStageRepo {
	return &ReportingStageRepo{Database: db}
}
