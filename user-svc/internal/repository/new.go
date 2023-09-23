package repository

import (
	"context"
	"database/sql"
)

type IUserRepo interface {
	// CreateUser creates a new user in db given by user model
	CreateUser(ctx context.Context, u UserInputRepo) error
	// GetUser returns a user in db given by id
	GetUser(ctx context.Context, id string) (UserOutputRepo, error)
	// CheckUserExists checks whether the specified user exists by name
	IsUserExists(ctx context.Context, email, id string) (bool, error)
	// UpdateUser updates the specified user by id
	UpdateUser(ctx context.Context, id string, user UserInputRepo) error
	// DeleteUser deletes a classroom in db given by id
	DeleteUser(ctx context.Context, id string) error
	// GetUsers returns a list of users in db with filter
	GetUsers(ctx context.Context) ([]UserOutputRepo, int, error)
	// GetAllUsersOfClassroom returns all users of the specified classroom given by classroom id
	GetAllUsersOfClassroom(ctx context.Context, classromID int) ([]UserOutputRepo, int, error)
}

type UserRepo struct {
	Database *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &UserRepo{Database: db}
}
