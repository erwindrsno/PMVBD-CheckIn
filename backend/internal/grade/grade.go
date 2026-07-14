package grade

import (
	"database/sql"
	"fmt"
	"log/slog"
)

type Grade struct {
	Id             int
	NumericalGrade int
	LevelId        int
	Label          string
}

type Store interface {
	Insert(g *Grade) error
	GetAll() ([]Grade, error)
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

// TODO: Should return unsupported operation
func (r *SQLiteStore) Insert(g *Grade) error {
	return nil
}

func (r *SQLiteStore) GetAll() ([]Grade, error) {
	query := `
		SELECT id, numerical_grade, level_id, label
		FROM grades
		ORDER BY id ASC
	`
	var grades []Grade
	rows, err := r.db.Query(query)
	if err != nil {
		slog.Error("Getting list view in repo", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var grade Grade
		if err := rows.Scan(&grade.Id, &grade.NumericalGrade, &grade.LevelId, &grade.Label); err != nil {
			slog.Error("Getting list view in repo", "error", err)
			return nil, err
		}
		grades = append(grades, grade)
	}
	fmt.Println("getting grades ok")
	return grades, nil
}
