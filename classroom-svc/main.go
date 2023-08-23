package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	classroompb "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/handler"
	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/repository"
	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/service"
	"github.com/qthuy2k1/thesis-management-backend/classroom-svc/pkg/db"
)

const (
	listenAddress = "0.0.0.0:9091"
)

func main() {
	log.Printf("Classrooms service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	dbUrl := "postgresql://postgres:root@127.0.0.1:5432/thesis_management_classrooms?sslmode=disable"
	database, err := db.Initialize(dbUrl)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	defer database.Close()

	repository := repository.NewClassroomRepo(database)
	service := service.NewClassroomSvc(repository)
	handler := handler.NewClassroomHdl(service)

	s := grpc.NewServer()

	classroompb.RegisterClassroomServiceServer(s, handler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
