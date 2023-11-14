package repository

import (
	"context"
	"database/sql"
)

type TimeSlotsInputRepo struct {
	ScheduleID int
}

type TimeSlotsOutputRepo struct {
	ID         int
	ScheduleID int
}

// CreateTimeSlots creates a new timeSlots in db given by timeSlots model
func (r *CommiteeRepo) CreateTimeSlots(ctx context.Context, p TimeSlotsInputRepo) (TimeSlotsOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateTimeSlots", "INSERT INTO time_slots (schedule_id) VALUES ($1) RETURNING id, schedule_id", p.ScheduleID)
	if err != nil {
		return TimeSlotsOutputRepo{}, err
	}

	var timeSlotsOutput TimeSlotsOutputRepo
	if err := row.Scan(&timeSlotsOutput.ID, &timeSlotsOutput.ScheduleID); err != nil {
		return TimeSlotsOutputRepo{}, err
	}

	return timeSlotsOutput, nil
}

// GetTimeSlots returns a timeSlots in db given by id
func (r *CommiteeRepo) GetTimeSlots(ctx context.Context, id int) (TimeSlotsOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetTimeSlots", "SELECT id, schedule_id FROM time_slots WHERE id=$1", id)
	if err != nil {
		return TimeSlotsOutputRepo{}, err
	}

	timeSlots := TimeSlotsOutputRepo{}
	if err = row.Scan(&timeSlots.ID, &timeSlots.ScheduleID); err != nil {
		if err == sql.ErrNoRows {
			return TimeSlotsOutputRepo{}, ErrTimeSlotsNotFound
		}
		return TimeSlotsOutputRepo{}, err
	}

	return timeSlots, nil
}

// // CheckTimeSlotsExists checks whether the specified timeSlots exists by title (true == exist)
// func (r *CommiteeRepo) IsTimeSlotsExists(ctx context.Context, title string, classroomID int) (bool, error) {
// 	var exists bool
// 	row, err := QueryRowSQL(ctx, r.Database, "IsTimeSlotsExists", "SELECT EXISTS(SELECT 1 FROM thesis_timeSlots WHERE title LIKE '%' || $1 || '%' AND classtime_slots_id=$2)", title, classroomID)
// 	if err != nil {
// 		return false, err
// 	}
// 	if err = row.Scan(&exists); err != nil {
// 		return false, err
// 	}
// 	return exists, nil
// }

// // UpdateTimeSlots updates the specified timeSlots by id
// func (r *CommiteeRepo) UpdateTimeSlots(ctx context.Context, id int, timeSlots TimeSlotsInputRepo) error {
// 	result, err := ExecSQL(ctx, r.Database, "UpdateTimeSlots", "UPDATE thesis_timeSlots SET start_date=$2, period=$3, time_slots_id=$4, time=$5 WHERE id=$1", id, timeSlots.StartDate, timeSlots.Period, timeSlots.TimeSlotsID, timeSlots.Time)
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
// 		return ErrTimeSlotsNotFound
// 	}

// 	return nil
// }

// // DeleteTimeSlots deletes a timeSlots in db given by id
// func (r *CommiteeRepo) DeleteTimeSlots(ctx context.Context, id int) error {
// 	result, err := ExecSQL(ctx, r.Database, "DeleteTimeSlots", "DELETE FROM thesis_timeSlots WHERE id=$1", id)
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
// 		return ErrTimeSlotsNotFound
// 	}

// 	return nil
// }

// GetTimeSlots returns a list of timeSlots in db with filter
func (r *CommiteeRepo) GetTimeSlotss(ctx context.Context) ([]TimeSlotsOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetTimeSlotss", "SELECT id, schedule_id FROM time_slots")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the timeSlots slice
	var timeSlotsList []TimeSlotsOutputRepo
	for rows.Next() {
		timeSlots := TimeSlotsOutputRepo{}
		err := rows.Scan(
			&timeSlots.ID,
			&timeSlots.ScheduleID,
		)
		if err != nil {
			return nil, 0, err
		}
		timeSlotsList = append(timeSlotsList, timeSlots)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getTimeSlotsCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return timeSlotsList, count, nil
}

func (r *CommiteeRepo) getTimeSlotsCount(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", "SELECT COUNT(*) FROM time_slots")
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// GetTimeSlots returns a timeSlots in db given by id
func (r *CommiteeRepo) GetTimeSlotsByScheduleID(ctx context.Context, scheduleID int) ([]TimeSlotsOutputRepo, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetTimeSlotsByScheduleID", "SELECT id, schedule_id FROM time_slots WHERE schedule_id = $1", scheduleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the timeSlots slice
	var timeSlotsList []TimeSlotsOutputRepo
	for rows.Next() {
		timeSlots := TimeSlotsOutputRepo{}
		err := rows.Scan(
			&timeSlots.ID,
			&timeSlots.ScheduleID,
		)
		if err != nil {
			return nil, err
		}
		timeSlotsList = append(timeSlotsList, timeSlots)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return timeSlotsList, nil
}
