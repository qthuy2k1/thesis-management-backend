package main

import (
	"context"
	"errors"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	authorizeSvcV1 "github.com/qthuy2k1/thesis-management-backend/authorization-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	"google.golang.org/grpc/metadata"
)

type authServiceGW struct {
	pb.UnimplementedAuthorizationServiceServer
	userClient          userSvcV1.UserServiceClient
	authorizationClient authorizeSvcV1.AuthorizationServiceClient
}

func NewAuthorizationService(userClient userSvcV1.UserServiceClient, authorizationClient authorizeSvcV1.AuthorizationServiceClient) *authServiceGW {
	return &authServiceGW{
		userClient:          userClient,
		authorizationClient: authorizationClient,
	}
}

func GetToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("error occured from incoming context")
	}

	authHeader := md.Get("authorization")
	if len(authHeader) > 0 {
		token := strings.Split(authHeader[0], " ")[1]
		if token == "" {
			return "", errors.New("token is empty")
		}
		return token, nil
	}
	return "", errors.New("authorization header not found")
}
