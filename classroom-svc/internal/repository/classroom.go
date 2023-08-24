package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	model "github.com/qthuy2k1/thesis-management-backend/classroom-svc/internal/model"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ClassroomInputRepo struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
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

	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "INSERT INTO posts (title, description, status) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the generated ID
	if _, err := stmt.ExecContext(ctx, clr.Title, clr.Description, clr.Status); err != nil {
		return err
	}

	return nil
}

type ClassroomOutputRepo struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// GetClassroom returns a classroom in db given by id
func (r *ClassroomRepo) GetClassroom(ctx context.Context, id int) (ClassroomOutputRepo, error) {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "SELECT id, title, description, status, created_at, updated_at FROM classrooms WHERE id=$1")
	if err != nil {
		return ClassroomOutputRepo{}, err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the classroom
	row := stmt.QueryRowContext(ctx, id)
	classroom := ClassroomOutputRepo{}

	if err = row.Scan(&classroom.ID, &classroom.Title, &classroom.Description, &classroom.Status, &classroom.CreatedAt, &classroom.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return ClassroomOutputRepo{}, ErrClassroomNotFound
		}
		return ClassroomOutputRepo{}, err
	}

	return classroom, nil
}

// CheckClassroomExists checks whether the specified classroom exists by title (true == exist)
func (r *ClassroomRepo) IsClassroomExists(ctx context.Context, title string) (bool, error) {
	count, err := model.Classrooms(qm.Where("title LIKE ?", "%"+title+"%")).Count(ctx, r.Database)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// UpdateClassroom updates the specified classroom by id
func (r *ClassroomRepo) UpdateClassroom(ctx context.Context, id int, classroom ClassroomInputRepo) error {
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "UPDATE classrooms SET title=$2, description=$3, status=$4, updated_at=$5 WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the ID of the updated classroom
	result, err := stmt.ExecContext(ctx, id, classroom.Title, classroom.Description, classroom.Status, time.Now())
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
	// Prepare the SQL statement
	stmt, err := r.Database.PrepareContext(ctx, "DELETE FROM classrooms WHERE id=$1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement and retrieve the deleted classroom's details
	result, err := stmt.ExecContext(ctx, id)
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
	query = append(query, "SELECT * FROM classrooms")

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

	// Iterate over the result rows and populate the classrooms slice
	var classrooms []ClassroomOutputRepo
	for rows.Next() {
		classroom := ClassroomOutputRepo{}
		err := rows.Scan(
			&classroom.ID,
			&classroom.Title,
			&classroom.Description,
			&classroom.Status,
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

	if err := r.Database.QueryRowContext(ctx, strings.Join(query, " ")).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
