package handler

import (
	reportingStagepb "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/reporting-stage"
)

type ReportingStageHdl struct {
	reportingStagepb.UnimplementedReportingStageServiceServer
	Service service.IReportingStageSvc
}

// NewReportingStageHdl returns the Handler struct that contains the Service
func NewReportingStageHdl(svc service.IReportingStageSvc) *ReportingStageHdl {
	return &ReportingStageHdl{Service: svc}
}
