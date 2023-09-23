package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	waitingListSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type waitingListServiceGW struct {
	pb.UnimplementedWaitingListServiceServer
	waitingListClient waitingListSvcV1.WaitingListServiceClient
	classroomClient   classroomSvcV1.ClassroomServiceClient
	userClient        userSvcV1.UserServiceClient
}

func NewWaitingListsService(waitingListClient waitingListSvcV1.WaitingListServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, userClient userSvcV1.UserServiceClient) *waitingListServiceGW {
	return &waitingListServiceGW{
		waitingListClient: waitingListClient,
		classroomClient:   classroomClient,
		userClient:        userClient,
	}
}

func (u *waitingListServiceGW) CreateWaitingList(ctx context.Context, req *pb.CreateWaitingListRequest) (*pb.CreateWaitingListResponse, error) {
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

	userExists, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.GetWaitingList().GetUserID()})
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
		UserID: req.GetWaitingList().GetUserID(),
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
			UserID:      req.GetWaitingList().GetUserID(),
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

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: res.WaitingList.UserID,
	})
	if err != nil {
		return nil, err
	}

	lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: clrRes.Classroom.LecturerId,
	})

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
					Id:          lecturerRes.User.Id,
					Class:       lecturerRes.User.Class,
					Major:       lecturerRes.User.Major,
					Phone:       lecturerRes.User.Phone,
					PhotoSrc:    lecturerRes.User.PhotoSrc,
					Role:        lecturerRes.User.Role,
					Name:        lecturerRes.User.Name,
					Email:       lecturerRes.User.Email,
					ClassroomID: &lecturerRes.User.ClassroomID,
				},
				CodeClassroom: clrRes.Classroom.CodeClassroom,
				TopicTags:     clrRes.Classroom.TopicTags,
				Quantity:      clrRes.Classroom.Quantity,
				CreatedAt:     clrRes.Classroom.CreatedAt,
				UpdatedAt:     clrRes.Classroom.UpdatedAt,
			},
			User: &pb.UserWaitingListResponse{
				Id:          userRes.User.Id,
				Class:       userRes.User.Class,
				Major:       userRes.User.Major,
				Phone:       userRes.User.Phone,
				PhotoSrc:    userRes.User.PhotoSrc,
				Role:        userRes.User.Role,
				Name:        userRes.User.Name,
				Email:       userRes.User.Email,
				ClassroomID: &userRes.User.ClassroomID,
			},
			CreatedAt: res.GetWaitingList().GetCreatedAt(),
		},
	}, nil
}

func (u *waitingListServiceGW) UpdateWaitingList(ctx context.Context, req *pb.UpdateWaitingListRequest) (*pb.UpdateWaitingListResponse, error) {
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

	userExists, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.GetWaitingList().GetUserID()})
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
		UserID: req.GetWaitingList().GetUserID(),
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
			UserID:      req.GetWaitingList().GetUserID(),
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

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: wtl.WaitingList.UserID,
	})
	if err != nil {
		return nil, err
	}

	classroomID := "0"
	userUpdateRes, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: userRes.User.Id,
		User: &userSvcV1.UserInput{
			Class:       userRes.User.Class,
			Major:       userRes.User.Major,
			Phone:       userRes.User.Phone,
			PhotoSrc:    userRes.User.PhotoSrc,
			Role:        userRes.User.Role,
			Email:       userRes.User.Email,
			Name:        userRes.User.Name,
			ClassroomID: &classroomID,
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

		userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: p.UserID,
		})
		if err != nil {
			return nil, err
		}

		lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: clrRes.Classroom.LecturerId,
		})

		waitingLists = append(waitingLists, &pb.WaitingListResponse{
			Id: p.Id,
			Classroom: &pb.ClassroomWTLResponse{
				Id:          clrRes.Classroom.Id,
				Title:       clrRes.Classroom.Title,
				Description: clrRes.Classroom.Description,
				Status:      clrRes.Classroom.Status,
				Lecturer: &pb.LecturerWaitingListResponse{
					Id:          lecturerRes.User.Id,
					Class:       lecturerRes.User.Class,
					Major:       lecturerRes.User.Major,
					Phone:       lecturerRes.User.Phone,
					PhotoSrc:    lecturerRes.User.PhotoSrc,
					Role:        lecturerRes.User.Role,
					Name:        lecturerRes.User.Name,
					Email:       lecturerRes.User.Email,
					ClassroomID: &lecturerRes.User.ClassroomID,
				},
				CodeClassroom: clrRes.Classroom.CodeClassroom,
				TopicTags:     clrRes.Classroom.TopicTags,
				Quantity:      clrRes.Classroom.Quantity,
				CreatedAt:     clrRes.Classroom.CreatedAt,
				UpdatedAt:     clrRes.Classroom.UpdatedAt,
			},
			User: &pb.UserWaitingListResponse{
				Id:          userRes.User.Id,
				Class:       userRes.User.Class,
				Major:       userRes.User.Major,
				Phone:       userRes.User.Phone,
				PhotoSrc:    userRes.User.PhotoSrc,
				Role:        userRes.User.Role,
				Name:        userRes.User.Name,
				Email:       userRes.User.Email,
				ClassroomID: &userRes.User.ClassroomID,
			},
			CreatedAt: p.GetCreatedAt(),
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

		lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: clrRes.Classroom.LecturerId,
		})

		return &pb.CheckUserInWaitingListClassroomResponse{
			Status: "WAITING",
			Classroom: &pb.ClassroomWTLResponse{
				Id:          clrRes.GetClassroom().GetId(),
				Title:       clrRes.GetClassroom().GetTitle(),
				Description: clrRes.GetClassroom().GetDescription(),
				Status:      clrRes.GetClassroom().GetStatus(),
				Lecturer: &pb.LecturerWaitingListResponse{
					Id:          lecturerRes.User.Id,
					Class:       lecturerRes.User.Class,
					Major:       lecturerRes.User.Major,
					Phone:       lecturerRes.User.Phone,
					PhotoSrc:    lecturerRes.User.PhotoSrc,
					Role:        lecturerRes.User.Role,
					Name:        lecturerRes.User.Name,
					Email:       lecturerRes.User.Email,
					ClassroomID: &lecturerRes.User.ClassroomID,
				},
				CodeClassroom: clrRes.GetClassroom().GetCodeClassroom(),
				TopicTags:     clrRes.GetClassroom().GetTopicTags(),
				Quantity:      clrRes.GetClassroom().GetQuantity(),
				CreatedAt:     clrRes.GetClassroom().GetCreatedAt(),
				UpdatedAt:     clrRes.GetClassroom().GetUpdatedAt(),
			},
		}, nil
	}

	return &pb.CheckUserInWaitingListClassroomResponse{
		Status:    "NOT REGISTERED",
		Classroom: &pb.ClassroomWTLResponse{},
	}, nil
}
