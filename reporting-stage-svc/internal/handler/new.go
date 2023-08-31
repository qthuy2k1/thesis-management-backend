package handler

import (
	reportingStagepb "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/internal/service"
)

type ReportingStageHdl struct {
	reportingStagepb.UnimplementedReportingStageServiceServer
	Service service.IReportingStageSvc
}

// NewReportingStageHdl returns the Handler struct that contains the Service
func NewReportingStageHdl(svc service.IReportingStageSvc) *ReportingStageHdl {
	return &ReportingStageHdl{Service: svc}
}
