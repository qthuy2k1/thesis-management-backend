package repository

import (
	"context"
	"database/sql"
	"log"
)

// QueryRowSQL is a wrapper function that logs the SQL command before executing it.
func QueryRowSQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (*sql.Row, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided arguments
	row := stmt.QueryRowContext(ctx, args...)

	return row, nil
}

// QuerySQL is a wrapper function that logs the SQL command before executing it.
func QuerySQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (*sql.Rows, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)

	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL statement with the provided arguments
	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		log.Printf("Error executing SQL command: %s", err.Error())
		return nil, err
	}

	return rows, nil
}

// ExecSQL is a wrapper function that logs the SQL command before executing it.
func ExecSQL(ctx context.Context, db *sql.DB, funcName string, query string, args ...interface{}) (sql.Result, error) {
	log.Printf("Function \"%s\" is executing SQL command: %s", funcName, query)
	// Prepare the SQL statement
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error preparing SQL statement: %s", err.Error())
		return nil, err
	}
	defer stmt.Close()

	// Execute the SQL command with the provided arguments
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Printf("Error executing SQL command: %s", err.Error())
		return nil, err
	}

	return result, nil
}

type TopicInputRepo struct {
	Title          string
	TypeTopic      string
	MemberQuantity int
	StudentID      string
	MemberEmail    string
	Description    string
}

// CreateClasroom creates a new topic in db given by topic model
func (r *TopicRepo) CreateTopic(ctx context.Context, topic TopicInputRepo) error {
	if _, err := ExecSQL(ctx, r.Database, "CreateTopic", "INSERT INTO topics (title, type_topic, member_quantity, student_id, member_email, description) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", topic.Title, topic.TypeTopic, topic.MemberQuantity, topic.StudentID, topic.MemberEmail, topic.Description); err != nil {
		return err
	}

	return nil
}

type TopicOutputRepo struct {
	ID             int
	Title          string
	TypeTopic      string
	MemberQuantity int
	StudentID      string
	MemberEmail    string
	Description    string
}

// GetTopic returns a topic in db given by id
func (r *TopicRepo) GetTopic(ctx context.Context, id int) (TopicOutputRepo, error) {
	row, err := QueryRowSQL(ctx, r.Database, "GetTopic", "SELECT id, title, type_topic, member_quantity, student_id, member_email, description FROM topics WHERE id=$1", id)
	if err != nil {
		return TopicOutputRepo{}, err
	}
	topic := TopicOutputRepo{}

	if err = row.Scan(&topic.ID, &topic.Title, &topic.TypeTopic, &topic.MemberQuantity, &topic.StudentID, &topic.MemberEmail, &topic.Description); err != nil {
		if err == sql.ErrNoRows {
			return TopicOutputRepo{}, ErrTopicNotFound
		}
		return TopicOutputRepo{}, err
	}

	return topic, nil
}

// UpdateTopic updates the specified topic by id
func (r *TopicRepo) UpdateTopic(ctx context.Context, id int, topic TopicInputRepo) error {
	result, err := ExecSQL(ctx, r.Database, "UpdateTopic", "UPDATE topics SET title=$2, type_topic=$3, member_quantity=$4, student_id=$5, member_email=$6, description=$7 WHERE id=$1", id, topic.Title, topic.TypeTopic, topic.MemberQuantity, topic.StudentID, topic.MemberEmail, topic.Description)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrTopicNotFound
	}

	return nil
}

// DeleteTopic deletes a topic in db given by id
func (r *TopicRepo) DeleteTopic(ctx context.Context, id int) error {
	result, err := ExecSQL(ctx, r.Database, "DeleteTopic", "DELETE FROM topics WHERE id=$1", id)
	if err != nil {
		return err
	}

	if rowsAff, _ := result.RowsAffected(); rowsAff == 0 {
		return ErrTopicNotFound
	}

	return nil
}
