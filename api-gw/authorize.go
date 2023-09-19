package main

import (
	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type authServiceGW struct {
	pb.UnimplementedPostServiceServer
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
