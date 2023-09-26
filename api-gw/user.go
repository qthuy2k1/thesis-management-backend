package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	waitingListSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type userServiceGW struct {
	pb.UnimplementedUserServiceServer
	userClient        userSvcV1.UserServiceClient
	classroomClient   classroomSvcV1.ClassroomServiceClient
	waitingListClient waitingListSvcV1.WaitingListServiceClient
}

func NewUsersService(userClient userSvcV1.UserServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, waitingListClient waitingListSvcV1.WaitingListServiceClient) *userServiceGW {
	return &userServiceGW{
		userClient:        userClient,
		classroomClient:   classroomClient,
		waitingListClient: waitingListClient,
		// jwtManager:      jwtManager,
		// accessibleRoles: accessibleRoles,
	}
}

func (u *userServiceGW) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.GetUser().GetClassroomID() != 0 {
		classroomID := req.GetUser().GetClassroomID()

		exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: int64(classroomID)})
		if err != nil {
			return nil, err
		}

		if !exists.GetExists() {
			return &pb.CreateUserResponse{
				Response: &pb.CommonUserResponse{
					StatusCode: 404,
					Message:    "Classroom does not exist",
				},
			}, nil
		}
	}

	major := req.GetUser().GetMajor()
	phone := req.GetUser().GetPhone()
	classroomID := req.GetUser().GetClassroomID()

	res, err := u.userClient.CreateUser(ctx, &userSvcV1.CreateUserRequest{
		User: &userSvcV1.UserInput{
			Id:          req.GetUser().GetId(),
			Class:       req.GetUser().GetClass(),
			Major:       &major,
			Phone:       &phone,
			PhotoSrc:    req.GetUser().GetPhotoSrc(),
			Role:        req.GetUser().GetRole(),
			Name:        req.GetUser().GetName(),
			Email:       req.GetUser().GetEmail(),
			ClassroomID: &classroomID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *userServiceGW) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	major := res.GetUser().GetMajor()
	phone := res.GetUser().GetPhone()
	classroomID := res.GetUser().GetClassroomID()

	return &pb.GetUserResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		User: &pb.UserResponse{
			Id:          res.GetUser().Id,
			Class:       res.GetUser().GetClass(),
			Major:       &major,
			Phone:       &phone,
			PhotoSrc:    res.GetUser().GetPhotoSrc(),
			Role:        res.GetUser().GetRole(),
			Name:        res.GetUser().GetName(),
			Email:       res.GetUser().GetEmail(),
			ClassroomID: &classroomID,
		},
	}, nil
}

func (u *userServiceGW) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if req.GetUser().GetClassroomID() != 0 {
		classroomID := req.GetUser().GetClassroomID()

		exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: int64(classroomID)})
		if err != nil {
			return nil, err
		}

		if !exists.GetExists() {
			return &pb.UpdateUserResponse{
				Response: &pb.CommonUserResponse{
					StatusCode: 404,
					Message:    "Classroom does not exist",
				},
			}, nil
		}
	}

	major := req.GetUser().GetMajor()
	phone := req.GetUser().GetPhone()
	classroomID := req.GetUser().GetClassroomID()

	res, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: req.GetId(),
		User: &userSvcV1.UserInput{
			Class:       req.GetUser().GetClass(),
			Major:       &major,
			Phone:       &phone,
			PhotoSrc:    req.GetUser().GetPhotoSrc(),
			Role:        req.GetUser().GetRole(),
			Name:        req.GetUser().GetName(),
			Email:       req.GetUser().GetEmail(),
			ClassroomID: &classroomID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *userServiceGW) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res, err := u.userClient.DeleteUser(ctx, &userSvcV1.DeleteUserRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *userServiceGW) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.GetUsers(ctx, &userSvcV1.GetUsersRequest{})
	if err != nil {
		return nil, err
	}

	var users []*pb.UserResponse
	for _, u := range res.GetUsers() {
		users = append(users, &pb.UserResponse{
			Id:          u.Id,
			Class:       u.Class,
			Major:       u.Major,
			Phone:       u.Phone,
			PhotoSrc:    u.PhotoSrc,
			Role:        u.Role,
			Name:        u.Name,
			Email:       u.Email,
			ClassroomID: &u.ClassroomID,
		})
	}

	return &pb.GetUsersResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Users:      users,
	}, nil
}

