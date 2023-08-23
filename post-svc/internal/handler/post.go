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

	if err := h.Service.CreatePost(ctx, p); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &postpb.CreatePostResponse{
		StatusCode: 201,
		Message:    "Created",
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
		Id:        int32(p.ID),
		Title:     p.Title,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
	}

	resp := &postpb.GetPostResponse{
		StatusCode: 200,
		Message:    "OK",
		Post:       &pResp,
	}
	return resp, nil
}

func validateAndConvertPost(pbPost *postpb.PostInput) (service.PostInputSvc, error) {
	if err := pbPost.Validate(); err != nil {
		return service.PostInputSvc{}, err
	}

	return service.PostInputSvc{
		Title: pbPost.Title,
		// Description: pbPost.Description,
		// Status:      pbPost.Status,
	}, nil
}
