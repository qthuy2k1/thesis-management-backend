package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	authorizationpb "github.com/qthuy2k1/thesis-management-backend/authorization-svc/api/goclient/v1"
)

const (
	listenAddress = "0.0.0.0:9091"
	serviceName   = "Authorization service"
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

	auth := NewAuthorization()

	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	authorizationpb.RegisterAuthorizationServiceServer(s, auth)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
