package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"

	authorizationpb "github.com/qthuy2k1/thesis-management-backend/authorization-svc/api/goclient/v1"
)

type Authorization struct {
	authorizationpb.UnimplementedAuthorizationServiceServer
}

func NewAuthorization() *Authorization {
	return &Authorization{}
}

func Authorize(ctx context.Context, req *authorizationpb.AuthorizeRequest) (*authorizationpb.AuthorizeResponse, error) {
	opt := option.WithCredentialsFile("thesis-course-registration-firebase-adminsdk-9o94i-5c3c81a7b0.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	authToken, err := auth.VerifyIDToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	var user struct {
		UID  string  `json:"uid"`
		Role *string `json:"role"`
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(authToken.Claims)
	if err != nil {
		fmt.Println(err)
	}

	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, &user); err != nil {
		fmt.Println(err)
	}

	log.Println(user)

	if user.Role != nil {
		accessibleRoles := accessibleRoles(*user.Role)
		log.Println(accessibleRoles)

		log.Println(req.Method)
		for _, m := range accessibleRoles {
			if m == req.Method {
				return &authorizationpb.AuthorizeResponse{
					AccessibleMethod: accessibleRoles,
				}, nil
			}
		}

	}
	return nil, errors.New("not allow")
}

func accessibleRoles(role string) []string {
	if role == "admin" {
		return []string{
			"CreateClassroom",
			"GetClassroom",
			"DeleteClassroom",
			"UpdateClassroom",
			"GetClassrooms",
			"CheckClassroomExists",
			"CreatePost",
			"GetPost",
			"DeletePost",
			"UpdatePost",
			"GetPosts",
			"GetAllPostsOfClassroom",
			"CreateExercise",
			"GetExercise",
			"DeleteExercise",
			"UpdateExercise",
			"GetExercises",
			"GetAllExercisesOfClassroom",
			"CreateSubmission",
			"GetSubmission",
			"DeleteSubmission",
			"UpdateSubmission",
			"GetSubmissions",
			"GetAllSubmissionsOfExercise",
			"CreateUser",
			"GetUser",
			"DeleteUser",
			"UpdateUser",
			"GetUsers",
			"GetAllUsersOfClassroom",
			"CreateReportingStage",
			"GetReportingStage",
			"DeleteReportingStage",
			"UpdateReportingStage",
			"GetReportingStages",
			"CreateWaitingList",
			"GetWaitingList",
			"DeleteWaitingList",
			"UpdateWaitingList",
			"GetWaitingLists",
		}
	} else if role == "lecturer" {
		return []string{
			"CreateClassroom",
			"GetClassroom",
			"DeleteClassroom",
			"UpdateClassroom",
			"GetClassrooms",
			"CheckClassroomExists",
			"CreatePost",
			"GetPost",
			"DeletePost",
			"UpdatePost",
			"GetPosts",
			"GetAllPostsOfClassroom",
			"CreateExercise",
			"GetExercise",
			"DeleteExercise",
			"UpdateExercise",
			"GetExercises",
			"GetAllExercisesOfClassroom",
			"CreateSubmission",
			"GetSubmission",
			"DeleteSubmission",
			"UpdateSubmission",
			"GetSubmissions",
			"GetAllSubmissionsOfExercise",
			"CreateUser",
			"GetUser",
			"UpdateUser",
			"GetAllUsersOfClassroom",
			"GetReportingStage",
			"GetReportingStages",
			"GetWaitingList",
			"GetWaitingLists",
		}
	} else if role == "student" {
		return []string{
			"GetClassroom",
			"GetClassrooms",
			"CheckClassroomExists",
			"GetPost",
			"GetPosts",
			"GetAllPostsOfClassroom",
			"GetExercise",
			"GetExercises",
			"GetAllExercisesOfClassroom",
			"GetSubmission",
			"GetSubmissions",
			"GetAllSubmissionsOfExercise",
			"CreateUser",
			"GetUser",
			"UpdateUser",
			"GetAllUsersOfClassroom",
			"GetReportingStage",
			"GetReportingStages",
			"GetWaitingList",
			"GetWaitingLists",
		}
	} else {
		return []string{"Not allowed :)"}
	}
}
