package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	redispb "github.com/qthuy2k1/thesis-management-backend/redis-svc/api/goclient/v1"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

const (
	listenAddress = "0.0.0.0:9091"
	serviceName   = "Redis service"
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

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	redisPort := os.Getenv("REDIS_PORT")
	redisPass := os.Getenv("REDIS_PASSWORD")

	client := redis.NewClient(&redis.Options{
		Addr:     "thesis-management-backend-redis-db-service:6379",
		Password: redisPass,
		DB:       0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Could not connect to redis server: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	redis := NewRedis(client)
	redispb.RegisterRedisServiceServer(s, redis)

	log.Printf("Started REDIS server on %s", redisPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
