package attendee

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type Attendee struct {
	Id                    uuid.UUID
	PublicId              string
	SchoolId              int
	Name                  string
	IsMale                Gender
	GradeId               int
	SubGradeId            int
	ContactNumber         string
	GuardianContactNumber string
}

type Store interface {
	Insert(a Attendee) error
	GetPaginatedListView(limit, offset int) ([]AttendeeViewItem, error)
	GetListView() ([]AttendeeViewItem, error)
	GetByPublicId(id string) (*AttendeeViewItem, error)
	Delete(publicId string) error
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (s *SQLiteStore) Insert(a Attendee) error {
	query := `
		INSERT INTO attendees (
			id, public_id, school_id, name, is_male, grade_id, subgrade_id, contact_number, guardian_contact_number, created_at)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, datetime('now', '+7 hours'))
	`
	_, err := s.db.Exec(query, a.Id, a.PublicId, a.SchoolId, a.Name, a.IsMale, a.GradeId, a.SubGradeId, a.ContactNumber, a.GuardianContactNumber)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteStore) GetPaginatedListView(limit, offset int) ([]AttendeeViewItem, error) {
	query := `
		SELECT a.name, a.public_id, s.name, g.label, sg.name, a.contact_number, a.guardian_contact_number
		FROM attendees a
		JOIN schools s ON s.id = a.school_id
		JOIN grades g ON g.id = a.grade_id 
		JOIN subgrades sg ON sg.id = a.subgrade_id
		ORDER BY a.name -- Always include ORDER BY when using LIMIT
		LIMIT ? OFFSET ?
	`

	var avis []AttendeeViewItem
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var avi AttendeeViewItem
		if err := rows.Scan(&avi.Name, &avi.PublicId, &avi.SchoolName, &avi.Grade, &avi.Subgrade, &avi.ContactNumber, &avi.GuardianContactNumber); err != nil {
			return nil, err
		}
		avis = append(avis, avi)
	}
	return avis, nil
}

func (s *SQLiteStore) GetListView() ([]AttendeeViewItem, error) {
	query := `
		SELECT a.name, a.public_id, s.name, g.label, sg.name, a.contact_number, a.guardian_contact_number
		FROM attendees a
		JOIN schools s ON s.id = a.school_id
		JOIN grades g ON g.id = a.grade_id 
		JOIN subgrades sg ON sg.id = a.subgrade_id
		ORDER BY a.name -- Always include ORDER BY when using LIMIT
	`

	var avis []AttendeeViewItem
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var avi AttendeeViewItem
		if err := rows.Scan(&avi.Name, &avi.PublicId, &avi.SchoolName, &avi.Grade, &avi.Subgrade, &avi.ContactNumber, &avi.GuardianContactNumber); err != nil {
			return nil, err
		}
		avis = append(avis, avi)
	}
	return avis, nil
}

func (s *SQLiteStore) GetByPublicId(publicId string) (*AttendeeViewItem, error) {
	query := `
			SELECT a.name, a.public_id, s.name, g.label, sg.name, a.contact_number, a.guardian_contact_number
			FROM attendees a
			JOIN schools s ON s.id = a.school_id
			JOIN grades g ON g.id = a.grade_id 
			JOIN subgrades sg ON sg.id = a.subgrade_id
			WHERE a.public_id = ?
    `

	var avi AttendeeViewItem
	// Use QueryRow because we expect exactly one (or zero) results
	err := s.db.QueryRow(query, publicId).Scan(
		&avi.Name,
		&avi.PublicId,
		&avi.SchoolName,
		&avi.Grade,
		&avi.Subgrade,
		&avi.ContactNumber,
		&avi.GuardianContactNumber,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// It's helpful to return a custom error or nil if the attendee isn't found
			slog.Error("database no rows", "error", err.Error())
			return nil, ErrAttendeeNotExist
		}
		return nil, err
	}

	return &avi, nil
}

func (s *SQLiteStore) Delete(publicId string) error {
	// 1. Prepare the SQL statement
	query := `DELETE FROM attendees WHERE public_id = ?`

	// 2. Execute the query
	// We convert the UUID to a string to match the SQLite storage format
	result, err := s.db.Exec(query, publicId)
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
		return fmt.Errorf("no event found with id %s", publicId)
	}

	return nil
}
