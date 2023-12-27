package service

import (
	"context"
	"errors"

	repository "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/reporting-stage"
)

type ReportingStageInputSvc struct {
	Label       string
	Description string
	Value       string
}

// CreateReportingStage creates a new reporting stage in db given by reporting stage model
func (s *ReportingStageSvc) CreateReportingStage(ctx context.Context, p ReportingStageInputSvc) error {
	pRepo := repository.ReportingStageInputRepo{
		Label:       p.Label,
		Description: p.Description,
		Value:       p.Value,
	}

	if err := s.Repository.CreateReportingStage(ctx, pRepo); err != nil {
		if errors.Is(err, repository.ErrReportingStageExisted) {
			return ErrReportingStageExisted
		}
		return err
	}

	return nil
}

// GetReportingStage returns a reporting stage in db given by id
func (s *ReportingStageSvc) GetReportingStage(ctx context.Context, id int) (ReportingStageOutputSvc, error) {
	p, err := s.Repository.GetReportingStage(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrReportingStageNotFound) {
			return ReportingStageOutputSvc{}, ErrReportingStageNotFound
		}
		return ReportingStageOutputSvc{}, err
	}

	return ReportingStageOutputSvc{
		ID:          p.ID,
		Label:       p.Label,
		Description: p.Description,
		Value:       p.Value,
	}, nil
}

// UpdateReportingStage updates the specified reporting stage by id
func (s *ReportingStageSvc) UpdateReportingStage(ctx context.Context, id int, reportingStage ReportingStageInputSvc) error {
	if err := s.Repository.UpdateReportingStage(ctx, id, repository.ReportingStageInputRepo{
		Label:       reportingStage.Label,
		Description: reportingStage.Description,
		Value:       reportingStage.Value,
	}); err != nil {
		if errors.Is(err, repository.ErrReportingStageNotFound) {
			return ErrReportingStageNotFound
		}
		return err
	}

	return nil
}

// DeleteReportingStage deletes a reporting stage in db given by id
func (s *ReportingStageSvc) DeleteReportingStage(ctx context.Context, id int) error {
	if err := s.Repository.DeleteReportingStage(ctx, id); err != nil {
		if errors.Is(err, repository.ErrReportingStageNotFound) {
			return ErrReportingStageNotFound
		}
		return err
	}

	return nil
}

type ReportingStageOutputSvc struct {
	ID          int
	Label       string
	Description string
	Value       string
}

// GetReportingStages returns a list of reporting stages in db with filter
func (s *ReportingStageSvc) GetReportingStages(ctx context.Context) ([]ReportingStageOutputSvc, error) {
	psRepo, err := s.Repository.GetReportingStages(ctx)
	if err != nil {
		return nil, err
	}

	var psSvc []ReportingStageOutputSvc
	for _, p := range psRepo {
		psSvc = append(psSvc, ReportingStageOutputSvc{
			ID:          p.ID,
			Label:       p.Label,
			Description: p.Description,
			Value:       p.Value,
		})
	}

	return psSvc, nil
}
