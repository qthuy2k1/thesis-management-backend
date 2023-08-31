package main

import (
	"log"
	"net"

	"github.com/qthuy2k1/thesis-management-backend/user-svc/pkg/db"
	"google.golang.org/grpc"
)

const (
	listenAddress = "0.0.0.0:9091"
)

func main() {
	log.Printf("Users service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// if err = godotenv.Load(); err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	dbUrl := "postgres://postgres:root@user-db:5432/thesis_management_users?sslmode=disable"
	_, err = db.Initialize(dbUrl)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	// log.Println(database)

	// repository := repository.NewPostRepo(database)
	// service := service.NewPostSvc(repository)
	// handler := handler.NewPostHdl(service)

	s := grpc.NewServer()

	// postpb.RegisterPostServiceServer(s, handler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
