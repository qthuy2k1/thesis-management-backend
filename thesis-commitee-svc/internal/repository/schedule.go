package repository

import (
	"context"
	"database/sql"
)

type ScheduleInputRepo struct {
	RoomID   int
	ThesisID int
}

type ScheduleOutputRepo struct {
	ID       int
	RoomID   int
	ThesisID int
}

// CreateSchedule creates a new schedule in db given by schedule model
func (r *CommiteeRepo) CreateSchedule(ctx context.Context, p ScheduleInputRepo) (ScheduleOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateSchedule", "INSERT INTO schedules (room_id, thesis_id) VALUES ($1, $2) RETURNING id, room_id", p.RoomID, p.ThesisID)
	if err != nil {
		return ScheduleOutputRepo{}, err
	}

	var scheduleOutput ScheduleOutputRepo
	if err := row.Scan(&scheduleOutput.ID, &scheduleOutput.RoomID, &scheduleOutput.ThesisID); err != nil {
		return ScheduleOutputRepo{}, err
	}

	return scheduleOutput, nil
}

// GetSchedule returns a schedule in db given by id
func (r *CommiteeRepo) GetSchedule(ctx context.Context, id int) (ScheduleOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetSchedule", "SELECT id, room_id, thesis_id FROM schedules WHERE id=$1", id)
	if err != nil {
		return ScheduleOutputRepo{}, err
	}

	schedule := ScheduleOutputRepo{}
	if err = row.Scan(&schedule.ID, &schedule.RoomID, &schedule.ThesisID); err != nil {
		if err == sql.ErrNoRows {
			return ScheduleOutputRepo{}, ErrScheduleNotFound
		}
		return ScheduleOutputRepo{}, err
	}

	return schedule, nil
}

// // CheckScheduleExists checks whether the specified schedule exists by title (true == exist)
// func (r *CommiteeRepo) IsScheduleExists(ctx context.Context, title string, classroomID int) (bool, error) {
// 	var exists bool
// 	row, err := QueryRowSQL(ctx, r.Database, "IsScheduleExists", "SELECT EXISTS(SELECT 1 FROM thesis_schedules WHERE title LIKE '%' || $1 || '%' AND classtime_slots_id=$2)", title, classroomID)
// 	if err != nil {
// 		return false, err
// 	}
// 	if err = row.Scan(&exists); err != nil {
// 		return false, err
// 	}
// 	return exists, nil
// }

// // UpdateSchedule updates the specified schedule by id
// func (r *CommiteeRepo) UpdateSchedule(ctx context.Context, id int, schedule ScheduleInputRepo) error {
// 	result, err := ExecSQL(ctx, r.Database, "UpdateSchedule", "UPDATE thesis_schedules SET start_date=$2, period=$3, time_slots_id=$4, time=$5 WHERE id=$1", id, schedule.StartDate, schedule.Period, schedule.TimeSlotsID, schedule.Time)
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
// 		return ErrScheduleNotFound
// 	}

// 	return nil
// }

// // DeleteSchedule deletes a schedule in db given by id
// func (r *CommiteeRepo) DeleteSchedule(ctx context.Context, id int) error {
// 	result, err := ExecSQL(ctx, r.Database, "DeleteSchedule", "DELETE FROM thesis_schedules WHERE id=$1", id)
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
// 		return ErrScheduleNotFound
// 	}

// 	return nil
// }

// GetSchedule returns a list of schedules in db with filter
func (r *CommiteeRepo) GetSchedules(ctx context.Context) ([]ScheduleOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetSchedules", "SELECT id, room_id, thesis_id FROM schedules")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the schedules slice
	var schedules []ScheduleOutputRepo
	for rows.Next() {
		schedule := ScheduleOutputRepo{}
		err := rows.Scan(
			&schedule.ID,
			&schedule.RoomID,
			&schedule.ThesisID,
		)
		if err != nil {
			return nil, 0, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getScheduleCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return schedules, count, nil
}

// GetSchedule returns a list of schedules in db with filter
func (r *CommiteeRepo) GetSchedulesByThesisID(ctx context.Context, thesisID int) ([]ScheduleOutputRepo, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetSchedules", "SELECT id, room_id, thesis_id FROM schedules WHERE thesis_id = $1", thesisID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the schedules slice
	var schedules []ScheduleOutputRepo
	for rows.Next() {
		schedule := ScheduleOutputRepo{}
		err := rows.Scan(
			&schedule.ID,
			&schedule.RoomID,
			&schedule.ThesisID,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *CommiteeRepo) getScheduleCount(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", "SELECT COUNT(*) FROM schedules")
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
