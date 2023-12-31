package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type SubmissionInputRepo struct {
	UserID     string
	ExerciseID int
	Status     string
}

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

// CreateSubmission creates a new exercise in db given by exercise model
func (r *SubmissionRepo) CreateSubmission(ctx context.Context, s SubmissionInputRepo) (int64, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateSubmission", "INSERT INTO submissions (user_id, exercise_id, status) VALUES ($1, $2, $3) RETURNING id", s.UserID, s.ExerciseID, s.Status)
	if err != nil {
		return 0, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

type SubmissionOutputRepo struct {
	ID         int
	UserID     string
	ExerciseID int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// GetSubmission returns a submission in db given by id
func (r *SubmissionRepo) GetSubmission(ctx context.Context, id int) (SubmissionOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetSubmission", "SELECT id, user_id, exercise_id, status, created_at, updated_at FROM submissions WHERE id=$1", id)
	if err != nil {
		return SubmissionOutputRepo{}, err
	}

	s := SubmissionOutputRepo{}
	if err = row.Scan(&s.ID, &s.UserID, &s.ExerciseID, &s.Status, &s.CreatedAt, &s.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return SubmissionOutputRepo{}, ErrSubmissionNotFound
		}
		return SubmissionOutputRepo{}, err
	}

	return s, nil
}

// UpdateSubmission updates the specified submission by id
func (r *SubmissionRepo) UpdateSubmission(ctx context.Context, id int, s SubmissionInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateSubmission", "UPDATE submissions SET user_id=$2, exercise_id=$3, status=$4 WHERE id=$1", id, s.UserID, s.ExerciseID, s.Status)
	if err != nil {
		return err
	}
	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrSubmissionNotFound
	}

	return nil
}

// DeleteSubmission deletes a exercise in db given by id
func (r *SubmissionRepo) DeleteSubmission(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteSubmission", "DELETE FROM submissions WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrSubmissionNotFound
	}

	return nil
}

// GetAllSubmissionsOfExercise returns all submissions of the specified exercise given by exercise id
func (r *SubmissionRepo) GetAllSubmissionsOfExercise(ctx context.Context, exerciseID int) ([]SubmissionOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetAllSubmissionsOfExercise", fmt.Sprintf("SELECT id, user_id, exercise_id, status, created_at, updated_at FROM submissions WHERE exercise_id=%d", exerciseID))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the submissions slice
	var submissions []SubmissionOutputRepo

	for rows.Next() {
		s := SubmissionOutputRepo{}
		err := rows.Scan(
			&s.ID,
			&s.UserID,
			&s.ExerciseID,
			&s.Status,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}

		submissions = append(submissions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCountInExercise(ctx, exerciseID)
	if err != nil {
		return nil, 0, err
	}

	return submissions, count, nil
}

func (r *SubmissionRepo) GetSubmissionOfUser(ctx context.Context, userID string, exerciseID int) ([]SubmissionOutputRepo, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetSubmissionOfUser", "SELECT id, user_id, exercise_id, status, created_at, updated_at FROM submissions WHERE user_id = $1 AND exercise_id = $2", userID, exerciseID)
	if err != nil {
		return nil, err
	}

	// Iterate over the result rows and populate the submissions slice
	var submissions []SubmissionOutputRepo

	for rows.Next() {
		s := SubmissionOutputRepo{}
		err := rows.Scan(
			&s.ID,
			&s.UserID,
			&s.ExerciseID,
			&s.Status,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		submissions = append(submissions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return submissions, nil
}

func (r *SubmissionRepo) GetAllSubmissionFromUser(ctx context.Context, userID string) ([]SubmissionOutputRepo, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetSubmissionOfUser", "SELECT id, user_id, exercise_id, status, created_at, updated_at FROM submissions WHERE user_id = $1 ", userID)
	if err != nil {
		return nil, err
	}

	// Iterate over the result rows and populate the submissions slice
	var submissions []SubmissionOutputRepo

	for rows.Next() {
		s := SubmissionOutputRepo{}
		err := rows.Scan(
			&s.ID,
			&s.UserID,
			&s.ExerciseID,
			&s.Status,
			&s.CreatedAt,
			&s.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		submissions = append(submissions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return submissions, nil
}

func (r *SubmissionRepo) getCountInExercise(ctx context.Context, exerciseID int) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", fmt.Sprintf("SELECT COUNT(*) FROM submissions WHERE exercise_id = %d", exerciseID))
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
