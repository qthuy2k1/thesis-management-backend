package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
)

type postServiceGW struct {
	pb.UnimplementedPostServiceServer
	postClient postSvcV1.PostServiceClient
}

func NewPostsService(postClient postSvcV1.PostServiceClient) *postServiceGW {
	return &postServiceGW{
		postClient: postClient,
	}
}

func (u *postServiceGW) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	res, err := u.postClient.CreatePost(ctx, &postSvcV1.CreatePostRequest{
		Post: &postSvcV1.PostInput{
			Title:       req.GetPost().Title,
			Content:     req.GetPost().Content,
			ClassroomID: req.GetPost().ClassroomID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreatePostResponse{
		StatusCode: res.StatusCode,
		Message:    res.Message,
	}, nil
}

func (u *postServiceGW) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	res, err := u.postClient.GetPost(ctx, &postSvcV1.GetPostRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetPostResponse{
		StatusCode: res.StatusCode,
		Message:    res.Message,
		Post: &pb.PostResponse{
			Id:          res.GetPost().Id,
			Title:       res.GetPost().Title,
			Content:     res.GetPost().Content,
			ClassroomID: res.GetPost().ClassroomID,
			CreatedAt:   res.GetPost().CreatedAt,
			UpdatedAt:   res.GetPost().UpdatedAt,
		},
	}, nil
}
