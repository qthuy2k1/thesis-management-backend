package service

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	waitingListSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-waiting-list-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type memberServiceGW struct {
	pb.UnimplementedMemberServiceServer
	userClient        userSvcV1.UserServiceClient
	classroomClient   classroomSvcV1.ClassroomServiceClient
	waitingListClient waitingListSvcV1.WaitingListServiceClient
}

func NewMembersService(userClient userSvcV1.UserServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, waitingListClient waitingListSvcV1.WaitingListServiceClient) *memberServiceGW {
	return &memberServiceGW{
		userClient:        userClient,
		classroomClient:   classroomClient,
		waitingListClient: waitingListClient,
	}
}

func (u *memberServiceGW) CreateMember(ctx context.Context, req *pb.CreateMemberRequest) (*pb.CreateMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.GetMember().GetClassroomID() != 0 {
		classroomID := req.GetMember().GetClassroomID()

		exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: int64(classroomID)})
		if err != nil {
			return nil, err
		}

		if !exists.GetExists() {
			return &pb.CreateMemberResponse{
				Response: &pb.CommonMemberResponse{
					StatusCode: 404,
					Message:    "Classroom does not exist",
				},
			}, nil
		}
	}

	if _, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.Member.MemberID,
	}); err != nil {
		return nil, err
	}

	res, err := u.userClient.CreateMember(ctx, &userSvcV1.CreateMemberRequest{
		Member: &userSvcV1.MemberInput{
			ClassroomID: req.GetMember().ClassroomID,
			MemberID:    req.GetMember().MemberID,
			Status:      req.GetMember().Status,
			IsDefense:   req.GetMember().RegisterDefense,
		},
	})
	if err != nil {
		return nil, err
	}

	wlt, err := u.waitingListClient.GetWaitingListByUser(ctx, &waitingListSvcV1.GetWaitingListByUserRequest{
		UserID: req.Member.GetMemberID(),
	})
	if err != nil {
		return nil, err
	}

	if _, err := u.waitingListClient.DeleteWaitingList(ctx, &waitingListSvcV1.DeleteWaitingListRequest{
		Id: wlt.WaitingList.Id,
	}); err != nil {
		return nil, err
	}

	return &pb.CreateMemberResponse{
		Response: &pb.CommonMemberResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *memberServiceGW) GetMember(ctx context.Context, req *pb.GetMemberRequest) (*pb.GetMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.GetMember(ctx, &userSvcV1.GetMemberRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
		Id: res.GetMember().ClassroomID,
	})
	if err != nil {
		return nil, err
	}

	if clrRes.GetResponse().StatusCode != 200 {
		return &pb.GetMemberResponse{
			Response: &pb.CommonMemberResponse{
				StatusCode: clrRes.GetResponse().StatusCode,
				Message:    clrRes.GetResponse().Message,
			},
		}, nil
	}

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: res.Member.MemberID})
	if err != nil {
		return nil, err
	}

	if userRes.GetResponse().StatusCode != 200 {
		return &pb.GetMemberResponse{
			Response: &pb.CommonMemberResponse{
				StatusCode: userRes.GetResponse().StatusCode,
				Message:    userRes.GetResponse().Message,
			},
		}, nil
	}

	lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: clrRes.Classroom.LecturerID})
	if err != nil {
		return nil, err
	}

	if userRes.GetResponse().StatusCode != 200 {
		return &pb.GetMemberResponse{
			Response: &pb.CommonMemberResponse{
				StatusCode: userRes.GetResponse().StatusCode,
				Message:    userRes.GetResponse().Message,
			},
		}, nil
	}

	topicTags := ""
	if clrRes.GetClassroom().TopicTags != nil {
		topicTags = clrRes.GetClassroom().GetTopicTags()
	}

	return &pb.GetMemberResponse{
		Response: &pb.CommonMemberResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Member: &pb.MemberResponse{
			Id: res.GetMember().Id,
			Classroom: &pb.ClassroomMemberResponse{
				Id:          clrRes.GetClassroom().GetId(),
				Title:       clrRes.GetClassroom().GetTitle(),
				Description: clrRes.GetClassroom().GetDescription(),
				Status:      clrRes.GetClassroom().GetStatus(),
				Lecturer: &pb.UserMemberResponse{
					Id:       lecturerRes.GetUser().GetId(),
					Class:    lecturerRes.GetUser().Class,
					Major:    lecturerRes.GetUser().Major,
					Phone:    lecturerRes.GetUser().Phone,
					PhotoSrc: lecturerRes.GetUser().GetPhotoSrc(),
					Name:     lecturerRes.GetUser().GetName(),
					Email:    lecturerRes.GetUser().GetEmail(),
					Role:     lecturerRes.GetUser().GetRole(),
				},
				ClassCourse:     clrRes.GetClassroom().GetClassCourse(),
				TopicTags:       &topicTags,
				QuantityStudent: clrRes.GetClassroom().GetQuantityStudent(),
				CreatedAt:       clrRes.GetClassroom().GetCreatedAt(),
				UpdatedAt:       clrRes.GetClassroom().GetUpdatedAt(),
			},
			Member: &pb.UserMemberResponse{
				Id:       userRes.GetUser().GetId(),
				Class:    userRes.GetUser().Class,
				Major:    userRes.GetUser().Major,
				Phone:    userRes.GetUser().Phone,
				PhotoSrc: userRes.GetUser().GetPhotoSrc(),
				Name:     userRes.GetUser().GetName(),
				Email:    userRes.GetUser().GetEmail(),
				Role:     userRes.GetUser().GetRole(),
			},
			Status:          res.GetMember().Status,
			RegisterDefense: res.GetMember().IsDefense,
			CreatedAt:       res.GetMember().CreatedAt,
		},
	}, nil
}

