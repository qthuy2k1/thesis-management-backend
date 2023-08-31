package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	rpspb "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/internal/handler"
	"github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/internal/repository"
	"github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/internal/service"
	"github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/pkg/db"
)

const (
	listenAddress = "0.0.0.0:9091"
	serviceName   = "Reporting stage service"
)

func logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("%s: method %q called\n", serviceName, info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("%s: method %q failed: %s\n", serviceName, info.FullMethod, err)
	}
	return resp, err
}

func main() {
	log.Printf("%s starting on %s", serviceName, listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// if err = godotenv.Load(); err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	dbUrl := "postgres://postgres:root@reporting-stage-db:5432/thesis_management_reporting_stages?sslmode=disable"
	database, err := db.Initialize(dbUrl)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	repository := repository.NewReportingStageRepo(database)
	service := service.NewReportingStageSvc(repository)
	handler := handler.NewReportingStageHdl(service)

	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	rpspb.RegisterReportingStageServiceServer(s, handler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
