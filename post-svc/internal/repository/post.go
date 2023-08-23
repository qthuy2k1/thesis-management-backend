package repository

import (
	"context"
	"database/sql"
	"time"
	// model "github.com/qthuy2k1/thesis-management-backend/post-svc/internal/model"
)

type PostInputRepo struct {
	ID          int
	Title       string
	Content     string
	ClassroomID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreatePost creates a new post in db given by post model
func (r *PostRepo) CreatePost(ctx context.Context, p PostInputRepo) error {
	// check classroom exists
	isExists, err := r.IsPostExists(ctx, p.Title)
	if err != nil {
		return err
	}

	if isExists {
		return ErrPostExisted
	}

	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "INSERT INTO posts (title, content, classroom_id) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the generated ID
	err = stmt.QueryRowContext(ctx, p.Title, p.Content, p.ClassroomID).Scan(&p.ID)
	if err != nil {
		return err
	}

	return nil
}

type PostOutputRepo struct {
	ID          int
	Title       string
	Content     string
	ClassroomID int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GetPost returns a post in db given by id
func (r *PostRepo) GetPost(ctx context.Context, id int) (PostOutputRepo, error) {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "SELECT id, title, content, classroom_id, created_at, updated_at FROM posts WHERE id=$1")
	if err != nil {
		return PostOutputRepo{}, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the post
	row := stmt.QueryRowContext(ctx, id)
	post := PostOutputRepo{}

	if err = row.Scan(&post.ID, &post.Title, &post.Content, &post.ClassroomID, &post.CreatedAt, &post.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return PostOutputRepo{}, ErrPostNotFound
		}
		return PostOutputRepo{}, err
	}

	return post, nil
}

// CheckPostExists checks whether the specified post exists by title (true == exist)
func (r *PostRepo) IsPostExists(ctx context.Context, title string) (bool, error) {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 FROM posts WHERE title LIKE '%' || $1 || '%')")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the result
	var exists bool
	err = stmt.QueryRowContext(ctx, title).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
