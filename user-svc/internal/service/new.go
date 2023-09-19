package service

import (
	"context"

	repository "github.com/qthuy2k1/thesis-management-backend/user-svc/internal/repository"
)

type IUserSvc interface {
	// CreateUser creates a new user in db given by user model
	CreateUser(ctx context.Context, p UserInputSvc) error
	// GetUser returns a user in db given by id
	GetUser(ctx context.Context, id string) (UserOutputSvc, error)
	// UpdateUser updates the specified user by id
	UpdateUser(ctx context.Context, id string, classroom UserInputSvc) error
	// DeleteUser deletes a user in db given by id
	DeleteUser(ctx context.Context, id string) error
	// GetUsers returns a list of users in db
	GetUsers(ctx context.Context) ([]UserOutputSvc, int, error)
	// GetAllUsersOfClassroom returns a list of users in a classroom
	GetAllUsersOfClassroom(ctx context.Context, classroomID int) ([]UserOutputSvc, int, error)
}

type UserSvc struct {
	Repository repository.IUserRepo
}

func NewUserSvc(pRepo repository.IUserRepo) IUserSvc {
	return &UserSvc{Repository: pRepo}
}
