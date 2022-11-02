package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func OpenDbConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgresql://postgres:postgres@user-mgmt-db:5432/user-mgmt-db?sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to reach database: %v", err)
	}

	return db
}
