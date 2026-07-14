package school

import (
	"database/sql"
)

type School struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Store interface {
	Insert(name string) error
	GetAll() ([]School, error)
	GetAllNames() ([]string, error)
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (s *SQLiteStore) Insert(name string) error {
	query := `
		INSERT INTO schools (name)
		VALUES(?)
	`
	_, err := s.db.Exec(query, name)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteStore) GetAll() ([]School, error) {
	query := `
		SELECT id, name
		FROM schools
		ORDER BY id ASC
	`
	var schools []School
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var school School
		if err := rows.Scan(&school.Id, &school.Name); err != nil {
			return nil, err
		}
		schools = append(schools, school)
	}
	return schools, nil
}

func (s *SQLiteStore) GetAllNames() ([]string, error) {
	query := `
		SELECT name
		FROM schools
		ORDER BY id ASC
	`
	var names []string
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}
