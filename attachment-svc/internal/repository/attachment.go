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

type AttachmentInputRepo struct {
	FileURL      string
	Status       string
	SubmissionID int
	ExerciseID   int
	AuthorID     string
}

// CreateClasroom creates a new attachment in db given by attachment model
func (r *AttachmentRepo) CreateAttachment(ctx context.Context, att AttachmentInputRepo) error {
	if _, err := ExecSQL(ctx, r.Database, "CreateAttachment", "INSERT INTO attachments (file_url, status, submission_id, exercise_id, author_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", att.FileURL, att.Status, att.SubmissionID, att.ExerciseID, att.AuthorID); err != nil {
		return err
	}

	return nil
}

type AttachmentOutputRepo struct {
	ID           int
	FileURL      string
	Status       string
	SubmissionID int
	ExerciseID   int
	AuthorID     string
	CreatedAt    time.Time
}

// GetAttachment returns a attachment in db given by id
func (r *AttachmentRepo) GetAttachment(ctx context.Context, id int) (AttachmentOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetAttachment", "SELECT id, file_url, status, submission_id, exercise_id, author_id, created_at FROM attachments WHERE id=$1", id)
	if err != nil {
		return AttachmentOutputRepo{}, err
	}
	attachment := AttachmentOutputRepo{}

	if err = row.Scan(&attachment.ID, attachment.FileURL, &attachment.Status, &attachment.SubmissionID, &attachment.ExerciseID, &attachment.AuthorID, &attachment.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return AttachmentOutputRepo{}, ErrAttachmentNotFound
		}
		return AttachmentOutputRepo{}, err
	}

	return attachment, nil
}

// UpdateAttachment updates the specified attachment by id
func (r *AttachmentRepo) UpdateAttachment(ctx context.Context, id int, attachment AttachmentInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateAttachment", "UPDATE attachments SET file_url=$2, status=$3, submission_id=$4, exercise_id=$5, author_id=$6 WHERE id=$1", id, attachment.FileURL, attachment.Status, attachment.SubmissionID, attachment.ExerciseID, attachment.AuthorID)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrAttachmentNotFound
	}

	return nil
}

// DeleteAttachment deletes a attachment in db given by id
func (r *AttachmentRepo) DeleteAttachment(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteAttachment", "DELETE FROM attachments WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrAttachmentNotFound
	}

	return nil
}

// GetAttachment returns a list of attachments of an exercise in db
func (r *AttachmentRepo) GetAttachmentsOfExercise(ctx context.Context, exerciseID int) ([]AttachmentOutputRepo, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetAttachments", fmt.Sprintf("SELECT id, file_url, status, submission_id, exercise_id, author_id, created_at FROM attachments WHERE exercise_id=%d", exerciseID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the attachments slice
	var attachments []AttachmentOutputRepo
	for rows.Next() {
		attachment := AttachmentOutputRepo{}
		err := rows.Scan(
			&attachment.ID,
			&attachment.FileURL,
			&attachment.Status,
			&attachment.SubmissionID,
			&attachment.ExerciseID,
			&attachment.AuthorID,
			&attachment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		attachments = append(attachments, attachment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return attachments, nil
}
