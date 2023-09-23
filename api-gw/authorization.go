package main

import (
	"context"
	"errors"
	"log"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type authServiceGW struct {
	pb.UnimplementedAuthorizationServiceServer
	userClient userSvcV1.UserServiceClient
}

func NewAuthService(userClient userSvcV1.UserServiceClient) *authServiceGW {
	return &authServiceGW{
		userClient: userClient,
	}
}

// func accessibleRoles(service string) map[string][]string {
// 	publicRoles := []string{"admin", "teacher", "user", ""}
// 	userRoles := []string{"admin", "teacher", "user"}
// 	teacherRoles := []string{"admin", "teacher"}
// 	adminRoles := []string{"admin"}

//		switch service {
//		case "classroom":
//			return map[string][]string{
//				"CreateClassroom":      teacherRoles,
//				"GetClassroom":         userRoles,
//				"DeleteClassroom":      teacherRoles,
//				"UpdateClassroom":      teacherRoles,
//				"GetClassrooms":        userRoles,
//				"CheckClassroomExists": userRoles,
//			}
//		case "post":
//			return map[string][]string{
//				"CreatePost":             teacherRoles,
//				"GetPost":                userRoles,
//				"DeletePost":             teacherRoles,
//				"UpdatePost":             teacherRoles,
//				"GetPosts":               userRoles,
//				"GetAllPostsOfClassroom": userRoles,
//			}
//		case "exercise":
//			return map[string][]string{
//				"CreateExercise":             teacherRoles,
//				"GetExercise":                userRoles,
//				"DeleteExercise":             teacherRoles,
//				"UpdateExercise":             teacherRoles,
//				"GetExercises":               userRoles,
//				"GetAllExercisesOfClassroom": userRoles,
//			}
//		case "submission":
//			return map[string][]string{
//				"CreateSubmission":            teacherRoles,
//				"GetSubmission":               userRoles,
//				"DeleteSubmission":            teacherRoles,
//				"UpdateSubmission":            teacherRoles,
//				"GetSubmissions":              userRoles,
//				"GetAllSubmissionsOfExercise": userRoles,
//			}
//		case "user":
//			return map[string][]string{
//				"CreateUser":             publicRoles,
//				"GetUser":                userRoles,
//				"DeleteUser":             adminRoles,
//				"UpdateUser":             userRoles,
//				"GetUsers":               adminRoles,
//				"GetAllUsersOfClassroom": userRoles,
//			}
//		case "reporting stage":
//			return map[string][]string{
//				"CreateReportingStage": adminRoles,
//				"GetReportingStage":    userRoles,
//				"DeleteReportingStage": adminRoles,
//				"UpdateReportingStage": adminRoles,
//				"GetReportingStages":   userRoles,
//			}
//		case "classroom waiting list":
//			return map[string][]string{
//				"CreateWaitingList": adminRoles,
//				"GetWaitingList":    userRoles,
//				"DeleteWaitingList": adminRoles,
//				"UpdateWaitingList": adminRoles,
//				"GetWaitingLists":   userRoles,
//			}
//		default:
//			return map[string][]string{}
//		}
//	}
func accessibleRoles() map[string][]string {
	publicRoles := []string{"admin", "teacher", "user", ""}
	userRoles := []string{"admin", "teacher", "user"}
	teacherRoles := []string{"admin", "teacher"}
	adminRoles := []string{"admin"}

	return map[string][]string{
		"CreateClassroom":             teacherRoles,
		"GetClassroom":                userRoles,
		"DeleteClassroom":             teacherRoles,
		"UpdateClassroom":             teacherRoles,
		"GetClassrooms":               userRoles,
		"CheckClassroomExists":        userRoles,
		"CreatePost":                  teacherRoles,
		"GetPost":                     userRoles,
		"DeletePost":                  teacherRoles,
		"UpdatePost":                  teacherRoles,
		"GetPosts":                    userRoles,
		"GetAllPostsOfClassroom":      userRoles,
		"CreateExercise":              teacherRoles,
		"GetExercise":                 userRoles,
		"DeleteExercise":              teacherRoles,
		"UpdateExercise":              teacherRoles,
		"GetExercises":                userRoles,
		"GetAllExercisesOfClassroom":  userRoles,
		"CreateSubmission":            teacherRoles,
		"GetSubmission":               userRoles,
		"DeleteSubmission":            teacherRoles,
		"UpdateSubmission":            teacherRoles,
		"GetSubmissions":              userRoles,
		"GetAllSubmissionsOfExercise": userRoles,
		"CreateUser":                  publicRoles,
		"GetUser":                     userRoles,
		"DeleteUser":                  adminRoles,
		"UpdateUser":                  userRoles,
		"GetUsers":                    adminRoles,
		"GetAllUsersOfClassroom":      userRoles,
		"CreateReportingStage":        adminRoles,
		"GetReportingStage":           userRoles,
		"DeleteReportingStage":        adminRoles,
		"UpdateReportingStage":        adminRoles,
		"GetReportingStages":          userRoles,
		"CreateWaitingList":           adminRoles,
		"GetWaitingList":              userRoles,
		"DeleteWaitingList":           adminRoles,
		"UpdateWaitingList":           adminRoles,
		"GetWaitingLists":             userRoles,
	}
}

func GetToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("error occured from incoming context")
	}

	authHeader := md.Get("authorization")
	token := strings.Split(authHeader[0], " ")[1]

	return token, nil
}

func Authorize(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// accessibleRoles := accessibleRoles()

	// jwtManager := NewJWTManager("something", time.Hour*24)
	authApp, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID: "thesis-course-registration",
	})
	if err != nil {
		return nil, err
	}

	client, err := authApp.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := GetToken(ctx)
	if err != nil {
		return nil, err
	}
	authToken, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		return nil, err
	}

	log.Println(authToken)
	return nil, nil

}

func GetUserFromToken(ctx context.Context, token string, authToken *auth.Token) (string, error) {
	log.Println(authToken.Claims)
	id, ok := authToken.Claims["id"]
	if !ok {
		return "", errors.New("invalid token claim")
	}

	return id.(string), nil
}
