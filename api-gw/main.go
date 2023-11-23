package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	attachmentSvcV1 "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
	authorizationSvcV1 "github.com/qthuy2k1/thesis-management-backend/authorization-svc/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	waitingListSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	commentSvcV1 "github.com/qthuy2k1/thesis-management-backend/comment-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	redisSvcV1 "github.com/qthuy2k1/thesis-management-backend/redis-svc/api/goclient/v1"
	rpsSvcV1 "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	scheduleSvcV1 "github.com/qthuy2k1/thesis-management-backend/schedule-svc/api/goclient/v1"
	submissionSvcV1 "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
	commiteeSvcV1 "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	topicSvcV1 "github.com/qthuy2k1/thesis-management-backend/topic-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

var address = map[string]string{
	"listenAddress":         "0.0.0.0:9091",
	"classroomAddress":      "thesis-management-backend-classroom-service:9091",
	"postAddress":           "thesis-management-backend-post-service:9091",
	"exerciseAddress":       "thesis-management-backend-exercise-service:9091",
	"reportingStageAddress": "thesis-management-backend-reporting-stage-service:9091",
	"submissionAddress":     "thesis-management-backend-submission-service:9091",
	"userAddress":           "thesis-management-backend-user-service:9091",
	"waitingListAddress":    "thesis-management-backend-classroom-waiting-list-service:9091",
	"commentAddress":        "thesis-management-backend-comment-service:9091",
	"attachmentAddress":     "thesis-management-backend-attachment-service:9091",
	"topicAddress":          "thesis-management-backend-topic-service:9091",
	"authorizationAddress":  "thesis-management-backend-authorization-service:9091",
	"commiteeAddress":       "thesis-management-backend-thesis-commitee-service:9091",
	"scheduleAddress":       "thesis-management-backend-schedule-service:9091",
	"redisAddress":          "thesis-management-backend-redis-service:9091",
}

func newClassroomSvcClient() (classroomSvcV1.ClassroomServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["classroomAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("classroom client: %w", err)
	}

	return classroomSvcV1.NewClassroomServiceClient(conn), nil
}

func newPostSvcClient() (postSvcV1.PostServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["postAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post client: %w", err)
	}

	return postSvcV1.NewPostServiceClient(conn), nil
}

func newExerciseSvcClient() (exerciseSvcV1.ExerciseServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["exerciseAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("exercise client: %w", err)
	}

	return exerciseSvcV1.NewExerciseServiceClient(conn), nil
}

func newReportingStageSvcClient() (rpsSvcV1.ReportingStageServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["reportingStageAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("reporting stage client: %w", err)
	}

	return rpsSvcV1.NewReportingStageServiceClient(conn), nil
}

func newSubmissionSvcClient() (submissionSvcV1.SubmissionServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["submissionAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("submission client: %w", err)
	}

	return submissionSvcV1.NewSubmissionServiceClient(conn), nil
}

func newUserSvcClient() (userSvcV1.UserServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["userAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("user client: %w", err)
	}

	return userSvcV1.NewUserServiceClient(conn), nil
}

func newWaitingListSvcClient() (waitingListSvcV1.WaitingListServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["waitingListAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("clasroom waiting list client: %w", err)
	}

	return waitingListSvcV1.NewWaitingListServiceClient(conn), nil
}

func newCommentSvcClient() (commentSvcV1.CommentServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["commentAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("comment client: %w", err)
	}

	return commentSvcV1.NewCommentServiceClient(conn), nil
}

func newAttachmentSvcClient() (attachmentSvcV1.AttachmentServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["attachmentAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("attachment client: %w", err)
	}

	return attachmentSvcV1.NewAttachmentServiceClient(conn), nil
}

func newTopicSvcClient() (topicSvcV1.TopicServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["topicAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("topic client: %w", err)
	}

	return topicSvcV1.NewTopicServiceClient(conn), nil
}

func newAuthorizationSvcClient() (authorizationSvcV1.AuthorizationServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["authorizationAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("authorize client: %w", err)
	}

	return authorizationSvcV1.NewAuthorizationServiceClient(conn), nil
}

func newCommiteeSvcClient() (commiteeSvcV1.CommiteeServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["commiteeAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("commitee client: %w", err)
	}

	return commiteeSvcV1.NewCommiteeServiceClient(conn), nil
}

func newThesisSvcClient() (commiteeSvcV1.ScheduleServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["commiteeAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("thesis client: %w", err)
	}

	return commiteeSvcV1.NewScheduleServiceClient(conn), nil
}

