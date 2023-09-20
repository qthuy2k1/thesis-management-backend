package handler

import (
	"context"
	"log"

	commentpb "github.com/qthuy2k1/thesis-management-backend/comment-svc/api/goclient/v1"
	service "github.com/qthuy2k1/thesis-management-backend/comment-svc/internal/service"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// CreateComment retrieves a comment request from gRPC-gateway and calls to the Service layer, then returns the response and status code.
func (h *CommentHdl) CreateComment(ctx context.Context, req *commentpb.CreateCommentRequest) (*commentpb.CreateCommentResponse, error) {
	log.Println("calling insert comment...")
	comment, err := validateAndConvertComment(req.Comment)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Service.CreateComment(ctx, comment); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &commentpb.CreateCommentResponse{
		Response: &commentpb.CommonCommentResponse{
			StatusCode: 200,
			Message:    "OK",
		},
	}

	return resp, nil
}

// GetComment returns a comment in db given by id
func (h *CommentHdl) GetComment(ctx context.Context, req *commentpb.GetCommentRequest) (*commentpb.GetCommentResponse, error) {
	log.Println("calling get comment...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	comment, err := h.Service.GetComment(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	commentResp := commentpb.CommentResponse{
		Id: int32(comment.ID),

		CreatedAt: timestamppb.New(comment.CreatedAt),
	}

	resp := &commentpb.GetCommentResponse{
		Response: &commentpb.CommonCommentResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Comment: &commentResp,
	}
	return resp, nil
}

func (h *CommentHdl) GetCommentsOfAPost(ctx context.Context, req *commentpb.GetCommentsOfAPostRequest) (*commentpb.GetCommentsResponse, error) {
	log.Println("calling get all comments of a post...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	comments, err := h.Service.GetCommentsOfAPost(ctx, int(req.GetPostID()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var commentsResp []*commentpb.CommentResponse
	for _, c := range comments {
		postID := int32(*c.PostID)
		exerciseID := int32(*c.ExerciseID)

		commentsResp = append(commentsResp, &commentpb.CommentResponse{
			Id:         int32(c.ID),
			UserID:     c.UserID,
			PostID:     &postID,
			ExerciseID: &exerciseID,
			Content:    c.Content,
			CreatedAt:  timestamppb.New(c.CreatedAt),
		})
	}

	return &commentpb.GetCommentsResponse{
		Response: &commentpb.CommonCommentResponse{
			StatusCode: 200,
			Message:    "Created",
		},
		Comments: commentsResp,
	}, nil
}

func validateAndConvertComment(pbComment *commentpb.CommentInput) (service.CommentInputSvc, error) {
	if err := pbComment.Validate(); err != nil {
		return service.CommentInputSvc{}, err
	}

	var postID int
	if pbComment.PostID != nil && *pbComment.PostID != 0 {
		postID = int(*pbComment.PostID)
	}

	var exerciseID int
	if pbComment.ExerciseID != nil && *pbComment.ExerciseID != 0 {
		exerciseID = int(*pbComment.ExerciseID)
	}

	return service.CommentInputSvc{
		UserID:     pbComment.UserID,
		PostID:     &postID,
		ExerciseID: &exerciseID,
		Content:    pbComment.Content,
	}, nil
}
