package repository

import (
	"context"
	"database/sql"
)

type ITopicRepo interface {
	// CreateTopic creates a new topic in db given by topic model
	CreateTopic(ctx context.Context, clr TopicInputRepo) error
	// GetTopic returns a topic in db given by id
	GetTopic(ctx context.Context, id int) (TopicOutputRepo, error)
	// GetTopic returns a topic in db given by id
	GetTopicFromUser(ctx context.Context, id string) (TopicOutputRepo, error)
	// UpdateTopic updates the specified topic by id
	UpdateTopic(ctx context.Context, id int, topic TopicInputRepo) error
	// DeleteTopic deletes a topic in db given by id
	DeleteTopic(ctx context.Context, id int) error

	GetTopics(ctx context.Context) ([]TopicOutputRepo, error)

	GetAllTopicOfListUser(ctx context.Context, userListID []string) ([]TopicOutputRepo, error)
}

type TopicRepo struct {
	Database *sql.DB
}

func NewTopicRepo(db *sql.DB) ITopicRepo {
	return &TopicRepo{Database: db}
}
