package attendance

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
)

type Store interface {
	Insert(payload Attendance) error
	GetListViewByEventId(eventId uuid.UUID) ([]AttendanceViewItem, error)
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

func (s *SQLiteStore) GetListViewByEventId(eventId uuid.UUID) ([]AttendanceViewItem, error) {
	query := `
		SELECT 
		    at.id AS attendance_id, 
		    at.event_id AS event_id, 
		    at.scanned_at AS scanned_at,
		    a.name AS attendee_name, 
		    s.name AS school_name, 
		    g.label AS grade_label, 
		    sg.name AS subgrade_name,
		    e.name AS event_name
		FROM attendances at
		JOIN attendees a ON at.attendee_id = a.public_id
		JOIN events e ON e.id = at.event_id
		JOIN schools s ON s.id = a.school_id 
		JOIN grades g ON g.id = a.grade_id 
		JOIN subgrades sg ON sg.id = a.subgrade_id 
		WHERE at.event_id = ?
		ORDER BY at.scanned_at DESC`

	rows, err := s.db.Query(query, eventId.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var atvis []AttendanceViewItem
	for rows.Next() {
		var atvi AttendanceViewItem
		// Scan into your struct fields
		// Assuming your Attendance struct has these fields or nested objects
		err := rows.Scan(
			&atvi.AttendanceId, &atvi.EventId, &atvi.ScannedAt, &atvi.AttendeeName,
			&atvi.SchoolName, &atvi.GradeLabel, &atvi.SubgradeName, &atvi.EventName,
		)
		if err != nil {
			return nil, err
		}
		atvis = append(atvis, atvi)
	}

	if err = rows.Err(); err != nil {
		return nil, err // This catches errors that happened during the loop
	}

	return atvis, nil
}

func (s *SQLiteStore) Delete(id uuid.UUID) error {
	query := `
		DELETE FROM attendances
		WHERE id = ?
	`

	result, err := s.db.Exec(query, id.String())
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrResourceNotFound
	}

	return nil
}
