package main

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
	// commiteeClient commiteeSvcV1.CommiteeServiceClient
	// thesisClient   commiteeSvcV1.PointServiceClient
	userClient userSvcV1.UserServiceClient
}

func NewPointsService(pointClient pointSvcV1.ScheduleServiceClient, userClient userSvcV1.UserServiceClient) *pointServiceGW {
	return &pointServiceGW{
		pointClient: pointClient,
		userClient:  userClient,
	}
}

// func (u *pointServiceGW) GetAllPointDef(ctx context.Context, req *pb.GetAllPointDefRequest) (*pb.GetAllPointDefResponse, error) {
// 	if err := req.Validate(); err != nil {
// 		return nil, err
// 	}

// 	res, err := u.pointClient.Get(ctx, &pointSvcV1.GetPointsRequest{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var thesisRes []*pb.Thesis
// 	for _, t := range res.Thesis {
// 		var councilsPointResponse []*pb.Council
// 		for _, c := range t.Council {
// 			councilsPointResponse = append(councilsPointResponse, &pb.Council{
// 				Id:       c.Id,
// 				Class:    c.Class,
// 				Major:    c.Major,
// 				Phone:    c.Phone,
// 				PhotoSrc: c.PhotoSrc,
// 				Role:     c.Role,
// 				Name:     c.Name,
// 				Email:    c.Email,
// 			})
// 		}

// 		var timeSlotsResponse []*pb.TimeSlots
// 		for _, t := range t.Point.TimeSlots {
// 			timeSlotsResponse = append(timeSlotsResponse, &pb.TimeSlots{
// 				Student: &pb.StudentDefPointResponse{
// 					Id: t.Student.Id,
// 					Infor: &pb.UserPointResponse{
// 						Id:       t.Student.Infor.GetId(),
// 						Class:    t.Student.Infor.GetClass(),
// 						Major:    t.Student.Infor.Major,
// 						Phone:    t.Student.Infor.Phone,
// 						PhotoSrc: t.Student.Infor.GetPhotoSrc(),
// 						Name:     t.Student.Infor.GetName(),
// 						Email:    t.Student.Infor.GetEmail(),
// 						Role:     t.Student.Infor.GetRole(),
// 					},
// 					Instructor: &pb.UserPointResponse{
// 						Id:       t.Student.Instructor.GetId(),
// 						Class:    t.Student.Instructor.GetClass(),
// 						Major:    t.Student.Instructor.Major,
// 						Phone:    t.Student.Instructor.Phone,
// 						PhotoSrc: t.Student.Instructor.GetPhotoSrc(),
// 						Name:     t.Student.Instructor.GetName(),
// 						Email:    t.Student.Instructor.GetEmail(),
// 						Role:     t.Student.Instructor.GetRole(),
// 					},
// 				},
// 				TimeSlot: &pb.TimeSlot{
// 					Date:  t.TimeSlot.Date,
// 					Shift: t.TimeSlot.Shift,
// 					Id:    t.TimeSlot.Id,
// 					Time:  t.TimeSlot.Time,
// 				},
// 			})
// 		}

// 		thesisRes = append(thesisRes, &pb.Thesis{
// 			Point: &pb.Point{
// 				TimeSlots: timeSlotsResponse,
// 				Room: &pb.RoomPoint{
// 					Id:          t.Point.Room.Id,
// 					Name:        t.Point.Room.Name,
// 					School:      t.Point.Room.School,
// 					Type:        t.Point.Room.Type,
// 					Description: t.Point.Room.Description,
// 				},
// 			},
// 			Council: councilsPointResponse,
// 			Id:      t.Id,
// 		})
// 	}

// 	return &pb.GetPointsResponse{
// 		Thesis: thesisRes,
// 	}, nil
// }

func (u *pointServiceGW) CreateOrUpdatePointDef(ctx context.Context, req *pb.CreateOrUpdatePointDefRequest) (*pb.CreateOrUpdatePointDefResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	student, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: req.Point.StudentID,
	})
	if err != nil {
		return nil, err
	}

	var assess []*pointSvcV1.AssessItem
	for _, a := range req.Point.Assesses {
		lecturer, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: a.LecturerID,
		})
		if err != nil {
			return nil, err
		}

		assess = append(assess, &pointSvcV1.AssessItem{
			Id: a.Id,
			Lecturer: &pointSvcV1.UserScheduleResponse{
				Id:       lecturer.User.Id,
				Class:    lecturer.User.Class,
				Major:    lecturer.User.Major,
				Phone:    lecturer.User.Phone,
				PhotoSrc: lecturer.User.PhotoSrc,
				Role:     lecturer.User.Role,
				Name:     lecturer.User.Name,
				Email:    lecturer.User.Email,
			},
			Point:   a.Point,
			Comment: a.Comment,
		})
	}

	res, err := u.pointClient.CreateOrUpdatePointDef(ctx, &pointSvcV1.CreateOrUpdatePointDefRequest{
		Point: &pointSvcV1.Point{
			Id: req.Point.Id,
			Student: &pointSvcV1.UserScheduleResponse{
				Id:       student.User.Id,
				Class:    student.User.Class,
				Major:    student.User.Major,
				Phone:    student.User.Phone,
				PhotoSrc: student.User.PhotoSrc,
				Role:     student.User.Role,
				Name:     student.User.Name,
				Email:    student.User.Email,
			},
			Assesses: assess,
		},
	})
	if err != nil {
		log.Println("create error", err)
		return nil, err
	}

	var assessRes []*pb.AssessItemResponse
	for _, a := range res.Point.Assesses {
		assessRes = append(assessRes, &pb.AssessItemResponse{
			Id: a.Id,
			Lecturer: &pb.UserScheduleResponse{
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
			Student: &pb.UserScheduleResponse{
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