func (u *userServiceGW) GetAllUsersOfClassroom(ctx context.Context, req *pb.GetAllUsersOfClassroomRequest) (*pb.GetAllUsersOfClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var classroomID int64

	if req.GetClassroomID() > 0 {
		classroomID = req.GetClassroomID()
	}

	res, err := u.userClient.GetAllUsersOfClassroom(ctx, &userSvcV1.GetAllUsersOfClassroomRequest{
		ClassroomID: classroomID,
	})
	if err != nil {
		return nil, err
	}

	var users []*pb.UserResponse
	for _, u := range res.GetUsers() {
		users = append(users, &pb.UserResponse{
			Id:          u.Id,
			Class:       u.Class,
			Major:       u.Major,
			Phone:       u.Phone,
			PhotoSrc:    u.PhotoSrc,
			Role:        u.Role,
			Name:        u.Name,
			Email:       u.Email,
			ClassroomID: &u.ClassroomID,
		})
	}

	return &pb.GetAllUsersOfClassroomResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Users:      users,
	}, nil
}

func (u *userServiceGW) ApproveUserJoinClassroom(ctx context.Context, req *pb.ApproveUserJoinClassroomRequest) (*pb.ApproveUserJoinClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	clrRes, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{
		ClassroomID: req.GetClassroomID(),
	})
	if err != nil {
		return nil, err
	}

	if !clrRes.GetExists() {
		return &pb.ApproveUserJoinClassroomResponse{
			Response: &pb.CommonUserResponse{
				StatusCode: 404,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	if userRes.Response.StatusCode == 404 {
		return &pb.ApproveUserJoinClassroomResponse{
			Response: &pb.CommonUserResponse{
				StatusCode: 404,
				Message:    "User does not exist",
			},
		}, nil
	}

	classroomID := req.GetClassroomID()
	res, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: req.GetUserID(),
		User: &userSvcV1.UserInput{
			Class:       userRes.User.Class,
			Major:       userRes.User.Major,
			Phone:       userRes.User.Phone,
			PhotoSrc:    userRes.User.PhotoSrc,
			Role:        userRes.User.Role,
			Name:        userRes.User.Name,
			Email:       userRes.User.Email,
			ClassroomID: &classroomID,
		},
	})
	if err != nil {
		return nil, err
	}

	if res.Response.StatusCode != 200 {
		return &pb.ApproveUserJoinClassroomResponse{
			Response: &pb.CommonUserResponse{
				StatusCode: res.GetResponse().GetStatusCode(),
				Message:    res.GetResponse().GetMessage(),
			},
		}, nil
	}

	return &pb.ApproveUserJoinClassroomResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: 200,
			Message:    "OK",
		},
	}, nil
}

// func (u *userServiceGW) Authorize(ctx context.Context, method string) error {
// 	accessibleRoles, ok := u.accessibleRoles[method]
// 	if !ok {
// 		// everyone can access
// 		return nil
// 	}

// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
// 	}

// 	values := md["authorization"]
// 	if len(values) == 0 {
// 		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
// 	}

// 	accessToken := values[0]
// 	claims, err := u.jwtManager.Verify(accessToken)
// 	if err != nil {
// 		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
// 	}

// 	for _, role := range accessibleRoles {
// 		if role == claims.Role {
// 			return nil
// 		}
// 	}

// 	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
// }

func (u *userServiceGW) CheckStatusUserJoinClassroom(ctx context.Context, req *pb.CheckStatusUserJoinClassroomRequest) (*pb.CheckStatusUserJoinClassroomResponse, error) {
	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	classroomID := userRes.GetUser().GetClassroomID()
	if err != nil {
		return nil, err
	}

	if classroomID == 0 {
		wtlRes, err := u.waitingListClient.CheckUserInWaitingListOfClassroom(ctx, &waitingListSvcV1.CheckUserInWaitingListClassroomRequest{
			UserID: req.GetId(),
		})
		if err != nil {
			return nil, err
		}

		if wtlRes.IsIn {
			clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
				Id: wtlRes.GetClassroomID(),
			})
			if err != nil {
				return nil, err
			}

			return &pb.CheckStatusUserJoinClassroomResponse{
				Status: "WAITING",
				Classroom: &pb.ClassroomUserResponse{
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
		} else {
			return &pb.CheckStatusUserJoinClassroomResponse{
				Status:    "NOT REGISTERED",
				Classroom: &pb.ClassroomUserResponse{},
			}, nil
		}
	}

	// classroomID != 0
	clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
		Id: int64(classroomID),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CheckStatusUserJoinClassroomResponse{
		Status: "ADDED",
		Classroom: &pb.ClassroomUserResponse{
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
