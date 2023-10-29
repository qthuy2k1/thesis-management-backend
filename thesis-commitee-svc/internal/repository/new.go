package repository

import (
	"context"
	"database/sql"
)

type ICommiteeRepo interface {
	// CreateCommitee creates a new commitee in db given by commitee model
	CreateCommitee(ctx context.Context, clr CommiteeInputRepo) (CommiteeOutputRepo, error)
	// GetCommitee returns a commitee in db given by id
	GetCommitee(ctx context.Context, id int) (CommiteeOutputRepo, error)
	// CheckCommiteeExists checks whether the specified commitee exists by name
	IsCommiteeExists(ctx context.Context, title string, classroomID int) (bool, error)
	// UpdateCommitee updates the specified commitee by id
	UpdateCommitee(ctx context.Context, id int, commitee CommiteeInputRepo) error
	// DeleteCommitee deletes a commitee in db given by id
	DeleteCommitee(ctx context.Context, id int) error
	// GetCommitees returns a list of commitees in db with filter
	GetCommitees(ctx context.Context) ([]CommiteeOutputRepo, int, error)

	// CreateCommiteeUserDetail creates a new commitee in db given by commitee model
	CreateCommiteeUserDetail(ctx context.Context, clr CommiteeUserDetailInputRepo) (CommiteeUserDetailOutputRepo, error)
	// GetCommiteeUserDetail returns a commitee in db given by id
	GetCommiteeUserDetail(ctx context.Context, id int) (CommiteeUserDetailOutputRepo, error)
	// CheckCommiteeUserDetailExists checks whether the specified commitee exists by name
	IsCommiteeUserDetailExists(ctx context.Context, commiteeID int, lecturerID string, studentID []string) (bool, error)
	// UpdateCommiteeUserDetail updates the specified commitee by id
	UpdateCommiteeUserDetail(ctx context.Context, commitee CommiteeUserDetailInputRepo) error
	// DeleteCommiteeUserDetail deletes a commitee in db given by id
	DeleteCommiteeUserDetail(ctx context.Context, commiteeID int, lecturerID string, studentID []string) error
	// GetCommiteeUserDetails returns a list of commitees in db with filter
	GetCommiteeUserDetails(ctx context.Context) ([]CommiteeUserDetailOutputRepo, int, error)
	// GetAllCommiteeUserDetailsFromCommitee returns a list of all commitee user details from a commitee
	GetAllCommiteeUserDetailsFromCommitee(ctx context.Context, commiteeID int) ([]CommiteeUserDetailOutputRepo, error)

	// CreateRoom creates a new commitee in db given by commitee model
	CreateRoom(ctx context.Context, r RoomInputRepo) (RoomOutputRepo, error)
	// GetRoom returns a commitee in db given by id
	GetRoom(ctx context.Context, id int) (RoomOutputRepo, error)
	// CheckRoomExists checks whether the specified commitee exists by name
	IsRoomExists(ctx context.Context, name string, school string) (bool, error)
	// UpdateRoom updates the specified commitee by id
	UpdateRoom(ctx context.Context, id int, commitee RoomInputRepo) error
	// DeleteRoom deletes a commitee in db given by id
	DeleteRoom(ctx context.Context, id int) error
	// GetRooms returns a list of commitees in db with filter
	GetRooms(ctx context.Context, filter RoomFilter) ([]RoomOutputRepo, int, error)
}

type CommiteeRepo struct {
	Database *sql.DB
}

func NewCommiteeRepo(db *sql.DB) ICommiteeRepo {
	return &CommiteeRepo{Database: db}
}
