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
	ID          int
	Title       string
	Content     string
	ClassroomID int
	Deadline    time.Time
	Score       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// CreateExercise creates a new exercise in db given by exercise model
func (r *ExerciseRepo) CreateExercise(ctx context.Context, p ExerciseInputRepo) error {
	// check exercise exists
	isExists, err := r.IsExerciseExists(ctx, p.Title)
	if err != nil {
		return err
	}

	if isExists {
		return ErrExerciseExisted
	}

	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "INSERT INTO exercises (title, content, classroom_id, deadline, score) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the generated ID
	if _, err := stmt.ExecContext(ctx, p.Title, p.Content, p.ClassroomID, p.Deadline, p.Score); err != nil {
		return err
	}

	return nil
}

type ExerciseOutputRepo struct {
	ID          int
	Title       string
	Content     string
	ClassroomID int
	Deadline    time.Time
	Score       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GetExercise returns a exercise in db given by id
func (r *ExerciseRepo) GetExercise(ctx context.Context, id int) (ExerciseOutputRepo, error) {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "SELECT id, title, content, classroom_id, deadline, score, created_at, updated_at FROM exercises WHERE id=$1")
	if err != nil {
		return ExerciseOutputRepo{}, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the exercise
	row := stmt.QueryRowContext(ctx, id)
	exercise := ExerciseOutputRepo{}

	if err = row.Scan(&exercise.ID, &exercise.Title, &exercise.Content, &exercise.ClassroomID, &exercise.Deadline, &exercise.Score, &exercise.CreatedAt, &exercise.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return ExerciseOutputRepo{}, ErrExerciseNotFound
		}
		return ExerciseOutputRepo{}, err
	}

	return exercise, nil
}

// CheckExerciseExists checks whether the specified exercise exists by title (true == exist)
func (r *ExerciseRepo) IsExerciseExists(ctx context.Context, title string) (bool, error) {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "SELECT EXISTS(SELECT 1 FROM exercises WHERE title LIKE '%' || $1 || '%')")
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

// UpdateExercise updates the specified exercise by id
func (r *ExerciseRepo) UpdateExercise(ctx context.Context, id int, exercise ExerciseInputRepo) error {
	log.Println(id, exercise)
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "UPDATE exercises SET title=$2, content=$3, classroom_id=$4, deadline=$5, score=$6, updated_at=$7 WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the ID of the updated exercise
	result, err := stmt.ExecContext(ctx, id, exercise.Title, exercise.Content, exercise.ClassroomID, exercise.Deadline, exercise.Score, time.Now())
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
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "DELETE FROM exercises WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the deleted exercise's details
	result, err := stmt.ExecContext(ctx, id)
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
	query := []string{"SELECT * FROM exercises"}

	if filter.TitleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE UPPER(title) LIKE UPPER('%s')", "%"+filter.TitleSearch+"%"))
	}

	query = append(query, fmt.Sprintf("ORDER BY %s %s", filter.SortColumn, filter.SortOrder),
		fmt.Sprintf("LIMIT %d OFFSET %d", filter.Limit, (filter.Page-1)*filter.Limit))

	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, strings.Join(query, " "))
	if err != nil {
		return nil, 0, err
	}
	defer stmt.Close()

	// Execute the SQL statement
	rows, err := stmt.QueryContext(ctx)
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
			&exercise.Score,
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

func (r *ExerciseRepo) getCount(ctx context.Context, titleSearch string) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM exercises"}
	if titleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE UPPER(title) LIKE UPPER('%s')", "%"+titleSearch+"%"))
	}

	if err := r.Database.QueryRowContext(ctx, strings.Join(query, " ")).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
