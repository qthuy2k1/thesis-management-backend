package handler

import (
	"context"
	"log"

	topicpb "github.com/qthuy2k1/thesis-management-backend/topic-svc/api/goclient/v1"
	repository "github.com/qthuy2k1/thesis-management-backend/topic-svc/internal/repository"
	"google.golang.org/grpc/status"
)

// CreateTopic retrieves a topic request from gRPC-gateway and calls to the Repository layer, then returns the response and status code.
func (h *TopicHdl) CreateTopic(ctx context.Context, req *topicpb.CreateTopicRequest) (*topicpb.CreateTopicResponse, error) {
	log.Println("calling insert topic...")
	topic, err := validateAndConvertTopic(req.Topic)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Repository.CreateTopic(ctx, topic); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	resp := &topicpb.CreateTopicResponse{
		Response: &topicpb.CommonTopicResponse{
			StatusCode: 200,
			Message:    "OK",
		},
	}

	return resp, nil
}

// GetTopic returns a topic in db given by id
func (h *TopicHdl) GetTopic(ctx context.Context, req *topicpb.GetTopicRequest) (*topicpb.GetTopicResponse, error) {
	log.Println("calling get topic...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	topic, err := h.Repository.GetTopic(ctx, int(req.GetId()))
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	topicResp := topicpb.TopicResponse{
		Id:            int64(topic.ID),
		Title:         topic.Title,
		TypeTopic:     topic.Title,
		MemberQuanity: int64(topic.MemberQuantity),
		StudentId:     topic.StudentID,
		MemberEmail:   topic.MemberEmail,
		Description:   topic.Description,
	}

	resp := &topicpb.GetTopicResponse{
		Response: &topicpb.CommonTopicResponse{
			StatusCode: 200,
			Message:    "OK",
		},
		Topic: &topicResp,
	}
	return resp, nil
}

func (c *TopicHdl) UpdateTopic(ctx context.Context, req *topicpb.UpdateTopicRequest) (*topicpb.UpdateTopicResponse, error) {
	log.Println("calling update topic...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	topic, err := validateAndConvertTopic(req.Topic)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := c.Repository.UpdateTopic(ctx, int(req.GetId()), repository.TopicInputRepo{
		Title:          topic.Title,
		TypeTopic:      topic.TypeTopic,
		MemberQuantity: topic.MemberQuantity,
		StudentID:      topic.StudentID,
		MemberEmail:    topic.MemberEmail,
		Description:    topic.Description,
	}); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &topicpb.UpdateTopicResponse{
		Response: &topicpb.CommonTopicResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func (h *TopicHdl) DeleteTopic(ctx context.Context, req *topicpb.DeleteTopicRequest) (*topicpb.DeleteTopicResponse, error) {
	log.Println("calling delete topic...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	if err := h.Repository.DeleteTopic(ctx, int(req.GetId())); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	return &topicpb.DeleteTopicResponse{
		Response: &topicpb.CommonTopicResponse{
			StatusCode: 200,
			Message:    "Success",
		},
	}, nil
}

func validateAndConvertTopic(pbTopic *topicpb.TopicInput) (repository.TopicInputRepo, error) {
	if err := pbTopic.Validate(); err != nil {
		return repository.TopicInputRepo{}, err
	}

	return repository.TopicInputRepo{
		Title:          pbTopic.Title,
		TypeTopic:      pbTopic.TypeTopic,
		MemberQuantity: int(pbTopic.MemberQuanity),
		StudentID:      pbTopic.StudentId,
		MemberEmail:    pbTopic.MemberEmail,
		Description:    pbTopic.Description,
	}, nil
}
