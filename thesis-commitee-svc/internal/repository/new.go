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
	// GetCommitees returns a list of commitees in db with filter
	GetCommiteeByTimeSlotsID(ctx context.Context, timeSlotsID int) (CommiteeOutputRepo, error)

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
	// GetRooms returns a list of commitees in db with filter
	GetRoomsByID(ctx context.Context, id []string) ([]RoomOutputRepo, error)

	// CreateCouncil creates a new commitee in db given by commitee model
	CreateCouncil(ctx context.Context, r CouncilInputRepo) (CouncilOutputRepo, error)
	// GetCouncil returns a commitee in db given by id
	GetCouncil(ctx context.Context, id int) (CouncilOutputRepo, error)
	// GetCouncils returns a list of commitees in db with filter
	GetCouncils(ctx context.Context) ([]CouncilOutputRepo, int, error)
	// GetCouncils returns a list of commitees in db with filter
	GetCouncilsByThesisID(ctx context.Context, thesisID int) ([]CouncilOutputRepo, int, error)

	// CreateTimeSlots creates a new commitee in db given by commitee model
	CreateTimeSlots(ctx context.Context, r TimeSlotsInputRepo) (TimeSlotsOutputRepo, error)
	// GetTimeSlots returns a commitee in db given by id
	GetTimeSlots(ctx context.Context, id int) (TimeSlotsOutputRepo, error)
	// GetTimeSlotss returns a list of commitees in db with filter
	GetTimeSlotss(ctx context.Context) ([]TimeSlotsOutputRepo, int, error)
	// GetTimeSlots returns a commitee in db given by id
	GetTimeSlotsByScheduleID(ctx context.Context, scheduleID int) ([]TimeSlotsOutputRepo, error)

	// CreateSchedule creates a new commitee in db given by commitee model
	CreateSchedule(ctx context.Context, r ScheduleInputRepo) (ScheduleOutputRepo, error)
	// GetSchedule returns a commitee in db given by id
	GetSchedule(ctx context.Context, id int) (ScheduleOutputRepo, error)
	// GetSchedules returns a list of commitees in db with filter
	GetSchedules(ctx context.Context) ([]ScheduleOutputRepo, int, error)
	// GetSchedules returns a list of commitees in db with filter
	GetSchedulesByThesisID(ctx context.Context, thesisID int) ([]ScheduleOutputRepo, error)

	// CreateThesis creates a new commitee in db given by commitee model
	CreateThesis(ctx context.Context, r ThesisInputRepo) (ThesisOutputRepo, error)
	// GetThesis returns a commitee in db given by id
	GetThesis(ctx context.Context, id int) (ThesisOutputRepo, error)
	// GetThesiss returns a list of commitees in db with filter
	GetThesiss(ctx context.Context) ([]ThesisOutputRepo, int, error)
}

type CommiteeRepo struct {
	Database *sql.DB
}

func NewCommiteeRepo(db *sql.DB) ICommiteeRepo {
	return &CommiteeRepo{Database: db}
}
