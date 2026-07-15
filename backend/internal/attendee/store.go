package attendee

import (
	"database/sql"

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
	ParentId              int
}

type Store interface {
	Insert(a Attendee) error
	GetPaginatedListView(limit, offset int) ([]AttendeeViewItem, error)
	GetListView() ([]AttendeeViewItem, error)
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (r *SQLiteStore) Insert(a Attendee) error {
	query := `
		INSERT INTO attendees (
			id, public_id, school_id, name, is_male, grade_id, subgrade_id, contact_number, guardian_contact_number, created_at)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, datetime('now', '+7 hours'))
	`
	_, err := r.db.Exec(query, a.Id, a.PublicId, a.SchoolId, a.Name, a.IsMale, a.GradeId, a.SubGradeId, a.ContactNumber, a.GuardianContactNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *SQLiteStore) GetPaginatedListView(limit, offset int) ([]AttendeeViewItem, error) {
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
	rows, err := r.db.Query(query, limit, offset)
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

func (r *SQLiteStore) GetListView() ([]AttendeeViewItem, error) {
	query := `
		SELECT a.name, a.public_id, s.name, g.label, sg.name, a.contact_number, a.guardian_contact_number
		FROM attendees a
		JOIN schools s ON s.id = a.school_id
		JOIN grades g ON g.id = a.grade_id 
		JOIN subgrades sg ON sg.id = a.subgrade_id
		ORDER BY a.name -- Always include ORDER BY when using LIMIT
	`

	var avis []AttendeeViewItem
	rows, err := r.db.Query(query)
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
