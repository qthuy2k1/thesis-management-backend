package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	authorizationSvcV1 "github.com/qthuy2k1/thesis-management-backend/authorization-svc/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	scheduleSvcV1 "github.com/qthuy2k1/thesis-management-backend/schedule-svc/api/goclient/v1"
	commiteeSvcV1 "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"

	service "github.com/qthuy2k1/thesis-management-backend/api-gw/service"
)

var address = map[string]string{
	"listenAddress":        "0.0.0.0:9091",
	"classroomAddress":     "thesis-management-backend-classroom-service:9091",
	"userAddress":          "thesis-management-backend-user-service:9091",
	"authorizationAddress": "thesis-management-backend-authorization-service:9091",
	"commiteeAddress":      "thesis-management-backend-thesis-commitee-service:9091",
	"scheduleAddress":      "thesis-management-backend-schedule-service:9091",
}

func newClassroomSvcClient() (classroomSvcV1.ClassroomServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["classroomAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("classroom client: %w", err)
	}

	return classroomSvcV1.NewClassroomServiceClient(conn), nil
}

func newPostSvcClient() (classroomSvcV1.PostServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["classroomAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post client: %w", err)
	}

	return classroomSvcV1.NewPostServiceClient(conn), nil
}

func newExerciseSvcClient() (classroomSvcV1.ExerciseServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["classroomAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("exercise client: %w", err)
	}

	return classroomSvcV1.NewExerciseServiceClient(conn), nil
}

func newReportingStageSvcClient() (classroomSvcV1.ReportingStageServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["classroomAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("reporting stage client: %w", err)
	}

	return classroomSvcV1.NewReportingStageServiceClient(conn), nil
}

func newSubmissionSvcClient() (classroomSvcV1.SubmissionServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["classroomAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("submission client: %w", err)
	}

	return classroomSvcV1.NewSubmissionServiceClient(conn), nil
}

func newUserSvcClient() (userSvcV1.UserServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["userAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("user client: %w", err)
	}

	return userSvcV1.NewUserServiceClient(conn), nil
}

func newWaitingListSvcClient() (classroomSvcV1.WaitingListServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["classroomAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("clasroom waiting list client: %w", err)
	}

	return classroomSvcV1.NewWaitingListServiceClient(conn), nil
}

func newCommentSvcClient() (userSvcV1.CommentServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["userAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("comment client: %w", err)
	}

	return userSvcV1.NewCommentServiceClient(conn), nil
}

func newAttachmentSvcClient() (classroomSvcV1.AttachmentServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["classroomAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("attachment client: %w", err)
	}

	return classroomSvcV1.NewAttachmentServiceClient(conn), nil
}

func newTopicSvcClient() (userSvcV1.TopicServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), address["userAddress"], grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("topic client: %w", err)
	}

	return userSvcV1.NewTopicServiceClient(conn), nil
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

// func newRedisSvcClient() (redisSvcV1.RedisServiceClient, error) {
// 	conn, err := grpc.DialContext(context.TODO(), address["redisAddress"], grpc.WithInsecure())
// 	if err != nil {
// 		return nil, fmt.Errorf("redis client: %w", err)
// 	}

// 	return redisSvcV1.NewRedisServiceClient(conn), nil
// }

func logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("APIGW service: method %q called\n", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("APIGW serivce: method %q failed: %s\n", info.FullMethod, err)
	}
	return resp, err
}

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
	// s := grpc.NewServer(grpc.UnaryInterceptor(service.NewAuthorizationService(userClient, authorizationClient).Authorize))
	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	pb.RegisterClassroomServiceServer(s, service.NewClassroomsService(classroomClient, postClient, exerciseClient, rpsClient, userClient, topicClient, waitingListClient, attachmentClient, submissionClient))
	pb.RegisterPostServiceServer(s, service.NewPostsService(postClient, classroomClient, rpsClient, commentClient, userClient, attachmentClient))
	pb.RegisterExerciseServiceServer(s, service.NewExercisesService(exerciseClient, classroomClient, rpsClient, commentClient, userClient, submissionClient, attachmentClient, scheduleClient))
	pb.RegisterReportingStageServiceServer(s, service.NewReportingStagesService(rpsClient))
	pb.RegisterSubmissionServiceServer(s, service.NewSubmissionsService(submissionClient, classroomClient, exerciseClient, attachmentClient, userClient))
	pb.RegisterUserServiceServer(s, service.NewUsersService(userClient, classroomClient, waitingListClient, topicClient, attachmentClient))
	pb.RegisterWaitingListServiceServer(s, service.NewWaitingListsService(waitingListClient, classroomClient, userClient))
	pb.RegisterCommentServiceServer(s, service.NewCommentsService(commentClient, postClient, exerciseClient, userClient))
	pb.RegisterAttachmentServiceServer(s, service.NewAttachmentsService(attachmentClient, userClient, submissionClient))
	pb.RegisterTopicServiceServer(s, service.NewTopicsService(topicClient, userClient))
	pb.RegisterAuthorizationServiceServer(s, service.NewAuthorizationService(userClient, authorizationClient))
	pb.RegisterMemberServiceServer(s, service.NewMembersService(userClient, classroomClient, waitingListClient))
	pb.RegisterCommiteeServiceServer(s, service.NewCommiteesService(commiteeClient, userClient))
	pb.RegisterRoomServiceServer(s, service.NewRoomsService(commiteeClient))
	pb.RegisterStudentDefServiceServer(s, service.NewStudentDefsService(userClient))
	pb.RegisterScheduleServiceServer(s, service.NewSchedulesService(scheduleClient, commiteeClient, userClient, thesisClient))
	pb.RegisterNotificationServiceServer(s, service.NewNotificationsService(scheduleClient, userClient))
	pb.RegisterPointServiceServer(s, service.NewPointsService(scheduleClient, userClient))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
