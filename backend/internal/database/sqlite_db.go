package database

import (
	"database/sql"
	"fmt"

	// "fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	// 2. Create the tables (The "Schema")
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
				username TEXT unique,
				password TEXT
    );

		CREATE TABLE IF NOT EXISTS events(
			id TEXT PRIMARY KEY,
			name TEXT,
			status INTEGER,
			created_at TEXT,
			started_at TEXT
		);

    CREATE TABLE IF NOT EXISTS schools (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			UNIQUE(name)
    );

		CREATE TABLE IF NOT EXISTS attendees(
			id TEXT PRIMARY KEY,
			public_id TEXT UNIQUE,
			school_id INTEGER,
			name TEXT,
			is_male INTEGER,
			grade_id INTEGER,
			subgrade_id INTEGER,
			contact_number TEXT,
			guardian_contact_number TEXT,
			created_at TEXT,
			updated_at TEXT,
      FOREIGN KEY(school_id) REFERENCES schools(id)
      FOREIGN KEY(grade_id) REFERENCES grades(id)
      FOREIGN KEY(subgrade_id) REFERENCES subgrades(id)
		);

    CREATE TABLE IF NOT EXISTS attendances (
      id TEXT PRIMARY KEY,
			attendee_id TEXT,
			event_id TEXT,
			scanned_at TEXT,
      FOREIGN KEY(attendee_id) REFERENCES attendees(id),
      FOREIGN KEY(event_id) REFERENCES events(id),
			UNIQUE(attendee_id, event_id)
    );

		CREATE TABLE IF NOT EXISTS levels (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    name TEXT NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS grades (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    numerical_grade INTEGER NOT NULL,
		    level_id INTEGER NOT NULL,
		    FOREIGN KEY (level_id) REFERENCES levels(id)
		);
		
		CREATE TABLE IF NOT EXISTS subgrades (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    level_id INTEGER NOT NULL,
		    name TEXT NOT NULL,
		    FOREIGN KEY (level_id) REFERENCES levels(id)
		);

		
	`

	_, err = db.Exec(query)

	db.SetMaxOpenConns(1)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}
