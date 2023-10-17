package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type MemberInputRepo struct {
	ClassroomID int
	MemberID    string
	Status      string
	IsDefense   bool
}

type MemberOutputRepo struct {
	ID          int
	ClassroomID int
	MemberID    string
	Status      string
	IsDefense   bool
	CreatedAt   time.Time
}

// CreateMember creates a new user in db given by user model
func (r *UserRepo) CreateMember(ctx context.Context, u MemberInputRepo) error {
	exists, err := r.IsMemberExists(ctx, u.MemberID)
	if err != nil {
		logger(err, "in IsMemberExists function", "CreateMember")
		return err
	}

	if exists {
		return ErrMemberExisted
	}

	if _, err := ExecSQL(ctx, r.Database, "CreateMember", "INSERT INTO members (classroom_id, member_id, status, is_defense) VALUES ($1, $2, $3) RETURNING id", u.ClassroomID, u.MemberID, u.Status, u.IsDefense); err != nil {
		logger(err, "execute SQL statement", "CreateMember")
		return err
	}

	return nil
}

// GetMember returns a member in db given by id
func (r *UserRepo) GetMember(ctx context.Context, id int) (MemberOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetMember", "SELECT id, classroom_id, member_id, status, is_defense, created_at FROM members WHERE id=$1", id)
	if err != nil {
		logger(err, "query row sql", "GetMember")
		return MemberOutputRepo{}, err
	}

	member := MemberOutputRepo{}

	if err = row.Scan(&member.ID, &member.ClassroomID, &member.MemberID, &member.Status, &member.CreatedAt, &member.IsDefense); err != nil {
		if err == sql.ErrNoRows {
			logger(err, "member not found", "GetMember")
			return MemberOutputRepo{}, ErrMemberNotFound
		}
		logger(err, "row scan err", "GetMember")
		return MemberOutputRepo{}, err
	}

	return member, nil
}

// CheckMemberExists checks whether the specified user exists by title (true == exist)
func (r *UserRepo) IsMemberExists(ctx context.Context, memberID string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsUserExists", "SELECT EXISTS(SELECT 1 FROM members WHERE member_id = $1)", memberID)
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

// UpdateMember updates the specified member by id
func (r *UserRepo) UpdateMember(ctx context.Context, id int, member MemberInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateMember", "UPDATE members SET classroom_id = $2, member_id = $3, status = $4, is_defense = $5 WHERE id=$1", id, member.ClassroomID, member.MemberID, member.Status, member.IsDefense)
	if err != nil {
		logger(err, "Exec SQL", "UpdateMember")
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		logger(ErrMemberNotFound, "rows affected equal 0", "UpdateMember")
		return ErrMemberNotFound
	}

	return nil
}

// DeleteMember deletes a member in db given by id
func (r *UserRepo) DeleteMember(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteMember", "DELETE FROM members WHERE id=$1", id)
	if err != nil {
		logger(err, "Exec SQL", "DeleteMember")
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		logger(ErrMemberNotFound, "rows affected equal to 0", "DeleteMember")
		return ErrMemberNotFound
	}

	return nil
}

// GetMember returns a list of members in db with filter
func (r *UserRepo) GetMembers(ctx context.Context) ([]MemberOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetMembers", "SELECT id, classroom_id, member_id, status, is_defense, created_at FROM members")
	if err != nil {
		logger(err, "Query SQL", "GetMembers")
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the members slice
	var members []MemberOutputRepo
	for rows.Next() {
		member := MemberOutputRepo{}
		err := rows.Scan(
			&member.ID,
			&member.ClassroomID,
			&member.MemberID,
			&member.Status,
			&member.IsDefense,
			&member.CreatedAt,
		)
		if err != nil {
			logger(err, "rows scan", "GetMembers")
			return nil, 0, err
		}
		members = append(members, member)
	}

	if err := rows.Err(); err != nil {
		logger(err, "rows err", "GetMembers")
		return nil, 0, err
	}

	count, err := r.getMemberCount(ctx)
	if err != nil {
		logger(err, "rows count", "GetMembers")
		return nil, 0, err
	}

	return members, count, nil
}

// GetAllMemberOfClassroom returns all users of the specified classroom given by classroom id
func (r *UserRepo) GetAllMembersOfClassroom(ctx context.Context, classroomID int) ([]MemberOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetAllMembersOfClassroom", fmt.Sprintf("SELECT id, classroom_id, member_id, status, is_defense, created_at FROM members WHERE classroom_id=%d", classroomID))
	if err != nil {
		logger(err, "query SQL", "GetAllMembersOfClassroom")
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the users slice
	var users []MemberOutputRepo
	for rows.Next() {
		user := MemberOutputRepo{}
		err := rows.Scan(
			&user.ID,
			&user.ClassroomID,
			&user.MemberID,
			&user.Status,
			&user.IsDefense,
			&user.CreatedAt,
		)
		if err != nil {
			logger(err, "rows scan", "GetAllMembersOfClassroom")
			return nil, 0, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		logger(err, "rows err", "GetAllMembersOfClassroo")
		return nil, 0, err
	}

	count, err := r.getCountInClassroom(ctx, classroomID)
	if err != nil {
		logger(err, "rows count", "GetAllMembersOfClassroom")
		return nil, 0, err
	}

	return users, count, nil
}

func (r *UserRepo) getMemberCount(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getMemberCount", "SELECT COUNT(*) FROM members")
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

	query := []string{"SELECT COUNT(*) FROM members", fmt.Sprintf("WHERE classroom_id=%d", classroomID)}

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

// IsUserJoinedClassroom returns a member if exists
func (r *UserRepo) IsUserJoinedClassroom(ctx context.Context, userID string) (MemberOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "IsUserJoinedClassroom", fmt.Sprintf("SELECT id, classroom_id, member_id, status, is_defense, created_at FROM members WHERE member_id = '%s'", userID))
	if err != nil {
		return MemberOutputRepo{}, err
	}

	var member MemberOutputRepo
	if err = row.Scan(&member.ID, &member.ClassroomID, &member.MemberID, &member.Status, &member.IsDefense, &member.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return MemberOutputRepo{}, ErrMemberNotFound
		}
		return MemberOutputRepo{}, err
	}

	return member, nil
}

// UnsubscribeClassroom returns an error if delete occurs any errors
func (r *UserRepo) UnsubscribeClassroom(ctx context.Context, userID string, classroomID int) error {
	result, err := ExecSQL(ctx, r.Database, "UnsubscribeClassroom", fmt.Sprintf("DELETE FROM members WHERE member_id='%s' AND classroom_id='%d'", userID, classroomID))
	if err != nil {
		logger(err, "Exec SQL", "UnsubscribeClassroom")
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		logger(ErrMemberNotFound, "rows affected equal to 0", "UnsubscribeClassroom")
		return ErrMemberNotFound
	}

	return nil
}
