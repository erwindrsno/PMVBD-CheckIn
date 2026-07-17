package event

import (
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type Store interface {
	Insert(payload Event) error
	GetListView(filter EventFilter) ([]EventViewItem, error)
	GetPaginatedListView(limit, offset int) ([]EventViewItem, error)
	UpdateStatus(id uuid.UUID) error
	Delete(id uuid.UUID) error
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (s *SQLiteStore) Insert(payload Event) error {
	query := `
		INSERT INTO events (id, name, status, created_at, started_at)
		VALUES(?, ?, ?, ?, ?)
	`
	_, err := s.db.Exec(query, payload.Id.String(), payload.Name, payload.Status, time.Now().Format("2006-01-02 15:04:05"), nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteStore) GetPaginatedListView(limit, offset int) ([]EventViewItem, error) {
	query := `
		SELECT e.id, e.name, e.created_at, e.started_at, e.status
		FROM events e
		ORDER BY e.created_at -- Always include ORDER BY when using LIMIT
		LIMIT ? OFFSET ?
	`

	var es []EventViewItem
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		slog.Error("Getting list view in repo", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id uuid.UUID
		var name, createdAtStr string
		var status int
		var startedAtNull sql.NullString

		if err := rows.Scan(&id, &name, &createdAtStr, &startedAtNull, &status); err != nil {
			return nil, err
		}

		// Store now only handles data retrieval, not transformation
		es = append(es, mapRowToEvent(id, name, createdAtStr, startedAtNull, status))
	}
	return es, nil
}

func (s *SQLiteStore) GetListView(filter EventFilter) ([]EventViewItem, error) {
	query := `
        SELECT e.id, e.name, e.created_at, e.started_at, e.status
        FROM events e
    `
	var args []interface{}

	// 1. Add WHERE clause if filter exists
	if filter.Status != nil {
		query += " WHERE e.status = ?"
		args = append(args, filter.Status)
	}

	// 2. Add ORDER BY (put this after the WHERE clause)
	query += " ORDER BY e.created_at DESC" // Added DESC for better UX

	// 3. PASS ARGS HERE!
	rows, err := s.db.Query(query, args...) // <-- This is the missing piece
	if err != nil {
		slog.Error("Getting list view in repo", "error", err)
		return nil, err
	}
	defer rows.Close()

	var es []EventViewItem
	for rows.Next() {
		var id uuid.UUID
		var name, createdAtStr string
		var status int
		var startedAtNull sql.NullString

		if err := rows.Scan(&id, &name, &createdAtStr, &startedAtNull, &status); err != nil {
			return nil, err
		}

		// Store now only handles data retrieval, not transformation
		es = append(es, mapRowToEvent(id, name, createdAtStr, startedAtNull, status))
	}

	return es, nil
}

// func (s *SQLiteStore) GetListViewByOpen() ([]Event, error) {
// 	query := `
// 		SELECT e.id, e.name, e.created_at, e.started_at, e.status
// 		FROM events e
// 		WHERE e.status == 1
// 		ORDER BY e.created_at -- Always include ORDER BY when using LIMIT
// 	`
//
// 	var es []Event
// 	rows, err := s.db.Query(query)
// 	if err != nil {
// 		slog.Error("Getting list view in repo", "error", err)
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var id uuid.UUID
// 		var name, createdAtStr, status string
// 		var startedAtNull sql.NullString
//
// 		if err := rows.Scan(&id, &name, &createdAtStr, &startedAtNull, &status); err != nil {
// 			return nil, err
// 		}
//
// 		// Store now only handles data retrieval, not transformation
// 		es = append(es, mapRowToEvent(id, name, createdAtStr, startedAtNull, status))
// 	}
// 	return es, nil
// }

func (s *SQLiteStore) UpdateStatus(id uuid.UUID) error {
	// 1. Get current status
	var currentStatus Status
	err := s.db.QueryRow("SELECT status FROM events WHERE id = ?", id.String()).Scan(&currentStatus)
	if err != nil {
		return err
	}

	// 2. Increment status
	nextStatus := currentStatus + 1

	// 3. Prepare query
	// If next status is Open (assuming 1), we set started_at
	if nextStatus == Open {
		query := `
            UPDATE events 
            SET status = ?, started_at = ? 
            WHERE id = ?
        `
		_, err = s.db.Exec(query, int(nextStatus), time.Now().Format("2006-01-02 15:04:05"), id.String())
	} else {
		query := `
            UPDATE events 
            SET status = ? 
            WHERE id = ?
        `
		_, err = s.db.Exec(query, int(nextStatus), id.String())
	}

	return err
}

func (s *SQLiteStore) Delete(id uuid.UUID) error {
	// 1. Prepare the SQL statement
	query := `DELETE FROM events WHERE id = ?`

	// 2. Execute the query
	// We convert the UUID to a string to match the SQLite storage format
	result, err := s.db.Exec(query, id.String())
	if err != nil {
		return err
	}

	// 3. Optional: Check if a row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		// This can happen if the ID doesn't exist
		return fmt.Errorf("no event found with id %s", id.String())
	}

	return nil
}
