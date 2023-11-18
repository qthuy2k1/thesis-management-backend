package main

import (
	"context"
	"errors"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	redisSvcV1 "github.com/qthuy2k1/thesis-management-backend/redis-svc/api/goclient/v1"
	topicSvcV1 "github.com/qthuy2k1/thesis-management-backend/topic-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type topicServiceGW struct {
	pb.UnimplementedTopicServiceServer
	topicClient topicSvcV1.TopicServiceClient
	userClient  userSvcV1.UserServiceClient
	redisClient redisSvcV1.RedisServiceClient
}

func NewTopicsService(topicClient topicSvcV1.TopicServiceClient, userClient userSvcV1.UserServiceClient, redisClient redisSvcV1.RedisServiceClient) *topicServiceGW {
	return &topicServiceGW{
		topicClient: topicClient,
		userClient:  userClient,
		redisClient: redisClient,
	}
}

func (u *topicServiceGW) CreateTopic(ctx context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
		Id: req.Topic.StudentID,
	})
	if err != nil {
		return nil, err
	}

	var studentRes *userSvcV1.GetUserResponse
	if redis.User != nil && redis.GetResponse().StatusCode == 200 {
		studentRes = &userSvcV1.GetUserResponse{
			Response: &userSvcV1.CommonUserResponse{
				StatusCode: 200,
				Message:    "OK",
			},
			User: &userSvcV1.UserResponse{
				Id:       redis.User.GetId(),
				Class:    redis.User.Class,
				Major:    redis.User.Major,
				Phone:    redis.User.Phone,
				PhotoSrc: redis.User.GetPhotoSrc(),
				Role:     redis.User.GetRole(),
				Name:     redis.User.GetName(),
				Email:    redis.User.GetEmail(),
			},
		}
	} else {
		studentRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.Topic.StudentID})
		if err != nil {
			return nil, err
		}

		if studentRes.Response.StatusCode != 200 {
			return nil, errors.New("error getting user")
		}

		cache, err := u.redisClient.SetUser(ctx, &redisSvcV1.SetUserRequest{
			User: &redisSvcV1.User{
				Id:       studentRes.User.GetId(),
				Class:    studentRes.User.Class,
				Major:    studentRes.User.Major,
				Phone:    studentRes.User.Major,
				PhotoSrc: studentRes.User.GetPhotoSrc(),
				Role:     studentRes.User.GetRole(),
				Name:     studentRes.User.GetName(),
				Email:    studentRes.User.GetEmail(),
			},
		})
		if err != nil {
			return nil, err
		}

		if cache.Response.StatusCode != 200 {
			return nil, errors.New("error set user cache")
		}
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

	res, err := u.topicClient.GetTopicFromUser(ctx, &topicSvcV1.GetTopicFromUserRequest{UserID: req.StudentID})
	if err != nil {
		return nil, err
	}

	redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
		Id: res.Topic.StudentID,
	})
	if err != nil {
		return nil, err
	}

	var studentRes *userSvcV1.GetUserResponse
	if redis.User != nil && redis.GetResponse().StatusCode == 200 {
		studentRes = &userSvcV1.GetUserResponse{
			Response: &userSvcV1.CommonUserResponse{
				StatusCode: 200,
				Message:    "OK",
			},
			User: &userSvcV1.UserResponse{
				Id:       redis.User.GetId(),
				Class:    redis.User.Class,
				Major:    redis.User.Major,
				Phone:    redis.User.Phone,
				PhotoSrc: redis.User.GetPhotoSrc(),
				Role:     redis.User.GetRole(),
				Name:     redis.User.GetName(),
				Email:    redis.User.GetEmail(),
			},
		}
	} else {
		studentRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: res.Topic.StudentID})
		if err != nil {
			return nil, err
		}

		if studentRes.Response.StatusCode != 200 {
			return nil, errors.New("error getting user")
		}

		cache, err := u.redisClient.SetUser(ctx, &redisSvcV1.SetUserRequest{
			User: &redisSvcV1.User{
				Id:       studentRes.User.GetId(),
				Class:    studentRes.User.Class,
				Major:    studentRes.User.Major,
				Phone:    studentRes.User.Major,
				PhotoSrc: studentRes.User.GetPhotoSrc(),
				Role:     studentRes.User.GetRole(),
				Name:     studentRes.User.GetName(),
				Email:    studentRes.User.GetEmail(),
			},
		})
		if err != nil {
			return nil, err
		}

		if cache.Response.StatusCode != 200 {
			return nil, errors.New("error set user cache")
		}
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

func (u *topicServiceGW) GetTopics(ctx context.Context, req *pb.GetTopicsRequest) (*pb.GetTopicsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.topicClient.GetTopics(ctx, &topicSvcV1.GetTopicsRequest{})
	if err != nil {
		return nil, err
	}

	var topics []*pb.TopicResponse
	for _, p := range res.Topic {
		redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
			Id: p.StudentID,
		})
		if err != nil {
			return nil, err
		}

		var studentRes *userSvcV1.GetUserResponse
		if redis.User != nil && redis.GetResponse().StatusCode == 200 {
			studentRes = &userSvcV1.GetUserResponse{
				Response: &userSvcV1.CommonUserResponse{
					StatusCode: 200,
					Message:    "OK",
				},
				User: &userSvcV1.UserResponse{
					Id:       redis.User.GetId(),
					Class:    redis.User.Class,
					Major:    redis.User.Major,
					Phone:    redis.User.Phone,
					PhotoSrc: redis.User.GetPhotoSrc(),
					Role:     redis.User.GetRole(),
					Name:     redis.User.GetName(),
					Email:    redis.User.GetEmail(),
				},
			}
		} else {
			studentRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: p.StudentID})
			if err != nil {
				return nil, err
			}

			if studentRes.Response.StatusCode != 200 {
				return nil, errors.New("error getting user")
			}

			cache, err := u.redisClient.SetUser(ctx, &redisSvcV1.SetUserRequest{
				User: &redisSvcV1.User{
					Id:       studentRes.User.GetId(),
					Class:    studentRes.User.Class,
					Major:    studentRes.User.Major,
					Phone:    studentRes.User.Major,
					PhotoSrc: studentRes.User.GetPhotoSrc(),
					Role:     studentRes.User.GetRole(),
					Name:     studentRes.User.GetName(),
					Email:    studentRes.User.GetEmail(),
				},
			})
			if err != nil {
				return nil, err
			}

			if cache.Response.StatusCode != 200 {
				return nil, errors.New("error set user cache")
			}
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
