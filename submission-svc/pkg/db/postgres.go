package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

// Initialize initializes the connection to db in docker with given db url and returns the db connection
func Initialize(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return db, err
	}

	if err = db.Ping(); err != nil {
		return db, err
	}

	runDBMigration("file:///migrations", dbUrl)

	log.Println("Submissions database connection etablished")
	return db, nil
}

func runDBMigration(migrationURL string, dbUrl string) {
	migration, err := migrate.New(migrationURL, dbUrl)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		if errForce := migration.Force(1); errForce != nil {
			log.Fatal("failed to run migrate force:", errForce)
		}

		if errDown := migration.Down(); errDown != nil {
			log.Fatal("failed to run migrate down:", errDown)
		}

		if errUpSecond := migration.Up(); errUpSecond != nil && errUpSecond != migrate.ErrNoChange {
			log.Fatal("failed to run migrate up:", err)
		}
	}

	log.Println("db migrated successfully")
}
