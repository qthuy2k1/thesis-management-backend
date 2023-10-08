package handler

import (
	"context"
	"log"

	reportingStagepb "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/internal/service"
	"google.golang.org/grpc/status"
)

// CreateReportingStage retrieves a reportingStage request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *ReportingStageHdl) CreateReportingStage(ctx context.Context, req *reportingStagepb.CreateReportingStageRequest) (*reportingStagepb.CreateReportingStageResponse, error) {
	log.Println("calling insert reportingStage...")
	p, err := validateAndConvertReportingStage(req.ReportingStage)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.CreateReportingStage(ctx, p); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &reportingStagepb.CreateReportingStageResponse{
		Response: &reportingStagepb.CommonReportingStageResponse{
			StatusCode: 201,
			Message:    "Created",
		},
	}

	return resp, nil
}

// GetReportingStage returns a reportingStage in db given by id
func (h *ReportingStageHdl) GetReportingStage(ctx context.Context, req *reportingStagepb.GetReportingStageRequest) (*reportingStagepb.GetReportingStageResponse, error) {
	log.Println("calling get reportingStage...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	p, err := h.Service.GetReportingStage(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := reportingStagepb.ReportingStageResponse{
		Id:          int64(p.ID),
		Label:       p.Label,
		Description: p.Description,
		Value:       p.Value,
	}

	resp := &reportingStagepb.GetReportingStageResponse{
		Response: &reportingStagepb.CommonReportingStageResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		ReportingStage: &pResp,
	}
	return resp, nil
}

func (c *ReportingStageHdl) UpdateReportingStage(ctx context.Context, req *reportingStagepb.UpdateReportingStageRequest) (*reportingStagepb.UpdateReportingStageResponse, error) {
	log.Println("calling update reportingStage...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	p, err := validateAndConvertReportingStage(req.ReportingStage)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdateReportingStage(ctx, int(req.GetId()), service.ReportingStageInputSvc{
		Label:       p.Label,
		Description: p.Description,
		Value:       p.Value,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &reportingStagepb.UpdateReportingStageResponse{
		Response: &reportingStagepb.CommonReportingStageResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *ReportingStageHdl) DeleteReportingStage(ctx context.Context, req *reportingStagepb.DeleteReportingStageRequest) (*reportingStagepb.DeleteReportingStageResponse, error) {
	log.Println("calling delete reportingStage...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeleteReportingStage(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &reportingStagepb.DeleteReportingStageResponse{
		Response: &reportingStagepb.CommonReportingStageResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *ReportingStageHdl) GetReportingStages(ctx context.Context, req *reportingStagepb.GetReportingStagesRequest) (*reportingStagepb.GetReportingStagesResponse, error) {
	log.Println("calling get all reportingStages...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	ps, err := h.Service.GetReportingStages(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*reportingStagepb.ReportingStageResponse
	for _, p := range ps {
		psResp = append(psResp, &reportingStagepb.ReportingStageResponse{
			Id:          int64(p.ID),
			Label:       p.Label,
			Description: p.Description,
			Value:       p.Value,
		})
	}

	return &reportingStagepb.GetReportingStagesResponse{
		Response: &reportingStagepb.CommonReportingStageResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		ReportingStages: psResp,
	}, nil
}

func validateAndConvertReportingStage(pbReportingStage *reportingStagepb.ReportingStageInput) (service.ReportingStageInputSvc, error) {
	if err := pbReportingStage.Validate(); err != nil {
		return service.ReportingStageInputSvc{}, err
	}

	return service.ReportingStageInputSvc{
		Label:       pbReportingStage.Label,
		Description: pbReportingStage.Description,
		Value:       pbReportingStage.Value,
	}, nil
}