func (u *memberServiceGW) UpdateMember(ctx context.Context, req *pb.UpdateMemberRequest) (*pb.UpdateMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: int64(req.GetMember().ClassroomID)})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.UpdateMemberResponse{
			Response: &pb.CommonMemberResponse{
				StatusCode: 404,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	res, err := u.userClient.UpdateMember(ctx, &userSvcV1.UpdateMemberRequest{
		Id: req.GetId(),
		Member: &userSvcV1.MemberInput{
			ClassroomID: req.GetMember().ClassroomID,
			MemberID:    req.GetMember().MemberID,
			Status:      req.GetMember().Status,
			IsDefense:   req.GetMember().RegisterDefense,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateMemberResponse{
		Response: &pb.CommonMemberResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *memberServiceGW) DeleteMember(ctx context.Context, req *pb.DeleteMemberRequest) (*pb.DeleteMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.DeleteMember(ctx, &userSvcV1.DeleteMemberRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteMemberResponse{
		Response: &pb.CommonMemberResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *memberServiceGW) GetMembers(ctx context.Context, req *pb.GetMembersRequest) (*pb.GetMembersResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.GetMembers(ctx, &userSvcV1.GetMembersRequest{})
	if err != nil {
		return nil, err
	}

	var members []*pb.MemberResponse
	for _, m := range res.GetMembers() {
		clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
			Id: m.ClassroomID,
		})
		if err != nil {
			return nil, err
		}

		if clrRes.GetResponse().StatusCode != 200 {
			return &pb.GetMembersResponse{
				Response: &pb.CommonMemberResponse{
					StatusCode: clrRes.GetResponse().StatusCode,
					Message:    clrRes.GetResponse().Message,
				},
			}, nil
		}

		userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: m.MemberID})
		if err != nil {
			return nil, err
		}

		lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: clrRes.Classroom.LecturerID})
		if err != nil {
			return nil, err
		}

		topicTags := ""
		if clrRes.GetClassroom().TopicTags != nil {
			topicTags = clrRes.GetClassroom().GetTopicTags()
		}

		members = append(members, &pb.MemberResponse{
			Id: m.Id,
			Classroom: &pb.ClassroomMemberResponse{
				Id:          clrRes.GetClassroom().GetId(),
				Title:       clrRes.GetClassroom().GetTitle(),
				Description: clrRes.GetClassroom().GetDescription(),
				Status:      clrRes.GetClassroom().GetStatus(),
				Lecturer: &pb.UserMemberResponse{
					Id:       lecturerRes.GetUser().GetId(),
					Class:    lecturerRes.GetUser().Class,
					Major:    lecturerRes.GetUser().Major,
					Phone:    lecturerRes.GetUser().Phone,
					PhotoSrc: lecturerRes.GetUser().GetPhotoSrc(),
					Name:     lecturerRes.GetUser().GetName(),
					Email:    lecturerRes.GetUser().GetEmail(),
					Role:     lecturerRes.GetUser().GetRole(),
				},
				ClassCourse:     clrRes.GetClassroom().GetClassCourse(),
				TopicTags:       &topicTags,
				QuantityStudent: clrRes.GetClassroom().GetQuantityStudent(),
				CreatedAt:       clrRes.GetClassroom().GetCreatedAt(),
				UpdatedAt:       clrRes.GetClassroom().GetUpdatedAt(),
			},
			Member: &pb.UserMemberResponse{
				Id:       userRes.GetUser().GetId(),
				Class:    userRes.GetUser().Class,
				Major:    userRes.GetUser().Major,
				Phone:    userRes.GetUser().Phone,
				PhotoSrc: userRes.GetUser().GetPhotoSrc(),
				Name:     userRes.GetUser().GetName(),
				Email:    userRes.GetUser().GetEmail(),
				Role:     userRes.GetUser().GetRole(),
			},
			Status:          m.Status,
			RegisterDefense: m.IsDefense,
			CreatedAt:       m.CreatedAt,
		})
	}

	return &pb.GetMembersResponse{
		TotalCount: res.GetTotalCount(),
		Members:    members,
	}, nil
}

func (u *memberServiceGW) GetAllMembersOfClassroom(ctx context.Context, req *pb.GetAllMembersOfClassroomRequest) (*pb.GetAllMembersOfClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
		Id: req.GetClassroomID(),
	})
	if err != nil {
		return nil, err
	}

	if clrRes.GetResponse().StatusCode != 200 {
		return &pb.GetAllMembersOfClassroomResponse{
			Response: &pb.CommonMemberResponse{
				StatusCode: clrRes.GetResponse().StatusCode,
				Message:    "classroom does not exist",
			},
		}, nil
	}

	res, err := u.userClient.GetAllMembersOfClassroom(ctx, &userSvcV1.GetAllMembersOfClassroomRequest{
		ClassroomID: req.ClassroomID,
	})
	if err != nil {
		return nil, err
	}

	var members []*pb.MemberResponse
	for _, m := range res.GetMembers() {
		clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
			Id: m.ClassroomID,
		})
		if err != nil {
			return nil, err
		}

		if clrRes.GetResponse().StatusCode != 200 {
			return &pb.GetAllMembersOfClassroomResponse{
				Response: &pb.CommonMemberResponse{
					StatusCode: clrRes.GetResponse().StatusCode,
					Message:    clrRes.GetResponse().Message,
				},
			}, nil
		}

		userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: m.MemberID,
		})
		if err != nil {
			return nil, err
		}

		if userRes.GetResponse().StatusCode != 200 {
			return &pb.GetAllMembersOfClassroomResponse{
				Response: &pb.CommonMemberResponse{
					StatusCode: userRes.GetResponse().StatusCode,
					Message:    userRes.GetResponse().Message,
				},
			}, nil
		}

		lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: clrRes.GetClassroom().LecturerID,
		})
		if err != nil {
			return nil, err
		}

		if lecturerRes.GetResponse().StatusCode != 200 {
			return &pb.GetAllMembersOfClassroomResponse{
				Response: &pb.CommonMemberResponse{
					StatusCode: lecturerRes.GetResponse().StatusCode,
					Message:    lecturerRes.GetResponse().Message,
				},
			}, nil
		}

		topicTags := ""
		if clrRes.GetClassroom().TopicTags != nil {
			topicTags = clrRes.GetClassroom().GetTopicTags()
		}

		members = append(members, &pb.MemberResponse{
			Id: m.Id,
			Classroom: &pb.ClassroomMemberResponse{
				Id:          clrRes.GetClassroom().GetId(),
				Title:       clrRes.GetClassroom().GetTitle(),
				Description: clrRes.GetClassroom().GetDescription(),
				Status:      clrRes.GetClassroom().GetStatus(),
				Lecturer: &pb.UserMemberResponse{
					Id:       lecturerRes.GetUser().GetId(),
					Class:    lecturerRes.GetUser().Class,
					Major:    lecturerRes.GetUser().Major,
					Phone:    lecturerRes.GetUser().Phone,
					PhotoSrc: lecturerRes.GetUser().GetPhotoSrc(),
					Name:     lecturerRes.GetUser().GetName(),
					Email:    lecturerRes.GetUser().GetEmail(),
					Role:     lecturerRes.GetUser().GetRole(),
				},
				ClassCourse:     clrRes.GetClassroom().GetClassCourse(),
				TopicTags:       &topicTags,
				QuantityStudent: clrRes.GetClassroom().GetQuantityStudent(),
				CreatedAt:       clrRes.GetClassroom().GetCreatedAt(),
				UpdatedAt:       clrRes.GetClassroom().GetUpdatedAt(),
			},
			Member: &pb.UserMemberResponse{
				Id:       userRes.GetUser().GetId(),
				Class:    userRes.GetUser().Class,
				Major:    userRes.GetUser().Major,
				Phone:    userRes.GetUser().Phone,
				PhotoSrc: userRes.GetUser().GetPhotoSrc(),
				Name:     userRes.GetUser().GetName(),
				Email:    userRes.GetUser().GetEmail(),
				Role:     userRes.GetUser().GetRole(),
			},
			Status:          m.Status,
			RegisterDefense: m.IsDefense,
			CreatedAt:       m.CreatedAt,
		})
	}

	return &pb.GetAllMembersOfClassroomResponse{
		TotalCount: res.GetTotalCount(),
		Members:    members,
	}, nil
}

