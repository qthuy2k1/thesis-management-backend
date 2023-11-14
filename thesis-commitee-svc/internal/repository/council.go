package repository

import (
	"context"
	"database/sql"
)

type CouncilInputRepo struct {
	LecturerID string
	ThesisID   int
}

type CouncilOutputRepo struct {
	ID         int
	LecturerID string
	ThesisID   int
}

// CreateCouncil creates a new council in db given by council model
func (r *CommiteeRepo) CreateCouncil(ctx context.Context, p CouncilInputRepo) (CouncilOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateCouncil", "INSERT INTO councils (lecturer_id, thesis_id) VALUES ($1, $2) RETURNING id, lecturer_id, thesis_id", p.LecturerID, p.ThesisID)
	if err != nil {
		return CouncilOutputRepo{}, err
	}

	var councilOutput CouncilOutputRepo
	if err := row.Scan(&councilOutput.ID, &councilOutput.LecturerID, &councilOutput.ThesisID); err != nil {
		return CouncilOutputRepo{}, err
	}

	return councilOutput, nil
}

// GetCouncil returns a council in db given by id
func (r *CommiteeRepo) GetCouncil(ctx context.Context, id int) (CouncilOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetCouncil", "SELECT id, lecturer_id, thesis_id FROM councils WHERE id=$1", id)
	if err != nil {
		return CouncilOutputRepo{}, err
	}

	council := CouncilOutputRepo{}
	if err = row.Scan(&council.ID, &council.LecturerID, &council.ThesisID); err != nil {
		if err == sql.ErrNoRows {
			return CouncilOutputRepo{}, ErrCouncilNotFound
		}
		return CouncilOutputRepo{}, err
	}

	return council, nil
}

// // CheckCouncilExists checks whether the specified council exists by title (true == exist)
// func (r *CommiteeRepo) IsCouncilExists(ctx context.Context, title string, classroomID int) (bool, error) {
// 	var exists bool
// 	row, err := QueryRowSQL(ctx, r.Database, "IsCouncilExists", "SELECT EXISTS(SELECT 1 FROM thesis_councils WHERE title LIKE '%' || $1 || '%' AND classtime_slots_id=$2)", title, classroomID)
// 	if err != nil {
// 		return false, err
// 	}
// 	if err = row.Scan(&exists); err != nil {
// 		return false, err
// 	}
// 	return exists, nil
// }

// // UpdateCouncil updates the specified council by id
// func (r *CommiteeRepo) UpdateCouncil(ctx context.Context, id int, council CouncilInputRepo) error {
// 	result, err := ExecSQL(ctx, r.Database, "UpdateCouncil", "UPDATE thesis_councils SET start_date=$2, period=$3, time_slots_id=$4, time=$5 WHERE id=$1", id, council.StartDate, council.Period, council.TimeSlotsID, council.Time)
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
// 		return ErrCouncilNotFound
// 	}

// 	return nil
// }

// // DeleteCouncil deletes a council in db given by id
// func (r *CommiteeRepo) DeleteCouncil(ctx context.Context, id int) error {
// 	result, err := ExecSQL(ctx, r.Database, "DeleteCouncil", "DELETE FROM thesis_councils WHERE id=$1", id)
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
// 		return ErrCouncilNotFound
// 	}

// 	return nil
// }

// GetCouncil returns a list of councils in db with filter
func (r *CommiteeRepo) GetCouncils(ctx context.Context) ([]CouncilOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetCouncils", "SELECT id, lecturer_id, thesis_id FROM councils")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the councils slice
	var councils []CouncilOutputRepo
	for rows.Next() {
		council := CouncilOutputRepo{}
		err := rows.Scan(
			&council.ID,
			&council.LecturerID,
			&council.ThesisID,
		)
		if err != nil {
			return nil, 0, err
		}
		councils = append(councils, council)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCouncilCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return councils, count, nil
}

func (r *CommiteeRepo) getCouncilCount(ctx context.Context) (int, error) {
	var count int

	rows, err := QueryRowSQL(ctx, r.Database, "getCount", "SELECT COUNT(*) FROM councils")
	if err != nil {
		return 0, err
	}

	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

// GetCouncil returns a list of councils in db with filter
func (r *CommiteeRepo) GetCouncilsByThesisID(ctx context.Context, thesisID int) ([]CouncilOutputRepo, int, error) {
	rows, err := QuerySQL(ctx, r.Database, "GetCouncils", "SELECT id, lecturer_id, thesis_id FROM councils WHERE thesis_id=$1", thesisID)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the councils slice
	var councils []CouncilOutputRepo
	for rows.Next() {
		council := CouncilOutputRepo{}
		err := rows.Scan(
			&council.ID,
			&council.LecturerID,
			&council.ThesisID,
		)
		if err != nil {
			return nil, 0, err
		}
		councils = append(councils, council)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCouncilCount(ctx)
	if err != nil {
		return nil, 0, err
	}

	return councils, count, nil
}
