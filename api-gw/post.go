package main

import (
	"context"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	classroomSvcV1 "github.com/qthuy2k1/thesis-management-backend/classroom-svc/api/goclient/v1"
	commentSvcV1 "github.com/qthuy2k1/thesis-management-backend/comment-svc/api/goclient/v1"
	postSvcV1 "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	rpsSvcV1 "github.com/qthuy2k1/thesis-management-backend/reporting-stage-svc/api/goclient/v1"
	userSvcV1 "github.com/qthuy2k1/thesis-management-backend/user-svc/api/goclient/v1"
)

type postServiceGW struct {
	pb.UnimplementedPostServiceServer
	postClient           postSvcV1.PostServiceClient
	classroomClient      classroomSvcV1.ClassroomServiceClient
	reportingStageClient rpsSvcV1.ReportingStageServiceClient
	commentClient        commentSvcV1.CommentServiceClient
	userClient           userSvcV1.UserServiceClient
}

func NewPostsService(postClient postSvcV1.PostServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, reportingStageClient rpsSvcV1.ReportingStageServiceClient, commentCLient commentSvcV1.CommentServiceClient, userClient userSvcV1.UserServiceClient) *postServiceGW {
	return &postServiceGW{
		postClient:           postClient,
		classroomClient:      classroomClient,
		reportingStageClient: reportingStageClient,
		commentClient:        commentCLient,
		userClient:           userClient,
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
				StatusCode: 404,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{Id: req.GetPost().GetReportingStageID()})
	if err != nil {
		return nil, err
	}

	if rpsRes.GetResponse().GetStatusCode() == 404 {
		return &pb.CreatePostResponse{
			Response: &pb.CommonPostResponse{
				StatusCode: 404,
				Message:    "Reporting stage does not exist",
			},
		}, nil
	}

	res, err := u.postClient.CreatePost(ctx, &postSvcV1.CreatePostRequest{
		Post: &postSvcV1.PostInput{
			Title:            req.GetPost().Title,
			Content:          req.GetPost().Content,
			ClassroomID:      req.GetPost().ClassroomID,
			ReportingStageID: req.GetPost().ReportingStageID,
			AuthorID:         req.GetPost().AuthorID,
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

	commentRes, err := u.commentClient.GetCommentsOfAPost(ctx, &commentSvcV1.GetCommentsOfAPostRequest{
		PostID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}

	var comments []*pb.CommentPostResponse
	for _, c := range commentRes.GetComments() {
		userRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: c.UserID,
		})
		if err != nil {
			return nil, err
		}

		comments = append(comments, &pb.CommentPostResponse{
			Id: c.Id,
			User: &pb.AuthorPostResponse{
				Id:          userRes.User.Id,
				Class:       userRes.User.Class,
				Major:       userRes.User.Major,
				Phone:       userRes.User.Phone,
				PhotoSrc:    userRes.User.PhotoSrc,
				Role:        userRes.User.Role,
				Name:        userRes.User.Name,
				Email:       userRes.User.Email,
				ClassroomID: &userRes.User.ClassroomID,
			},
			PostID:    *c.PostID,
			Content:   c.Content,
			CreatedAt: c.CreatedAt,
		})
	}

	reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{
		Id: res.Post.ReportingStageID,
	})
	if err != nil {
		return nil, err
	}

	authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
		Id: res.Post.AuthorID,
	})
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
			ReportingStage: &pb.ReportingStagePostResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Name:        reportingStageRes.ReportingStage.Name,
				Description: reportingStageRes.ReportingStage.Description,
			},
			Author: &pb.AuthorPostResponse{
				Id:          authorRes.User.Id,
				Class:       authorRes.User.Class,
				Major:       authorRes.User.Major,
				Phone:       authorRes.User.Phone,
				PhotoSrc:    authorRes.User.PhotoSrc,
				Role:        authorRes.User.Role,
				Name:        authorRes.User.Name,
				Email:       authorRes.User.Email,
				ClassroomID: &authorRes.User.ClassroomID,
			},
			CreatedAt: res.GetPost().CreatedAt,
			UpdatedAt: res.GetPost().UpdatedAt,
		},
		Comments: comments,
	}, nil
}

