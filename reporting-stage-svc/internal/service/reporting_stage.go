package service

import (
	"context"
	"errors"

	repository "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/internal/repository"
)

type ReportingStageInputSvc struct {
	Name        string
	Description string
}

// CreateReportingStage creates a new reporting stage in db given by reporting stage model
func (s *ReportingStageSvc) CreateReportingStage(ctx context.Context, p ReportingStageInputSvc) error {
	pRepo := repository.ReportingStageInputRepo{
		Name:        p.Name,
		Description: p.Description,
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
		Name:        p.Name,
		Description: p.Description,
	}, nil
}

// UpdateReportingStage updates the specified reporting stage by id
func (s *ReportingStageSvc) UpdateReportingStage(ctx context.Context, id int, reportingStage ReportingStageInputSvc) error {
	if err := s.Repository.UpdateReportingStage(ctx, id, repository.ReportingStageInputRepo{
		Name:        reportingStage.Name,
		Description: reportingStage.Description,
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
	Name        string
	Description string
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
			Name:        p.Name,
			Description: p.Description,
		})
	}

	return psSvc, nil
}
