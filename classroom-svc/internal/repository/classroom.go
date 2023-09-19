package repository

import (
	"context"
	"database/sql"
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

type ClassroomInputRepo struct {
	ID            int
	Title         string
	Description   string
	Status        string
	LecturerID    string
	CodeClassroom string
	TopicTags     string
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// CreateClasroom creates a new classroom in db given by classroom model
func (r *ClassroomRepo) CreateClassroom(ctx context.Context, clr ClassroomInputRepo) error {
	// check post exists
	isExists, err := r.IsClassroomExists(ctx, clr.Title)
	if err != nil {
		return err
	}

	if isExists {
		return ErrClassroomExisted
	}

	if _, err := ExecSQL(ctx, r.Database, "CreateClassroom", "INSERT INTO classrooms (title, description, status, lecturer_id, code_classroom, topic_tags, quantity) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", clr.Title, clr.Description, clr.Status, clr.LecturerID, clr.CodeClassroom, clr.TopicTags, clr.Quantity); err != nil {
		return err
	}

	return nil
}

type ClassroomOutputRepo struct {
	ID            int
	Title         string
	Description   string
	Status        string
	LecturerID    string
	CodeClassroom string
	TopicTags     string
	Quantity      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// GetClassroom returns a classroom in db given by id
func (r *ClassroomRepo) GetClassroom(ctx context.Context, id int) (ClassroomOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetClassroom", "SELECT id, title, description, status, lecturer_id, code_classroom, topic_tags, quantity, created_at, updated_at FROM classrooms WHERE id=$1", id)
	if err != nil {
		return ClassroomOutputRepo{}, err
	}
	classroom := ClassroomOutputRepo{}

	if err = row.Scan(&classroom.ID, &classroom.Title, &classroom.Description, &classroom.Status, &classroom.LecturerID, &classroom.CodeClassroom, &classroom.TopicTags, &classroom.Quantity, &classroom.CreatedAt, &classroom.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return ClassroomOutputRepo{}, ErrClassroomNotFound
		}
		return ClassroomOutputRepo{}, err
	}

	return classroom, nil
}

// CheckClassroomExists checks whether the specified classroom exists by title (true == exist)
func (r *ClassroomRepo) IsClassroomExists(ctx context.Context, title string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsClassroomExists", "SELECT EXISTS(SELECT 1 FROM classrooms WHERE title LIKE '%' || $1 || '%')", title)
	if err != nil {
		return false, err
	}
	if err = row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil

}

// UpdateClassroom updates the specified classroom by id
func (r *ClassroomRepo) UpdateClassroom(ctx context.Context, id int, classroom ClassroomInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateClassroom", "UPDATE classrooms SET title=$2, description=$3, status=$4, lecturer_id=$5, code_classroom=$6, topic_tags=$7, quantity=$8, updated_at=$9 WHERE id=$1", id, classroom.Title, classroom.Description, classroom.Status, classroom.LecturerID, classroom.CodeClassroom, classroom.TopicTags, classroom.Quantity, time.Now())
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrClassroomNotFound
	}

	return nil
}

// DeleteClassroom deletes a classroom in db given by id
func (r *ClassroomRepo) DeleteClassroom(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteClassroom", "DELETE FROM classrooms WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrClassroomNotFound
	}

	return nil
}

type ClassroomFilterRepo struct {
	Limit       int
	Page        int
	TitleSearch string
	SortColumn  string
	SortOrder   string
}

// GetClassroom returns a list of classrooms in db with filter
func (r *ClassroomRepo) GetClassrooms(ctx context.Context, filter ClassroomFilterRepo) ([]ClassroomOutputRepo, int, error) {
	var query []string
	query = append(query, "SELECT id, title, description, status, lecturer_id, code_classroom, topic_tags, quantity, created_at, updated_at FROM classrooms")

	if filter.TitleSearch != "" {
		query = append(query, fmt.Sprintf("WHERE UPPER(title) LIKE UPPER('%s')", "%"+filter.TitleSearch+"%"))
	}

	query = append(query, fmt.Sprintf("ORDER BY %s %s", filter.SortColumn, filter.SortOrder),
		fmt.Sprintf("LIMIT %d OFFSET %d", filter.Limit, (filter.Page-1)*filter.Limit))

	rows, err := QuerySQL(ctx, r.Database, "GetClassrooms", strings.Join(query, " "))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the classrooms slice
	var classrooms []ClassroomOutputRepo
	for rows.Next() {
		classroom := ClassroomOutputRepo{}
		err := rows.Scan(
			&classroom.ID,
			&classroom.Title,
			&classroom.Description,
			&classroom.Status,
			&classroom.LecturerID,
			&classroom.CodeClassroom,
			&classroom.TopicTags,
			&classroom.Quantity,
			&classroom.CreatedAt,
			&classroom.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		classrooms = append(classrooms, classroom)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCount(ctx, filter.TitleSearch)
	if err != nil {
		return nil, 0, err
	}

	return classrooms, count, nil
}

func (r *ClassroomRepo) getCount(ctx context.Context, titleSearch string) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM classrooms"}
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
