package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type StudentDefInputRepo struct {
	UserID       string
	InstructorID string
}

type StudentDefOutputRepo struct {
	ID           int
	UserID       string
	InstructorID string
}

// CreateStudentDef creates a new user in db given by user model
func (r *UserRepo) CreateStudentDef(ctx context.Context, u StudentDefInputRepo) error {
	exists, err := r.IsStudentDefExists(ctx, u.UserID)
	if err != nil {
		logger(err, "in IsStudentDefExists function", "CreateStudentDef")
		return err
	}

	if exists {
		return ErrStudentDefExisted
	}

	if _, err := ExecSQL(ctx, r.Database, "CreateStudentDef", "INSERT INTO student_defs (user_id, instructor_id) VALUES ($1, $2) RETURNING id", u.UserID, u.InstructorID); err != nil {
		logger(err, "execute SQL statement", "CreateStudentDef")
		return err
	}

	return nil
}

// GetStudentDef returns a student_def in db given by id
func (r *UserRepo) GetStudentDef(ctx context.Context, id int) (StudentDefOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetStudentDef", "SELECT id, user_id, instructor_id FROM student_defs WHERE id=$1", id)
	if err != nil {
		logger(err, "query row sql", "GetStudentDef")
		return StudentDefOutputRepo{}, err
	}

	studentDef := StudentDefOutputRepo{}

	if err = row.Scan(&studentDef.ID, &studentDef.UserID, &studentDef.InstructorID); err != nil {
		if err == sql.ErrNoRows {
			logger(err, "student def not found", "GetStudentDef")
			return StudentDefOutputRepo{}, ErrStudentDefNotFound
		}
		logger(err, "row scan err", "GetStudentDef")
		return StudentDefOutputRepo{}, err
	}

	return studentDef, nil
}

// CheckStudentDefExists checks whether the specified user exists by title (true == exist)
func (r *UserRepo) IsStudentDefExists(ctx context.Context, userID string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsUserExists", "SELECT EXISTS(SELECT 1 FROM student_defs WHERE user_id = $1)", userID)
	if err != nil {
		logger(err, "Query sql", "IsUserExists")
		return false, err
	}
	if err = row.Scan(&exists); err != nil {
		logger(err, "Scan sql", "IsUserExists")
		return false, err
	}
	return exists, nil
}

// UpdateStudentDef updates the specified student_def by id
func (r *UserRepo) UpdateStudentDef(ctx context.Context, id int, studentDef StudentDefInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateStudentDef", "UPDATE student_defs SET user_id=$2, instructor_id=$3 WHERE id=$1", id, studentDef.UserID, studentDef.InstructorID)
	if err != nil {
		logger(err, "Exec SQL", "UpdateStudentDef")
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		logger(ErrStudentDefNotFound, "rows affected equal 0", "UpdateStudentDef")
		return ErrStudentDefNotFound
	}

	return nil
}

// DeleteStudentDef deletes a student_def in db given by id
func (r *UserRepo) DeleteStudentDef(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteStudentDef", "DELETE FROM student_defs WHERE id=$1", id)
	if err != nil {
		logger(err, "Exec SQL", "DeleteStudentDef")
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		logger(ErrStudentDefNotFound, "rows affected equal to 0", "DeleteStudentDef")
		return ErrStudentDefNotFound
	}

	return nil
}

// GetStudentDef returns a list of student_defs in db with filter
func (r *UserRepo) GetStudentDefs(ctx context.Context) ([]StudentDefOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetStudentDefs", "SELECT id, user_id, instructor_id FROM student_defs")
	if err != nil {
		logger(err, "Query SQL", "GetStudentDefs")
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the student_defs slice
	var studentDefs []StudentDefOutputRepo
	for rows.Next() {
		studentDef := StudentDefOutputRepo{}
		err := rows.Scan(
			&studentDef.ID,
			&studentDef.UserID,
			&studentDef.InstructorID,
		)
		if err != nil {
			logger(err, "rows scan", "GetStudentDefs")
			return nil, 0, err
		}
		studentDefs = append(studentDefs, studentDef)
	}

	if err := rows.Err(); err != nil {
		logger(err, "rows err", "GetStudentDefs")
		return nil, 0, err
	}

	count, err := r.getCountStudentDef(ctx)
	if err != nil {
		logger(err, "rows count", "GetStudentDefs")
		return nil, 0, err
	}

	return studentDefs, count, nil
}

// GetAllStudentDefOfClassroom returns all users of the specified classroom given by classroom id
func (r *UserRepo) GetAllStudentDefsOfInstructor(ctx context.Context, instructorID string) ([]StudentDefOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetAllStudentDefsOfClassroom", fmt.Sprintf("SELECT id, user_id, instructor_id FROM student_defs WHERE instructor_id LIKE '%s'", instructorID))
	if err != nil {
		logger(err, "query SQL", "GetAllStudentDefsOfClassroom")
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the users slice
	var studentDefs []StudentDefOutputRepo
	for rows.Next() {
		studentDef := StudentDefOutputRepo{}
		err := rows.Scan(
			&studentDef.ID,
			&studentDef.UserID,
			&studentDef.InstructorID,
		)
		if err != nil {
			logger(err, "rows scan", "GetAllStudentDefsOfClassroom")
			return nil, 0, err
		}
		studentDefs = append(studentDefs, studentDef)
	}

	if err := rows.Err(); err != nil {
		logger(err, "rows err", "GetAllStudentDefsOfClassroo")
		return nil, 0, err
	}

	count, err := r.getCountStudentDefOfInstructor(ctx, instructorID)
	if err != nil {
		logger(err, "rows count", "GetAllStudentDefsOfClassroom")
		return nil, 0, err
	}

	return studentDefs, count, nil
}

func (r *UserRepo) getCountStudentDef(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getStudentDefCount", "SELECT COUNT(*) FROM student_defs")
	if err != nil {
		logger(err, "Query sql", "getCount")
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		logger(err, "rows scan", "getCount")
		return 0, err
	}

	return count, nil
}

func (r *UserRepo) getCountStudentDefOfInstructor(ctx context.Context, instructorID string) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM student_defs", fmt.Sprintf("WHERE instructor_id LIKE '%s'", instructorID)}

	rows, err := QueryRowSQL(ctx, r.Database, "getCountIntClassroom", strings.Join(query, " "))
	if err != nil {
		logger(err, "Query SQL", "getCountInClassroom")
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		logger(err, "rows scan", "getCountInClassroom")
		return 0, err
	}

	return count, nil
}
