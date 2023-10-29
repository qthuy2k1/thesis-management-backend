package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	commiteeSvcV1 "github.com/qthuy2k1/thesis-management-backend/thesis-commitee-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type commiteeServiceGW struct {
	pb.UnimplementedCommiteeServiceServer
	commiteeClient commiteeSvcV1.CommiteeServiceClient
	userClient     userSvcV1.UserServiceClient
}

func NewCommiteesService(commiteeClient commiteeSvcV1.CommiteeServiceClient, userClient userSvcV1.UserServiceClient) *commiteeServiceGW {
	return &commiteeServiceGW{
		commiteeClient: commiteeClient,
		userClient:     userClient,
	}
}

func (u *commiteeServiceGW) CreateCommitee(ctx context.Context, req *pb.CreateCommiteeRequest) (*pb.CreateCommiteeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.commiteeClient.CreateCommitee(ctx, &commiteeSvcV1.CreateCommiteeRequest{
		Commitee: &commiteeSvcV1.CommiteeInput{
			StartDate: req.GetCommitee().StartDate,
			Shift:     req.GetCommitee().Shift,
			RoomID:    req.GetCommitee().RoomID,
		},
	})
	if err != nil {
		return nil, err
	}

	var details []*commiteeSvcV1.CreateCommiteeUserDetailResponse
	for _, lecID := range req.Commitee.LecturerID {
		detailRes, err := u.commiteeClient.CreateCommiteeUserDetail(ctx, &commiteeSvcV1.CreateCommiteeUserDetailRequest{
			CommiteeUserDetail: &commiteeSvcV1.CommiteeUserDetail{
				CommiteeID: res.GetCommitee().Id,
				LecturerID: lecID,
				StudentID:  req.Commitee.StudentID,
			},
		})
		if err != nil {
			if len(details) > 0 {
				for _, d := range details {
					if _, err := u.commiteeClient.DeleteCommiteeUserDetail(ctx, &commiteeSvcV1.DeleteCommiteeUserDetailRequest{
						CommiteeID: d.CommiteeUserDetail.CommiteeID,
						LecturerID: d.CommiteeUserDetail.LecturerID,
						StudentID:  d.CommiteeUserDetail.StudentID,
					}); err != nil {
						return nil, err
					}
				}

				if _, err := u.commiteeClient.DeleteCommitee(ctx, &commiteeSvcV1.DeleteCommiteeRequest{
					Id: res.Commitee.Id,
				}); err != nil {
					return nil, err
				}
			}
			return nil, err
		}

		details = append(details, detailRes)

		if detailRes.GetResponse().StatusCode != 201 {
			return &pb.CreateCommiteeResponse{
				Response: &pb.CommonCommiteeResponse{
					StatusCode: detailRes.GetResponse().StatusCode,
					Message:    detailRes.GetResponse().Message,
				},
			}, nil
		}
	}

	return &pb.CreateCommiteeResponse{
		Response: &pb.CommonCommiteeResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *commiteeServiceGW) GetCommitee(ctx context.Context, req *pb.GetCommiteeRequest) (*pb.GetCommiteeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.commiteeClient.GetCommitee(ctx, &commiteeSvcV1.GetCommiteeRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	detailsRes, err := u.commiteeClient.GetAllCommiteeUserDetailsFromCommitee(ctx, &commiteeSvcV1.GetAllCommiteeUserDetailsFromCommiteeRequest{
		CommiteeID: res.Commitee.Id,
	})
	if err != nil {
		return nil, err
	}

	var lecturers []*pb.UserCommiteeResponse
	for _, l := range detailsRes.CommiteeUserDetails {
		lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: l.LecturerID,
		})
		if err != nil {
			return nil, err
		}

		lecturers = append(lecturers, &pb.UserCommiteeResponse{
			Id:       lecturerRes.User.Id,
			Class:    lecturerRes.User.Class,
			Major:    lecturerRes.User.Major,
			Phone:    lecturerRes.User.Phone,
			PhotoSrc: lecturerRes.User.PhotoSrc,
			Role:     lecturerRes.User.Role,
			Name:     lecturerRes.User.Name,
			Email:    lecturerRes.User.Email,
		})
	}

	var studentList []*pb.UserCommiteeResponse
	for _, userID := range detailsRes.CommiteeUserDetails[0].StudentID {
		studentRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: userID,
		})
		if err != nil {
			return nil, err
		}

		studentList = append(studentList, &pb.UserCommiteeResponse{
			Id:       studentRes.User.Id,
			Class:    studentRes.User.Class,
			Major:    studentRes.User.Major,
			Phone:    studentRes.User.Phone,
			PhotoSrc: studentRes.User.PhotoSrc,
			Role:     studentRes.User.Role,
			Name:     studentRes.User.Name,
			Email:    studentRes.User.Email,
		})
	}

	return &pb.GetCommiteeResponse{
		Response: &pb.CommonCommiteeResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Commitee: &pb.CommiteeResponse{
			Id:        res.GetCommitee().Id,
			StartDate: res.GetCommitee().StartDate,
			Shift:     res.GetCommitee().Shift,
			RoomID:    res.GetCommitee().RoomID,
			Lecturers: lecturers,
			Student:   studentList,
		},
	}, nil
}

