package service

import (
	"context"
	"errors"
	"log"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/crypto/bcrypt"
)

type userServiceGW struct {
	pb.UnimplementedUserServiceServer
	userClient        userSvcV1.UserServiceClient
	classroomClient   classroomSvcV1.ClassroomServiceClient
	waitingListClient classroomSvcV1.WaitingListServiceClient
	topicClient       userSvcV1.TopicServiceClient
	attachmentClient  classroomSvcV1.AttachmentServiceClient
}

func NewUsersService(userClient userSvcV1.UserServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, waitingListClient classroomSvcV1.WaitingListServiceClient, topicClient userSvcV1.TopicServiceClient, attachmentClient classroomSvcV1.AttachmentServiceClient) *userServiceGW {
	return &userServiceGW{
		userClient:        userClient,
		classroomClient:   classroomClient,
		waitingListClient: waitingListClient,
		topicClient:       topicClient,
		attachmentClient:  attachmentClient,
	}
}

func (u *userServiceGW) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	class := req.GetUser().GetClass()
	major := req.GetUser().GetMajor()
	phone := req.GetUser().GetPhone()
	password := req.GetUser().GetPassword()

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	passwordHashedStr := string(passwordHashed)

	res, err := u.userClient.CreateUser(ctx, &userSvcV1.CreateUserRequest{
		User: &userSvcV1.UserInput{
			Id:             req.GetUser().GetId(),
			Class:          &class,
			Major:          &major,
			Phone:          &phone,
			PhotoSrc:       req.GetUser().GetPhotoSrc(),
			Role:           req.GetUser().GetRole(),
			Name:           req.GetUser().GetName(),
			Email:          req.GetUser().GetEmail(),
			HashedPassword: &passwordHashedStr,
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
		HashedPassword: string(passwordHashed),
	}, nil
}

func (u *userServiceGW) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, ok := ctx.Value("user").(*pb.UserResponse)
	if !ok {
		return nil, status.Errorf(codes.Internal, "User not found in context")
	}

	// Check if the user has the required role to update user data.
	if !hasRequiredRole(user.Role, "admin") {
		return nil, status.Errorf(codes.PermissionDenied, "Insufficient privileges to update user")
	}

	res, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	topic, err := u.topicClient.GetTopicFromUser(ctx, &userSvcV1.GetTopicFromUserRequest{
		UserID: res.User.Id,
	})
	if err != nil {
		log.Println("topic err: ", err)
	}

	class := res.GetUser().GetClass()
	major := res.GetUser().GetMajor()
	phone := res.GetUser().GetPhone()

	var topicRes *pb.TopicUserResponse
	if topic != nil {
		if topic.Topic != nil && topic.Response != nil {
			if topic.Response.StatusCode == 200 {
				topicRes = &pb.TopicUserResponse{
					Id:             topic.Topic.GetId(),
					Title:          topic.Topic.GetTitle(),
					TypeTopic:      topic.Topic.GetTypeTopic(),
					MemberQuantity: topic.Topic.GetMemberQuantity(),
					Student: &pb.UserResponse{
						Id:       res.GetUser().Id,
						Class:    &class,
						Major:    &major,
						Phone:    &phone,
						PhotoSrc: res.GetUser().GetPhotoSrc(),
						Role:     res.GetUser().GetRole(),
						Name:     res.GetUser().GetName(),
						Email:    res.GetUser().GetEmail(),
					},
					MemberEmail: topic.Topic.GetMemberEmail(),
					Description: topic.Topic.GetDescription(),
				}
			}
		}
	}

	return &pb.GetUserResponse{
		Response: &pb.CommonUserResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		User: &pb.UserResponse{
			Id:       res.GetUser().Id,
			Class:    &class,
			Major:    &major,
			Phone:    &phone,
			PhotoSrc: res.GetUser().GetPhotoSrc(),
			Role:     res.GetUser().GetRole(),
			Name:     res.GetUser().GetName(),
			Email:    res.GetUser().GetEmail(),
			Topic:    topicRes,
		},
	}, nil
}

