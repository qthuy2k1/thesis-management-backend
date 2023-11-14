package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type RoomInputRepo struct {
	Name        string
	Type        string
	School      string
	Description string
}

type RoomOutputRepo struct {
	ID          int
	Name        string
	Type        string
	School      string
	Description string
}

// CreateRoom creates a new room in db given by room model
func (r *CommiteeRepo) CreateRoom(ctx context.Context, p RoomInputRepo) (RoomOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "CreateRoom", "INSERT INTO rooms (name, type, school, description) VALUES ($1, $2, $3, $4) RETURNING id, name, type, school, description", p.Name, p.Type, p.School, p.Description)
	if err != nil {
		return RoomOutputRepo{}, err
	}

	var roomOutput RoomOutputRepo
	if err := row.Scan(&roomOutput.ID, &roomOutput.Name, &roomOutput.Type, &roomOutput.School, &roomOutput.Description); err != nil {
		return RoomOutputRepo{}, err
	}

	return roomOutput, nil
}

// GetRoom returns a room in db given by id
func (r *CommiteeRepo) GetRoom(ctx context.Context, id int) (RoomOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetRoom", "SELECT id, name, type, school, description FROM rooms WHERE id=$1", id)
	if err != nil {
		return RoomOutputRepo{}, err
	}

	room := RoomOutputRepo{}
	if err = row.Scan(&room.ID, &room.Name, &room.Type, &room.School, &room.Description); err != nil {
		if err == sql.ErrNoRows {
			return RoomOutputRepo{}, ErrRoomNotFound
		}
		return RoomOutputRepo{}, err
	}

	return room, nil
}

// CheckRoomExists checks whether the specified room exists by title (true == exist)
func (r *CommiteeRepo) IsRoomExists(ctx context.Context, name string, school string) (bool, error) {
	var exists bool
	row, err := QueryRowSQL(ctx, r.Database, "IsRoomExists", "SELECT EXISTS(SELECT 1 FROM rooms WHERE name LIKE '%' || $1 || '%' AND school LIKE '%' || $2 || '%')", name, school)
	if err != nil {
		return false, err
	}
	if err = row.Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

// UpdateRoom updates the specified room by id
func (r *CommiteeRepo) UpdateRoom(ctx context.Context, id int, room RoomInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateRoom", "UPDATE rooms SET name=$2, type=$3, school=$4, description=$5 WHERE id=$1", id, room.Name, room.Type, room.School, room.Description)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrRoomNotFound
	}

	return nil
}

// DeleteRoom deletes a room in db given by id
func (r *CommiteeRepo) DeleteRoom(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteRoom", "DELETE FROM rooms WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrRoomNotFound
	}

	return nil
}

type RoomFilter struct {
	Name   *string
	Type   *string
	School *string
}

// GetRoom returns a list of rooms in db with filter
func (r *CommiteeRepo) GetRooms(ctx context.Context, filter RoomFilter) ([]RoomOutputRepo, int, error) {
	query := []string{"SELECT id, name, type, school, description FROM rooms"}

	isFilter := false
	var conditions []string
	if filter.Name != nil && strings.TrimSpace(*filter.Name) != "" {
		isFilter = true
		conditions = append(conditions, fmt.Sprintf("name LIKE '%%%s%%'", *filter.Name))
	}
	if filter.Type != nil && strings.TrimSpace(*filter.Type) != "" {
		isFilter = true
		conditions = append(conditions, fmt.Sprintf("type LIKE '%%%s%%'", *filter.Type))
	}
	if filter.School != nil && strings.TrimSpace(*filter.School) != "" {
		isFilter = true
		conditions = append(conditions, fmt.Sprintf("school LIKE '%%%s%%'", *filter.School))
	}
	log.Println(conditions)

	if isFilter {
		query = append(query, "WHERE ", strings.Join(conditions, " AND "))
	}

	rows, err := QuerySQL(ctx, r.Database, "GetRooms", strings.Join(query, " "))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the rooms slice
	var rooms []RoomOutputRepo
	for rows.Next() {
		room := RoomOutputRepo{}
		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.Type,
			&room.School,
			&room.Description,
		)
		if err != nil {
			return nil, 0, err
		}
		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	count, err := r.getCountRoom(ctx, conditions, isFilter)
	if err != nil {
		return nil, 0, err
	}

	return rooms, count, nil
}

func (r *CommiteeRepo) getCountRoom(ctx context.Context, conditions []string, isFilter bool) (int, error) {
	var count int

	if isFilter {
		rows, err := QueryRowSQL(ctx, r.Database, "getCount", fmt.Sprintf("SELECT COUNT(*) FROM rooms WHERE %s", strings.Join(conditions, " AND ")))
		if err != nil {
			return 0, err
		}
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}

	} else {
		rows, err := QueryRowSQL(ctx, r.Database, "getCount", "SELECT COUNT(*) FROM rooms")
		if err != nil {
			return 0, err
		}
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
	}

	return count, nil
}

// GetRoom returns a list of rooms in db with filter
func (r *CommiteeRepo) GetRoomsByID(ctx context.Context, id []string) ([]RoomOutputRepo, error) {
	query := []string{"SELECT id, name, type, school, description FROM rooms IN (%s)", strings.Join(id, ",")}

	rows, err := QuerySQL(ctx, r.Database, "GetRooms", strings.Join(query, " "))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the rooms slice
	var rooms []RoomOutputRepo
	for rows.Next() {
		room := RoomOutputRepo{}
		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.Type,
			&room.School,
			&room.Description,
		)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// count, err := r.getCountRoom(ctx)
	// if err != nil {
	// 	return nil, 0, err
	// }

	return rooms, nil
}
