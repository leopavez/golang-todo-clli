package database

import "database/sql"

func Migrate(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		isCompleted BOOLEAN,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}
