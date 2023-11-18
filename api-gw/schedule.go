package main

import (
	"context"
	"errors"
	"log"
	"strconv"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	redisSvcV1 "github.com/qthuy2k1/thesis-management-backend/redis-svc/api/goclient/v1"
	scheduleSvcV1 "github.com/qthuy2k1/thesis-management-backend/schedule-svc/api/goclient/v1"
	commiteeSvcV1 "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type scheduleServiceGW struct {
	pb.UnimplementedScheduleServiceServer
	scheduleClient scheduleSvcV1.ScheduleServiceClient
	commiteeClient commiteeSvcV1.CommiteeServiceClient
	thesisClient   commiteeSvcV1.ScheduleServiceClient
	userClient     userSvcV1.UserServiceClient
	redisClient    redisSvcV1.RedisServiceClient
}

func NewSchedulesService(scheduleClient scheduleSvcV1.ScheduleServiceClient, commiteeClient commiteeSvcV1.CommiteeServiceClient, userClient userSvcV1.UserServiceClient, thesisClient commiteeSvcV1.ScheduleServiceClient, redisClient redisSvcV1.RedisServiceClient) *scheduleServiceGW {
	return &scheduleServiceGW{
		scheduleClient: scheduleClient,
		commiteeClient: commiteeClient,
		userClient:     userClient,
		thesisClient:   thesisClient,
		redisClient:    redisClient,
	}
}

func (u *scheduleServiceGW) GetSchedules(ctx context.Context, req *pb.GetSchedulesRequest) (*pb.GetSchedulesResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.scheduleClient.GetSchedules(ctx, &scheduleSvcV1.GetSchedulesRequest{})
	if err != nil {
		return nil, err
	}

	var thesisRes []*pb.Thesis
	for _, t := range res.Thesis {
		var councilsScheduleResponse []*pb.Council
		for _, c := range t.Council {
			councilsScheduleResponse = append(councilsScheduleResponse, &pb.Council{
				Id:       c.Id,
				Class:    c.Class,
				Major:    c.Major,
				Phone:    c.Phone,
				PhotoSrc: c.PhotoSrc,
				Role:     c.Role,
				Name:     c.Name,
				Email:    c.Email,
			})
		}

		var timeSlotsResponse []*pb.TimeSlots
		for _, t := range t.Schedule.TimeSlots {
			timeSlotsResponse = append(timeSlotsResponse, &pb.TimeSlots{
				Student: &pb.StudentDefScheduleResponse{
					Id: t.Student.Id,
					Infor: &pb.UserScheduleResponse{
						Id:       t.Student.Infor.GetId(),
						Class:    &t.Student.Infor.Class,
						Major:    t.Student.Infor.Major,
						Phone:    t.Student.Infor.Phone,
						PhotoSrc: t.Student.Infor.GetPhotoSrc(),
						Name:     t.Student.Infor.GetName(),
						Email:    t.Student.Infor.GetEmail(),
						Role:     t.Student.Infor.GetRole(),
					},
					Instructor: &pb.UserScheduleResponse{
						Id:       t.Student.Instructor.GetId(),
						Class:    &t.Student.Instructor.Class,
						Major:    t.Student.Instructor.Major,
						Phone:    t.Student.Instructor.Phone,
						PhotoSrc: t.Student.Instructor.GetPhotoSrc(),
						Name:     t.Student.Instructor.GetName(),
						Email:    t.Student.Instructor.GetEmail(),
						Role:     t.Student.Instructor.GetRole(),
					},
				},
				TimeSlot: &pb.TimeSlot{
					Date:  t.TimeSlot.Date,
					Shift: t.TimeSlot.Shift,
					Id:    t.TimeSlot.Id,
					Time:  t.TimeSlot.Time,
				},
			})
		}

		thesisRes = append(thesisRes, &pb.Thesis{
			Schedule: &pb.Schedule{
				TimeSlots: timeSlotsResponse,
				Room: &pb.RoomSchedule{
					Id:          t.Schedule.Room.Id,
					Name:        t.Schedule.Room.Name,
					School:      t.Schedule.Room.School,
					Type:        t.Schedule.Room.Type,
					Description: t.Schedule.Room.Description,
				},
			},
			Council: councilsScheduleResponse,
			Id:      t.Id,
		})
	}

	return &pb.GetSchedulesResponse{
		Thesis: thesisRes,
	}, nil
}

