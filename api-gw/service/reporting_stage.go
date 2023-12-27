package service

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
)

type reportingStageServiceGW struct {
	pb.UnimplementedReportingStageServiceServer
	reportingStageClient classroomSvcV1.ReportingStageServiceClient
}

func NewReportingStagesService(reportingStageClient classroomSvcV1.ReportingStageServiceClient) *reportingStageServiceGW {
	return &reportingStageServiceGW{
		reportingStageClient: reportingStageClient,
	}
}

func (u *reportingStageServiceGW) CreateReportingStage(ctx context.Context, req *pb.CreateReportingStageRequest) (*pb.CreateReportingStageResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.reportingStageClient.CreateReportingStage(ctx, &classroomSvcV1.CreateReportingStageRequest{
		ReportingStage: &classroomSvcV1.ReportingStageInput{
			Label:       req.GetCategory().Label,
			Description: req.GetCategory().Description,
			Value:       req.GetCategory().Value,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateReportingStageResponse{
		Response: &pb.CommonReportingStageResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *reportingStageServiceGW) GetReportingStage(ctx context.Context, req *pb.GetReportingStageRequest) (*pb.GetReportingStageResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.reportingStageClient.GetReportingStage(ctx, &classroomSvcV1.GetReportingStageRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetReportingStageResponse{
		Response: &pb.CommonReportingStageResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Category: &pb.ReportingStageResponse{
			Id:          res.GetReportingStage().Id,
			Label:       res.GetReportingStage().Label,
			Description: res.GetReportingStage().Description,
			Value:       res.GetReportingStage().Value,
		},
	}, nil
}

func (u *reportingStageServiceGW) UpdateReportingStage(ctx context.Context, req *pb.UpdateReportingStageRequest) (*pb.UpdateReportingStageResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.reportingStageClient.UpdateReportingStage(ctx, &classroomSvcV1.UpdateReportingStageRequest{
		Id: req.GetId(),
		ReportingStage: &classroomSvcV1.ReportingStageInput{
			Label:       req.GetCategory().Label,
			Description: req.GetCategory().Description,
			Value:       req.GetCategory().Value,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateReportingStageResponse{
		Response: &pb.CommonReportingStageResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *reportingStageServiceGW) DeleteReportingStage(ctx context.Context, req *pb.DeleteReportingStageRequest) (*pb.DeleteReportingStageResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.reportingStageClient.DeleteReportingStage(ctx, &classroomSvcV1.DeleteReportingStageRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteReportingStageResponse{
		Response: &pb.CommonReportingStageResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *reportingStageServiceGW) GetReportingStages(ctx context.Context, req *pb.GetReportingStagesRequest) (*pb.GetReportingStagesResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.reportingStageClient.GetReportingStages(ctx, &classroomSvcV1.GetReportingStagesRequest{})
	if err != nil {
		return nil, err
	}

	var reportingStages []*pb.ReportingStageResponse
	for _, p := range res.GetReportingStages() {
		reportingStages = append(reportingStages, &pb.ReportingStageResponse{
			Id:          p.Id,
			Label:       p.Label,
			Description: p.Description,
			Value:       p.Value,
		})
	}

	return &pb.GetReportingStagesResponse{
		Response: &pb.CommonReportingStageResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Categorys: reportingStages,
	}, nil
}
