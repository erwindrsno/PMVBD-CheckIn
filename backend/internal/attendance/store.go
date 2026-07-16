package attendance

import (
	"database/sql"

	"github.com/mattn/go-sqlite3"
)

type Store interface {
	Insert(payload Attendance) error
	// GetListView() ([]Attendance, error)
	// GetListView(filter EventFilter) ([]Event, error)
	// GetPaginatedListView(limit, offset int) ([]Event, error)
	// UpdateStatus(id uuid.UUID) error
	// Delete(id uuid.UUID) error
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (s *SQLiteStore) Insert(payload Attendance) error {
	query := `
		INSERT INTO attendances (id, attendee_id, event_id, scanned_at)
		VALUES(?, ?, ?, ?)
	`
	_, err := s.db.Exec(query, payload.Id.String(), payload.AttendeeId, payload.EventId, payload.ScannedAt)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.Code == sqlite3.ErrConstraint {
				return ErrDuplicateAttendance
			}
		}
		return err
	}
	return nil
}
