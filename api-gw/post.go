package main

import (
	"context"
	"strings"

	pb "github.com/qthuy2k1/thesis-management-backend/api-gw/api/goclient/v1"
	attachmentSvcV1 "github.com/qthuy2k1/thesis-management-backend/attachment-svc/api/goclient/v1"
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
	attachmentClient     attachmentSvcV1.AttachmentServiceClient
}

func NewPostsService(postClient postSvcV1.PostServiceClient, classroomClient classroomSvcV1.ClassroomServiceClient, reportingStageClient rpsSvcV1.ReportingStageServiceClient, commentCLient commentSvcV1.CommentServiceClient, userClient userSvcV1.UserServiceClient, attachmentClient attachmentSvcV1.AttachmentServiceClient) *postServiceGW {
	return &postServiceGW{
		postClient:           postClient,
		classroomClient:      classroomClient,
		reportingStageClient: reportingStageClient,
		commentClient:        commentCLient,
		userClient:           userClient,
		attachmentClient:     attachmentClient,
	}
}

func (u *postServiceGW) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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

	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{Id: req.GetPost().GetCategoryID()})
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
			Content:          req.GetPost().Description,
			ClassroomID:      req.GetPost().ClassroomID,
			ReportingStageID: req.GetPost().CategoryID,
			AuthorID:         req.GetPost().AuthorID,
		},
	})
	if err != nil {
		return nil, err
	}

	var attCreated []int64
	if req.Post.Attachments != nil && len(req.Post.Attachments) > 0 {
		for _, att := range req.Post.Attachments {
			postID := res.Post.Id

			if _, err = u.attachmentClient.CreateAttachment(ctx, &attachmentSvcV1.CreateAttachmentRequest{
				Attachment: &attachmentSvcV1.AttachmentInput{
					FileURL:  att.FileURL,
					Status:   att.FileURL,
					PostID:   &postID,
					AuthorID: res.Post.AuthorID,
				},
			}); err != nil {
				if len(attCreated) > 0 {
					for _, aErr := range attCreated {
						if _, err := u.attachmentClient.DeleteAttachment(ctx, &attachmentSvcV1.DeleteAttachmentRequest{
							Id: aErr,
						}); err != nil {
							return nil, err
						}
					}
				}
				return nil, err
			}
		}
	}

	return &pb.CreatePostResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
	}, nil
}

func (u *postServiceGW) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

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
				Id:       userRes.User.Id,
				Class:    userRes.User.Class,
				Major:    userRes.User.Major,
				Phone:    userRes.User.Phone,
				PhotoSrc: userRes.User.PhotoSrc,
				Role:     userRes.User.Role,
				Name:     userRes.User.Name,
				Email:    userRes.User.Email,
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

	attachment, err := u.attachmentClient.GetAttachmentsOfPost(ctx, &attachmentSvcV1.GetAttachmentsOfPostRequest{
		PostID: res.Post.Id,
	})
	if err != nil {
		return nil, err
	}

	var attachments []*pb.AttachmentPostResponse
	for _, a := range attachment.Attachments {
		author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: a.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, &pb.AttachmentPostResponse{
			Id:      a.Id,
			FileURL: a.FileURL,
			Status:  a.Status,
			Author: &pb.AuthorPostResponse{
				Id:       author.User.Id,
				Class:    author.User.Class,
				Major:    author.User.Major,
				Phone:    author.User.Phone,
				PhotoSrc: author.User.PhotoSrc,
				Role:     author.User.Role,
				Name:     author.User.Name,
				Email:    author.User.Email,
			},
			CreatedAt: a.CreatedAt,
		})
	}

	return &pb.GetPostResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		Post: &pb.PostResponse{
			Id:          res.GetPost().Id,
			Title:       res.GetPost().Title,
			Description: res.GetPost().Content,
			ClassroomID: res.GetPost().ClassroomID,
			Category: &pb.ReportingStagePostResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Label:       reportingStageRes.ReportingStage.Label,
				Description: reportingStageRes.ReportingStage.Description,
				Value:       reportingStageRes.ReportingStage.Value,
			},
			Author: &pb.AuthorPostResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt:   res.GetPost().CreatedAt,
			UpdatedAt:   res.GetPost().UpdatedAt,
			Attachments: attachments,
		},
		Comments: comments,
	}, nil
}

