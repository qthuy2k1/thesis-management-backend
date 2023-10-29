package main

import (
	"context"
	"log"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type studentDefServiceGW struct {
	pb.UnimplementedStudentDefServiceServer
	userClient userSvcV1.UserServiceClient
}

func NewStudentDefsService(userClient userSvcV1.UserServiceClient) *studentDefServiceGW {
	return &studentDefServiceGW{
		userClient: userClient,
	}
}

func (u *studentDefServiceGW) CreateStudentDef(ctx context.Context, req *pb.CreateStudentDefRequest) (*pb.CreateStudentDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	log.Println(u.userClient)

	studentRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetStudentDef().UserID,
	})
	if err != nil {
		return nil, err
	}

	if studentRes.GetResponse().StatusCode != 200 {
		return &pb.CreateStudentDefResponse{
			Response: &pb.CommonStudentDefResponse{
				StatusCode: studentRes.GetResponse().StatusCode,
				Message:    studentRes.GetResponse().Message,
			},
		}, nil
	}

	instructorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetStudentDef().InstructorID,
	})
	if err != nil {
		return nil, err
	}

	if instructorRes.GetResponse().StatusCode != 200 {
		return &pb.CreateStudentDefResponse{
			Response: &pb.CommonStudentDefResponse{
				StatusCode: instructorRes.GetResponse().StatusCode,
				Message:    instructorRes.GetResponse().Message,
			},
		}, nil
	}

	res, err := u.userClient.CreateStudentDef(ctx, &userSvcV1.CreateStudentDefRequest{
		StudentDef: &userSvcV1.StudentDefInput{
			UserID:       req.GetStudentDef().UserID,
			InstructorID: req.GetStudentDef().InstructorID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateStudentDefResponse{
		Response: &pb.CommonStudentDefResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *studentDefServiceGW) GetStudentDef(ctx context.Context, req *pb.GetStudentDefRequest) (*pb.GetStudentDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.GetStudentDef(ctx, &userSvcV1.GetStudentDefRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: res.GetStudentDef().UserID,
	})
	if err != nil {
		return nil, err
	}

	if userRes.GetResponse().StatusCode != 200 {
		return &pb.GetStudentDefResponse{
			Response: &pb.CommonStudentDefResponse{
				StatusCode: userRes.GetResponse().StatusCode,
				Message:    userRes.GetResponse().Message,
			},
		}, nil
	}

	instructorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: res.GetStudentDef().InstructorID,
	})
	if err != nil {
		return nil, err
	}

	if userRes.GetResponse().StatusCode != 200 {
		return &pb.GetStudentDefResponse{
			Response: &pb.CommonStudentDefResponse{
				StatusCode: userRes.GetResponse().StatusCode,
				Message:    userRes.GetResponse().Message,
			},
		}, nil
	}

	return &pb.GetStudentDefResponse{
		Response: &pb.CommonStudentDefResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		StudentDef: &pb.StudentDefResponse{
			Id: res.GetStudentDef().Id,
			Infor: &pb.StudentDefUserResponse{
				Id:       userRes.GetUser().GetId(),
				Class:    userRes.GetUser().GetClass(),
				Major:    userRes.GetUser().Major,
				Phone:    userRes.GetUser().Phone,
				PhotoSrc: userRes.GetUser().GetPhotoSrc(),
				Name:     userRes.GetUser().GetName(),
				Email:    userRes.GetUser().GetEmail(),
				Role:     userRes.GetUser().GetRole(),
			},
			Instructor: &pb.StudentDefUserResponse{
				Id:       instructorRes.GetUser().GetId(),
				Class:    instructorRes.GetUser().GetClass(),
				Major:    instructorRes.GetUser().Major,
				Phone:    instructorRes.GetUser().Phone,
				PhotoSrc: instructorRes.GetUser().GetPhotoSrc(),
				Name:     instructorRes.GetUser().GetName(),
				Email:    instructorRes.GetUser().GetEmail(),
				Role:     instructorRes.GetUser().GetRole(),
			},
		},
	}, nil
}

