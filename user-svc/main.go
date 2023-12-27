package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	userpb "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"

	userHdl "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/handler/user"
	userRepo "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository/user"
	userSvc "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service/user"

	topicHdl "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/handler/topic"
	topicRepo "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository/topic"

	commentHdl "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/handler/comment"
	commentRepo "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository/comment"
	commentSvc "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/service/comment"

	"github.com/qthuy2k1/thesis-management-backend/user-svc/pkg/db"
	"google.golang.org/grpc"
)

const (
	listenAddress = "0.0.0.0:9091"
	serviceName   = "User service"
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
	log.Printf("Users service starting on %s", listenAddress)
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

	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PORT")

	redis := db.RedisInitialize(redisPort, redisPassword)

	uRepo := userRepo.NewUserRepo(database, redis)
	uSvc := userSvc.NewUserSvc(uRepo)
	uHdl := userHdl.NewUserHdl(uSvc)

	cRepo := commentRepo.NewCommentRepo(database)
	cSvc := commentSvc.NewCommentSvc(cRepo)
	cHdl := commentHdl.NewCommentHdl(cSvc)

	tRepo := topicRepo.NewTopicRepo(database)
	tHdl := topicHdl.NewTopicHdl(tRepo)

	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	userpb.RegisterUserServiceServer(s, uHdl)
	userpb.RegisterCommentServiceServer(s, cHdl)
	userpb.RegisterTopicServiceServer(s, tHdl)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
