package repository

import (
	"context"
	"database/sql"
	"log"
	"strings"
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

type ReportingStageInputRepo struct {
	Name        string
	Description string
}

// CreateReportingStage creates a new reporting_stage in db given by reporting_stage model
func (r *ReportingStageRepo) CreateReportingStage(ctx context.Context, p ReportingStageInputRepo) error {
	// check reporting_stage exists
	isExists, err := r.IsReportingStageExists(ctx, p.Name)
	if err != nil {
		return err
	}

	if isExists {
		return ErrReportingStageExisted
	}

	if _, err := ExecSQL(ctx, r.Database, "CreateReportingStage", "INSERT INTO reporting_stages (name, description) VALUES ($1, $2) RETURNING id", p.Name, p.Description); err != nil {
		return err
	}

	return nil
}

type ReportingStageOutputRepo struct {
	ID          int
	Name        string
	Description string
}

// GetReportingStage returns a reporting_stage in db given by id
func (r *ReportingStageRepo) GetReportingStage(ctx context.Context, id int) (ReportingStageOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetReportingStage", "SELECT id, name, description FROM reporting_stages WHERE id=$1", id)
	if err != nil {
		return ReportingStageOutputRepo{}, err
	}

	reporting_stage := ReportingStageOutputRepo{}
	if err = row.Scan(&reporting_stage.ID, &reporting_stage.Name, &reporting_stage.Description); err != nil {
		if err == sql.ErrNoRows {
			return ReportingStageOutputRepo{}, ErrReportingStageNotFound
		}
		return ReportingStageOutputRepo{}, err
	}

	return reporting_stage, nil
}

// CheckReportingStageExists checks whether the specified reporting_stage exists by name (true == exist)
func (r *ReportingStageRepo) IsReportingStageExists(ctx context.Context, name string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsReportingStageExists", "SELECT EXISTS(SELECT 1 FROM reporting_stages WHERE name LIKE '%' || $1 || '%')", name)
	if err != nil {
		return false, err
	}
	if err = row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

// UpdateReportingStage updates the specified reporting_stage by id
func (r *ReportingStageRepo) UpdateReportingStage(ctx context.Context, id int, reporting_stage ReportingStageInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateReportingStage", "UPDATE reporting_stages SET name=$2, description=$3 WHERE id=$1", id, reporting_stage.Name, reporting_stage.Description)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrReportingStageNotFound
	}

	return nil
}

// DeleteReportingStage deletes a reporting_stage in db given by id
func (r *ReportingStageRepo) DeleteReportingStage(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteReportingStage", "DELETE FROM reporting_stages WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrReportingStageNotFound
	}

	return nil
}

// GetReportingStage returns a list of reporting_stages in db with filter
func (r *ReportingStageRepo) GetReportingStages(ctx context.Context) ([]ReportingStageOutputRepo, error) {
	query := []string{"SELECT * FROM reporting_stages"}

	rows, err := QuerySQL(ctx, r.Database, "GetReportingStages", strings.Join(query, " "))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the reporting_stages slice
	var reporting_stages []ReportingStageOutputRepo
	for rows.Next() {
		reporting_stage := ReportingStageOutputRepo{}
		err := rows.Scan(
			&reporting_stage.ID,
			&reporting_stage.Name,
			&reporting_stage.Description,
		)
		if err != nil {
			return nil, err
		}
		reporting_stages = append(reporting_stages, reporting_stage)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reporting_stages, nil
}
