package main

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	classroompb "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler"
	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository"
	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service"
	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/pkg/db"

	"github.com/joho/godotenv"
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

	repository := repository.NewClassroomRepo(database)
	service := service.NewClassroomSvc(repository)
	handler := handler.NewClassroomHdl(service)

	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	classroompb.RegisterClassroomServiceServer(s, handler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
