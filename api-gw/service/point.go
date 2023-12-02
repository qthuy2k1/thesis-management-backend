package service

import (
	"context"
	"log"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	pointSvcV1 "github.com/qthuy2k1/thesis-management-backend/schedule-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type pointServiceGW struct {
	pb.UnimplementedPointServiceServer
	pointClient pointSvcV1.ScheduleServiceClient
	userClient  userSvcV1.UserServiceClient
}

func NewPointsService(pointClient pointSvcV1.ScheduleServiceClient, userClient userSvcV1.UserServiceClient) *pointServiceGW {
	return &pointServiceGW{
		pointClient: pointClient,
		userClient:  userClient,
	}
}

func (u *pointServiceGW) GetAllPointDef(ctx context.Context, req *pb.GetAllPointDefRequest) (*pb.GetAllPointDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	res, err := u.pointClient.GetAllPointDefs(ctx, &pointSvcV1.GetAllPointDefsRequest{
		Id: req.GetUserID(),
	})
	if err != nil {
		return nil, err
	}

	var pointRes []*pb.PointResponse
	for _, t := range res.Points {
		var assessItems []*pb.AssessItemResponse
		for _, c := range t.Assesses {
			assessItems = append(assessItems, &pb.AssessItemResponse{
				Id: c.Id,
				Lecturer: &pb.UserPointResponse{
					Id:       c.Lecturer.GetId(),
					Class:    c.Lecturer.Class,
					Major:    c.Lecturer.Major,
					Phone:    c.Lecturer.Phone,
					PhotoSrc: c.Lecturer.GetPhotoSrc(),
					Name:     c.Lecturer.GetName(),
					Email:    c.Lecturer.GetEmail(),
					Role:     c.Lecturer.GetRole(),
				},
				Point:   c.Point,
				Comment: c.Comment,
			})
		}

		pointRes = append(pointRes, &pb.PointResponse{
			Id: t.Id,
			Student: &pb.UserPointResponse{
				Id:       t.Student.GetId(),
				Class:    t.Student.Class,
				Major:    t.Student.Major,
				Phone:    t.Student.Phone,
				PhotoSrc: t.Student.GetPhotoSrc(),
				Name:     t.Student.GetName(),
				Email:    t.Student.GetEmail(),
				Role:     t.Student.GetRole(),
			},
			Assesses: assessItems,
		})
	}

	return &pb.GetAllPointDefResponse{
		Points: pointRes,
	}, nil
}

func (u *pointServiceGW) CreateOrUpdatePointDef(ctx context.Context, req *pb.CreateOrUpdatePointDefRequest) (*pb.CreateOrUpdatePointDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	studentRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.Point.StudentID,
	})
	if err != nil {
		return nil, err
	}

	var assess []*pointSvcV1.AssessItem
	for _, a := range req.Point.Assesses {
		lecturerRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{Id: a.LecturerID})
		if err != nil {
			return nil, err
		}

		assess = append(assess, &pointSvcV1.AssessItem{
			Id: a.Id,
			Lecturer: &pointSvcV1.UserScheduleResponse{
				Id:       lecturerRes.User.Id,
				Class:    lecturerRes.User.Class,
				Major:    lecturerRes.User.Major,
				Phone:    lecturerRes.User.Phone,
				PhotoSrc: lecturerRes.User.PhotoSrc,
				Role:     lecturerRes.User.Role,
				Name:     lecturerRes.User.Name,
				Email:    lecturerRes.User.Email,
			},
			Point:   a.Point,
			Comment: a.Comment,
		})
	}

	res, err := u.pointClient.CreateOrUpdatePointDef(ctx, &pointSvcV1.CreateOrUpdatePointDefRequest{
		Point: &pointSvcV1.Point{
			Id: req.GetPoint().Id,
			Student: &pointSvcV1.UserScheduleResponse{
				Id:       studentRes.User.Id,
				Class:    studentRes.User.Class,
				Major:    studentRes.User.Major,
				Phone:    studentRes.User.Phone,
				PhotoSrc: studentRes.User.PhotoSrc,
				Role:     studentRes.User.Role,
				Name:     studentRes.User.Name,
				Email:    studentRes.User.Email,
			},
			Assesses: assess,
		},
	})
	if err != nil {
		log.Println("create error", err)
		return nil, err
	}

	var assessRes []*pb.AssessItemResponse
	for _, a := range res.GetPoint().Assesses {
		assessRes = append(assessRes, &pb.AssessItemResponse{
			Id: a.Id,
			Lecturer: &pb.UserPointResponse{
				Id:       a.Lecturer.Id,
				Class:    a.Lecturer.Class,
				Major:    a.Lecturer.Major,
				Phone:    a.Lecturer.Phone,
				PhotoSrc: a.Lecturer.PhotoSrc,
				Role:     a.Lecturer.Role,
				Name:     a.Lecturer.Name,
				Email:    a.Lecturer.Email,
			},
			Point:   a.Point,
			Comment: a.Comment,
		})
	}

	return &pb.CreateOrUpdatePointDefResponse{
		Point: &pb.PointResponse{
			Id: res.Point.Id,
			Student: &pb.UserPointResponse{
				Id:       res.Point.Student.Id,
				Class:    res.Point.Student.Class,
				Major:    res.Point.Student.Major,
				Phone:    res.Point.Student.Phone,
				PhotoSrc: res.Point.Student.PhotoSrc,
				Role:     res.Point.Student.Role,
				Name:     res.Point.Student.Name,
				Email:    res.Point.Student.Email,
			},
			Assesses: assessRes,
		},
		Message: res.Message,
	}, nil

}
