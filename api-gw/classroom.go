package main

import (
	"context"
	"strings"

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
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *classroomServiceGW) GetClassroom(ctx context.Context, req *pb.GetClassroomRequest) (*pb.GetClassroomResponse, error) {
	res, err := u.classroomClient.GetClassroom(ctx, &classroomSvcV1.GetClassroomRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetClassroomResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
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

func (u *classroomServiceGW) UpdateClassroom(ctx context.Context, req *pb.UpdateClassroomRequest) (*pb.UpdateClassroomResponse, error) {
	res, err := u.classroomClient.UpdateClassroom(ctx, &classroomSvcV1.UpdateClassroomRequest{
		Id: req.GetId(),
		Classroom: &classroomSvcV1.ClassroomInput{
			Title:       req.GetClassroom().Title,
			Description: req.GetClassroom().Description,
			Status:      req.Classroom.GetStatus(),
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateClassroomResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *classroomServiceGW) DeleteClassroom(ctx context.Context, req *pb.DeleteClassroomRequest) (*pb.DeleteClassroomResponse, error) {
	res, err := u.classroomClient.DeleteClassroom(ctx, &classroomSvcV1.DeleteClassroomRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteClassroomResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *classroomServiceGW) GetClassrooms(ctx context.Context, req *pb.GetClassroomsRequest) (*pb.GetClassroomsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	filter := &classroomSvcV1.GetClassroomsRequest{}

	if req.GetLimit() > 0 {
		filter.Limit = req.GetLimit()
	} else {
		filter.Limit = 5
	}

	if req.GetPage() > 0 {
		filter.Page = req.GetPage()
	} else {
		filter.Page = 1
	}

	titleSearchTrim := strings.TrimSpace(req.GetTitleSearch())
	if len(titleSearchTrim) > 0 {
		filter.TitleSearch = titleSearchTrim
	}

	sortColumnTrim := strings.TrimSpace(req.GetSortColumn())
	if len(sortColumnTrim) > 0 {
		columns := map[string]string{
			"id":          "id",
			"title":       "title",
			"description": "description",
			"status":      "status",
			"created_at":  "created_at",
			"updated_at":  "updated_at",
		}
		if stringInMap(sortColumnTrim, columns) {
			filter.SortColumn = sortColumnTrim
		} else {
			filter.SortColumn = "id"
		}
	} else {
		filter.SortColumn = "id"
	}

	sortOrder := "asc"
	if req.IsDesc {
		sortOrder = "desc"
	}

	res, err := u.classroomClient.GetClassrooms(ctx, &classroomSvcV1.GetClassroomsRequest{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   sortOrder,
	})
	if err != nil {
		return nil, err
	}

	var classrooms []*pb.ClassroomResponse
	for _, c := range res.GetClassrooms() {
		classrooms = append(classrooms, &pb.ClassroomResponse{
			Id:          c.Id,
			Title:       c.Title,
			Description: c.Description,
			Status:      c.Status,
			CreatedAt:   c.CreatedAt,
			UpdatedAt:   c.UpdatedAt,
		})
	}

	return &pb.GetClassroomsResponse{
		Response: &pb.CommonClassroomResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Classrooms: classrooms,
	}, nil
}

func stringInMap(s string, m map[string]string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}
