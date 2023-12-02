package service

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

	return &pb.GetStudentDefResponse{
		Response: &pb.CommonStudentDefResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		StudentDef: &pb.StudentDefResponse{
			Id: res.GetStudentDef().Id,
			Infor: &pb.StudentDefUserResponse{
				Id:       res.StudentDef.User.GetId(),
				Class:    res.StudentDef.User.Class,
				Major:    res.StudentDef.User.Major,
				Phone:    res.StudentDef.User.Phone,
				PhotoSrc: res.StudentDef.User.GetPhotoSrc(),
				Name:     res.StudentDef.User.GetName(),
				Email:    res.StudentDef.User.GetEmail(),
				Role:     res.StudentDef.User.GetRole(),
			},
			Instructor: &pb.StudentDefUserResponse{
				Id:       res.StudentDef.Instructor.GetId(),
				Class:    res.StudentDef.Instructor.Class,
				Major:    res.StudentDef.Instructor.Major,
				Phone:    res.StudentDef.Instructor.Phone,
				PhotoSrc: res.StudentDef.Instructor.GetPhotoSrc(),
				Name:     res.StudentDef.Instructor.GetName(),
				Email:    res.StudentDef.Instructor.GetEmail(),
				Role:     res.StudentDef.Instructor.GetRole(),
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
		studentDefs = append(studentDefs, &pb.StudentDefResponse{
			Id: s.Id,
			Infor: &pb.StudentDefUserResponse{
				Id:       s.User.GetId(),
				Class:    s.User.Class,
				Major:    s.User.Major,
				Phone:    s.User.Phone,
				PhotoSrc: s.User.GetPhotoSrc(),
				Name:     s.User.GetName(),
				Email:    s.User.GetEmail(),
				Role:     s.User.GetRole(),
			},
			Instructor: &pb.StudentDefUserResponse{
				Id:       s.Instructor.GetId(),
				Class:    s.Instructor.Class,
				Major:    s.Instructor.Major,
				Phone:    s.Instructor.Phone,
				PhotoSrc: s.Instructor.GetPhotoSrc(),
				Name:     s.Instructor.GetName(),
				Email:    s.Instructor.GetEmail(),
				Role:     s.Instructor.GetRole(),
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
		studentDefs = append(studentDefs, &pb.StudentDefResponse{
			Id: s.Id,
			Infor: &pb.StudentDefUserResponse{
				Id:       s.User.GetId(),
				Class:    s.User.Class,
				Major:    s.User.Major,
				Phone:    s.User.Phone,
				PhotoSrc: s.User.GetPhotoSrc(),
				Name:     s.User.GetName(),
				Email:    s.User.GetEmail(),
				Role:     s.User.GetRole(),
			},
			Instructor: &pb.StudentDefUserResponse{
				Id:       s.Instructor.GetId(),
				Class:    s.Instructor.Class,
				Major:    s.Instructor.Major,
				Phone:    s.Instructor.Phone,
				PhotoSrc: s.Instructor.GetPhotoSrc(),
				Name:     s.Instructor.GetName(),
				Email:    s.Instructor.GetEmail(),
				Role:     s.Instructor.GetRole(),
			},
		})
	}

	return &pb.GetAllStudentDefsOfInstructorResponse{
		TotalCount:  res.GetTotalCount(),
		StudentDefs: studentDefs,
	}, nil
}
