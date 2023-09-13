package repository

import (
	"context"
	"database/sql"
)

type IPostRepo interface {
	// CreatePost creates a new post in db given by post model
	CreatePost(ctx context.Context, clr PostInputRepo) error
	// GetPost returns a post in db given by id
	GetPost(ctx context.Context, id int) (PostOutputRepo, error)
	// CheckPostExists checks whether the specified post exists by name
	IsPostExists(ctx context.Context, title string, classroomID int) (bool, error)
	// UpdatePost updates the specified post by id
	UpdatePost(ctx context.Context, id int, post PostInputRepo) error
	// DeletePost deletes a post in db given by id
	DeletePost(ctx context.Context, id int) error
	// GetPosts returns a list of posts in db with filter
	GetPosts(ctx context.Context, filter PostFilterRepo) ([]PostOutputRepo, int, error)
	// GetAllPostsOfClassroom returns all posts of the specified classroom given by classroom id
	GetAllPostsOfClassroom(ctx context.Context, filter PostFilterRepo, classromID int) ([]PostOutputRepo, int, error)
}

type PostRepo struct {
	Database *sql.DB
}

func NewPostRepo(db *sql.DB) IPostRepo {
	return &PostRepo{Database: db}
}
