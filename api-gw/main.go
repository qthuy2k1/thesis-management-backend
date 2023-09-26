package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	attachmentSvcV1 "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	waitingListSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	commentSvcV1 "github.com/qthuy2k1/thesis-management-backend/comment-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	rpsSvcV1 "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	submissionSvcV1 "github.com/qthuy2k1/thesis-management-backend/submission-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

const (
	listenAddress         = "0.0.0.0:9091"
	classroomAddress      = "classroom:9091"
	postAddress           = "post:9091"
	exerciseAddress       = "exercise:9091"
	reportingStageAddress = "reporting-stage:9091"
	submissionAddress     = "submission:9091"
	userAddress           = "user:9091"
	waitingListAddress    = "classroom-waiting-list:9091"
	commentAddress        = "comment:9091"
	attachmentAddress     = "attachment:9091"
)

func newClassroomSvcClient() (classroomSvcV1.ClassroomServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), classroomAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("classroom client: %w", err)
	}

	return classroomSvcV1.NewClassroomServiceClient(conn), nil
}

func newPostSvcClient() (postSvcV1.PostServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), postAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post client: %w", err)
	}

	return postSvcV1.NewPostServiceClient(conn), nil
}

func newExerciseSvcClient() (exerciseSvcV1.ExerciseServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), exerciseAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("exercise client: %w", err)
	}

	return exerciseSvcV1.NewExerciseServiceClient(conn), nil
}

func newReportingStageSvcClient() (rpsSvcV1.ReportingStageServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), reportingStageAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("reporting stage client: %w", err)
	}

	return rpsSvcV1.NewReportingStageServiceClient(conn), nil
}

func newSubmissionSvcClient() (submissionSvcV1.SubmissionServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), submissionAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("submission client: %w", err)
	}

	return submissionSvcV1.NewSubmissionServiceClient(conn), nil
}

func newUserSvcClient() (userSvcV1.UserServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), userAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("user client: %w", err)
	}

	return userSvcV1.NewUserServiceClient(conn), nil
}

func newWaitingListSvcClient() (waitingListSvcV1.WaitingListServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), waitingListAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("clasroom waiting list client: %w", err)
	}

	return waitingListSvcV1.NewWaitingListServiceClient(conn), nil
}

func newCommentSvcClient() (commentSvcV1.CommentServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), commentAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("comment client: %w", err)
	}

	return commentSvcV1.NewCommentServiceClient(conn), nil
}

func newAttachmentSvcClient() (attachmentSvcV1.AttachmentServiceClient, error) {
	conn, err := grpc.DialContext(context.TODO(), attachmentAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("attachment client: %w", err)
	}

	return attachmentSvcV1.NewAttachmentServiceClient(conn), nil
}

func logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("APIGW service: method %q called\n", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("APIGW serivce: method %q failed: %s\n", info.FullMethod, err)
	}
	return resp, err
}

func main() {
	log.Printf("APIGW service starting on %s", listenAddress)

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

	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(logger))

	pb.RegisterClassroomServiceServer(s, NewClassroomsService(classroomClient, postClient, exerciseClient, rpsClient, userClient))
	pb.RegisterPostServiceServer(s, NewPostsService(postClient, classroomClient, rpsClient, commentClient, userClient))
	pb.RegisterExerciseServiceServer(s, NewExercisesService(exerciseClient, classroomClient, rpsClient, commentClient, userClient, submissionClient, attachmentClient))
	pb.RegisterReportingStageServiceServer(s, NewReportingStagesService(rpsClient))
	pb.RegisterSubmissionServiceServer(s, NewSubmissionsService(submissionClient, classroomClient, exerciseClient))
	pb.RegisterUserServiceServer(s, NewUsersService(userClient, classroomClient, waitingListClient))
	pb.RegisterWaitingListServiceServer(s, NewWaitingListsService(waitingListClient, classroomClient, userClient))
	pb.RegisterCommentServiceServer(s, NewCommentsService(commentClient, postClient, exerciseClient, userClient))
	pb.RegisterAuthorizationServiceServer(s, NewAuthService(userClient))
	pb.RegisterAttachmentServiceServer(s, NewAttachmentsService(attachmentClient, userClient, submissionClient))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
