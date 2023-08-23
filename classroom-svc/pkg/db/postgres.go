package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Initialize initializes the connection to db in docker with given db url and returns the db connection
func Initialize(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=postgres user=postgres password=root dbname=thesis_management_classrooms port=5432 sslmode=disable")
	if err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	log.Println("Database connection etablished")
	return db, nil
}
