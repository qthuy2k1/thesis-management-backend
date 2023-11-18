package main

import (
	"context"
	"errors"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	waitingListSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	redisSvcV1 "github.com/qthuy2k1/thesis-management-backend/redis-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type waitingListServiceGW struct {
	pb.UnimplementedWaitingListServiceServer
	waitingListClient waitingListSvcV1.WaitingListServiceClient
	classroomClient   classroomSvcV1.ClassroomServiceClient
	userClient        userSvcV1.UserServiceClient
	redisClient       redisSvcV1.RedisServiceClient
}

func NewWaitingListsService(waitingListClient waitingListSvcV1.WaitingListServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, userClient userSvcV1.UserServiceClient, redisClient redisSvcV1.RedisServiceClient) *waitingListServiceGW {
	return &waitingListServiceGW{
		waitingListClient: waitingListClient,
		classroomClient:   classroomClient,
		userClient:        userClient,
		redisClient:       redisClient,
	}
}

func (u *waitingListServiceGW) CreateWaitingList(ctx context.Context, req *pb.CreateWaitingListRequest) (*pb.CreateWaitingListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	clrExists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{
		ClassroomID: req.GetWaitingList().GetClassroomID(),
	})
	if err != nil {
		return nil, err
	}

	if !clrExists.GetExists() {
		return &pb.CreateWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: 404,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	userExists, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.GetWaitingList().GetMemberID()})
	if err != nil {
		return nil, err
	}

	if userExists.GetResponse().GetStatusCode() == 404 {
		return &pb.CreateWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: 404,
				Message:    "User does not exist",
			},
		}, nil
	}

	wtlExistRes, err := u.waitingListClient.CheckUserInWaitingListOfClassroom(ctx, &waitingListSvcV1.CheckUserInWaitingListClassroomRequest{
		UserID: req.GetWaitingList().GetMemberID(),
	})
	if err != nil {
		return nil, err
	}

	if wtlExistRes.IsIn {
		return &pb.CreateWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: 400,
				Message:    "User already requested to join a classroom",
			},
		}, nil
	}

	res, err := u.waitingListClient.CreateWaitingList(ctx, &waitingListSvcV1.CreateWaitingListRequest{
		WaitingList: &waitingListSvcV1.WaitingListInput{
			ClassroomID: req.GetWaitingList().GetClassroomID(),
			UserID:      req.GetWaitingList().GetMemberID(),
			IsDefense:   req.GetWaitingList().GetRegisterDefense(),
			Status:      req.GetWaitingList().GetStatus(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateWaitingListResponse{
		Response: &pb.CommonWaitingListResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *waitingListServiceGW) GetWaitingList(ctx context.Context, req *pb.GetWaitingListRequest) (*pb.GetWaitingListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.waitingListClient.GetWaitingList(ctx, &waitingListSvcV1.GetWaitingListRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
		Id: res.WaitingList.ClassroomID,
	})
	if err != nil {
		return nil, err
	}

	redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
		Id: res.WaitingList.UserID,
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
		studentRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: res.WaitingList.UserID})
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

	redis, err = u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
		Id: clrRes.Classroom.LecturerID,
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
		lecturerRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: clrRes.Classroom.LecturerID})
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

	return &pb.GetWaitingListResponse{
		Response: &pb.CommonWaitingListResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		WaitingList: &pb.WaitingListResponse{
			Id: res.GetWaitingList().Id,
			Classroom: &pb.ClassroomWTLResponse{
				Id:          clrRes.Classroom.Id,
				Title:       clrRes.Classroom.Title,
				Description: clrRes.Classroom.Description,
				Status:      clrRes.Classroom.Status,
				Lecturer: &pb.LecturerWaitingListResponse{
					Id:       lecturerRes.User.Id,
					Class:    lecturerRes.User.GetClass(),
					Major:    lecturerRes.User.Major,
					Phone:    lecturerRes.User.Phone,
					PhotoSrc: lecturerRes.User.PhotoSrc,
					Role:     lecturerRes.User.Role,
					Name:     lecturerRes.User.Name,
					Email:    lecturerRes.User.Email,
				},
				ClassCourse:     clrRes.Classroom.ClassCourse,
				TopicTags:       clrRes.Classroom.TopicTags,
				QuantityStudent: clrRes.Classroom.QuantityStudent,
				CreatedAt:       clrRes.Classroom.CreatedAt,
				UpdatedAt:       clrRes.Classroom.UpdatedAt,
			},
			Member: &pb.UserWaitingListResponse{
				Id:       studentRes.User.Id,
				Class:    studentRes.User.Class,
				Major:    studentRes.User.Major,
				Phone:    studentRes.User.Phone,
				PhotoSrc: studentRes.User.PhotoSrc,
				Role:     studentRes.User.Role,
				Name:     studentRes.User.Name,
				Email:    studentRes.User.Email,
			},
			RegisterDefense: res.GetWaitingList().GetIsDefense(),
			Status:          res.GetWaitingList().GetStatus(),
			CreatedAt:       res.GetWaitingList().GetCreatedAt(),
		},
	}, nil
}

