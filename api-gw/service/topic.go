package service

import (
	"context"
	"log"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type topicServiceGW struct {
	pb.UnimplementedTopicServiceServer
	topicClient userSvcV1.TopicServiceClient
	userClient  userSvcV1.UserServiceClient
}

func NewTopicsService(topicClient userSvcV1.TopicServiceClient, userClient userSvcV1.UserServiceClient) *topicServiceGW {
	return &topicServiceGW{
		topicClient: topicClient,
		userClient:  userClient,
	}
}

func (u *topicServiceGW) CreateTopic(ctx context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	log.Println("TOPIC", req)

	studentRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.Topic.StudentID})
	if err != nil {
		return nil, err
	}

	if studentRes.Response.StatusCode == 404 {
		return &pb.CreateTopicResponse{
			Response: &pb.CommonTopicResponse{
				StatusCode: studentRes.Response.StatusCode,
				Message:    studentRes.Response.Message,
			},
		}, nil
	}

	res, err := u.topicClient.CreateTopic(ctx, &userSvcV1.CreateTopicRequest{
		Topic: &userSvcV1.TopicInput{
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

	res, err := u.topicClient.GetTopicFromUser(ctx, &userSvcV1.GetTopicFromUserRequest{UserID: req.StudentID})
	if err != nil {
		return nil, err
	}

	studentRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: res.Topic.StudentID})
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
			Student: &pb.UserTopicResponse{
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

	res, err := u.topicClient.UpdateTopic(ctx, &userSvcV1.UpdateTopicRequest{
		Id: req.GetId(),
		Topic: &userSvcV1.TopicInput{
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

	res, err := u.topicClient.DeleteTopic(ctx, &userSvcV1.DeleteTopicRequest{
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

func (u *topicServiceGW) GetTopics(ctx context.Context, req *pb.GetTopicsRequest) (*pb.GetTopicsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.topicClient.GetTopics(ctx, &userSvcV1.GetTopicsRequest{})
	if err != nil {
		return nil, err
	}

	var topics []*pb.TopicResponse
	for _, p := range res.Topic {

		studentRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: p.StudentID})
		if err != nil {
			return nil, err
		}

		major := studentRes.GetUser().GetMajor()
		phone := studentRes.GetUser().GetPhone()
		topics = append(topics, &pb.TopicResponse{
			Id:             p.Id,
			Title:          p.GetTitle(),
			TypeTopic:      p.GetTypeTopic(),
			MemberQuantity: p.GetMemberQuantity(),
			Student: &pb.UserTopicResponse{
				Id:       studentRes.GetUser().GetId(),
				Class:    studentRes.GetUser().GetClass(),
				Major:    &major,
				Phone:    &phone,
				PhotoSrc: studentRes.GetUser().GetPhotoSrc(),
				Role:     studentRes.GetUser().GetRole(),
				Name:     studentRes.GetUser().GetName(),
				Email:    studentRes.GetUser().GetEmail(),
			},
			MemberEmail: p.GetMemberEmail(),
			Description: p.GetDescription(),
		})
	}

	return &pb.GetTopicsResponse{
		Topics: topics,
	}, nil
}
