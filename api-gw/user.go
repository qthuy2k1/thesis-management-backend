package main

import (
	"context"
	"strconv"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type userServiceGW struct {
	pb.UnimplementedUserServiceServer
	userClient      userSvcV1.UserServiceClient
	classroomClient classroomSvcV1.ClassroomServiceClient
}

func NewUsersService(userClient userSvcV1.UserServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient) *userServiceGW {
	return &userServiceGW{
		userClient:      userClient,
		classroomClient: classroomClient,
	}
}

func (u *userServiceGW) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.GetUser().GetClassroomID() != "" {
		classroomID, err := strconv.Atoi(req.GetUser().GetClassroomID())
		if err != nil {
			return nil, err
		}

		exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: int32(classroomID)})
		if err != nil {
			return nil, err
		}

		if !exists.GetExists() {
			return &pb.CreateUserResponse{
				Response: &pb.CommonUserResponse{
					StatusCode: 400,
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
			ClassroomID: res.GetUser().GetClassroomID(),
		},
	}, nil
}

func (u *userServiceGW) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if req.GetUser().GetClassroomID() != "" {
		classroomID, err := strconv.Atoi(req.GetUser().GetClassroomID())
		if err != nil {
			return nil, err
		}

		exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: int32(classroomID)})
		if err != nil {
			return nil, err
		}

		if !exists.GetExists() {
			return &pb.UpdateUserResponse{
				Response: &pb.CommonUserResponse{
					StatusCode: 400,
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
			ClassroomID: u.ClassroomID,
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

	var classroomID int32

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
			ClassroomID: u.ClassroomID,
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

	classroomID := strconv.Itoa(int(req.GetClassroomID()))
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
