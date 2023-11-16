package handler

import (
	"context"
	"log"

	postpb "github.com/qthuy2k1/thesis-management-backend/post-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/post-svc/internal/service"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreatePost retrieves a post request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *PostHdl) CreatePost(ctx context.Context, req *postpb.CreatePostRequest) (*postpb.CreatePostResponse, error) {
	log.Println("calling insert post...")
	p, err := validateAndConvertPost(req.Post)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pRes, err := h.Service.CreatePost(ctx, p)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &postpb.CreatePostResponse{
		Response: &postpb.CommonPostResponse{
			StatusCode: 201,
			Message:    "Created",
		},
		Post: &postpb.PostResponse{
			Id:               int64(pRes.ID),
			Title:            pRes.Title,
			Content:          pRes.Content,
			ClassroomID:      int64(pRes.ClassroomID),
			ReportingStageID: int64(pRes.ReportingStageID),
			AuthorID:         pRes.AuthorID,
			CreatedAt:        timestamppb.New(pRes.CreatedAt),
			UpdatedAt:        timestamppb.New(pRes.UpdatedAt),
		},
	}

	return resp, nil
}

// GetPost returns a post in db given by id
func (h *PostHdl) GetPost(ctx context.Context, req *postpb.GetPostRequest) (*postpb.GetPostResponse, error) {
	log.Println("calling get post...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	p, err := h.Service.GetPost(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	pResp := postpb.PostResponse{
		Id:               int64(p.ID),
		Title:            p.Title,
		Content:          p.Content,
		ClassroomID:      int64(p.ClassroomID),
		ReportingStageID: int64(p.ReportingStageID),
		AuthorID:         p.AuthorID,
		CreatedAt:        timestamppb.New(p.CreatedAt),
		UpdatedAt:        timestamppb.New(p.UpdatedAt),
	}

	resp := &postpb.GetPostResponse{
		Response: &postpb.CommonPostResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Post: &pResp,
	}
	return resp, nil
}

func (c *PostHdl) UpdatePost(ctx context.Context, req *postpb.UpdatePostRequest) (*postpb.UpdatePostResponse, error) {
	log.Println("calling update post...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	p, err := validateAndConvertPost(req.Post)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Service.UpdatePost(ctx, int(req.GetId()), service.PostInputSvc{
		Title:            p.Title,
		Content:          p.Content,
		ClassroomID:      p.ClassroomID,
		ReportingStageID: p.ReportingStageID,
		AuthorID:         p.AuthorID,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &postpb.UpdatePostResponse{
		Response: &postpb.CommonPostResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *PostHdl) DeletePost(ctx context.Context, req *postpb.DeletePostRequest) (*postpb.DeletePostResponse, error) {
	log.Println("calling delete post...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.DeletePost(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &postpb.DeletePostResponse{
		Response: &postpb.CommonPostResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *PostHdl) GetPosts(ctx context.Context, req *postpb.GetPostsRequest) (*postpb.GetPostsResponse, error) {
	log.Println("calling get all posts...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	filter := service.PostFilterSvc{
		Limit:       int(req.GetLimit()),
		Page:        int(req.GetPage()),
		TitleSearch: req.GetTitleSearch(),
		SortColumn:  req.GetSortColumn(),
		SortOrder:   req.GetSortOrder(),
	}

	ps, count, err := h.Service.GetPosts(ctx, filter)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*postpb.PostResponse
	for _, p := range ps {
		psResp = append(psResp, &postpb.PostResponse{
			Id:               int64(p.ID),
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      int64(p.ClassroomID),
			ReportingStageID: int64(p.ReportingStageID),
			AuthorID:         p.AuthorID,
			CreatedAt:        timestamppb.New(p.CreatedAt),
			UpdatedAt:        timestamppb.New(p.UpdatedAt),
		})
	}

	return &postpb.GetPostsResponse{
		Response: &postpb.CommonPostResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Posts:      psResp,
		TotalCount: int64(count),
	}, nil
}

func (h *PostHdl) GetAllPostsOfClassroom(ctx context.Context, req *postpb.GetAllPostsOfClassroomRequest) (*postpb.GetAllPostsOfClassroomResponse, error) {
	log.Println("calling get all posts of a classroom...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	filter := service.PostFilterSvc{
		Limit:       int(req.GetLimit()),
		Page:        int(req.GetPage()),
		TitleSearch: req.GetTitleSearch(),
		SortColumn:  req.GetSortColumn(),
		SortOrder:   req.GetSortOrder(),
	}

	ps, count, err := h.Service.GetAllPostsOfClassroom(ctx, filter, int(req.GetClassroomID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*postpb.PostResponse
	for _, p := range ps {
		psResp = append(psResp, &postpb.PostResponse{
			Id:               int64(p.ID),
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      int64(p.ClassroomID),
			ReportingStageID: int64(p.ReportingStageID),
			AuthorID:         p.AuthorID,
			CreatedAt:        timestamppb.New(p.CreatedAt),
			UpdatedAt:        timestamppb.New(p.UpdatedAt),
		})
	}

	return &postpb.GetAllPostsOfClassroomResponse{
		Response: &postpb.CommonPostResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Posts:      psResp,
		TotalCount: int64(count),
	}, nil
}

func (h *PostHdl) GetAllPostsInReportingStage(ctx context.Context, req *postpb.GetAllPostsInReportingStageRequest) (*postpb.GetAllPostsInReportingStageResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ps, count, err := h.Service.GetAllPostsInReportingStage(ctx, int(req.GetReportingStageID()), int(req.GetClassroomID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var psResp []*postpb.PostResponse
	for _, p := range ps {
		psResp = append(psResp, &postpb.PostResponse{
			Id:               int64(p.ID),
			Title:            p.Title,
			Content:          p.Content,
			ClassroomID:      int64(p.ClassroomID),
			ReportingStageID: int64(p.ReportingStageID),
			AuthorID:         p.AuthorID,
			CreatedAt:        timestamppb.New(p.CreatedAt),
			UpdatedAt:        timestamppb.New(p.UpdatedAt),
		})
	}

	return &postpb.GetAllPostsInReportingStageResponse{
		Response: &postpb.CommonPostResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Posts:      psResp,
		TotalCount: int64(count),
	}, nil
}

func validateAndConvertPost(pbPost *postpb.PostInput) (service.PostInputSvc, error) {
	if err := pbPost.Validate(); err != nil {
		return service.PostInputSvc{}, err
	}

	return service.PostInputSvc{
		Title:            pbPost.Title,
		Content:          pbPost.Content,
		ClassroomID:      int(pbPost.ClassroomID),
		ReportingStageID: int(pbPost.ReportingStageID),
		AuthorID:         pbPost.AuthorID,
	}, nil
}