func (u *scheduleServiceGW) CreateSchedule(ctx context.Context, req *pb.CreateScheduleRequest) (*pb.CreateScheduleResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var rooms []*scheduleSvcV1.RoomSchedule
	roomRes, err := u.commiteeClient.GetRooms(ctx, &commiteeSvcV1.GetRoomsRequest{})
	if err != nil {
		return nil, err
	}

	for _, r := range roomRes.Rooms {
		rooms = append(rooms, &scheduleSvcV1.RoomSchedule{
			Id:          strconv.Itoa(int(r.Id)),
			Name:        r.Name,
			School:      r.School,
			Type:        r.Type,
			Description: r.Description,
			// CreatedAt: r.CreatedAt,
		})
	}

	var councils []*scheduleSvcV1.UserScheduleResponse
	councilRes, err := u.thesisClient.GetCouncils(ctx, &commiteeSvcV1.GetCouncilsRequest{})
	if err != nil {
		return nil, err
	}

	for _, r := range councilRes.Councils {
		// userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		// 	Id: r.LecturerID,
		// })
		// if err != nil {
		// 	return nil, err
		// }
		redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
			Id: r.LecturerID,
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
			lecturerRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: r.LecturerID})
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

		councils = append(councils, &scheduleSvcV1.UserScheduleResponse{
			Id:       lecturerRes.User.Id,
			Class:    lecturerRes.User.GetClass(),
			Major:    lecturerRes.User.Major,
			Phone:    lecturerRes.User.Phone,
			PhotoSrc: lecturerRes.User.PhotoSrc,
			Role:     lecturerRes.User.Role,
			Name:     lecturerRes.User.Name,
			Email:    lecturerRes.User.Email,
		})
	}

	var studentDefs []*scheduleSvcV1.StudentDefScheduleResponse
	stRes, err := u.userClient.GetStudentDefs(ctx, &userSvcV1.GetStudentDefsRequest{})
	if err != nil {
		return nil, err
	}

	for _, r := range stRes.StudentDefs {
		studentDefs = append(studentDefs, &scheduleSvcV1.StudentDefScheduleResponse{
			Id: strconv.Itoa(int(r.Id)),
			Infor: &scheduleSvcV1.UserScheduleResponse{
				Id:       r.User.GetId(),
				Class:    r.User.GetClass(),
				Major:    r.User.Major,
				Phone:    r.User.Phone,
				PhotoSrc: r.User.GetPhotoSrc(),
				Name:     r.User.GetName(),
				Email:    r.User.GetEmail(),
				Role:     r.User.GetRole(),
			},
			Instructor: &scheduleSvcV1.UserScheduleResponse{
				Id:       r.Instructor.GetId(),
				Class:    r.Instructor.GetClass(),
				Major:    r.Instructor.Major,
				Phone:    r.Instructor.Phone,
				PhotoSrc: r.Instructor.GetPhotoSrc(),
				Name:     r.Instructor.GetName(),
				Email:    r.Instructor.GetEmail(),
				Role:     r.Instructor.GetRole(),
			},
		})
	}

	res, err := u.scheduleClient.CreateSchedule(ctx, &scheduleSvcV1.CreateScheduleRequest{
		StartDate:    req.StartDate,
		QuantityWeek: req.QuantityWeek,
		Rooms:        rooms,
		Councils:     councils,
		StudentDefs:  studentDefs,
	})
	if err != nil {
		log.Println("create error", err)
		return nil, err
	}

	var thesisRes []*pb.Thesis
	for _, t := range res.Thesis {
		var councilsScheduleResponse []*pb.Council
		for _, c := range t.Council {
			councilsScheduleResponse = append(councilsScheduleResponse, &pb.Council{
				Id:       c.Id,
				Class:    c.Class,
				Major:    c.Major,
				Phone:    c.Phone,
				PhotoSrc: c.PhotoSrc,
				Role:     c.Role,
				Name:     c.Name,
				Email:    c.Email,
			})
		}

		var timeSlotsResponse []*pb.TimeSlots
		for _, t := range t.Schedule.TimeSlots {
			timeSlotsResponse = append(timeSlotsResponse, &pb.TimeSlots{
				Student: &pb.StudentDefScheduleResponse{
					Id: t.Student.Id,
					Infor: &pb.UserScheduleResponse{
						Id:       t.Student.Infor.GetId(),
						Class:    &t.Student.Infor.Class,
						Major:    t.Student.Infor.Major,
						Phone:    t.Student.Infor.Phone,
						PhotoSrc: t.Student.Infor.GetPhotoSrc(),
						Name:     t.Student.Infor.GetName(),
						Email:    t.Student.Infor.GetEmail(),
						Role:     t.Student.Infor.GetRole(),
					},
					Instructor: &pb.UserScheduleResponse{
						Id:       t.Student.Instructor.GetId(),
						Class:    &t.Student.Instructor.Class,
						Major:    t.Student.Instructor.Major,
						Phone:    t.Student.Instructor.Phone,
						PhotoSrc: t.Student.Instructor.GetPhotoSrc(),
						Name:     t.Student.Instructor.GetName(),
						Email:    t.Student.Instructor.GetEmail(),
						Role:     t.Student.Instructor.GetRole(),
					},
				},
				TimeSlot: &pb.TimeSlot{
					Date:  t.TimeSlot.Date,
					Shift: t.TimeSlot.Shift,
					Id:    t.TimeSlot.Id,
					Time:  t.TimeSlot.Time,
				},
			})
		}

		thesisRes = append(thesisRes, &pb.Thesis{
			Schedule: &pb.Schedule{
				TimeSlots: timeSlotsResponse,
				Room: &pb.RoomSchedule{
					Id:          t.Schedule.Room.Id,
					Name:        t.Schedule.Room.Name,
					School:      t.Schedule.Room.School,
					Type:        t.Schedule.Room.Type,
					Description: t.Schedule.Room.Description,
				},
			},
			Council: councilsScheduleResponse,
			Id:      t.Id,
		})
	}

	return &pb.CreateScheduleResponse{
		Thesis: thesisRes,
	}, nil

}
