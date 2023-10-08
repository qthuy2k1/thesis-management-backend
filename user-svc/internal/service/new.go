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
	// GetAllLecturers returns all members who has the role named "lecturer"
	GetAllLecturers(ctx context.Context) ([]UserOutputSvc, int, error)
	// UnsubscribeClassroom returns an error if delete occurs any errors
	UnsubscribeClassroom(ctx context.Context, userID string, classroomID int) error

	// CreateMember creates a new user in db given by user model
	CreateMember(ctx context.Context, p MemberInputSvc) error
	// GetMember returns a user in db given by id
	GetMember(ctx context.Context, id int) (MemberOutputSvc, error)
	// UpdateMember updates the specified user by id
	UpdateMember(ctx context.Context, id int, classroom MemberInputSvc) error
	// DeleteMember deletes a user in db given by id
	DeleteMember(ctx context.Context, id int) error
	// GetMembers returns a list of users in db
	GetMembers(ctx context.Context) ([]MemberOutputSvc, int, error)
	// GetAllMembersOfClassroom returns a list of users in a classroom
	GetAllMembersOfClassroom(ctx context.Context, classroomID int) ([]MemberOutputSvc, int, error)
	// IsUserJoinedClassroom returns a member if exists
	IsUserJoinedClassroom(ctx context.Context, userID string) (MemberOutputSvc, error)
}

type UserSvc struct {
	Repository repository.IUserRepo
}

func NewUserSvc(pRepo repository.IUserRepo) IUserSvc {
	return &UserSvc{Repository: pRepo}
}
