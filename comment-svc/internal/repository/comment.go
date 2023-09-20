package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// QueryRowSQL is a wrapper function that logs the SQL command before executing it.
func QueryRowSQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (*sql.Row, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided arguments
	row := stmt.QueryRowContext(ctx, args...)

	return row, nil
}

// QuerySQL is a wrapper function that logs the SQL command before executing it.
func QuerySQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (*sql.Rows, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided arguments
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.Printf("Error executing SQL command: %s", err.Error())
		return nil, err
	}

	return rows, nil
}

// ExecSQL is a wrapper function that logs the SQL command before executing it.
func ExecSQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (sql.Result, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)
	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL command with the provided arguments
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Printf("Error executing SQL command: %s", err.Error())
		return nil, err
	}

	return result, nil
}

type CommentInputRepo struct {
	UserID     string
	PostID     *int
	ExerciseID *int
	Content    string
}

// CreateClasroom creates a new comment in db given by comment model
func (r *CommentRepo) CreateComment(ctx context.Context, comment CommentInputRepo) error {
	if _, err := ExecSQL(ctx, r.Database, "CreateComment", "INSERT INTO comments (user_id, post_id, exercise_id, content) VALUES ($1, $2, $3, $4) RETURNING id", comment.UserID, comment.PostID, comment.ExerciseID, comment.Content); err != nil {
		return err
	}

	return nil
}

type CommentOutputRepo struct {
	ID         int
	UserID     string
	PostID     *int
	ExerciseID *int
	Content    string
	CreatedAt  time.Time
}

// GetComment returns a comment in db given by id
func (r *CommentRepo) GetComment(ctx context.Context, id int) (CommentOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetComment", "SELECT id, user_id, post_id, exercise_id, content, created_at FROM comments WHERE id=$1", id)
	if err != nil {
		return CommentOutputRepo{}, err
	}
	comment := CommentOutputRepo{}

	if err = row.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.ExerciseID, &comment.Content, &comment.Content); err != nil {
		if err == sql.ErrNoRows {
			return CommentOutputRepo{}, ErrCommentNotFound
		}
		return CommentOutputRepo{}, err
	}

	return comment, nil
}

// UpdateComment updates the specified comment by id
func (r *CommentRepo) UpdateComment(ctx context.Context, id int, comment CommentInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateComment", "UPDATE comments SET user_id=$2, post_id=$3, exercise_id=$4, content=$5 WHERE id=$1", id, comment.UserID, comment.PostID, comment.ExerciseID, comment.Content)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrCommentNotFound
	}

	return nil
}

// DeleteComment deletes a comment in db given by id
func (r *CommentRepo) DeleteComment(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteComment", "DELETE FROM comments WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrCommentNotFound
	}

	return nil
}

// GetComment returns a list of comments in db with filter
func (r *CommentRepo) GetCommentsOfAPost(ctx context.Context, postID int) ([]CommentOutputRepo, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetComments", fmt.Sprintf("SELECT id, user_id, post_id, exercise_id, content, created_at FROM comments WHERE post_id = %d ORDER BY created_at DESC", postID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the comments slice
	var comments []CommentOutputRepo
	for rows.Next() {
		comment := CommentOutputRepo{}
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.ExerciseID,
			&comment.Content,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

// GetComment returns a list of comments in db with filter
func (r *CommentRepo) GetCommentsOfAExercise(ctx context.Context, exerciseID int) ([]CommentOutputRepo, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetComments", fmt.Sprintf("SELECT id, user_id, post_id, exercise_id, content, created_at FROM comments WHERE post_id = %d ORDER BY created_at DESC", exerciseID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the comments slice
	var comments []CommentOutputRepo
	for rows.Next() {
		comment := CommentOutputRepo{}
		err := rows.Scan(
			&comment.ID,
			&comment.UserID,
			&comment.PostID,
			&comment.ExerciseID,
			&comment.Content,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