func (u *studentDefServiceGW) UpdateStudentDef(ctx context.Context, req *pb.UpdateStudentDefRequest) (*pb.UpdateStudentDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.GetStudentDef().UserID,
	})
	if err != nil {
		return nil, err
	}

	if userRes.GetResponse().StatusCode != 200 {
		return &pb.UpdateStudentDefResponse{
			Response: &pb.CommonStudentDefResponse{
				StatusCode: userRes.GetResponse().StatusCode,
				Message:    userRes.GetResponse().Message,
			},
		}, nil
	}

	res, err := u.userClient.UpdateStudentDef(ctx, &userSvcV1.UpdateStudentDefRequest{
		Id: req.GetId(),
		StudentDef: &userSvcV1.StudentDefInput{
			UserID:       req.GetStudentDef().UserID,
			InstructorID: req.GetStudentDef().InstructorID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateStudentDefResponse{
		Response: &pb.CommonStudentDefResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *studentDefServiceGW) DeleteStudentDef(ctx context.Context, req *pb.DeleteStudentDefRequest) (*pb.DeleteStudentDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.DeleteStudentDef(ctx, &userSvcV1.DeleteStudentDefRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteStudentDefResponse{
		Response: &pb.CommonStudentDefResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *studentDefServiceGW) GetStudentDefs(ctx context.Context, req *pb.GetStudentDefsRequest) (*pb.GetStudentDefsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.GetStudentDefs(ctx, &userSvcV1.GetStudentDefsRequest{})
	if err != nil {
		return nil, err
	}

	var studentDefs []*pb.StudentDefResponse
	for _, s := range res.GetStudentDefs() {
		userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: s.UserID,
		})
		if err != nil {
			return nil, err
		}

		if userRes.GetResponse().StatusCode != 200 {
			return &pb.GetStudentDefsResponse{
				Response: &pb.CommonStudentDefResponse{
					StatusCode: userRes.GetResponse().StatusCode,
					Message:    userRes.GetResponse().Message,
				},
			}, nil
		}

		instructorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: s.InstructorID,
		})
		if err != nil {
			return nil, err
		}

		if userRes.GetResponse().StatusCode != 200 {
			return &pb.GetStudentDefsResponse{
				Response: &pb.CommonStudentDefResponse{
					StatusCode: userRes.GetResponse().StatusCode,
					Message:    userRes.GetResponse().Message,
				},
			}, nil
		}

		studentDefs = append(studentDefs, &pb.StudentDefResponse{
			Id: s.Id,
			Infor: &pb.StudentDefUserResponse{
				Id:       userRes.GetUser().GetId(),
				Class:    userRes.GetUser().GetClass(),
				Major:    userRes.GetUser().Major,
				Phone:    userRes.GetUser().Phone,
				PhotoSrc: userRes.GetUser().GetPhotoSrc(),
				Name:     userRes.GetUser().GetName(),
				Email:    userRes.GetUser().GetEmail(),
				Role:     userRes.GetUser().GetRole(),
			},
			Instructor: &pb.StudentDefUserResponse{
				Id:       instructorRes.GetUser().GetId(),
				Class:    instructorRes.GetUser().GetClass(),
				Major:    instructorRes.GetUser().Major,
				Phone:    instructorRes.GetUser().Phone,
				PhotoSrc: instructorRes.GetUser().GetPhotoSrc(),
				Name:     instructorRes.GetUser().GetName(),
				Email:    instructorRes.GetUser().GetEmail(),
				Role:     instructorRes.GetUser().GetRole(),
			},
		})
	}

	return &pb.GetStudentDefsResponse{
		TotalCount:  res.GetTotalCount(),
		StudentDefs: studentDefs,
	}, nil
}

func (u *studentDefServiceGW) GetAllStudentDefsOfInstructor(ctx context.Context, req *pb.GetAllStudentDefsOfInstructorRequest) (*pb.GetAllStudentDefsOfInstructorResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.userClient.GetAllStudentDefsOfInstructor(ctx, &userSvcV1.GetAllStudentDefsOfInstructorRequest{
		InstructorID: req.GetInstructorID(),
	})
	if err != nil {
		return nil, err
	}

	var studentDefs []*pb.StudentDefResponse
	for _, s := range res.GetStudentDefs() {
		userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: s.UserID,
		})
		if err != nil {
			return nil, err
		}

		if userRes.GetResponse().StatusCode != 200 {
			return &pb.GetAllStudentDefsOfInstructorResponse{
				Response: &pb.CommonStudentDefResponse{
					StatusCode: userRes.GetResponse().StatusCode,
					Message:    userRes.GetResponse().Message,
				},
			}, nil
		}

		instructorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: s.InstructorID,
		})
		if err != nil {
			return nil, err
		}

		if userRes.GetResponse().StatusCode != 200 {
			return &pb.GetAllStudentDefsOfInstructorResponse{
				Response: &pb.CommonStudentDefResponse{
					StatusCode: userRes.GetResponse().StatusCode,
					Message:    userRes.GetResponse().Message,
				},
			}, nil
		}

		studentDefs = append(studentDefs, &pb.StudentDefResponse{
			Id: s.Id,
			Infor: &pb.StudentDefUserResponse{
				Id:       userRes.GetUser().GetId(),
				Class:    userRes.GetUser().GetClass(),
				Major:    userRes.GetUser().Major,
				Phone:    userRes.GetUser().Phone,
				PhotoSrc: userRes.GetUser().GetPhotoSrc(),
				Name:     userRes.GetUser().GetName(),
				Email:    userRes.GetUser().GetEmail(),
				Role:     userRes.GetUser().GetRole(),
			},
			Instructor: &pb.StudentDefUserResponse{
				Id:       instructorRes.GetUser().GetId(),
				Class:    instructorRes.GetUser().GetClass(),
				Major:    instructorRes.GetUser().Major,
				Phone:    instructorRes.GetUser().Phone,
				PhotoSrc: instructorRes.GetUser().GetPhotoSrc(),
				Name:     instructorRes.GetUser().GetName(),
				Email:    instructorRes.GetUser().GetEmail(),
				Role:     instructorRes.GetUser().GetRole(),
			},
		})
	}

	return &pb.GetAllStudentDefsOfInstructorResponse{
		TotalCount:  res.GetTotalCount(),
		StudentDefs: studentDefs,
	}, nil
}
