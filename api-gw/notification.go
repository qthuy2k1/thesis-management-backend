package main

import (
	"context"
	"errors"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	redisSvcV1 "github.com/qthuy2k1/thesis-management-backend/redis-svc/api/goclient/v1"
	notificationSvcV1 "github.com/qthuy2k1/thesis-management-backend/schedule-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type notificationServiceGW struct {
	pb.UnimplementedNotificationServiceServer
	notificationClient notificationSvcV1.ScheduleServiceClient
	userClient         userSvcV1.UserServiceClient
	redisClient        redisSvcV1.RedisServiceClient
}

func NewNotificationsService(notificationClient notificationSvcV1.ScheduleServiceClient, userClient userSvcV1.UserServiceClient, redisClient redisSvcV1.RedisServiceClient) *notificationServiceGW {
	return &notificationServiceGW{
		notificationClient: notificationClient,
		userClient:         userClient,
		redisClient:        redisClient,
	}
}

func (u *notificationServiceGW) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
		Id: req.Noti.SenderUserID,
	})
	if err != nil {
		return nil, err
	}

	var sender *userSvcV1.GetUserResponse
	if redis.User != nil && redis.GetResponse().StatusCode == 200 {
		sender = &userSvcV1.GetUserResponse{
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
		sender, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.Noti.SenderUserID})
		if err != nil {
			return nil, err
		}

		if sender.Response.StatusCode != 200 {
			return nil, errors.New("error getting user")
		}

		cache, err := u.redisClient.SetUser(ctx, &redisSvcV1.SetUserRequest{
			User: &redisSvcV1.User{
				Id:       sender.User.GetId(),
				Class:    sender.User.Class,
				Major:    sender.User.Major,
				Phone:    sender.User.Major,
				PhotoSrc: sender.User.GetPhotoSrc(),
				Role:     sender.User.GetRole(),
				Name:     sender.User.GetName(),
				Email:    sender.User.GetEmail(),
			},
		})
		if err != nil {
			return nil, err
		}

		if cache.Response.StatusCode != 200 {
			return nil, errors.New("error set user cache")
		}
	}

	redis, err = u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
		Id: req.Noti.ReceiverAuthorID,
	})
	if err != nil {
		return nil, err
	}

	var receiver *userSvcV1.GetUserResponse
	if redis.User != nil && redis.GetResponse().StatusCode == 200 {
		receiver = &userSvcV1.GetUserResponse{
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
		receiver, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.Noti.ReceiverAuthorID})
		if err != nil {
			return nil, err
		}

		if receiver.Response.StatusCode != 200 {
			return nil, errors.New("error getting user")
		}

		cache, err := u.redisClient.SetUser(ctx, &redisSvcV1.SetUserRequest{
			User: &redisSvcV1.User{
				Id:       receiver.User.GetId(),
				Class:    receiver.User.Class,
				Major:    receiver.User.Major,
				Phone:    receiver.User.Major,
				PhotoSrc: receiver.User.GetPhotoSrc(),
				Role:     receiver.User.GetRole(),
				Name:     receiver.User.GetName(),
				Email:    receiver.User.GetEmail(),
			},
		})
		if err != nil {
			return nil, err
		}

		if cache.Response.StatusCode != 200 {
			return nil, errors.New("error set user cache")
		}
	}

	res, err := u.notificationClient.CreateNotification(ctx, &notificationSvcV1.CreateNotificationRequest{
		Noti: &notificationSvcV1.Notification{
			Id: req.Noti.Id,
			SenderUser: &notificationSvcV1.UserScheduleResponse{
				Id:       sender.User.GetId(),
				Class:    sender.User.GetClass(),
				Major:    sender.User.Major,
				Phone:    sender.User.Phone,
				PhotoSrc: sender.User.GetPhotoSrc(),
				Name:     sender.User.GetName(),
				Email:    sender.User.GetEmail(),
				Role:     sender.User.GetRole(),
			},
			ReceiverAuthor: &notificationSvcV1.UserScheduleResponse{
				Id:       receiver.User.GetId(),
				Class:    receiver.User.GetClass(),
				Major:    receiver.User.Major,
				Phone:    receiver.User.Phone,
				PhotoSrc: receiver.User.GetPhotoSrc(),
				Name:     receiver.User.GetName(),
				Email:    receiver.User.GetEmail(),
				Role:     receiver.User.GetRole(),
			},
			Type: req.Noti.Type,
		},
	})
	if err != nil {
		return nil, err
	}

	var notificationsRes []*pb.NotificationResponse
	for _, n := range res.Notifications {
		notificationsRes = append(notificationsRes, &pb.NotificationResponse{
			Id: n.Id,
			SenderUser: &pb.UserNotificationResponse{
				Id:       n.SenderUser.GetId(),
				Class:    &n.SenderUser.Class,
				Major:    n.SenderUser.Major,
				Phone:    n.SenderUser.Phone,
				PhotoSrc: n.SenderUser.GetPhotoSrc(),
				Name:     n.SenderUser.GetName(),
				Email:    n.SenderUser.GetEmail(),
				Role:     n.SenderUser.GetRole(),
			},
			ReceiverAuthor: &pb.UserNotificationResponse{
				Id:       n.ReceiverAuthor.GetId(),
				Class:    &n.ReceiverAuthor.Class,
				Major:    n.ReceiverAuthor.Major,
				Phone:    n.ReceiverAuthor.Phone,
				PhotoSrc: n.ReceiverAuthor.GetPhotoSrc(),
				Name:     n.ReceiverAuthor.GetName(),
				Email:    n.ReceiverAuthor.GetEmail(),
				Role:     n.ReceiverAuthor.GetRole(),
			},
			Type:      res.Notification.Type,
			CreatedAt: res.Notification.CreatedAt,
		})
	}

	return &pb.CreateNotificationResponse{
		Notification: &pb.NotificationResponse{
			Id: res.Notification.Id,
			SenderUser: &pb.UserNotificationResponse{
				Id:       res.Notification.SenderUser.GetId(),
				Class:    &res.Notification.SenderUser.Class,
				Major:    res.Notification.SenderUser.Major,
				Phone:    res.Notification.SenderUser.Phone,
				PhotoSrc: res.Notification.SenderUser.GetPhotoSrc(),
				Name:     res.Notification.SenderUser.GetName(),
				Email:    res.Notification.SenderUser.GetEmail(),
				Role:     res.Notification.SenderUser.GetRole(),
			},
			ReceiverAuthor: &pb.UserNotificationResponse{
				Id:       res.Notification.ReceiverAuthor.GetId(),
				Class:    &res.Notification.ReceiverAuthor.Class,
				Major:    res.Notification.ReceiverAuthor.Major,
				Phone:    res.Notification.ReceiverAuthor.Phone,
				PhotoSrc: res.Notification.ReceiverAuthor.GetPhotoSrc(),
				Name:     res.Notification.ReceiverAuthor.GetName(),
				Email:    res.Notification.ReceiverAuthor.GetEmail(),
				Role:     res.Notification.ReceiverAuthor.GetRole(),
			},
			Type:      res.Notification.Type,
			CreatedAt: res.Notification.CreatedAt,
		},
		Message:       res.Message,
		Notifications: notificationsRes,
	}, nil
}

func (u *notificationServiceGW) GetNotifications(ctx context.Context, req *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	panic("not implemented")
}
