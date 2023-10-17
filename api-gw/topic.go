package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	topicSvcV1 "github.com/qthuy2k1/thesis-management-backend/topic-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type topicServiceGW struct {
	pb.UnimplementedTopicServiceServer
	topicClient topicSvcV1.TopicServiceClient
	userClient  userSvcV1.UserServiceClient
}

func NewTopicsService(topicClient topicSvcV1.TopicServiceClient, userClient userSvcV1.UserServiceClient) *topicServiceGW {
	return &topicServiceGW{
		topicClient: topicClient,
		userClient:  userClient,
	}
}

func (u *topicServiceGW) CreateTopic(ctx context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.topicClient.CreateTopic(ctx, &topicSvcV1.CreateTopicRequest{
		Topic: &topicSvcV1.TopicInput{
			Title:          req.GetTopic().GetTitle(),
			TypeTopic:      req.GetTopic().GetTypeTopic(),
			MemberQuantity: req.GetTopic().GetMemberQuantity(),
			StudentID:      req.GetTopic().GetStudentID(),
			MemberEmail:    req.GetTopic().GetMemberEmail(),
			Description:    req.GetTopic().GetDescription(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateTopicResponse{
		Response: &pb.CommonTopicResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
	}, nil
}

func (u *topicServiceGW) GetTopic(ctx context.Context, req *pb.GetTopicRequest) (*pb.GetTopicResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.topicClient.GetTopic(ctx, &topicSvcV1.GetTopicRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	studentRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: res.GetTopic().StudentID,
	})
	if err != nil {
		return nil, err
	}

	major := studentRes.GetUser().GetMajor()
	phone := studentRes.GetUser().GetPhone()

	return &pb.GetTopicResponse{
		Response: &pb.CommonTopicResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
		Topic: &pb.TopicResponse{
			Id:             res.GetTopic().GetId(),
			Title:          res.GetTopic().GetTitle(),
			TypeTopic:      res.GetTopic().GetTypeTopic(),
			MemberQuantity: res.GetTopic().GetMemberQuantity(),
			StudentID: &pb.UserTopicResponse{
				Id:       studentRes.GetUser().GetId(),
				Class:    studentRes.GetUser().GetClass(),
				Major:    &major,
				Phone:    &phone,
				PhotoSrc: studentRes.GetUser().GetPhotoSrc(),
				Role:     studentRes.GetUser().GetRole(),
				Name:     studentRes.GetUser().GetName(),
				Email:    studentRes.GetUser().GetEmail(),
			},
			MemberEmail: res.GetTopic().GetMemberEmail(),
			Description: res.GetTopic().GetDescription(),
		},
	}, nil
}

func (u *topicServiceGW) UpdateTopic(ctx context.Context, req *pb.UpdateTopicRequest) (*pb.UpdateTopicResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.topicClient.UpdateTopic(ctx, &topicSvcV1.UpdateTopicRequest{
		Id: req.GetId(),
		Topic: &topicSvcV1.TopicInput{
			Title:          req.GetTopic().GetTitle(),
			TypeTopic:      req.GetTopic().GetTypeTopic(),
			MemberQuantity: req.GetTopic().GetMemberQuantity(),
			StudentID:      req.GetTopic().GetStudentID(),
			MemberEmail:    req.GetTopic().GetMemberEmail(),
			Description:    req.GetTopic().GetDescription(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTopicResponse{
		Response: &pb.CommonTopicResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
	}, nil
}

func (u *topicServiceGW) DeleteTopic(ctx context.Context, req *pb.DeleteTopicRequest) (*pb.DeleteTopicResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.topicClient.DeleteTopic(ctx, &topicSvcV1.DeleteTopicRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTopicResponse{
		Response: &pb.CommonTopicResponse{
			StatusCode: res.GetResponse().GetStatusCode(),
			Message:    res.GetResponse().GetMessage(),
		},
	}, nil
}
