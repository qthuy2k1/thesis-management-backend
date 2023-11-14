package main

import (
	"context"
	"log"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	waitingListSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"

	"golang.org/x/crypto/bcrypt"
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
	}
}

func (u *userServiceGW) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	major := req.GetUser().GetMajor()
	phone := req.GetUser().GetPhone()

	res, err := u.userClient.CreateUser(ctx, &userSvcV1.CreateUserRequest{
		User: &userSvcV1.UserInput{
			Id:       req.GetUser().GetId(),
			Class:    req.GetUser().GetClass(),
			Major:    &major,
			Phone:    &phone,
			PhotoSrc: req.GetUser().GetPhotoSrc(),
			Role:     req.GetUser().GetRole(),
			Name:     req.GetUser().GetName(),
			Email:    req.GetUser().GetEmail(),
		},
	})
	if err != nil {
		return nil, err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.GetUser().Password), 14)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		HashedPassword: string(password),
	}, nil
}

func (u *userServiceGW) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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
			Id:       res.GetUser().Id,
			Class:    res.GetUser().GetClass(),
			Major:    &major,
			Phone:    &phone,
			PhotoSrc: res.GetUser().GetPhotoSrc(),
			Role:     res.GetUser().GetRole(),
			Name:     res.GetUser().GetName(),
			Email:    res.GetUser().GetEmail(),
		},
	}, nil
}

func (u *userServiceGW) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	major := req.GetUser().GetMajor()
	phone := req.GetUser().GetPhone()

	res, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: req.GetId(),
		User: &userSvcV1.UserInput{
			Class:    req.GetUser().GetClass(),
			Major:    &major,
			Phone:    &phone,
			PhotoSrc: req.GetUser().GetPhotoSrc(),
			Role:     req.GetUser().GetRole(),
			Name:     req.GetUser().GetName(),
			Email:    req.GetUser().GetEmail(),
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
	if err := req.Validate(); err != nil {
		return nil, err
	}

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
			Id:       u.Id,
			Class:    u.Class,
			Major:    u.Major,
			Phone:    u.Phone,
			PhotoSrc: u.PhotoSrc,
			Role:     u.Role,
			Name:     u.Name,
			Email:    u.Email,
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

	res, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: req.GetUserID(),
		User: &userSvcV1.UserInput{
			Class:    userRes.User.Class,
			Major:    userRes.User.Major,
			Phone:    userRes.User.Phone,
			PhotoSrc: userRes.User.PhotoSrc,
			Role:     userRes.User.Role,
			Name:     userRes.User.Name,
			Email:    userRes.User.Email,
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

func (u *userServiceGW) CheckStatusUserJoinClassroom(ctx context.Context, req *pb.CheckStatusUserJoinClassroomRequest) (*pb.CheckStatusUserJoinClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	if userRes.GetResponse().StatusCode == 404 {
		return &pb.CheckStatusUserJoinClassroomResponse{
			Response: &pb.CommonUserResponse{
				StatusCode: userRes.GetResponse().StatusCode,
				Message:    userRes.GetResponse().Message,
			},
		}, nil
	}

	memberRes, err := u.userClient.IsUserJoinedClassroom(ctx, &userSvcV1.IsUserJoinedClassroomRequest{
		UserID: req.GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
		Id: memberRes.GetMember().ClassroomID,
	})
	if err != nil {
		return nil, err
	}

	if clrRes.GetResponse().StatusCode == 404 {
		return &pb.CheckStatusUserJoinClassroomResponse{
			Response: &pb.CommonUserResponse{
				StatusCode: clrRes.GetResponse().StatusCode,
				Message:    clrRes.GetResponse().Message,
			},
		}, nil
	}

	return &pb.CheckStatusUserJoinClassroomResponse{
		Member: &pb.MemberUserResponse{
			Id: memberRes.GetMember().Id,
			Classroom: &pb.ClassroomUserResponse{
				Id:              clrRes.GetClassroom().GetId(),
				Title:           clrRes.GetClassroom().GetTitle(),
				Description:     clrRes.GetClassroom().GetDescription(),
				Status:          clrRes.GetClassroom().GetStatus(),
				LecturerID:      clrRes.GetClassroom().GetLecturerID(),
				ClassCourse:     clrRes.GetClassroom().GetClassCourse(),
				TopicTags:       clrRes.GetClassroom().GetTopicTags(),
				QuantityStudent: clrRes.GetClassroom().GetQuantityStudent(),
				CreatedAt:       clrRes.GetClassroom().GetCreatedAt(),
				UpdatedAt:       clrRes.GetClassroom().GetUpdatedAt(),
			},
			Member: &pb.UserResponse{
				Class:    userRes.User.Class,
				Major:    userRes.User.Major,
				Phone:    userRes.User.Phone,
				PhotoSrc: userRes.User.PhotoSrc,
				Role:     userRes.User.Role,
				Name:     userRes.User.Name,
				Email:    userRes.User.Email,
			},
			Status:    memberRes.GetMember().Status,
			IsDefense: memberRes.GetMember().IsDefense,
			CreatedAt: memberRes.GetMember().CreatedAt,
		},
	}, nil

}

func (u *userServiceGW) UpdateBasicUser(ctx context.Context, req *pb.UpdateBasicUserRequest) (*pb.UpdateBasicUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	log.Println(req.User)

	userGetRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	if userGetRes.Response.StatusCode != 200 {
		return &pb.UpdateBasicUserResponse{
			Response: &pb.CommonUserResponse{
				StatusCode: userGetRes.Response.StatusCode,
				Message:    userGetRes.Response.Message,
			},
		}, nil
	}

	major := req.GetUser().GetMajor()
	phone := req.GetUser().GetPhone()

	res, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: req.GetId(),
		User: &userSvcV1.UserInput{
			Class:    req.GetUser().GetClass(),
			Major:    &major,
			Phone:    &phone,
			PhotoSrc: req.GetUser().GetPhotoSrc(),
			Name:     req.GetUser().GetName(),
			Email:    req.GetUser().GetEmail(),
			Role:     userGetRes.GetUser().GetRole(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateBasicUserResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *userServiceGW) GetAllLecturers(ctx context.Context, req *pb.GetAllLecturerRequest) (*pb.GetAllLecturerResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.GetAllLecturers(ctx, &userSvcV1.GetAllLecturersRequest{})
	if err != nil {
		return nil, err
	}

	var lecturers []*pb.UserResponse
	for _, l := range res.GetLecturers() {
		lecturers = append(lecturers, &pb.UserResponse{
			Id:       l.Id,
			Class:    l.Class,
			Major:    l.Major,
			Phone:    l.Phone,
			PhotoSrc: l.PhotoSrc,
			Role:     l.Role,
			Name:     l.Name,
			Email:    l.Email,
		})
	}

	return &pb.GetAllLecturerResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Lecturers:  lecturers,
	}, nil
}

func (u *userServiceGW) UnsubscribeClassroom(ctx context.Context, req *pb.UnsubscribeClassroomRequest) (*pb.UnsubscribeClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.UnsubscribeClassroom(ctx, &userSvcV1.UnsubscribeClassroomRequest{
		MemberID:    req.GetMemberID(),
		ClassroomID: req.GetClassroomID(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UnsubscribeClassroomResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}
