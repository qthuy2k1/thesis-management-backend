package repository

import (
	"context"
	"database/sql"
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

type CommiteeInputRepo struct {
	StartDate   time.Time
	Period      string
	Time        string
	TimeSlotsID int
}

type CommiteeOutputRepo struct {
	ID          int
	StartDate   time.Time
	Period      string
	Time        string
	TimeSlotsID int
}

// CreateCommitee creates a new commitee in db given by commitee model
func (r *CommiteeRepo) CreateCommitee(ctx context.Context, p CommiteeInputRepo) (CommiteeOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateCommitee", "INSERT INTO thesis_commitees (start_date, period, time_slots_id, time) VALUES ($1, $2, $3, $4) RETURNING id, start_date, period, room", p.StartDate, p.Period, p.TimeSlotsID)
	if err != nil {
		return CommiteeOutputRepo{}, err
	}

	var commiteeOutput CommiteeOutputRepo
	if err := row.Scan(&commiteeOutput.ID, &commiteeOutput.StartDate, &commiteeOutput.Period, &commiteeOutput.Time); err != nil {
		return CommiteeOutputRepo{}, err
	}

	return commiteeOutput, nil
}

// GetCommitee returns a commitee in db given by id
func (r *CommiteeRepo) GetCommitee(ctx context.Context, id int) (CommiteeOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetCommitee", "SELECT id, start_date, period, time_slots_id, time FROM thesis_commitees WHERE id=$1", id)
	if err != nil {
		return CommiteeOutputRepo{}, err
	}

	commitee := CommiteeOutputRepo{}
	if err = row.Scan(&commitee.ID, &commitee.StartDate, &commitee.Period, &commitee.TimeSlotsID, &commitee.Time); err != nil {
		if err == sql.ErrNoRows {
			return CommiteeOutputRepo{}, ErrCommiteeNotFound
		}
		return CommiteeOutputRepo{}, err
	}

	return commitee, nil
}

// CheckCommiteeExists checks whether the specified commitee exists by title (true == exist)
func (r *CommiteeRepo) IsCommiteeExists(ctx context.Context, title string, classroomID int) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsCommiteeExists", "SELECT EXISTS(SELECT 1 FROM thesis_commitees WHERE title LIKE '%' || $1 || '%' AND classtime_slots_id=$2)", title, classroomID)
	if err != nil {
		return false, err
	}
	if err = row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

// UpdateCommitee updates the specified commitee by id
func (r *CommiteeRepo) UpdateCommitee(ctx context.Context, id int, commitee CommiteeInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateCommitee", "UPDATE thesis_commitees SET start_date=$2, period=$3, time_slots_id=$4, time=$5 WHERE id=$1", id, commitee.StartDate, commitee.Period, commitee.TimeSlotsID, commitee.Time)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrCommiteeNotFound
	}

	return nil
}

// DeleteCommitee deletes a commitee in db given by id
func (r *CommiteeRepo) DeleteCommitee(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteCommitee", "DELETE FROM thesis_commitees WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrCommiteeNotFound
	}

	return nil
}

// GetCommitee returns a list of commitees in db with filter
func (r *CommiteeRepo) GetCommitees(ctx context.Context) ([]CommiteeOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetCommitees", "SELECT id, start_date, period, time_slots_id, time FROM thesis_commitees")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the commitees slice
	var commitees []CommiteeOutputRepo
	for rows.Next() {
		commitee := CommiteeOutputRepo{}
		err := rows.Scan(
			&commitee.ID,
			&commitee.StartDate,
			&commitee.Period,
			&commitee.TimeSlotsID,
			&commitee.Time,
		)
		if err != nil {
			return nil, 0, err
		}
		commitees = append(commitees, commitee)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return commitees, count, nil
}

func (r *CommiteeRepo) getCount(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", "SELECT COUNT(*) FROM thesis_commitees")
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// GetCommitee returns a commitee in db given by id
func (r *CommiteeRepo) GetCommiteeByTimeSlotsID(ctx context.Context, timeSlotsID int) (CommiteeOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetCommitee", "SELECT id, start_date, period, time_slots_id, time FROM thesis_commitees WHERE time_slots_id=$1", timeSlotsID)
	if err != nil {
		return CommiteeOutputRepo{}, err
	}

	commitee := CommiteeOutputRepo{}
	if err = row.Scan(&commitee.ID, &commitee.StartDate, &commitee.Period, &commitee.TimeSlotsID, &commitee.Time); err != nil {
		if err == sql.ErrNoRows {
			return CommiteeOutputRepo{}, ErrCommiteeNotFound
		}
		return CommiteeOutputRepo{}, err
	}

	return commitee, nil
}
