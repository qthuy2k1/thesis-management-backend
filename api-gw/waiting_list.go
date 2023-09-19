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
				StatusCode: 400,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	userExists, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.GetWaitingList().GetUserID()})
	if err != nil {
		return nil, err
	}

	if userExists.GetResponse().GetStatusCode() == 400 {
		return &pb.CreateWaitingListResponse{
			Response: &pb.CommonWaitingListResponse{
				StatusCode: 400,
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

	return &pb.GetWaitingListResponse{
		Response: &pb.CommonWaitingListResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		WaitingList: &pb.WaitingListResponse{
			Id:          res.GetWaitingList().Id,
			ClassroomID: res.GetWaitingList().GetClassroomID(),
			UserID:      res.GetWaitingList().GetUserID(),
			CreatedAt:   res.GetWaitingList().GetCreatedAt(),
		},
	}, nil
}

func (u *waitingListServiceGW) UpdateWaitingList(ctx context.Context, req *pb.UpdateWaitingListRequest) (*pb.UpdateWaitingListResponse, error) {
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
	res, err := u.waitingListClient.DeleteWaitingList(ctx, &waitingListSvcV1.DeleteWaitingListRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
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
		waitingLists = append(waitingLists, &pb.WaitingListResponse{
			Id:          p.Id,
			ClassroomID: p.GetClassroomID(),
			UserID:      p.GetUserID(),
			CreatedAt:   p.GetCreatedAt(),
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

		return &pb.CheckUserInWaitingListClassroomResponse{
			Status: "WAITING",
			Classroom: &pb.ClassroomWTLResponse{
				Id:            clrRes.GetClassroom().GetId(),
				Title:         clrRes.GetClassroom().GetTitle(),
				Description:   clrRes.GetClassroom().GetDescription(),
				Status:        clrRes.GetClassroom().GetStatus(),
				LecturerId:    clrRes.GetClassroom().GetLecturerId(),
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
