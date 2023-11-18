package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	redisSvcV1 "github.com/qthuy2k1/thesis-management-backend/redis-svc/api/goclient/v1"
	pointSvcV1 "github.com/qthuy2k1/thesis-management-backend/schedule-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type pointServiceGW struct {
	pb.UnimplementedPointServiceServer
	pointClient pointSvcV1.ScheduleServiceClient
	userClient  userSvcV1.UserServiceClient
	redisClient redisSvcV1.RedisServiceClient
}

func NewPointsService(pointClient pointSvcV1.ScheduleServiceClient, userClient userSvcV1.UserServiceClient, redisClient redisSvcV1.RedisServiceClient) *pointServiceGW {
	return &pointServiceGW{
		pointClient: pointClient,
		userClient:  userClient,
		redisClient: redisClient,
	}
}

func (u *pointServiceGW) GetAllPointDef(ctx context.Context, req *pb.GetAllPointDefRequest) (*pb.GetAllPointDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.pointClient.GetAllPointDefs(ctx, &pointSvcV1.GetAllPointDefsRequest{})
	if err != nil {
		return nil, err
	}

	var pointRes []*pb.PointResponse
	for _, t := range res.Points {
		var assessItems []*pb.AssessItemResponse
		for _, c := range t.Assesses {
			assessItems = append(assessItems, &pb.AssessItemResponse{
				Id: c.Id,
				Lecturer: &pb.UserPointResponse{
					Id:       c.Lecturer.GetId(),
					Class:    &c.Lecturer.Class,
					Major:    c.Lecturer.Major,
					Phone:    c.Lecturer.Phone,
					PhotoSrc: c.Lecturer.GetPhotoSrc(),
					Name:     c.Lecturer.GetName(),
					Email:    c.Lecturer.GetEmail(),
					Role:     c.Lecturer.GetRole(),
				},
				Point:   c.Point,
				Comment: c.Comment,
			})
		}

		pointRes = append(pointRes, &pb.PointResponse{
			Id: t.Id,
			Student: &pb.UserPointResponse{
				Id:       t.Student.GetId(),
				Class:    &t.Student.Class,
				Major:    t.Student.Major,
				Phone:    t.Student.Phone,
				PhotoSrc: t.Student.GetPhotoSrc(),
				Name:     t.Student.GetName(),
				Email:    t.Student.GetEmail(),
				Role:     t.Student.GetRole(),
			},
			Assesses: assessItems,
		})
	}

	return &pb.GetAllPointDefResponse{
		Points: pointRes,
	}, nil
}

func (u *pointServiceGW) CreateOrUpdatePointDef(ctx context.Context, req *pb.CreateOrUpdatePointDefRequest) (*pb.CreateOrUpdatePointDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
		Id: req.Point.StudentID,
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
		studentRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: req.Point.StudentID,
		})
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

	if studentRes.Response.StatusCode == 404 {
		return nil, errors.New("user not found")
	}

	var assess []*pointSvcV1.AssessItem
	for _, a := range req.Point.Assesses {
		redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
			Id: a.LecturerID,
		})
		if err != nil {
			return nil, err
		}

		var lecturerRes *userSvcV1.GetUserResponse
		if redis.User != nil && redis.GetResponse().StatusCode == 200 {
			lecturerRes = &userSvcV1.GetUserResponse{
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
			lecturerRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: a.LecturerID})
			if err != nil {
				return nil, err
			}

			if lecturerRes.Response.StatusCode != 200 {
				return nil, errors.New("error getting user")
			}

			cache, err := u.redisClient.SetUser(ctx, &redisSvcV1.SetUserRequest{
				User: &redisSvcV1.User{
					Id:       lecturerRes.User.GetId(),
					Class:    lecturerRes.User.Class,
					Major:    lecturerRes.User.Major,
					Phone:    lecturerRes.User.Major,
					PhotoSrc: lecturerRes.User.GetPhotoSrc(),
					Role:     lecturerRes.User.GetRole(),
					Name:     lecturerRes.User.GetName(),
					Email:    lecturerRes.User.GetEmail(),
				},
			})
			if err != nil {
				return nil, err
			}

			if cache.Response.StatusCode != 200 {
				return nil, errors.New("error set user cache")
			}
		}

		assess = append(assess, &pointSvcV1.AssessItem{
			Id: a.Id,
			Lecturer: &pointSvcV1.UserScheduleResponse{
				Id:       lecturerRes.User.Id,
				Class:    lecturerRes.User.GetClass(),
				Major:    lecturerRes.User.Major,
				Phone:    lecturerRes.User.Phone,
				PhotoSrc: lecturerRes.User.PhotoSrc,
				Role:     lecturerRes.User.Role,
				Name:     lecturerRes.User.Name,
				Email:    lecturerRes.User.Email,
			},
			Point:   a.Point,
			Comment: a.Comment,
		})
	}

	res, err := u.pointClient.CreateOrUpdatePointDef(ctx, &pointSvcV1.CreateOrUpdatePointDefRequest{
		Point: &pointSvcV1.Point{
			Id: req.Point.Id,
			Student: &pointSvcV1.UserScheduleResponse{
				Id:       studentRes.User.Id,
				Class:    studentRes.User.GetClass(),
				Major:    studentRes.User.Major,
				Phone:    studentRes.User.Phone,
				PhotoSrc: studentRes.User.PhotoSrc,
				Role:     studentRes.User.Role,
				Name:     studentRes.User.Name,
				Email:    studentRes.User.Email,
			},
			Assesses: assess,
		},
	})
	if err != nil {
		log.Println("create error", err)
		return nil, err
	}

	var assessRes []*pb.AssessItemResponse
	for _, a := range res.Point.Assesses {
		assessRes = append(assessRes, &pb.AssessItemResponse{
			Id: a.Id,
			Lecturer: &pb.UserPointResponse{
				Id:       a.Lecturer.Id,
				Class:    &a.Lecturer.Class,
				Major:    a.Lecturer.Major,
				Phone:    a.Lecturer.Phone,
				PhotoSrc: a.Lecturer.PhotoSrc,
				Role:     a.Lecturer.Role,
				Name:     a.Lecturer.Name,
				Email:    a.Lecturer.Email,
			},
			Point:   a.Point,
			Comment: a.Comment,
		})
	}

	return &pb.CreateOrUpdatePointDefResponse{
		Point: &pb.PointResponse{
			Id: res.Point.Id,
			Student: &pb.UserPointResponse{
				Id:       res.Point.Student.Id,
				Class:    &res.Point.Student.Class,
				Major:    res.Point.Student.Major,
				Phone:    res.Point.Student.Phone,
				PhotoSrc: res.Point.Student.PhotoSrc,
				Role:     res.Point.Student.Role,
				Name:     res.Point.Student.Name,
				Email:    res.Point.Student.Email,
			},
			Assesses: assessRes,
		},
		Message: res.Message,
	}, nil

}
