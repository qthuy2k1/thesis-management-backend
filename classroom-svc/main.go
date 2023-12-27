package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/joho/godotenv"
	classroompb "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"

	classroomHdl "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler/classroom"
	classroomRepo "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/classroom"
	classroomSvc "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/classroom"

	classroomWaitingListHdl "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler/classroom-waiting-list"
	classroomWaitingListRepo "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/classroom-waiting-list"
	classroomWaitingListSvc "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/classroom-waiting-list"

	exerciseHdl "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler/exercise"
	exerciseRepo "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/exercise"
	exerciseSvc "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/exercise"

	postHdl "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler/post"
	postRepo "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/post"
	postSvc "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/post"

	reportingStageHdl "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler/reporting-stage"
	reportingStageRepo "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/reporting-stage"
	reportingStageSvc "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/reporting-stage"

	submissionHdl "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler/submission"
	submissionRepo "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/submission"
	submissionSvc "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/submission"

	attachmentHdl "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler/attachment"
	attachmentRepo "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository/attachment"
	attachmentSvc "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service/attachment"

	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/pkg/db"
)

const (
	listenAddress = "0.0.0.0:9091"
	serviceName   = "Classroom service"
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

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dbUrl := os.Getenv("DB_URL")
	database, err := db.Initialize(dbUrl)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	defer database.Close()

	clrRepo := classroomRepo.NewClassroomRepo(database)
	clrSvc := classroomSvc.NewClassroomSvc(clrRepo)
	clrHdl := classroomHdl.NewClassroomHdl(clrSvc)

	clrWTLRepo := classroomWaitingListRepo.NewWaitingListRepo(database)
	clrWTLSvc := classroomWaitingListSvc.NewWaitingListSvc(clrWTLRepo)
	clrWTLHdl := classroomWaitingListHdl.NewWaitingListHdl(clrWTLSvc)

	exRepo := exerciseRepo.NewExerciseRepo(database)
	exSvc := exerciseSvc.NewExerciseSvc(exRepo)
	exHdl := exerciseHdl.NewExerciseHdl(exSvc)

	pRepo := postRepo.NewPostRepo(database)
	pSvc := postSvc.NewPostSvc(pRepo)
	pHdl := postHdl.NewPostHdl(pSvc)

	rpsRepo := reportingStageRepo.NewReportingStageRepo(database)
	rpsSvc := reportingStageSvc.NewReportingStageSvc(rpsRepo)
	rpsHdl := reportingStageHdl.NewReportingStageHdl(rpsSvc)

	sRepo := submissionRepo.NewSubmissionRepo(database)
	sSvc := submissionSvc.NewSubmissionSvc(sRepo)
	sHdl := submissionHdl.NewSubmissionHdl(sSvc)

	aRepo := attachmentRepo.NewAttachmentRepo(database)
	aSvc := attachmentSvc.NewAttachmentSvc(aRepo)
	aHdl := attachmentHdl.NewAttachmentHdl(aSvc)

	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	classroompb.RegisterClassroomServiceServer(s, clrHdl)
	classroompb.RegisterWaitingListServiceServer(s, clrWTLHdl)
	classroompb.RegisterExerciseServiceServer(s, exHdl)
	classroompb.RegisterPostServiceServer(s, pHdl)
	classroompb.RegisterReportingStageServiceServer(s, rpsHdl)
	classroompb.RegisterSubmissionServiceServer(s, sHdl)
	classroompb.RegisterAttachmentServiceServer(s, aHdl)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
