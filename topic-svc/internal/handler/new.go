package handler

import (
	topicpb "github.com/qthuy2k1/thesis-management-backend/topic-svc/api/goclient/v1"
	repository "github.com/qthuy2k1/thesis-management-backend/topic-svc/internal/repository"
)

type TopicHdl struct {
	topicpb.UnimplementedTopicServiceServer
	Repository repository.ITopicRepo
}

// NewTopicHdl returns the Handler struct that contains the Repository
func NewTopicHdl(svc repository.ITopicRepo) *TopicHdl {
	return &TopicHdl{Repository: svc}
}