func (u *userServiceGW) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	class := req.GetUser().GetClass()
	major := req.GetUser().GetMajor()
	phone := req.GetUser().GetPhone()
	password := req.GetUser().GetPassword()

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	passwordHashedStr := string(passwordHashed)

	res, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: req.GetId(),
		User: &userSvcV1.UserInput{
			Class:          &class,
			Major:          &major,
			Phone:          &phone,
			PhotoSrc:       req.GetUser().GetPhotoSrc(),
			Role:           req.GetUser().GetRole(),
			Name:           req.GetUser().GetName(),
			Email:          req.GetUser().GetEmail(),
			HashedPassword: &passwordHashedStr,
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

	memberGetRes, err := u.userClient.GetUserMember(ctx, &userSvcV1.GetUserMemberRequest{
		UserID: req.GetId(),
	})
	if err != nil {
		log.Println("err get member from delete user", err)
	}

	if memberGetRes.GetMember() != nil {
		if _, err := u.userClient.DeleteMember(ctx, &userSvcV1.DeleteMemberRequest{
			Id: memberGetRes.Member.Id,
		}); err != nil {
			log.Println("err delete member from delete user", err)
			return nil, err
		}
	}

	topicGetRes, err := u.topicClient.GetTopicFromUser(ctx, &userSvcV1.GetTopicFromUserRequest{
		UserID: req.GetId(),
	})
	if err != nil {
		log.Println("err get topic from delete user", err)
	}

	if topicGetRes.GetTopic() != nil {
		if _, err := u.topicClient.DeleteTopic(ctx, &userSvcV1.DeleteTopicRequest{
			Id: topicGetRes.Topic.Id,
		}); err != nil {
			log.Println("err delete topic from delete user", err)
			return nil, err
		}
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
	for _, us := range res.GetUsers() {
		var isNil bool
		topic, err := u.topicClient.GetTopicFromUser(ctx, &userSvcV1.GetTopicFromUserRequest{
			UserID: us.Id,
		})
		if err != nil {
			isNil = true
			log.Println(err)
		}

		if isNil {
			users = append(users, &pb.UserResponse{
				Id:       us.Id,
				Class:    us.Class,
				Major:    us.Major,
				Phone:    us.Phone,
				PhotoSrc: us.PhotoSrc,
				Role:     us.Role,
				Name:     us.Name,
				Email:    us.Email,
				Topic:    nil,
			})
		} else {
			users = append(users, &pb.UserResponse{
				Id:       us.Id,
				Class:    us.Class,
				Major:    us.Major,
				Phone:    us.Phone,
				PhotoSrc: us.PhotoSrc,
				Role:     us.Role,
				Name:     us.Name,
				Email:    us.Email,
				Topic: &pb.TopicUserResponse{
					Id:             topic.Topic.GetId(),
					Title:          topic.Topic.GetTitle(),
					TypeTopic:      topic.Topic.GetTypeTopic(),
					MemberQuantity: topic.Topic.GetMemberQuantity(),
					Student: &pb.UserResponse{
						Id:       us.Id,
						Class:    us.Class,
						Major:    us.Major,
						Phone:    us.Phone,
						PhotoSrc: us.PhotoSrc,
						Role:     us.Role,
						Name:     us.Name,
						Email:    us.Email,
					},
					MemberEmail: topic.Topic.GetMemberEmail(),
					Description: topic.Topic.GetDescription(),
				},
			})
		}
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
		Id: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	res, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: req.GetUserID(),
		User: &userSvcV1.UserInput{
			Class:          userRes.User.Class,
			Major:          userRes.User.Major,
			Phone:          userRes.User.Phone,
			PhotoSrc:       userRes.User.PhotoSrc,
			Role:           userRes.User.Role,
			Name:           userRes.User.Name,
			Email:          userRes.User.Email,
			HashedPassword: userRes.User.HashedPassword,
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
		Id: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	memberRes, err := u.userClient.IsUserJoinedClassroom(ctx, &userSvcV1.IsUserJoinedClassroomRequest{
		UserID: req.GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	if memberRes.Member != nil {
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

		topicTags := ""
		if clrRes.GetClassroom().TopicTags != nil {
			topicTags = clrRes.GetClassroom().GetTopicTags()
		}

		lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: clrRes.GetClassroom().GetLecturerID(),
		})
		if err != nil {
			return nil, err
		}

		return &pb.CheckStatusUserJoinClassroomResponse{
			Member: []*pb.MemberUserResponse{
				{
					Id: memberRes.GetMember().Id,
					Classroom: &pb.ClassroomUserResponse{
						Id:          clrRes.GetClassroom().GetId(),
						Title:       clrRes.GetClassroom().GetTitle(),
						Description: clrRes.GetClassroom().GetDescription(),
						Status:      clrRes.GetClassroom().GetStatus(),
						Lecturer: &pb.UserResponse{
							Class:    lecturerRes.User.Class,
							Major:    lecturerRes.User.Major,
							Phone:    lecturerRes.User.Phone,
							PhotoSrc: lecturerRes.User.PhotoSrc,
							Role:     lecturerRes.User.Role,
							Name:     lecturerRes.User.Name,
							Email:    lecturerRes.User.Email,
							Id:       lecturerRes.User.Id,
						},
						ClassCourse:     clrRes.GetClassroom().GetClassCourse(),
						TopicTags:       &topicTags,
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
						Id:       userRes.User.Id,
					},
					Status:    memberRes.GetMember().Status,
					IsDefense: memberRes.GetMember().IsDefense,
					CreatedAt: memberRes.GetMember().CreatedAt,
				},
			},
			Status: "SUBSCRIBED",
		}, nil
	} else {
		wlt, err := u.waitingListClient.GetWaitingLists(ctx, &classroomSvcV1.GetWaitingListsRequest{})
		if err != nil {
			return nil, err
		}

		var userInWLT []*classroomSvcV1.WaitingListResponse
		for _, w := range wlt.WaitingLists {
			if w.UserID == userRes.User.Id {
				userInWLT = append(userInWLT, w)
			}
		}

		if len(userInWLT) > 0 {
			var members []*pb.MemberUserResponse
			for _, m := range userInWLT {
				clrRes, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{
					Id: m.ClassroomID,
				})
				if err != nil {
					return nil, err
				}

				topicTags := ""
				if clrRes.GetClassroom().TopicTags != nil {
					topicTags = clrRes.GetClassroom().GetTopicTags()
				}

				lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
					Id: clrRes.GetClassroom().GetLecturerID(),
				})
				if err != nil {
					return nil, err
				}

				members = append(members, &pb.MemberUserResponse{
					Classroom: &pb.ClassroomUserResponse{
						Id:          clrRes.GetClassroom().GetId(),
						Title:       clrRes.GetClassroom().GetTitle(),
						Description: clrRes.GetClassroom().GetDescription(),
						Status:      clrRes.GetClassroom().GetStatus(),
						Lecturer: &pb.UserResponse{
							Class:    lecturerRes.User.Class,
							Major:    lecturerRes.User.Major,
							Phone:    lecturerRes.User.Phone,
							PhotoSrc: lecturerRes.User.PhotoSrc,
							Role:     lecturerRes.User.Role,
							Name:     lecturerRes.User.Name,
							Email:    lecturerRes.User.Email,
							Id:       lecturerRes.User.Id,
						},
						ClassCourse:     clrRes.GetClassroom().GetClassCourse(),
						TopicTags:       &topicTags,
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
						Id:       userRes.User.Id,
					},
					Status:    m.Status,
					IsDefense: m.IsDefense,
				})
			}

			return &pb.CheckStatusUserJoinClassroomResponse{
				Member: members,
				Status: "WAITING",
			}, nil
		} else {
			return &pb.CheckStatusUserJoinClassroomResponse{
				Response: &pb.CommonUserResponse{
					StatusCode: memberRes.GetResponse().StatusCode,
					Message:    memberRes.GetResponse().Message,
				},
				Member: nil,
				Status: "NOT SUBSCRIBED YET",
			}, nil
		}
	}
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

	class := req.GetUser().GetClass()
	major := req.GetUser().GetMajor()
	phone := req.GetUser().GetPhone()

	res, err := u.userClient.UpdateUser(ctx, &userSvcV1.UpdateUserRequest{
		Id: req.GetId(),
		User: &userSvcV1.UserInput{
			Class:    &class,
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
		topic, err := u.topicClient.GetAllTopicsOfListUser(ctx, &userSvcV1.GetAllTopicsOfListUserRequest{
			UserID: []string{l.Id},
		})
		if err != nil {
			return nil, err
		}

		var topicRes *pb.TopicUserResponse
		if len(topic.Topic) > 0 {
			topicRes = &pb.TopicUserResponse{
				Id:             topic.Topic[0].GetId(),
				Title:          topic.Topic[0].GetTitle(),
				TypeTopic:      topic.Topic[0].GetTypeTopic(),
				MemberQuantity: topic.Topic[0].GetMemberQuantity(),
				Student:        &pb.UserResponse{},
				MemberEmail:    topic.Topic[0].GetMemberEmail(),
				Description:    topic.Topic[0].GetDescription(),
			}
		}
		lecturers = append(lecturers, &pb.UserResponse{
			Id:       l.Id,
			Class:    l.Class,
			Major:    l.Major,
			Phone:    l.Phone,
			PhotoSrc: l.PhotoSrc,
			Role:     l.Role,
			Name:     l.Name,
			Email:    l.Email,
			Topic:    topicRes,
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

func (u *userServiceGW) GetUserAttachment(ctx context.Context, req *pb.GetUserAttachmentRequest) (*pb.GetUserAttachmentResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	att, err := u.attachmentClient.GetFinalFile(ctx, &classroomSvcV1.GetFinalFileRequest{
		AuthorID: user.GetUser().GetId(),
	})
	if err != nil {
		log.Println(err)
		if errors.Is(err, status.Errorf(codes.NotFound, "err: attachment not found")) {
			return &pb.GetUserAttachmentResponse{
				FinalFile: &pb.FinalFile{
					Attachment: nil,
					Author: &pb.UserResponse{
						Id:       user.GetUser().Id,
						Class:    user.GetUser().Class,
						Major:    user.GetUser().Major,
						Phone:    user.GetUser().Phone,
						PhotoSrc: user.GetUser().PhotoSrc,
						Role:     user.GetUser().Role,
						Name:     user.GetUser().Name,
						Email:    user.GetUser().Email,
					},
				},
			}, nil
		}
	}

	if att.GetAttachment() == nil {
		return &pb.GetUserAttachmentResponse{
			FinalFile: &pb.FinalFile{
				Attachment: nil,
				Author: &pb.UserResponse{
					Id:       user.GetUser().Id,
					Class:    user.GetUser().Class,
					Major:    user.GetUser().Major,
					Phone:    user.GetUser().Phone,
					PhotoSrc: user.GetUser().PhotoSrc,
					Role:     user.GetUser().Role,
					Name:     user.GetUser().Name,
					Email:    user.GetUser().Email,
				},
			},
		}, nil
	}

	return &pb.GetUserAttachmentResponse{
		FinalFile: &pb.FinalFile{
			Attachment: &pb.AttachmentFinalFileResponse{
				Id:        att.GetAttachment().Id,
				FileURL:   att.GetAttachment().FileURL,
				Name:      att.GetAttachment().Name,
				Status:    att.GetAttachment().GetStatus(),
				Type:      att.GetAttachment().Type,
				Thumbnail: att.GetAttachment().Thumbnail,
				Size:      att.GetAttachment().Size,
				CreatedAt: att.GetAttachment().CreatedAt,
			},
			Author: &pb.UserResponse{
				Id:       user.GetUser().Id,
				Class:    user.GetUser().Class,
				Major:    user.GetUser().Major,
				Phone:    user.GetUser().Phone,
				PhotoSrc: user.GetUser().PhotoSrc,
				Role:     user.GetUser().Role,
				Name:     user.GetUser().Name,
				Email:    user.GetUser().Email,
			},
		},
	}, nil
}