func (u *commiteeServiceGW) UpdateCommitee(ctx context.Context, req *pb.UpdateCommiteeRequest) (*pb.UpdateCommiteeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.commiteeClient.UpdateCommitee(ctx, &commiteeSvcV1.UpdateCommiteeRequest{
		Id: req.GetId(),
		Commitee: &commiteeSvcV1.CommiteeInput{
			StartDate: req.GetCommitee().StartDate,
			Shift:     req.GetCommitee().Shift,
			RoomID:    req.GetCommitee().RoomID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCommiteeResponse{
		Response: &pb.CommonCommiteeResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *commiteeServiceGW) DeleteCommitee(ctx context.Context, req *pb.DeleteCommiteeRequest) (*pb.DeleteCommiteeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.commiteeClient.DeleteCommitee(ctx, &commiteeSvcV1.DeleteCommiteeRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteCommiteeResponse{
		Response: &pb.CommonCommiteeResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *commiteeServiceGW) GetCommitees(ctx context.Context, req *pb.GetCommiteesRequest) (*pb.GetCommiteesResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.commiteeClient.GetCommitees(ctx, &commiteeSvcV1.GetCommiteesRequest{})
	if err != nil {
		return nil, err
	}

	var commitees []*pb.CommiteeResponse
	for _, p := range res.GetCommitees() {
		detailsRes, err := u.commiteeClient.GetAllCommiteeUserDetailsFromCommitee(ctx, &commiteeSvcV1.GetAllCommiteeUserDetailsFromCommiteeRequest{
			CommiteeID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var lecturers []*pb.UserCommiteeResponse
		for _, l := range detailsRes.CommiteeUserDetails {
			lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: l.LecturerID,
			})
			if err != nil {
				return nil, err
			}

			lecturers = append(lecturers, &pb.UserCommiteeResponse{
				Id:       lecturerRes.User.Id,
				Class:    lecturerRes.User.Class,
				Major:    lecturerRes.User.Major,
				Phone:    lecturerRes.User.Phone,
				PhotoSrc: lecturerRes.User.PhotoSrc,
				Role:     lecturerRes.User.Role,
				Name:     lecturerRes.User.Name,
				Email:    lecturerRes.User.Email,
			})
		}

		var studentList []*pb.UserCommiteeResponse
		for _, userID := range detailsRes.CommiteeUserDetails[0].StudentID {
			studentRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: userID,
			})
			if err != nil {
				return nil, err
			}

			studentList = append(studentList, &pb.UserCommiteeResponse{
				Id:       studentRes.User.Id,
				Class:    studentRes.User.Class,
				Major:    studentRes.User.Major,
				Phone:    studentRes.User.Phone,
				PhotoSrc: studentRes.User.PhotoSrc,
				Role:     studentRes.User.Role,
				Name:     studentRes.User.Name,
				Email:    studentRes.User.Email,
			})
		}

		commitees = append(commitees, &pb.CommiteeResponse{
			Id:        p.Id,
			StartDate: p.GetStartDate(),
			Shift:     p.GetShift(),
			RoomID:    p.GetRoomID(),
			Lecturers: lecturers,
			Student:   studentList,
		})
	}

	return &pb.GetCommiteesResponse{
		Response: &pb.CommonCommiteeResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Commitees:  commitees,
	}, nil
}
