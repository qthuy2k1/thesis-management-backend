package service

import (
	"context"
	"errors"
	"log"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	authorizeSvcV1 "github.com/qthuy2k1/thesis-management-backend/authorization-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserContext string

const USER_CONTEXT UserContext = "user"

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

func (u *authServiceGW) GetToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("error occured from incoming context")
	}

	authHeader := md.Get("authorization")
	if len(authHeader) > 0 {
		token := strings.Split(authHeader[0], " ")[1]
		if token == "" {
			return "", status.Error(codes.Unauthenticated, "token is empty")
		}
		return token, nil
	}

	return "", status.Error(codes.Unauthenticated, "authorization header not found")
}

func (u *authServiceGW) Authorize(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("APIGW service: method %q called\n", info.FullMethod)

	token, err := u.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	methodArr := strings.Split(info.FullMethod, "/")

	res, err := u.authorizationClient.ExtractToken(ctx, &authorizeSvcV1.ExtractTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}

	user, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: res.UserID,
	})
	if err != nil {
		log.Println("authorize: GetUser", err)
	}

	// Add user data to the context.
	ctx = context.WithValue(ctx, USER_CONTEXT, user)

	auth, err := u.authorizationClient.Authorize(ctx, &authorizeSvcV1.AuthorizeRequest{
		Method: methodArr[2],
		Role:   user.GetUser().GetRole(),
	})
	if err != nil {
		return nil, err
	}

	if !auth.CanAccess {
		return nil, status.Errorf(codes.PermissionDenied, "you're not allowed to access")
	}

	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("APIGW serivce: method %q failed: %s\n", info.FullMethod, err)
	}
	return resp, err
}

// hasRequiredRole checks if the user has the required role.
func hasRequiredRole(role string, requiredRole string) bool {
	return role == requiredRole
}
