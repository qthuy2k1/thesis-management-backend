package main

import (
	"context"
	"log"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
)

type postServiceGW struct {
	pb.UnimplementedPostServiceServer
	postClient      postSvcV1.PostServiceClient
	classroomClient classroomSvcV1.ClassroomServiceClient
}

func NewPostsService(postClient postSvcV1.PostServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient) *postServiceGW {
	return &postServiceGW{
		postClient:      postClient,
		classroomClient: classroomClient,
	}
}

func (u *postServiceGW) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetPost().ClassroomID})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.CreatePostResponse{
			Response: &pb.CommonPostResponse{
				StatusCode: 400,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	res, err := u.postClient.CreatePost(ctx, &postSvcV1.CreatePostRequest{
		Post: &postSvcV1.PostInput{
			Title:       req.GetPost().Title,
			Content:     req.GetPost().Content,
			ClassroomID: req.GetPost().ClassroomID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreatePostResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *postServiceGW) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	res, err := u.postClient.GetPost(ctx, &postSvcV1.GetPostRequest{Id: req.GetId()})
	if err != nil {
		return nil, err
	}

	return &pb.GetPostResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Post: &pb.PostResponse{
			Id:          res.GetPost().Id,
			Title:       res.GetPost().Title,
			Content:     res.GetPost().Content,
			ClassroomID: res.GetPost().ClassroomID,
			CreatedAt:   res.GetPost().CreatedAt,
			UpdatedAt:   res.GetPost().UpdatedAt,
		},
	}, nil
}

func (u *postServiceGW) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	log.Println(req)
	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetPost().ClassroomID})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.UpdatePostResponse{
			Response: &pb.CommonPostResponse{
				StatusCode: 400,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	res, err := u.postClient.UpdatePost(ctx, &postSvcV1.UpdatePostRequest{
		Id: req.GetId(),
		Post: &postSvcV1.PostInput{
			Title:       req.GetPost().Title,
			Content:     req.GetPost().Content,
			ClassroomID: req.GetPost().ClassroomID,
		},
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePostResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *postServiceGW) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	res, err := u.postClient.DeletePost(ctx, &postSvcV1.DeletePostRequest{
		Id: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.DeletePostResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *postServiceGW) GetPosts(ctx context.Context, req *pb.GetPostsRequest) (*pb.GetPostsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	filter := &postSvcV1.GetPostsRequest{}

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
			"id":           "id",
			"title":        "title",
			"content":      "content",
			"classroom_id": "classroom_id",
			"created_at":   "created_at",
			"updated_at":   "updated_at",
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

	res, err := u.postClient.GetPosts(ctx, &postSvcV1.GetPostsRequest{
		Limit:       filter.Limit,
		Page:        filter.Page,
		TitleSearch: filter.TitleSearch,
		SortColumn:  filter.SortColumn,
		SortOrder:   sortOrder,
	})
	if err != nil {
		return nil, err
	}

	var posts []*pb.PostResponse
	for _, p := range res.GetPosts() {
		posts = append(posts, &pb.PostResponse{
			Id:          p.Id,
			Title:       p.Title,
			Content:     p.Content,
			ClassroomID: p.ClassroomID,
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
		})
	}

	return &pb.GetPostsResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Posts:      posts,
	}, nil
}
