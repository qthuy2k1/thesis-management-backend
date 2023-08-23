package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	postpb "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/post-svc/internal/handler"
	"github.com/qthuy2k1/thesis-management-backend/post-svc/internal/repository"
	"github.com/qthuy2k1/thesis-management-backend/post-svc/internal/service"
	"github.com/qthuy2k1/thesis-management-backend/post-svc/pkg/db"
)

const (
	listenAddress = "0.0.0.0:9091"
)

func main() {
	log.Printf("Posts service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// if err = godotenv.Load(); err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	dbUrl := "postgres://postgres:root@post-db:5432/thesis_management_posts?sslmode=disable"
	database, err := db.Initialize(dbUrl)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	repository := repository.NewPostRepo(database)
	service := service.NewPostSvc(repository)
	handler := handler.NewPostHdl(service)

	s := grpc.NewServer()

	postpb.RegisterPostServiceServer(s, handler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
