package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	reportingStageSvcV1 "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
)

type reportingStageServiceGW struct {
	pb.UnimplementedReportingStageServiceServer
	reportingStageClient reportingStageSvcV1.ReportingStageServiceClient
}

func NewReportingStagesService(reportingStageClient reportingStageSvcV1.ReportingStageServiceClient) *reportingStageServiceGW {
	return &reportingStageServiceGW{
		reportingStageClient: reportingStageClient,
	}
}

func (u *reportingStageServiceGW) CreateReportingStage(ctx context.Context, req *pb.CreateReportingStageRequest) (*pb.CreateReportingStageResponse, error) {
	res, err := u.reportingStageClient.CreateReportingStage(ctx, &reportingStageSvcV1.CreateReportingStageRequest{
		ReportingStage: &reportingStageSvcV1.ReportingStageInput{
			Name:        req.GetReportingStage().Name,
			Description: req.GetReportingStage().Description,
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
	res, err := u.reportingStageClient.GetReportingStage(ctx, &reportingStageSvcV1.GetReportingStageRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetReportingStageResponse{
		Response: &pb.CommonReportingStageResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		ReportingStage: &pb.ReportingStageResponse{
			Id:          res.GetReportingStage().Id,
			Name:        res.GetReportingStage().Name,
			Description: res.GetReportingStage().Description,
		},
	}, nil
}

func (u *reportingStageServiceGW) UpdateReportingStage(ctx context.Context, req *pb.UpdateReportingStageRequest) (*pb.UpdateReportingStageResponse, error) {
	res, err := u.reportingStageClient.UpdateReportingStage(ctx, &reportingStageSvcV1.UpdateReportingStageRequest{
		Id: req.GetId(),
		ReportingStage: &reportingStageSvcV1.ReportingStageInput{
			Name:        req.GetReportingStage().Name,
			Description: req.GetReportingStage().Description,
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
	res, err := u.reportingStageClient.DeleteReportingStage(ctx, &reportingStageSvcV1.DeleteReportingStageRequest{
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
	res, err := u.reportingStageClient.GetReportingStages(ctx, &reportingStageSvcV1.GetReportingStagesRequest{})
	if err != nil {
		return nil, err
	}

	var reportingStages []*pb.ReportingStageResponse
	for _, p := range res.GetReportingStages() {
		reportingStages = append(reportingStages, &pb.ReportingStageResponse{
			Id:          p.Id,
			Name:        p.Name,
			Description: p.Description,
		})
	}

	return &pb.GetReportingStagesResponse{
		Response: &pb.CommonReportingStageResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		ReportingStages: reportingStages,
	}, nil
}
