package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	commentSvcV1 "github.com/qthuy2k1/thesis-management-backend/comment-svc/api/goclient/v1"
	exerciseSvcV1 "github.com/qthuy2k1/thesis-management-backend/exercise-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type commentServiceGW struct {
	pb.UnimplementedCommentServiceServer
	commentClient  commentSvcV1.CommentServiceClient
	postClient     postSvcV1.PostServiceClient
	exerciseClient exerciseSvcV1.ExerciseServiceClient
	userClient     userSvcV1.UserServiceClient
}

func NewCommentsService(commentClient commentSvcV1.CommentServiceClient, postClient postSvcV1.PostServiceClient, exerciseClient exerciseSvcV1.ExerciseServiceClient, userClient userSvcV1.UserServiceClient) *commentServiceGW {
	return &commentServiceGW{
		commentClient:  commentClient,
		postClient:     postClient,
		exerciseClient: exerciseClient,
		userClient:     userClient,
	}
}

func (u *commentServiceGW) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.GetComment().PostID == nil && req.GetComment().ExerciseID == nil {
		return &pb.CreateCommentResponse{
			Response: &pb.CommonCommentResponse{
				StatusCode: 400,
				Message:    "Post ID or Exercise ID is missing, requiring at least one",
			},
		}, nil
	}

	if req.GetComment().PostID != nil {
		postRes, err := u.postClient.GetPost(ctx, &postSvcV1.GetPostRequest{
			Id: *req.GetComment().PostID,
		})
		if err != nil {
			return nil, err
		}

		if postRes.GetResponse().GetStatusCode() == 404 {
			return &pb.CreateCommentResponse{
				Response: &pb.CommonCommentResponse{
					StatusCode: 404,
					Message:    "Post is not found",
				},
			}, nil
		}
	} else {
		exerciseRes, err := u.exerciseClient.GetExercise(ctx, &exerciseSvcV1.GetExerciseRequest{
			Id: *req.GetComment().ExerciseID,
		})
		if err != nil {
			return nil, err
		}

		if exerciseRes.GetResponse().GetStatusCode() == 404 {
			return &pb.CreateCommentResponse{
				Response: &pb.CommonCommentResponse{
					StatusCode: 404,
					Message:    "Exercise is not found",
				},
			}, nil
		}
	}

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetComment().GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	if userRes.GetResponse().GetStatusCode() == 404 {
		return &pb.CreateCommentResponse{
			Response: &pb.CommonCommentResponse{
				StatusCode: 404,
				Message:    "User is not found",
			},
		}, nil
	}

	res, err := u.commentClient.CreateComment(ctx, &commentSvcV1.CreateCommentRequest{
		Comment: &commentSvcV1.CommentInput{
			UserID:     req.GetComment().GetUserID(),
			PostID:     req.GetComment().PostID,
			ExerciseID: req.GetComment().ExerciseID,
			Content:    req.GetComment().GetContent(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentResponse{
		Response: &pb.CommonCommentResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}