func (u *postServiceGW) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{Id: req.GetPost().GetReportingStageID()})
	if err != nil {
		return nil, err
	}

	if rpsRes.GetResponse().GetStatusCode() == 404 {
		return &pb.UpdatePostResponse{
			Response: &pb.CommonPostResponse{
				StatusCode: 404,
				Message:    "Reporting stage does not exist",
			},
		}, nil
	}

	exists, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{ClassroomID: req.GetPost().ClassroomID})
	if err != nil {
		return nil, err
	}

	if !exists.GetExists() {
		return &pb.UpdatePostResponse{
			Response: &pb.CommonPostResponse{
				StatusCode: 404,
				Message:    "Classroom does not exist",
			},
		}, nil
	}

	res, err := u.postClient.UpdatePost(ctx, &postSvcV1.UpdatePostRequest{
		Id: req.GetId(),
		Post: &postSvcV1.PostInput{
			Title:            req.GetPost().Title,
			Content:          req.GetPost().Content,
			ClassroomID:      req.GetPost().ClassroomID,
			ReportingStageID: req.GetPost().ReportingStageID,
			AuthorID:         req.GetPost().AuthorID,
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

	var limit int64 = 5
	var page int64 = 1
	titleSearch := ""
	sortColumn := "id"
	sortOrder := "asc"

	if req.GetLimit() > 0 {
		limit = req.GetLimit()
	}

	if req.GetPage() > 0 {
		page = req.GetPage()
	}

	titleSearchTrim := strings.TrimSpace(req.GetTitleSearch())
	if len(titleSearchTrim) > 0 {
		titleSearch = titleSearchTrim
	}

	sortColumnTrim := strings.TrimSpace(req.GetSortColumn())
	if len(sortColumnTrim) > 0 {
		columns := map[string]string{
			"id":                 "id",
			"title":              "title",
			"content":            "content",
			"classroom_id":       "classroom_id",
			"reporting_stage_id": "reporting_stage_id",
			"author_id":          "author_id",
			"created_at":         "created_at",
			"updated_at":         "updated_at",
		}
		if stringInMap(sortColumnTrim, columns) {
			sortColumn = sortColumnTrim
		}
	}

	if req.IsDesc {
		sortOrder = "desc"
	}

	res, err := u.postClient.GetPosts(ctx, &postSvcV1.GetPostsRequest{
		Limit:       limit,
		Page:        page,
		TitleSearch: titleSearch,
		SortColumn:  sortColumn,
		SortOrder:   sortOrder,
	})
	if err != nil {
		return nil, err
	}

	var posts []*pb.PostResponse
	for _, p := range res.GetPosts() {
		reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{
			Id: p.ReportingStageID,
		})
		if err != nil {
			return nil, err
		}

		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: p.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		posts = append(posts, &pb.PostResponse{
			Id:          p.Id,
			Title:       p.Title,
			Content:     p.Content,
			ClassroomID: p.ClassroomID,
			ReportingStage: &pb.ReportingStagePostResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Name:        reportingStageRes.ReportingStage.Name,
				Description: reportingStageRes.ReportingStage.Description,
			},
			Author: &pb.AuthorPostResponse{
				Id:          authorRes.User.Id,
				Class:       authorRes.User.Class,
				Major:       authorRes.User.Major,
				Phone:       authorRes.User.Phone,
				PhotoSrc:    authorRes.User.PhotoSrc,
				Role:        authorRes.User.Role,
				Name:        authorRes.User.Name,
				Email:       authorRes.User.Email,
				ClassroomID: &authorRes.User.ClassroomID,
			},
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
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

func (u *postServiceGW) GetAllPostsOfClassroom(ctx context.Context, req *pb.GetAllPostsOfClassroomRequest) (*pb.GetAllPostsOfClassroomResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	var limit int64 = 5
	var page int64 = 1
	titleSearch := ""
	sortColumn := "id"
	sortOrder := "asc"
	var classroomID int64

	if req.GetLimit() > 0 {
		limit = req.GetLimit()
	}

	if req.GetPage() > 0 {
		page = req.GetPage()
	}

	titleSearchTrim := strings.TrimSpace(req.GetTitleSearch())
	if len(titleSearchTrim) > 0 {
		titleSearch = titleSearchTrim
	}

	sortColumnTrim := strings.TrimSpace(req.GetSortColumn())
	if len(sortColumnTrim) > 0 {
		columns := map[string]string{
			"id":                 "id",
			"title":              "title",
			"content":            "content",
			"classroom_id":       "classroom_id",
			"reporting_stage_id": "reporting_stage_id",
			"author_id":          "author_id",
			"created_at":         "created_at",
			"updated_at":         "updated_at",
		}
		if stringInMap(sortColumnTrim, columns) {
			sortColumn = sortColumnTrim
		}
	}

	if req.IsDesc {
		sortOrder = "desc"
	}

	if req.GetClassroomID() > 0 {
		classroomID = req.GetClassroomID()
	}

	res, err := u.postClient.GetAllPostsOfClassroom(ctx, &postSvcV1.GetAllPostsOfClassroomRequest{
		Limit:       limit,
		Page:        page,
		TitleSearch: titleSearch,
		SortColumn:  sortColumn,
		SortOrder:   sortOrder,
		ClassroomID: classroomID,
	})
	if err != nil {
		return nil, err
	}

	var posts []*pb.PostResponse
	for _, p := range res.GetPosts() {
		reportingStageRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{
			Id: p.ReportingStageID,
		})
		if err != nil {
			return nil, err
		}

		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: p.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		posts = append(posts, &pb.PostResponse{
			Id:          p.Id,
			Title:       p.Title,
			Content:     p.Content,
			ClassroomID: p.ClassroomID,
			ReportingStage: &pb.ReportingStagePostResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Name:        reportingStageRes.ReportingStage.Name,
				Description: reportingStageRes.ReportingStage.Description,
			},
			Author: &pb.AuthorPostResponse{
				Id:          authorRes.User.Id,
				Class:       authorRes.User.Class,
				Major:       authorRes.User.Major,
				Phone:       authorRes.User.Phone,
				PhotoSrc:    authorRes.User.PhotoSrc,
				Role:        authorRes.User.Role,
				Name:        authorRes.User.Name,
				Email:       authorRes.User.Email,
				ClassroomID: &authorRes.User.ClassroomID,
			},
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	return &pb.GetAllPostsOfClassroomResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Posts:      posts,
	}, nil
}
