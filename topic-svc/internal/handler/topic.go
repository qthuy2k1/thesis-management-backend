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
		Id:             int64(topic.ID),
		Title:          topic.Title,
		TypeTopic:      topic.Title,
		MemberQuantity: int64(topic.MemberQuantity),
		StudentID:      topic.StudentID,
		MemberEmail:    topic.MemberEmail,
		Description:    topic.Description,
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

// GetTopic returns a topic in db given by id
func (h *TopicHdl) GetTopicFromUser(ctx context.Context, req *topicpb.GetTopicFromUserRequest) (*topicpb.GetTopicFromUserResponse, error) {
	log.Println("calling get topic...")
	if err := req.Validate(); err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}
	topic, err := h.Repository.GetTopicFromUser(ctx, req.GetUserID())
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	topicResp := topicpb.TopicResponse{
		Id:             int64(topic.ID),
		Title:          topic.Title,
		TypeTopic:      topic.Title,
		MemberQuantity: int64(topic.MemberQuantity),
		StudentID:      topic.StudentID,
		MemberEmail:    topic.MemberEmail,
		Description:    topic.Description,
	}

	resp := &topicpb.GetTopicFromUserResponse{
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

func (h *TopicHdl) GetTopics(ctx context.Context, req *topicpb.GetTopicsRequest) (*topicpb.GetTopicsResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ts, err := h.Repository.GetTopics(ctx)
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var tsResp []*topicpb.TopicResponse
	for _, t := range ts {
		tsResp = append(tsResp, &topicpb.TopicResponse{
			Id:             int64(t.ID),
			Title:          t.Title,
			TypeTopic:      t.Title,
			MemberQuantity: int64(t.MemberQuantity),
			StudentID:      t.StudentID,
			MemberEmail:    t.MemberEmail,
			Description:    t.Description,
		})
	}

	return &topicpb.GetTopicsResponse{
		Response: &topicpb.CommonTopicResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Topic: tsResp,
	}, nil
}

func (h *TopicHdl) GetAllTopicsOfListUser(ctx context.Context, req *topicpb.GetAllTopicsOfListUserRequest) (*topicpb.GetAllTopicsOfListUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ts, err := h.Repository.GetAllTopicOfListUser(ctx, req.GetUserID())
	if err != nil {
		code, err := convertCtrlError(err)
		return nil, status.Errorf(code, "err: %v", err)
	}

	var tsResp []*topicpb.TopicResponse
	for _, t := range ts {
		tsResp = append(tsResp, &topicpb.TopicResponse{
			Id:             int64(t.ID),
			Title:          t.Title,
			TypeTopic:      t.Title,
			MemberQuantity: int64(t.MemberQuantity),
			StudentID:      t.StudentID,
			MemberEmail:    t.MemberEmail,
			Description:    t.Description,
		})
	}

	return &topicpb.GetAllTopicsOfListUserResponse{
		Response: &topicpb.CommonTopicResponse{
			StatusCode: 200,
			Message:    "Success",
		},
		Topic: tsResp,
	}, nil
}

func validateAndConvertTopic(pbTopic *topicpb.TopicInput) (repository.TopicInputRepo, error) {
	if err := pbTopic.Validate(); err != nil {
		return repository.TopicInputRepo{}, err
	}

	return repository.TopicInputRepo{
		Title:          pbTopic.Title,
		TypeTopic:      pbTopic.TypeTopic,
		MemberQuantity: int(pbTopic.MemberQuantity),
		StudentID:      pbTopic.StudentID,
		MemberEmail:    pbTopic.MemberEmail,
		Description:    pbTopic.Description,
	}, nil
}
