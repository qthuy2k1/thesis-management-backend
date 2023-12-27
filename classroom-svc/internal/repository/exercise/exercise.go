package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

type ExerciseInputRepo struct {
	Title            string
	Content          string
	ClassroomID      int
	Deadline         time.Time
	ReportingStageID int
	AuthorID         string
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

// CreateExercise creates a new exercise in db given by exercise model
func (r *ExerciseRepo) CreateExercise(ctx context.Context, p ExerciseInputRepo) (int64, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateExercise", "INSERT INTO exercises (title, content, classroom_id, deadline, reporting_stage_id, author_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", p.Title, p.Content, p.ClassroomID, p.Deadline, p.ReportingStageID, p.AuthorID)
	if err != nil {
		return 0, err
	}

	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

type ExerciseOutputRepo struct {
	ID               int
	Title            string
	Content          string
	ClassroomID      int
	Deadline         time.Time
	ReportingStageID int
	AuthorID         string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// GetExercise returns a exercise in db given by id
func (r *ExerciseRepo) GetExercise(ctx context.Context, id int) (ExerciseOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetExercise", "SELECT id, title, content, classroom_id, deadline, reporting_stage_id, author_id, created_at, updated_at FROM exercises WHERE id=$1", id)
	if err != nil {
		return ExerciseOutputRepo{}, err
	}

	exercise := ExerciseOutputRepo{}

	if err = row.Scan(&exercise.ID, &exercise.Title, &exercise.Content, &exercise.ClassroomID, &exercise.Deadline, &exercise.ReportingStageID, &exercise.AuthorID, &exercise.CreatedAt, &exercise.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return ExerciseOutputRepo{}, ErrExerciseNotFound
		}
		return ExerciseOutputRepo{}, err
	}

	return exercise, nil
}

// CheckExerciseExists checks whether the specified exercise exists by title (true == exist)
func (r *ExerciseRepo) IsExerciseExists(ctx context.Context, classroomID int, title string) (bool, error) {
	row, err := QueryRowSQL(ctx, r.Database, "IsExerciseExists", "SELECT EXISTS(SELECT 1 FROM exercises WHERE title LIKE '%' || $1 || '%' AND classroom_id=$2)", title, classroomID)
	if err != nil {
		return false, err
	}

	var exists bool
	if err = row.Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

// UpdateExercise updates the specified exercise by id
func (r *ExerciseRepo) UpdateExercise(ctx context.Context, id int, exercise ExerciseInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateExercise", "UPDATE exercises SET title=$2, content=$3, classroom_id=$4, deadline=$5, reporting_stage_id=$6, author_id=$7, updated_at=$8 WHERE id=$1", id, exercise.Title, exercise.Content, exercise.ClassroomID, exercise.Deadline, exercise.ReportingStageID, exercise.AuthorID, time.Now())
	if err != nil {
		return err
	}
	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrExerciseNotFound
	}

	return nil
}

// DeleteExercise deletes a exercise in db given by id
func (r *ExerciseRepo) DeleteExercise(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteExercise", "DELETE FROM exercises WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrExerciseNotFound
	}

	return nil
}

type ExerciseFilterRepo struct {
	Limit       int
	Page        int
	TitleSearch string
	SortColumn  string
	SortOrder   string
}

// GetExercise returns a list of exercises in db with filter
func (r *ExerciseRepo) GetExercises(ctx context.Context, filter ExerciseFilterRepo) ([]ExerciseOutputRepo, int, error) {
	log.Println(filter)
	query := []string{"SELECT id, title, content, classroom_id, deadline, reporting_stage_id, author_id, created_at, updated_at FROM exercises"}

	if filter.TitleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE UPPER(title) LIKE UPPER('%s')", "%"+filter.TitleSearch+"%"))
	}

	query = append(query, fmt.Sprintf("ORDER BY %s %s", filter.SortColumn, filter.SortOrder),
		fmt.Sprintf("LIMIT %d OFFSET %d", filter.Limit, (filter.Page-1)*filter.Limit))

	rows, err := QuerySQL(ctx, r.Database, "GetExercises", strings.Join(query, " "))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the exercises slice
	var exercises []ExerciseOutputRepo
	for rows.Next() {
		exercise := ExerciseOutputRepo{}
		err := rows.Scan(
			&exercise.ID,
			&exercise.Title,
			&exercise.Content,
			&exercise.ClassroomID,
			&exercise.Deadline,
			&exercise.ReportingStageID,
			&exercise.AuthorID,
			&exercise.CreatedAt,
			&exercise.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		exercises = append(exercises, exercise)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCount(ctx, filter.TitleSearch)
	if err != nil {
		return nil, 0, err
	}

	return exercises, count, nil
}

// GetAllExercisesOfClassroom returns all exercises of the specified classroom given by classroom id
func (r *ExerciseRepo) GetAllExercisesOfClassroom(ctx context.Context, filter ExerciseFilterRepo, classromID int) ([]ExerciseOutputRepo, int, error) {
	query := []string{"SELECT id, title, content, classroom_id, deadline, reporting_stage_id, author_id, created_at, updated_at FROM exercises"}

	if filter.TitleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE classroom_id=%d AND UPPER(title) LIKE UPPER('%s')", classromID, "%"+filter.TitleSearch+"%"))
	} else {
		query = append(query, fmt.Sprintf("WHERE classroom_id=%d", classromID))
	}

	query = append(query, fmt.Sprintf("ORDER BY %s %s", filter.SortColumn, filter.SortOrder),
		fmt.Sprintf("LIMIT %d OFFSET %d", filter.Limit, (filter.Page-1)*filter.Limit))

	rows, err := QuerySQL(ctx, r.Database, "GetAllExercisesOfClassroom", strings.Join(query, " "))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the exercises slice
	var exercises []ExerciseOutputRepo
	for rows.Next() {
		exercise := ExerciseOutputRepo{}
		err := rows.Scan(
			&exercise.ID,
			&exercise.Title,
			&exercise.Content,
			&exercise.ClassroomID,
			&exercise.Deadline,
			&exercise.ReportingStageID,
			&exercise.AuthorID,
			&exercise.CreatedAt,
			&exercise.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		exercises = append(exercises, exercise)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCountInClassroom(ctx, filter.TitleSearch, classromID)
	if err != nil {
		return nil, 0, err
	}

	return exercises, count, nil
}

// GetAllExercisesInReportingStage returns all exercises of the specified reporting stage given by reporting stage id
func (r *ExerciseRepo) GetAllExercisesInReportingStage(ctx context.Context, reportingStageID, classroomID int) ([]ExerciseOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetAllExercisesInReportingStage", fmt.Sprintf("SELECT id, title, content, classroom_id, deadline, reporting_stage_id, author_id, created_at, updated_at FROM exercises WHERE reporting_stage_id = %d AND classroom_id = %d", reportingStageID, classroomID))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the exercises slice
	var exercises []ExerciseOutputRepo
	for rows.Next() {
		exercise := ExerciseOutputRepo{}
		err := rows.Scan(
			&exercise.ID,
			&exercise.Title,
			&exercise.Content,
			&exercise.ClassroomID,
			&exercise.Deadline,
			&exercise.ReportingStageID,
			&exercise.AuthorID,
			&exercise.CreatedAt,
			&exercise.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		exercises = append(exercises, exercise)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCountInReportingStage(ctx, reportingStageID, classroomID)
	if err != nil {
		return nil, 0, err
	}

	return exercises, count, nil
}

func (r *ExerciseRepo) getCount(ctx context.Context, titleSearch string) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM exercises"}
	if titleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE UPPER(title) LIKE UPPER('%s')", "%"+titleSearch+"%"))
	}

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", strings.Join(query, " "))
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *ExerciseRepo) getCountInClassroom(ctx context.Context, titleSearch string, classroomID int) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM exercises"}
	if titleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE classroom_id=%d AND UPPER(title) LIKE UPPER('%s')", classroomID, "%"+titleSearch+"%"))
	} else {
		query = append(query, fmt.Sprintf("WHERE classroom_id=%d", classroomID))
	}

	rows, err := QueryRowSQL(ctx, r.Database, "getCountIntClassroom", strings.Join(query, " "))
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *ExerciseRepo) getCountInReportingStage(ctx context.Context, reportingStageID, classroomID int) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCountIntClassroom", fmt.Sprintf("SELECT COUNT(*) FROM exercises WHERE reporting_stage_id = %d AND classroom_id = %d", reportingStageID, classroomID))
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
