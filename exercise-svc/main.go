package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	exercisepb "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	"github.com/qthuy2k1/thesis-management-backend/exercise-svc/internal/handler"
	"github.com/qthuy2k1/thesis-management-backend/exercise-svc/internal/repository"
	"github.com/qthuy2k1/thesis-management-backend/exercise-svc/internal/service"
	"github.com/qthuy2k1/thesis-management-backend/exercise-svc/pkg/db"
)

const (
	listenAddress = "0.0.0.0:9091"
)

func main() {
	log.Printf("Exercises service starting on %s", listenAddress)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// if err = godotenv.Load(); err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	dbUrl := "postgres://postgres:root@exercise-db:5432/thesis_management_exercises?sslmode=disable"
	database, err := db.Initialize(dbUrl)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	repository := repository.NewExerciseRepo(database)
	service := service.NewExerciseSvc(repository)
	handler := handler.NewExerciseHdl(service)

	s := grpc.NewServer()

	exercisepb.RegisterExerciseServiceServer(s, handler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