func newScheduleSvcClient() (scheduleSvcV1.ScheduleServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["scheduleAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("schedule client: %w", err)
	}

	return scheduleSvcV1.NewScheduleServiceClient(conn), nil
}

func newRedisSvcClient() (redisSvcV1.RedisServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["redisAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("redis client: %w", err)
	}

	return redisSvcV1.NewRedisServiceClient(conn), nil
}

func logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("APIGW service: method %q called\n", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("APIGW serivce: method %q failed: %s\n", info.FullMethod, err)
	}
	return resp, err
}

// func authorize(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	token, err := GetToken(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	authorizationClient, err := newAuthorizationSvcClient()
// 	if err != nil {
// 		return nil, err
// 	}

// 	methodArr := strings.Split(info.FullMethod, "/")
// 	if _, err := authorizationClient.Authorize(ctx, &authorizeSvcV1.AuthorizeRequest{
// 		Token:  token,
// 		Method: methodArr[2],
// 	}); err != nil {
// 		return nil, err
// 	}

// 	resp, err := handler(ctx, req)
// 	if err != nil {
// 		log.Printf("APIGW serivce: method %q failed: %s\n", info.FullMethod, err)
// 	}
// 	return resp, err
// }

func main() {
	fmt.Printf("APIGW service starting on %s", address["listenAddress"])

	// connect to classroom svc
	classroomClient, err := newClassroomSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to post svc
	postClient, err := newPostSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to exercise svc
	exerciseClient, err := newExerciseSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to reporting stage svc
	rpsClient, err := newReportingStageSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to submission svc
	submissionClient, err := newSubmissionSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to user svc
	userClient, err := newUserSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to classrom waiting list svc
	waitingListClient, err := newWaitingListSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to comment svc
	commentClient, err := newCommentSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to attachment svc
	attachmentClient, err := newAttachmentSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to topic svc
	topicClient, err := newTopicSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to authorization svc
	authorizationClient, err := newAuthorizationSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to commitee svc
	commiteeClient, err := newCommiteeSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to schedule svc
	scheduleClient, err := newScheduleSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to thesis svc
	thesisClient, err := newThesisSvcClient()
	if err != nil {
		panic(err)
	}

	// connect to redis svc
	// redisClient, err := newRedisSvcClient()
	// if err != nil {
	// 	panic(err)
	// }

	lis, err := net.Listen("tcp", address["listenAddress"])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	pb.RegisterClassroomServiceServer(s, NewClassroomsService(classroomClient, postClient, exerciseClient, rpsClient, userClient, topicClient))
	pb.RegisterPostServiceServer(s, NewPostsService(postClient, classroomClient, rpsClient, commentClient, userClient, attachmentClient))
	pb.RegisterExerciseServiceServer(s, NewExercisesService(exerciseClient, classroomClient, rpsClient, commentClient, userClient, submissionClient, attachmentClient))
	pb.RegisterReportingStageServiceServer(s, NewReportingStagesService(rpsClient))
	pb.RegisterSubmissionServiceServer(s, NewSubmissionsService(submissionClient, classroomClient, exerciseClient, attachmentClient, userClient))
	pb.RegisterUserServiceServer(s, NewUsersService(userClient, classroomClient, waitingListClient, topicClient))
	pb.RegisterWaitingListServiceServer(s, NewWaitingListsService(waitingListClient, classroomClient, userClient))
	pb.RegisterCommentServiceServer(s, NewCommentsService(commentClient, postClient, exerciseClient, userClient))
	pb.RegisterAttachmentServiceServer(s, NewAttachmentsService(attachmentClient, userClient, submissionClient))
	pb.RegisterTopicServiceServer(s, NewTopicsService(topicClient, userClient))
	pb.RegisterAuthorizationServiceServer(s, NewAuthorizationService(userClient, authorizationClient))
	pb.RegisterMemberServiceServer(s, NewMembersService(userClient, classroomClient, waitingListClient))
	pb.RegisterCommiteeServiceServer(s, NewCommiteesService(commiteeClient, userClient))
	pb.RegisterRoomServiceServer(s, NewRoomsService(commiteeClient))
	pb.RegisterStudentDefServiceServer(s, NewStudentDefsService(userClient))
	pb.RegisterScheduleServiceServer(s, NewSchedulesService(scheduleClient, commiteeClient, userClient, thesisClient))
	pb.RegisterNotificationServiceServer(s, NewNotificationsService(scheduleClient, userClient))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