func (u *waitingListServiceGW) UpdateWaitingList(ctx context.Context, req *pb.UpdateWaitingListRequest) (*pb.UpdateWaitingListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	clrExists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{
		ClassroomID: req.GetWaitingList().GetClassroomID(),
	})
	if err != nil {
		return nil, err
	}

	if !clrExists.GetExists() {
		return &pb.UpdateWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: 404,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	userExists, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.GetWaitingList().GetMemberID()})
	if err != nil {
		return nil, err
	}

	if userExists.GetResponse().GetStatusCode() == 404 {
		return &pb.UpdateWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: 404,
				Message:    "User does not exist",
			},
		}, nil
	}

	wtlExistRes, err := u.waitingListClient.CheckUserInWaitingListOfClassroom(ctx, &waitingListSvcV1.CheckUserInWaitingListClassroomRequest{
		UserID: req.GetWaitingList().GetMemberID(),
	})
	if err != nil {
		return nil, err
	}

	if wtlExistRes.IsIn {
		return &pb.UpdateWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: 400,
				Message:    "User already requested to join a classroom",
			},
		}, nil
	}
	res, err := u.waitingListClient.UpdateWaitingList(ctx, &waitingListSvcV1.UpdateWaitingListRequest{
		Id: req.GetId(),
		WaitingList: &waitingListSvcV1.WaitingListInput{
			ClassroomID: req.GetWaitingList().GetClassroomID(),
			UserID:      req.GetWaitingList().GetMemberID(),
			IsDefense:   req.GetWaitingList().GetRegisterDefense(),
			Status:      req.GetWaitingList().GetStatus(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateWaitingListResponse{
		Response: &pb.CommonWaitingListResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *waitingListServiceGW) DeleteWaitingList(ctx context.Context, req *pb.DeleteWaitingListRequest) (*pb.DeleteWaitingListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	wtl, err := u.waitingListClient.GetWaitingList(ctx, &waitingListSvcV1.GetWaitingListRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	if wtl.Response.StatusCode == 404 {
		return &pb.DeleteWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: 404,
				Message:    "waiting list is not found",
			},
		}, nil
	}

	res, err := u.waitingListClient.DeleteWaitingList(ctx, &waitingListSvcV1.DeleteWaitingListRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
		Id: wtl.WaitingList.UserID,
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
		studentRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: wtl.WaitingList.UserID})
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

	userUpdateRes, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: studentRes.User.Id,
		User: &userSvcV1.UserInput{
			Class:    studentRes.User.Class,
			Major:    studentRes.User.Major,
			Phone:    studentRes.User.Phone,
			PhotoSrc: studentRes.User.PhotoSrc,
			Role:     studentRes.User.Role,
			Email:    studentRes.User.Email,
			Name:     studentRes.User.Name,
		},
	})
	if err != nil {
		return nil, err
	}

	if userUpdateRes.Response.StatusCode != 200 {
		return &pb.DeleteWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: userUpdateRes.Response.StatusCode,
				Message:    userUpdateRes.Response.Message,
			},
		}, nil
	}

	return &pb.DeleteWaitingListResponse{
		Response: &pb.CommonWaitingListResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *waitingListServiceGW) GetWaitingListsOfClassroom(ctx context.Context, req *pb.GetWaitingListsRequest) (*pb.GetWaitingListsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.waitingListClient.GetWaitingListsOfClassroom(ctx, &waitingListSvcV1.GetWaitingListsRequest{
		ClassroomID: req.GetClassroomID(),
	})
	if err != nil {
		return nil, err
	}

	var waitingLists []*pb.WaitingListResponse
	for _, p := range res.GetWaitingLists() {
		clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
			Id: p.ClassroomID,
		})
		if err != nil {
			return nil, err
		}

		redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
			Id: p.UserID,
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
			studentRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: p.UserID})
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

		redis, err = u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
			Id: clrRes.Classroom.LecturerID,
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
			lecturerRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: clrRes.Classroom.LecturerID})
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

		waitingLists = append(waitingLists, &pb.WaitingListResponse{
			Id: p.Id,
			Classroom: &pb.ClassroomWTLResponse{
				Id:          clrRes.Classroom.Id,
				Title:       clrRes.Classroom.Title,
				Description: clrRes.Classroom.Description,
				Status:      clrRes.Classroom.Status,
				Lecturer: &pb.LecturerWaitingListResponse{
					Id:       lecturerRes.User.Id,
					Class:    lecturerRes.User.GetClass(),
					Major:    lecturerRes.User.Major,
					Phone:    lecturerRes.User.Phone,
					PhotoSrc: lecturerRes.User.PhotoSrc,
					Role:     lecturerRes.User.Role,
					Name:     lecturerRes.User.Name,
					Email:    lecturerRes.User.Email,
				},
				ClassCourse:     clrRes.Classroom.ClassCourse,
				TopicTags:       clrRes.Classroom.TopicTags,
				QuantityStudent: clrRes.Classroom.QuantityStudent,
				CreatedAt:       clrRes.Classroom.CreatedAt,
				UpdatedAt:       clrRes.Classroom.UpdatedAt,
			},
			Member: &pb.UserWaitingListResponse{
				Id:       studentRes.User.Id,
				Class:    studentRes.User.Class,
				Major:    studentRes.User.Major,
				Phone:    studentRes.User.Phone,
				PhotoSrc: studentRes.User.PhotoSrc,
				Role:     studentRes.User.Role,
				Name:     studentRes.User.Name,
				Email:    studentRes.User.Email,
			},
			RegisterDefense: p.IsDefense,
			Status:          p.Status,
			CreatedAt:       p.GetCreatedAt(),
		})
	}

	return &pb.GetWaitingListsResponse{
		Response: &pb.CommonWaitingListResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		WaitingLists: waitingLists,
	}, nil
}

