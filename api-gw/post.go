package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
)

type postServiceGW struct {
	pb.UnimplementedPostServiceServer
	postClient      postSvcV1.PostServiceClient
	classroomClient classroomSvcV1.ClassroomServiceClient
}

func NewPostsService(postClient postSvcV1.PostServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient) *postServiceGW {
	return &postServiceGW{
		postClient:      postClient,
		classroomClient: classroomClient,
	}
}

func (u *postServiceGW) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetPost().ClassroomID})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.CreatePostResponse{
			StatusCode: 400,
			Message:    "Classroom does not exist",
		}, nil
	}

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

// func (u *postServiceGW) CheckClassroomExists(ctx context.Context, req *pb.CheckClassroomExistsRequest) (*pb.CheckClassroomExistsResponse, error) {
// 	res, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetClassroomID()})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.CheckClassroomExistsResponse{
// 		Exists: res.GetExists(),
// 	}, nil
// }
