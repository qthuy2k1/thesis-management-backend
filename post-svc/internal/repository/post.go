package repository

import (
	"context"
	"database/sql"
	"log"
	"time"
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
	// check post exists
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

	if _, err := stmt.ExecContext(ctx, p.Title, p.Content, p.ClassroomID); err != nil {
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

// UpdatePost updates the specified post by id
func (r *PostRepo) UpdatePost(ctx context.Context, id int, post PostInputRepo) error {
	log.Println(id, post)
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "UPDATE posts SET title=$2, content=$3, classroom_id=$4, updated_at=$5 WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the ID of the updated post
	result, err := stmt.ExecContext(ctx, id, post.Title, post.Content, post.ClassroomID, time.Now())
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrPostNotFound
	}

	return nil
}

// DeletePost deletes a post in db given by id
func (r *PostRepo) DeletePost(ctx context.Context, id int) error {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "DELETE FROM posts WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the deleted post's details
	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrPostNotFound
	}

	return nil
}
