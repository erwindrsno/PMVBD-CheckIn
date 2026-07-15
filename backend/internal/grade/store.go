package grade

import (
	"database/sql"
)

type Grade struct {
	Id             int    `json:"id"`
	NumericalGrade int    `json:"numerical_grade"`
	LevelId        int    `json:"level_id"`
	Label          string `json:"label"`
}

type Store interface {
	Insert(g *Grade) error
	GetListView() ([]Grade, error)
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (r *SQLiteStore) Insert(g *Grade) error {
	return nil
}

func (r *SQLiteStore) GetListView() ([]Grade, error) {
	query := `
		SELECT id, numerical_grade, level_id, label
		FROM grades
		ORDER BY id ASC
	`
	var grades []Grade
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var grade Grade
		if err := rows.Scan(&grade.Id, &grade.NumericalGrade, &grade.LevelId, &grade.Label); err != nil {
			return nil, err
		}
		grades = append(grades, grade)
	}
	return grades, nil
}
