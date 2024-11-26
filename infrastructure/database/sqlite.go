package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteConnection(path string) *sql.DB {
	db, err := sql.Open(("sqlite3"), path)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	return db
}
