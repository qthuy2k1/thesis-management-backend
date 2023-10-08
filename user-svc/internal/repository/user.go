package repository

import (
	"context"
	"database/sql"
	"log"
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
	ID       string
	Email    string
	Class    string
	Major    *string
	Phone    *string
	PhotoSrc string
	Role     string
	Name     string
}

// CreateUser creates a new user in db given by user model
func (r *UserRepo) CreateUser(ctx context.Context, u UserInputRepo) error {
	exists, err := r.IsUserExists(ctx, u.Email, u.ID)
	if err != nil {
		logger(err, "in IsUserExists function", "CreateUser")
		return err
	}

	if exists {
		return ErrUserExisted
	}

	if _, err := ExecSQL(ctx, r.Database, "CreateUser", "INSERT INTO users (id, class, major, phone, photo_src, role, name, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", u.ID, u.Class, u.Major, u.Phone, u.PhotoSrc, u.Role, u.Name, u.Email); err != nil {
		logger(err, "execute SQL statement", "CreateUser")
		return err
	}

	return nil
}

type UserOutputRepo struct {
	ID       string
	Email    string
	Class    string
	Major    *string
	Phone    *string
	PhotoSrc string
	Role     string
	Name     string
}

// GetUser returns a user in db given by id
func (r *UserRepo) GetUser(ctx context.Context, id string) (UserOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetUser", "SELECT id, class, major, phone, photo_src, role, name, email FROM users WHERE id=$1", id)
	if err != nil {
		logger(err, "query row sql", "GetUser")
		return UserOutputRepo{}, err
	}

	user := UserOutputRepo{}

	if err = row.Scan(&user.ID, &user.Class, &user.Major, &user.Phone, &user.PhotoSrc, &user.Role, &user.Name, &user.Email); err != nil {
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
func (r *UserRepo) IsUserExists(ctx context.Context, email, id string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsUserExists", "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND id = $2)", email, id)
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
	result, err := ExecSQL(ctx, r.Database, "UpdateUser", "UPDATE users SET class=$2, major=$3, phone=$4, photo_src=$5, role=$6, name=$7, email=$8 WHERE id=$1", id, user.Class, user.Major, user.Phone, user.PhotoSrc, user.Role, user.Name, user.Email)
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
	rows, err := QuerySQL(ctx, r.Database, "GetUsers", "SELECT id, class, major, phone, photo_src, role, name, email FROM users")
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

	count, err := r.getUserCount(ctx)
	if err != nil {
		logger(err, "rows count", "GetUsers")
		return nil, 0, err
	}

	return users, count, nil
}

// GetAllLecturers returns all members who has the role named "lecturer"
func (r *UserRepo) GetAllLecturers(ctx context.Context) ([]UserOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetAllLecturers", "SELECT id, class, major, phone, photo_src, role, name, email FROM users WHERE role = 'lecturer'")
	if err != nil {
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
		)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	var count int
	lecturerRows, err := QueryRowSQL(ctx, r.Database, "GetAllLecturers", "SELECT COUNT(*) FROM users WHERE role = 'lecturer'")
	if err != nil {
		return nil, 0, err
	}

	if err := lecturerRows.Scan(&count); err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (r *UserRepo) getUserCount(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getUserCount", "SELECT COUNT(*) FROM users")
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