func (u *waitingListServiceGW) CheckUserInWaitingListClassroom(ctx context.Context, req *pb.CheckUserInWaitingListClassroomRequest) (*pb.CheckUserInWaitingListClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.waitingListClient.CheckUserInWaitingListOfClassroom(ctx, &waitingListSvcV1.CheckUserInWaitingListClassroomRequest{
		UserID: req.GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	if res.IsIn {
		clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
			Id: res.GetClassroomID(),
		})
		if err != nil {
			return nil, err
		}

		redis, err := u.redisClient.GetUser(ctx, &redisSvcV1.GetUserRequest{
			Id: clrRes.Classroom.LecturerID,
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
			lecturerRes, err = u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: clrRes.Classroom.LecturerID})
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

		return &pb.CheckUserInWaitingListClassroomResponse{
			Status: "WAITING",
			Classroom: &pb.ClassroomWTLResponse{
				Id:          clrRes.GetClassroom().GetId(),
				Title:       clrRes.GetClassroom().GetTitle(),
				Description: clrRes.GetClassroom().GetDescription(),
				Status:      clrRes.GetClassroom().GetStatus(),
				Lecturer: &pb.LecturerWaitingListResponse{
					Id:       lecturerRes.User.Id,
					Class:    lecturerRes.User.GetClass(),
					Major:    lecturerRes.User.Major,
					Phone:    lecturerRes.User.Phone,
					PhotoSrc: lecturerRes.User.PhotoSrc,
					Role:     lecturerRes.User.Role,
					Name:     lecturerRes.User.Name,
					Email:    lecturerRes.User.Email,
				},
				ClassCourse:     clrRes.GetClassroom().GetClassCourse(),
				TopicTags:       clrRes.GetClassroom().GetTopicTags(),
				QuantityStudent: clrRes.GetClassroom().GetQuantityStudent(),
				CreatedAt:       clrRes.GetClassroom().GetCreatedAt(),
				UpdatedAt:       clrRes.GetClassroom().GetUpdatedAt(),
			},
		}, nil
	}

	return &pb.CheckUserInWaitingListClassroomResponse{
		Status:    "NOT REGISTERED",
		Classroom: &pb.ClassroomWTLResponse{},
	}, nil
}
