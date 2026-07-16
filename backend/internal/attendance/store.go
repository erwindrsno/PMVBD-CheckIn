package attendance

import "database/sql"

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
		return err
	}
	return nil
}
