package main

import (
	"context"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
)

type classroomServiceGW struct {
	pb.UnimplementedClassroomServiceServer
	classroomClient classroomSvcV1.ClassroomServiceClient
}

func NewClassroomsService(classroomClient classroomSvcV1.ClassroomServiceClient) *classroomServiceGW {
	return &classroomServiceGW{
		classroomClient: classroomClient,
	}
}

func (u *classroomServiceGW) CreateClassroom(ctx context.Context, req *pb.CreateClassroomRequest) (*pb.CreateClassroomResponse, error) {
	res, err := u.classroomClient.CreateClassroom(ctx, &classroomSvcV1.CreateClassroomRequest{
		Classroom: &classroomSvcV1.ClassroomInput{
			Title:       req.GetClassroom().Title,
			Description: req.GetClassroom().Description,
			Status:      req.Classroom.GetStatus(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateClassroomResponse{
		StatusCode: res.StatusCode,
		Message:    res.Message,
	}, nil
}

func (u *classroomServiceGW) GetClassroom(ctx context.Context, req *pb.GetClassroomRequest) (*pb.GetClassroomResponse, error) {
	res, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetClassroomResponse{
		StatusCode: res.StatusCode,
		Message:    res.Message,
		Classroom: &pb.ClassroomResponse{
			Id:          res.GetClassroom().Id,
			Title:       res.GetClassroom().Title,
			Description: res.GetClassroom().Description,
			Status:      res.GetClassroom().Status,
			CreatedAt:   res.GetClassroom().CreatedAt,
			UpdatedAt:   res.GetClassroom().UpdatedAt,
		},
	}, nil
}