func (u *memberServiceGW) GetUserMember(ctx context.Context, req *pb.GetUserMemberRequest) (*pb.GetUserMemberResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	memberRes, err := u.userClient.GetUserMember(ctx, &userSvcV1.GetUserMemberRequest{
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	classroomRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
		Id: memberRes.GetMember().ClassroomID,
	})
	if err != nil {
		return nil, err
	}

	lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: classroomRes.Classroom.LecturerID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetUserMemberResponse{
		Member: &pb.MemberResponse{
			Id: memberRes.Member.Id,
			Classroom: &pb.ClassroomMemberResponse{
				Id:     classroomRes.Classroom.Id,
				Status: classroomRes.Classroom.Status,
				Lecturer: &pb.UserMemberResponse{
					Id:       lecturerRes.User.Id,
					Class:    lecturerRes.User.Class,
					Major:    lecturerRes.User.Major,
					Phone:    lecturerRes.User.Phone,
					PhotoSrc: lecturerRes.User.PhotoSrc,
					Role:     lecturerRes.User.Role,
					Name:     lecturerRes.User.Name,
					Email:    lecturerRes.User.Email,
				},
				ClassCourse:     classroomRes.Classroom.ClassCourse,
				TopicTags:       classroomRes.Classroom.TopicTags,
				QuantityStudent: classroomRes.Classroom.QuantityStudent,
				CreatedAt:       classroomRes.Classroom.CreatedAt,
				UpdatedAt:       classroomRes.Classroom.UpdatedAt,
			},
			Member: &pb.UserMemberResponse{
				Id:       userRes.User.Id,
				Class:    userRes.User.Class,
				Major:    userRes.User.Major,
				Phone:    userRes.User.Phone,
				PhotoSrc: userRes.User.PhotoSrc,
				Role:     userRes.User.Role,
				Name:     userRes.User.Name,
				Email:    userRes.User.Email,
			},
			Status:          memberRes.Member.Status,
			RegisterDefense: memberRes.Member.IsDefense,
			CreatedAt:       memberRes.Member.CreatedAt,
		},
	}, nil
}
