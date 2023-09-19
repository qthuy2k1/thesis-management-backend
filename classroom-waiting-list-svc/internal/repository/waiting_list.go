package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
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

type WaitingListInputRepo struct {
	ClassroomID int
	UserID      string
	CreatedAt   time.Time
}

// CreateWaitingList creates a new waiting_list in db given by waiting_list model
func (r *WaitingListRepo) CreateWaitingList(ctx context.Context, wt WaitingListInputRepo) error {
	// check waiting_list exists
	isExists, err := r.IsWaitingListExists(ctx, wt.ClassroomID, wt.UserID)
	if err != nil {
		return err
	}

	if isExists {
		return ErrWaitingListExisted
	}

	if _, err := ExecSQL(ctx, r.Database, "CreateWaitingList", "INSERT INTO waiting_lists (classroom_id, user_id) VALUES ($1, $2) RETURNING id", wt.ClassroomID, wt.UserID); err != nil {
		return err
	}

	return nil
}

type WaitingListOutputRepo struct {
	ID          int
	ClassroomID int
	UserID      string
	CreatedAt   time.Time
}

// GetWaitingList returns a waiting_list in db given by id
func (r *WaitingListRepo) GetWaitingList(ctx context.Context, id int) (WaitingListOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetWaitingList", "SELECT id, classroom_id, user_id, created_at FROM waiting_lists WHERE id=$1", id)
	if err != nil {
		return WaitingListOutputRepo{}, err
	}

	waiting_list := WaitingListOutputRepo{}
	if err = row.Scan(&waiting_list.ID, &waiting_list.ClassroomID, &waiting_list.UserID, &waiting_list.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return WaitingListOutputRepo{}, ErrWaitingListNotFound
		}
		return WaitingListOutputRepo{}, err
	}

	return waiting_list, nil
}

// CheckWaitingListExists checks whether the specified waiting_list exists (true == exist)
func (r *WaitingListRepo) IsWaitingListExists(ctx context.Context, classroomID int, userID string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsWaitingListExists", "SELECT EXISTS(SELECT 1 FROM waiting_lists WHERE classroom_id = $1 AND user_id = $2)", classroomID, userID)
	if err != nil {
		return false, err
	}
	if err = row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

// UpdateWaitingList updates the specified waiting_list by id
func (r *WaitingListRepo) UpdateWaitingList(ctx context.Context, id int, waiting_list WaitingListInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateWaitingList", "UPDATE waiting_lists SET classroom_id=$2, user_id=$3 WHERE id=$1", id, waiting_list.ClassroomID, waiting_list.UserID)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrWaitingListNotFound
	}

	return nil
}

// DeleteWaitingList deletes a waiting_list in db given by id
func (r *WaitingListRepo) DeleteWaitingList(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteWaitingList", "DELETE FROM waiting_lists WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrWaitingListNotFound
	}

	return nil
}

// GetWaitingList returns a list of waiting_lists in db with filter
func (r *WaitingListRepo) GetWaitingListsOfClassroom(ctx context.Context, classroomID int) ([]WaitingListOutputRepo, error) {
	query := []string{fmt.Sprintf("SELECT * FROM waiting_lists WHERE classroom_id = %d", classroomID)}

	rows, err := QuerySQL(ctx, r.Database, "GetWaitingLists", strings.Join(query, " "))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the waiting_lists slice
	var waiting_lists []WaitingListOutputRepo
	for rows.Next() {
		waiting_list := WaitingListOutputRepo{}
		err := rows.Scan(
			&waiting_list.ID,
			&waiting_list.ClassroomID,
			&waiting_list.UserID,
			&waiting_list.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		waiting_lists = append(waiting_lists, waiting_list)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return waiting_lists, nil
}

// CheckUserInWaitingListOfClassroom returns a boolean indicating whether user is in waiting list
func (r *WaitingListRepo) CheckUserInWaitingListOfClassroom(ctx context.Context, userID string) (bool, int, error) {
	query := []string{fmt.Sprintf("SELECT id, user_id, classroom_id FROM waiting_lists WHERE user_id = '%s' LIMIT 1", userID)}

	row, err := QueryRowSQL(ctx, r.Database, "CheckUserInWaitingListOfClassroom", strings.Join(query, " "))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, 0, nil
		}
		return false, 0, err
	}

	var waitingList WaitingListOutputRepo
	if err := row.Scan(&waitingList.ID, &waitingList.UserID, &waitingList.ClassroomID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, 0, nil
		}
		return false, 0, err
	}

	if waitingList.ClassroomID > 0 {
		return true, waitingList.ClassroomID, nil
	}

	return false, 0, nil
}
