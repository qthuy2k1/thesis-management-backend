package repository

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
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
	// GetAllLecturers returns all members who has the role named "lecturer"
	GetAllLecturers(ctx context.Context) ([]UserOutputRepo, int, error)
	// UnsubscribeClassroom returns an error if delete occurs any errors
	UnsubscribeClassroom(ctx context.Context, userID string, classroomID int) error

	// CreateMember creates a new member in db given by member model
	CreateMember(ctx context.Context, u MemberInputRepo) error
	// GetMember returns a member in db given by id
	GetMember(ctx context.Context, id int) (MemberOutputRepo, error)
	// CheckMemberExists checks whether the specified member exists by name
	IsMemberExists(ctx context.Context, memberID string) (bool, error)
	// UpdateMember updates the specified member by id
	UpdateMember(ctx context.Context, id int, member MemberInputRepo) error
	// DeleteMember deletes a classroom in db given by id
	DeleteMember(ctx context.Context, id int) error
	// GetMembers returns a list of members in db with filter
	GetMembers(ctx context.Context) ([]MemberOutputRepo, int, error)
	// GetAllMembersOfClassroom returns all users of the specified classroom given by classroom id
	GetAllMembersOfClassroom(ctx context.Context, classroomID int) ([]MemberOutputRepo, int, error)
	// IsUserJoinedClassroom returns a member if exists
	IsUserJoinedClassroom(ctx context.Context, userID string) (MemberOutputRepo, error)

	// CreateStudentDef creates a new member in db given by member model
	CreateStudentDef(ctx context.Context, u StudentDefInputRepo) error
	// GetStudentDef returns a member in db given by id
	GetStudentDef(ctx context.Context, id int) (StudentDefOutputRepo, error)
	// CheckStudentDefExists checks whether the specified member exists by name
	IsStudentDefExists(ctx context.Context, userID string) (bool, error)
	// UpdateStudentDef updates the specified member by id
	UpdateStudentDef(ctx context.Context, id int, member StudentDefInputRepo) error
	// DeleteStudentDef deletes a classroom in db given by id
	DeleteStudentDef(ctx context.Context, id int) error
	// GetStudentDefs returns a list of members in db with filter
	GetStudentDefs(ctx context.Context) ([]StudentDefOutputRepo, int, error)
	// GetStudentDefsOfInstructor returns a list of members in db with filter
	GetAllStudentDefsOfInstructor(ctx context.Context, instructorID string) ([]StudentDefOutputRepo, int, error)
	// GetStudentDefs returns a list of members in db with filter
	GetStudentDefByTimeSlotsID(ctx context.Context, timeSlotsID int) (StudentDefOutputRepo, error)
}

type UserRepo struct {
	Database *sql.DB
	Redis    *redis.Client
}

func NewUserRepo(db *sql.DB, redis *redis.Client) IUserRepo {
	return &UserRepo{
		Database: db,
		Redis:    redis,
	}
}