func (u *postServiceGW) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{Id: req.GetPost().GetCategoryID()})
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
			Content:          req.GetPost().Description,
			ClassroomID:      req.GetPost().ClassroomID,
			ReportingStageID: req.GetPost().CategoryID,
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
	if err := req.Validate(); err != nil {
		return nil, err
	}

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

		attachment, err := u.attachmentClient.GetAttachmentsOfPost(ctx, &attachmentSvcV1.GetAttachmentsOfPostRequest{
			PostID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentPostResponse
		for _, a := range attachment.Attachments {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentPostResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorPostResponse{
					Id:       author.User.Id,
					Class:    author.User.Class,
					Major:    author.User.Major,
					Phone:    author.User.Phone,
					PhotoSrc: author.User.PhotoSrc,
					Role:     author.User.Role,
					Name:     author.User.Name,
					Email:    author.User.Email,
				},
				CreatedAt: a.CreatedAt,
			})
		}

		posts = append(posts, &pb.PostResponse{
			Id:          p.Id,
			Title:       p.Title,
			Description: p.Content,
			ClassroomID: p.ClassroomID,
			Category: &pb.ReportingStagePostResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Label:       reportingStageRes.ReportingStage.Label,
				Description: reportingStageRes.ReportingStage.Description,
				Value:       reportingStageRes.ReportingStage.Value,
			},
			Author: &pb.AuthorPostResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Attachments: attachments,
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

		attachment, err := u.attachmentClient.GetAttachmentsOfPost(ctx, &attachmentSvcV1.GetAttachmentsOfPostRequest{
			PostID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentPostResponse
		for _, a := range attachment.Attachments {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentPostResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorPostResponse{
					Id:       author.User.Id,
					Class:    author.User.Class,
					Major:    author.User.Major,
					Phone:    author.User.Phone,
					PhotoSrc: author.User.PhotoSrc,
					Role:     author.User.Role,
					Name:     author.User.Name,
					Email:    author.User.Email,
				},
				CreatedAt: a.CreatedAt,
			})
		}

		posts = append(posts, &pb.PostResponse{
			Id:          p.Id,
			Title:       p.Title,
			Description: p.Content,
			ClassroomID: p.ClassroomID,
			Category: &pb.ReportingStagePostResponse{
				Id:          reportingStageRes.ReportingStage.Id,
				Label:       reportingStageRes.ReportingStage.Label,
				Description: reportingStageRes.ReportingStage.Description,
				Value:       reportingStageRes.ReportingStage.Value,
			},
			Author: &pb.AuthorPostResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Attachments: attachments,
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

func (u *postServiceGW) GetAllPostsInReportingStage(ctx context.Context, req *pb.GetAllPostsInReportingStageRequest) (*pb.GetAllPostsInReportingStageResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	classRes, err := u.classroomClient.CheckClassroomExists(ctx, &classroomSvcV1.CheckClassroomExistsRequest{
		ClassroomID: req.GetClassroomID(),
	})
	if err != nil {
		return nil, err
	}

	if !classRes.GetExists() {
		return &pb.GetAllPostsInReportingStageResponse{
			Response: &pb.CommonPostResponse{
				StatusCode: 404,
				Message:    "classroom does not exist",
			},
		}, nil
	}

	rpsRes, err := u.reportingStageClient.GetReportingStage(ctx, &rpsSvcV1.GetReportingStageRequest{
		Id: req.GetCategoryID(),
	})
	if err != nil {
		return nil, err
	}

	if rpsRes.GetResponse().StatusCode == 404 {
		return &pb.GetAllPostsInReportingStageResponse{
			Response: &pb.CommonPostResponse{
				StatusCode: rpsRes.Response.StatusCode,
				Message:    rpsRes.Response.Message,
			},
		}, nil
	}

	res, err := u.postClient.GetAllPostsInReportingStage(ctx, &postSvcV1.GetAllPostsInReportingStageRequest{
		ClassroomID:      req.GetClassroomID(),
		ReportingStageID: req.GetCategoryID(),
	})
	if err != nil {
		return nil, err
	}

	var posts []*pb.PostResponse
	for _, p := range res.GetPosts() {
		authorRes, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
			Id: p.AuthorID,
		})
		if err != nil {
			return nil, err
		}

		attachment, err := u.attachmentClient.GetAttachmentsOfPost(ctx, &attachmentSvcV1.GetAttachmentsOfPostRequest{
			PostID: p.Id,
		})
		if err != nil {
			return nil, err
		}

		var attachments []*pb.AttachmentPostResponse
		for _, a := range attachment.Attachments {
			author, err := u.userClient.GetUser(ctx, &userSvcV1.GetUserRequest{
				Id: a.AuthorID,
			})
			if err != nil {
				return nil, err
			}

			attachments = append(attachments, &pb.AttachmentPostResponse{
				Id:      a.Id,
				FileURL: a.FileURL,
				Status:  a.Status,
				Author: &pb.AuthorPostResponse{
					Id:       author.User.Id,
					Class:    author.User.Class,
					Major:    author.User.Major,
					Phone:    author.User.Phone,
					PhotoSrc: author.User.PhotoSrc,
					Role:     author.User.Role,
					Name:     author.User.Name,
					Email:    author.User.Email,
				},
				CreatedAt: a.CreatedAt,
			})
		}

		posts = append(posts, &pb.PostResponse{
			Id:          p.Id,
			Title:       p.Title,
			Description: p.Content,
			ClassroomID: p.ClassroomID,
			Category: &pb.ReportingStagePostResponse{
				Id:          rpsRes.ReportingStage.Id,
				Label:       rpsRes.ReportingStage.Label,
				Description: rpsRes.ReportingStage.Description,
				Value:       rpsRes.ReportingStage.Value,
			},
			Author: &pb.AuthorPostResponse{
				Id:       authorRes.User.Id,
				Class:    authorRes.User.Class,
				Major:    authorRes.User.Major,
				Phone:    authorRes.User.Phone,
				PhotoSrc: authorRes.User.PhotoSrc,
				Role:     authorRes.User.Role,
				Name:     authorRes.User.Name,
				Email:    authorRes.User.Email,
			},
			CreatedAt:   p.CreatedAt,
			UpdatedAt:   p.UpdatedAt,
			Attachments: attachments,
		})
	}

	return &pb.GetAllPostsInReportingStageResponse{
		Response: &pb.CommonPostResponse{
			StatusCode: res.GetResponse().StatusCode,
			Message:    res.GetResponse().Message,
		},
		TotalCount: res.GetTotalCount(),
		Posts:      posts,
	}, nil
}
