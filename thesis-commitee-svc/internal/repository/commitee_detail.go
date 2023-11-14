package repository

import (
	"context"
	"database/sql"
	"strings"
)

type CommiteeUserDetailInputRepo struct {
	CommiteeID   int
	InstructorID string
	StudentID    []string
}

type CommiteeUserDetailOutputRepo struct {
	CommiteeID   int
	InstructorID string
	StudentID    []string
}

// CreateCommiteeUserDetail creates a new commitee in db given by commitee model
func (r *CommiteeRepo) CreateCommiteeUserDetail(ctx context.Context, p CommiteeUserDetailInputRepo) (CommiteeUserDetailOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateCommiteeUserDetail", "INSERT INTO thesis_commitee_user_details (commitee_id, lecturer_id, student_id) VALUES ($1, $2, $3) RETURNING commitee_id, lecturer_id, student_id", p.CommiteeID, p.InstructorID, strings.Join(p.StudentID, ","))
	if err != nil {
		return CommiteeUserDetailOutputRepo{}, err
	}

	var commiteeUserDetailOutput CommiteeUserDetailOutputRepo
	var studentListStr string
	if err := row.Scan(&commiteeUserDetailOutput.CommiteeID, &commiteeUserDetailOutput.InstructorID, &studentListStr); err != nil {
		return CommiteeUserDetailOutputRepo{}, err
	}

	commiteeUserDetailOutput.StudentID = strings.Split(studentListStr, ",")

	return commiteeUserDetailOutput, nil
}

// GetCommiteeUserDetail returns a commitee in db given by id
func (r *CommiteeRepo) GetCommiteeUserDetail(ctx context.Context, id int) (CommiteeUserDetailOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetCommiteeUserDetail", "SELECT commitee_id, lecturer_id, student_id FROM thesis_commitee_user_details WHERE id=$1", id)
	if err != nil {
		return CommiteeUserDetailOutputRepo{}, err
	}

	commitee := CommiteeUserDetailOutputRepo{}
	var studentListStr string
	if err = row.Scan(&commitee.CommiteeID, &commitee.InstructorID, &studentListStr); err != nil {
		if err == sql.ErrNoRows {
			return CommiteeUserDetailOutputRepo{}, ErrCommiteeUserDetailNotFound
		}
		return CommiteeUserDetailOutputRepo{}, err
	}

	commitee.StudentID = strings.Split(studentListStr, ",")
	return commitee, nil
}

// CheckCommiteeUserDetailExists checks whether the specified commitee exists by title (true == exist)
func (r *CommiteeRepo) IsCommiteeUserDetailExists(ctx context.Context, commiteeID int, lecturerID string, studentID []string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsCommiteeUserDetailExists", "SELECT EXISTS(SELECT 1 FROM thesis_commitee_user_details WHERE commitee_id=$1 AND lecturer_id=$2 AND student_id=$3)", commiteeID, lecturerID, strings.Join(studentID, ","))
	if err != nil {
		return false, err
	}
	if err = row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

// UpdateCommiteeUserDetail updates the specified commitee by id
func (r *CommiteeRepo) UpdateCommiteeUserDetail(ctx context.Context, commitee CommiteeUserDetailInputRepo) error {
	if err := r.deleteAllCommiteeDetailByCommiteeID(ctx, commitee.CommiteeID); err != nil {
		return err
	}

	result, err := ExecSQL(ctx, r.Database, "UpdateCommiteeUserDetail", "UPDATE thesis_commitee_user_details SET commitee_id=$2, lecturer_id=$3, student_id=$4 WHERE id=$1", commitee.CommiteeID, commitee.InstructorID, strings.Join(commitee.StudentID, ","))
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrCommiteeUserDetailNotFound
	}

	return nil
}

// DeleteCommiteeUserDetail deletes a commitee in db given by id
func (r *CommiteeRepo) DeleteCommiteeUserDetail(ctx context.Context, commiteeID int, lecturerID string, studentID []string) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteCommiteeUserDetail", "DELETE FROM thesis_commitee_user_details WHERE commitee_id=$1, lecturer_id=$2, student_id=$3", commiteeID, lecturerID, strings.Join(studentID, ","))
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrCommiteeUserDetailNotFound
	}

	return nil
}

// GetCommiteeUserDetail returns a list of commitees in db with filter
func (r *CommiteeRepo) GetCommiteeUserDetails(ctx context.Context) ([]CommiteeUserDetailOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetCommiteeUserDetails", "SELECT commitee_id, lecturer_id, student_id FROM thesis_commitee_user_details")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the commitees slice
	var commitees []CommiteeUserDetailOutputRepo
	for rows.Next() {
		var studentListStr string

		commitee := CommiteeUserDetailOutputRepo{}
		err := rows.Scan(
			&commitee.CommiteeID,
			&commitee.InstructorID,
			&studentListStr,
		)
		if err != nil {
			return nil, 0, err
		}

		commitee.StudentID = strings.Split(studentListStr, ",")

		commitees = append(commitees, commitee)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCountComiteeDetail(ctx)
	if err != nil {
		return nil, 0, err
	}

	return commitees, count, nil
}

func (r *CommiteeRepo) getCountComiteeDetail(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", "SELECT COUNT(*) FROM thesis_commitee_user_details")
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *CommiteeRepo) deleteAllCommiteeDetailByCommiteeID(ctx context.Context, commiteeID int) error {
	if _, err := ExecSQL(ctx, r.Database, "deleteAllCommiteeDetailByCommiteeID", "DELETE FROM thesis_commitee_user_details WHERE commitee_id=$1", commiteeID); err != nil {
		return err
	}

	return nil
}

// GetAllCommiteeUserDetailsFromCommitee returns a list of all commitee user details from a commitee
func (r *CommiteeRepo) GetAllCommiteeUserDetailsFromCommitee(ctx context.Context, commiteeID int) ([]CommiteeUserDetailOutputRepo, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetAllCommiteeUserDetailsFromCommitee", "SELECT commitee_id, lecturer_id, student_id FROM thesis_commitee_user_details WHERE commitee_id=$1", commiteeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commiteeUsers []CommiteeUserDetailOutputRepo
	for rows.Next() {
		var studentListStr string

		var commiteeUser CommiteeUserDetailOutputRepo
		err := rows.Scan(
			&commiteeUser.CommiteeID,
			&commiteeUser.InstructorID,
			&studentListStr,
		)
		if err != nil {
			return nil, err
		}

		commiteeUser.StudentID = strings.Split(studentListStr, ",")

		commiteeUsers = append(commiteeUsers, commiteeUser)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return commiteeUsers, nil
}
