package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func logger(err error, describe string, functionName string) {
	layer := "Repository"
	log.Printf("Function %s in %s error(%s): %s\n", functionName, layer, describe, err.Error())
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

type UserInputRepo struct {
	ID          string
	Email       string
	Class       string
	Major       *string
	Phone       *string
	PhotoSrc    string
	Role        string
	Name        string
	ClassroomID *int
}

// CreateUser creates a new user in db given by user model
func (r *UserRepo) CreateUser(ctx context.Context, u UserInputRepo) error {
	exists, err := r.IsUserExists(ctx, u.Email)
	if err != nil {
		logger(err, "in IsUserExists function", "CreateUser")
		return err
	}

	if exists {
		return ErrUserExisted
	}

	if _, err := ExecSQL(ctx, r.Database, "CreateUser", "INSERT INTO users (id, class, major, phone, photo_src, role, name, email, classroom_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", u.ID, u.Class, u.Major, u.Phone, u.PhotoSrc, u.Role, u.Name, u.Email, u.ClassroomID); err != nil {
		logger(err, "execute SQL statement", "CreateUser")
		return err
	}

	return nil
}

type UserOutputRepo struct {
	ID          string
	Email       string
	Class       string
	Major       *string
	Phone       *string
	PhotoSrc    string
	Role        string
	Name        string
	ClassroomID *int
}

// GetUser returns a user in db given by id
func (r *UserRepo) GetUser(ctx context.Context, id string) (UserOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetUser", "SELECT id, class, major, phone, photo_src, role, name, email, classroom_id FROM users WHERE id=$1", id)
	if err != nil {
		logger(err, "query row sql", "GetUser")
		return UserOutputRepo{}, err
	}

	user := UserOutputRepo{}

	if err = row.Scan(&user.ID, &user.Class, &user.Major, &user.Phone, &user.PhotoSrc, &user.Role, &user.Name, &user.Email, &user.ClassroomID); err != nil {
		if err == sql.ErrNoRows {
			logger(err, "user not found", "GetUser")
			return UserOutputRepo{}, ErrUserNotFound
		}
		logger(err, "row scan err", "GetUser")
		return UserOutputRepo{}, err
	}

	return user, nil
}

// CheckUserExists checks whether the specified user exists by title (true == exist)
func (r *UserRepo) IsUserExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsUserExists", "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", email)
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

// UpdateUser updates the specified user by id
func (r *UserRepo) UpdateUser(ctx context.Context, id string, user UserInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateUser", "UPDATE users SET class=$2, major=$3, phone=$4, photo_src=$5, role=$6, name=$7, email=$8, classroom_id=$9 WHERE id=$1", id, user.Class, user.Major, user.Phone, user.PhotoSrc, user.Role, user.Name, user.Email, user.ClassroomID)
	if err != nil {
		logger(err, "Exec SQL", "UpdateUser")
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		logger(ErrUserNotFound, "rows affected equal 0", "UpdateUser")
		return ErrUserNotFound
	}

	return nil
}

// DeleteUser deletes a user in db given by id
func (r *UserRepo) DeleteUser(ctx context.Context, id string) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteUser", "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		logger(err, "Exec SQL", "DeleteUser")
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		logger(ErrUserNotFound, "rows affected equal to 0", "DeleteUser")
		return ErrUserNotFound
	}

	return nil
}

// GetUser returns a list of users in db with filter
func (r *UserRepo) GetUsers(ctx context.Context) ([]UserOutputRepo, int, error) {
	query := []string{"SELECT id, class, major, phone, photo_src, role, name, email, classroom_id FROM users"}

	rows, err := QuerySQL(ctx, r.Database, "GetUsers", strings.Join(query, " "))
	if err != nil {
		logger(err, "Query SQL", "GetUsers")
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the users slice
	var users []UserOutputRepo
	for rows.Next() {
		user := UserOutputRepo{}
		err := rows.Scan(
			&user.ID,
			&user.Class,
			&user.Major,
			&user.Phone,
			&user.PhotoSrc,
			&user.Role,
			&user.Name,
			&user.Email,
			&user.ClassroomID,
		)
		if err != nil {
			logger(err, "rows scan", "GetUsers")
			return nil, 0, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		logger(err, "rows err", "GetUsers")
		return nil, 0, err
	}

	count, err := r.getCount(ctx)
	if err != nil {
		logger(err, "rows count", "GetUsers")
		return nil, 0, err
	}

	return users, count, nil
}

// GetAllUsersOfClassroom returns all users of the specified classroom given by classroom id
func (r *UserRepo) GetAllUsersOfClassroom(ctx context.Context, classromID int) ([]UserOutputRepo, int, error) {
	query := []string{"SELECT id, class, major, phone, photo_src, role, name, email, classroom_id FROM users", fmt.Sprintf("WHERE classroom_id=%d", classromID)}

	rows, err := QuerySQL(ctx, r.Database, "GetAllUsersOfClassroom", strings.Join(query, " "))
	if err != nil {
		logger(err, "query SQL", "GetAllUsersOfClassroom")
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the users slice
	var users []UserOutputRepo
	for rows.Next() {
		user := UserOutputRepo{}
		err := rows.Scan(
			&user.ID,
			&user.Class,
			&user.Major,
			&user.Phone,
			&user.PhotoSrc,
			&user.Role,
			&user.Name,
			&user.Email,
			&user.ClassroomID,
		)
		if err != nil {
			logger(err, "rows scan", "GetAllUsersOfClassroom")
			return nil, 0, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		logger(err, "rows err", "GetAllUsersOfClassroo")
		return nil, 0, err
	}

	count, err := r.getCountInClassroom(ctx, classromID)
	if err != nil {
		logger(err, "rows count", "GetAllUsersOfClassroom")
		return nil, 0, err
	}

	return users, count, nil
}

func (r *UserRepo) getCount(ctx context.Context) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM users"}
	rows, err := QueryRowSQL(ctx, r.Database, "getCount", strings.Join(query, " "))
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

func (r *UserRepo) getCountInClassroom(ctx context.Context, classroomID int) (int, error) {
	var count int

	query := []string{"SELECT COUNT(*) FROM users", fmt.Sprintf("WHERE classroom_id=%d", classroomID)}

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
