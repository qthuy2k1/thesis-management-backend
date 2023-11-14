package repository

import (
	"context"
	"database/sql"
	"time"
)

type ThesisInputRepo struct {
}

type ThesisOutputRepo struct {
	ID        int
	CreatedAt time.Time
}

// CreateThesis creates a new schedule in db given by schedule model
func (r *CommiteeRepo) CreateThesis(ctx context.Context, p ThesisInputRepo) (ThesisOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateThesis", "INSERT INTO thesis (id) VALUES (DEFAULT) RETURNING id, created_at")
	if err != nil {
		return ThesisOutputRepo{}, err
	}

	var scheduleOutput ThesisOutputRepo
	if err := row.Scan(&scheduleOutput.ID, &scheduleOutput.CreatedAt); err != nil {
		return ThesisOutputRepo{}, err
	}

	return scheduleOutput, nil
}

// GetThesis returns a schedule in db given by id
func (r *CommiteeRepo) GetThesis(ctx context.Context, id int) (ThesisOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetThesis", "SELECT id, schedule_id FROM thesis WHERE id=$1", id)
	if err != nil {
		return ThesisOutputRepo{}, err
	}

	schedule := ThesisOutputRepo{}
	if err = row.Scan(&schedule.ID, &schedule.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return ThesisOutputRepo{}, ErrThesisNotFound
		}
		return ThesisOutputRepo{}, err
	}

	return schedule, nil
}

// // CheckThesisExists checks whether the specified schedule exists by title (true == exist)
// func (r *CommiteeRepo) IsThesisExists(ctx context.Context, title string, classroomID int) (bool, error) {
// 	var exists bool
// 	row, err := QueryRowSQL(ctx, r.Database, "IsThesisExists", "SELECT EXISTS(SELECT 1 FROM thesis_thesis WHERE title LIKE '%' || $1 || '%' AND classtime_slots_id=$2)", title, classroomID)
// 	if err != nil {
// 		return false, err
// 	}
// 	if err = row.Scan(&exists); err != nil {
// 		return false, err
// 	}
// 	return exists, nil
// }

// // UpdateThesis updates the specified schedule by id
// func (r *CommiteeRepo) UpdateThesis(ctx context.Context, id int, schedule ThesisInputRepo) error {
// 	result, err := ExecSQL(ctx, r.Database, "UpdateThesis", "UPDATE thesis_thesis SET start_date=$2, period=$3, time_slots_id=$4, time=$5 WHERE id=$1", id, schedule.StartDate, schedule.Period, schedule.TimeSlotsID, schedule.Time)
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
// 		return ErrThesisNotFound
// 	}

// 	return nil
// }

// // DeleteThesis deletes a schedule in db given by id
// func (r *CommiteeRepo) DeleteThesis(ctx context.Context, id int) error {
// 	result, err := ExecSQL(ctx, r.Database, "DeleteThesis", "DELETE FROM thesis_thesis WHERE id=$1", id)
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
// 		return ErrThesisNotFound
// 	}

// 	return nil
// }

// GetThesis returns a list of thesis in db with filter
func (r *CommiteeRepo) GetThesiss(ctx context.Context) ([]ThesisOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetThesiss", "SELECT id, created_at FROM thesis")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the thesis slice
	var thesis []ThesisOutputRepo
	for rows.Next() {
		schedule := ThesisOutputRepo{}
		err := rows.Scan(
			&schedule.ID,
			&schedule.CreatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		thesis = append(thesis, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getThesisCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return thesis, count, nil
}

func (r *CommiteeRepo) getThesisCount(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", "SELECT COUNT(*) FROM thesis")
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
