package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func OpenDbConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@0.0.0.0:5433/user-mgmt-db?sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to reach database: %v", err)
	}

	return db
}
