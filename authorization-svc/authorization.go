package main

import (
	"context"
	"encoding/json"
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

type UserClaim struct {
	ID    string `json:"user_id"`
	Email string `json:"email"`
}

func (a *Authorization) ExtractToken(ctx context.Context, req *authorizationpb.ExtractTokenRequest) (*authorizationpb.ExtractTokenResponse, error) {
	opt := option.WithCredentialsFile("thesis-course-registration-firebase-adminsdk-9o94i-5c3c81a7b0.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	authToken, err := auth.VerifyIDToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	// Convert map to json string
	jsonStr, err := json.Marshal(authToken.Claims)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user UserClaim
	// Convert json string to struct
	if err := json.Unmarshal(jsonStr, &user); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &authorizationpb.ExtractTokenResponse{
		UserID: user.ID,
		Email:  user.Email,
	}, nil
}

func (a *Authorization) Authorize(ctx context.Context, req *authorizationpb.AuthorizeRequest) (*authorizationpb.AuthorizeResponse, error) {
	accessibleRoles := accessibleRoles(req.Role)

	for _, m := range accessibleRoles {
		if m == req.Method {
			log.Println(req.Method)
			return &authorizationpb.AuthorizeResponse{
				CanAccess: true,
			}, nil
		}
	}

	return &authorizationpb.AuthorizeResponse{
		CanAccess: false,
	}, nil
}

func accessibleRoles(role string) []string {
	if role == "admin" {
		return []string{
			"CreateAttachment",
			"GetAttachment",
			"UpdateAttachment",
			"DeleteAttachment",
			"GetAttachmentsOfExercise",
			"CreateClassroom",
			"GetClassroom",
			"DeleteClassroom",
			"UpdateClassroom",
			"GetClassrooms",
			"CheckClassroomExists",
			"GetLecturerClassroom",
			"CreateComment",
			"CreateExercise",
			"GetExercise",
			"DeleteExercise",
			"UpdateExercise",
			"GetExercises",
			"GetAllExercisesOfClassroom",
			"GetAllExercisesInReportingStage",
			"CreateMemer",
			"GetMember",
			"UpdateMember",
			"DeleteMember",
			"GetMembers",
			"GetAllMembersOfClassroom",
			"GetUserMember",
			"CreatePost",
			"GetPost",
			"DeletePost",
			"UpdatePost",
			"GetPosts",
			"GetAllPostsOfClassroom",
			"GetAllPostsInReportingStage",
			"GetAllUsersOfClassroom",
			"CreateReportingStage",
			"GetReportingStage",
			"DeleteReportingStage",
			"UpdateReportingStage",
			"GetReportingStages",
			"CreateSubmission",
			"GetSubmission",
			"DeleteSubmission",
			"UpdateSubmission",
			"GetSubmissions",
			"GetAllSubmissionsOfExercise",
			"GetSubmissionsOfUser",
			"CreateTopic",
			"GetTopic",
			"UpdateTopic",
			"DeleteTopic",
			"GetTopics",
			"CreateUser",
			"GetUser",
			"DeleteUser",
			"UpdateUser",
			"GetUsers",
			"CheckStatusUserJoinClassroom",
			"UnsubscribeClassroom",
			"GetAllLecturers",
			"CreateWaitingList",
			"GetWaitingList",
			"DeleteWaitingList",
			"UpdateWaitingList",
			"GetWaitingLists",
			"GetWaitingListsOfClassroom",
			"CheckUserInWaitingListOfClassroom",
		}
	} else if role == "lecturer" {
		return []string{
			"CreateAttachment",
			"GetAttachment",
			"UpdateAttachment",
			"DeleteAttachment",
			"GetAttachmentsOfExercise",
			"CreateClassroom",
			"GetClassroom",
			"UpdateClassroom",
			"GetClassrooms",
			"CheckClassroomExists",
			"GetLecturerClassroom",
			"CreateComment",
			"CreateExercise",
			"GetExercise",
			"DeleteExercise",
			"UpdateExercise",
			"GetExercises",
			"GetAllExercisesOfClassroom",
			"GetAllExercisesInReportingStage",
			"CreateMemer",
			"GetMember",
			"UpdateMember",
			"DeleteMember",
			"GetMembers",
			"GetAllMembersOfClassroom",
			"GetUserMember",
			"CreatePost",
			"GetPost",
			"DeletePost",
			"UpdatePost",
			"GetPosts",
			"GetAllPostsOfClassroom",
			"GetAllPostsInReportingStage",
			"GetAllUsersOfClassroom",
			"CreateReportingStage",
			"GetReportingStage",
			"DeleteReportingStage",
			"UpdateReportingStage",
			"GetReportingStages",
			"CreateSubmission",
			"GetSubmission",
			"DeleteSubmission",
			"UpdateSubmission",
			"GetSubmissions",
			"GetAllSubmissionsOfExercise",
			"GetSubmissionsOfUser",
			"CreateTopic",
			"GetTopic",
			"UpdateTopic",
			"DeleteTopic",
			"GetTopics",
			"CreateUser",
			"GetUser",
			"UpdateUser",
			"GetUsers",
			"CheckStatusUserJoinClassroom",
			"UnsubscribeClassroom",
			"GetAllLecturers",
			"CreateWaitingList",
			"GetWaitingList",
			"DeleteWaitingList",
			"UpdateWaitingList",
			"GetWaitingLists",
			"GetWaitingListsOfClassroom",
			"CheckUserInWaitingListOfClassroom",
		}
	} else if role == "student" {
		return []string{
			"CreateAttachment",
			"GetAttachment",
			"UpdateAttachment",
			"DeleteAttachment",
			"GetClassroom",
			"GetClassrooms",
			"CheckClassroomExists",
			"GetLecturerClassroom",
			"CreateComment",
			"GetExercise",
			"GetExercises",
			"GetAllExercisesOfClassroom",
			"GetAllExercisesInReportingStage",
			"CreateMemer",
			"GetMember",
			"UpdateMember",
			"GetAllMembersOfClassroom",
			"GetUserMember",
			"GetPost",
			"GetPosts",
			"GetAllPostsOfClassroom",
			"GetAllPostsInReportingStage",
			"GetAllUsersOfClassroom",
			"GetReportingStage",
			"GetReportingStages",
			"CreateSubmission",
			"GetSubmission",
			"DeleteSubmission",
			"UpdateSubmission",
			"GetSubmissions",
			"GetSubmissionsOfUser",
			"CreateTopic",
			"GetTopic",
			"UpdateTopic",
			"DeleteTopic",
			"GetTopics",
			"CreateUser",
			"GetUser",
			"UpdateUser",
			"GetUsers",
			"CheckStatusUserJoinClassroom",
			"UnsubscribeClassroom",
			"GetAllLecturers",
			"CreateWaitingList",
			"GetWaitingList",
			"DeleteWaitingList",
			"UpdateWaitingList",
			"GetWaitingLists",
			"GetWaitingListsOfClassroom",
			"CheckUserInWaitingListOfClassroom",
		}

	} else {
		return []string{"CreateUser"}
	}
}

// func Producer() {
// 	// Configure Kafka producer for user requests
// 	producerConfig := &kafka.ConfigMap{
// 		"bootstrap.servers": "your_kafka_bootstrap_servers",
// 	}

// 	producer, err := kafka.NewProducer(producerConfig)
// 	if err != nil {
// 		fmt.Printf("Error creating Kafka producer: %v\n", err)
// 		return
// 	}
// 	defer producer.Close()

// 	// Configure Kafka consumer for user information
// 	config := &kafka.ConfigMap{
// 		"bootstrap.servers": "your_kafka_bootstrap_servers",
// 		"group.id":          "authorization_service_group",
// 		"auto.offset.reset": "earliest",
// 	}

// 	consumer, err := kafka.NewConsumer(config)
// 	if err != nil {
// 		fmt.Printf("Error creating Kafka consumer: %v\n", err)
// 		return
// 	}
// 	defer consumer.Close()

// 	// Subscribe to the user information topic
// 	topic := "user_info_topic"
// 	err = consumer.SubscribeTopics([]string{topic}, nil)
// 	if err != nil {
// 		fmt.Printf("Error subscribing to topic %s: %v\n", topic, err)
// 		return
// 	}

// 	// Set up a function to send user ID to Message Queue Service
// 	sendUserIDToMessageQueue := func(userID string) {
// 		userRequestTopic := "user_request_topic"

// 		err := producer.Produce(&kafka.Message{
// 			TopicPartition: kafka.TopicPartition{Topic: &userRequestTopic, Partition: kafka.PartitionAny},
// 			Value:          []byte(userID),
// 		}, nil)

// 		if err != nil {
// 			fmt.Printf("Error producing user ID to Message Queue Service: %v\n", err)
// 		}
// 	}

// 	// Consume user information and process accordingly
// 	for {
// 		ev := consumer.Poll(100)
// 		switch e := ev.(type) {
// 		case *kafka.Message:
// 			userInfo := string(e.Value)
// 			// Process user information (e.g., authorization logic)
// 			fmt.Printf("Received user information: %s\n", userInfo)
// 		case kafka.Error:
// 			fmt.Printf("Error: %v\n", e)
// 		}
// 	}
// }
