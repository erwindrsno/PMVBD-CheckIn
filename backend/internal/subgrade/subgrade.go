package subgrade

import (
	"database/sql"
	"log/slog"
)

type Subgrade struct {
	Id      int
	LevelId int
	Name    string
}

type Store interface {
	Insert(sg *Subgrade) error
	GetAll() ([]Subgrade, error)
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
func (r *SQLiteStore) Insert(g *Subgrade) error {
	return nil
}

func (r *SQLiteStore) GetAll() ([]Subgrade, error) {
	query := `
		SELECT id, level_id, name
		FROM subgrades
		ORDER BY id ASC
	`
	var subgrades []Subgrade
	rows, err := r.db.Query(query)
	if err != nil {
		slog.Error("Getting list view in repo", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var subgrade Subgrade
		if err := rows.Scan(&subgrade.Id, &subgrade.LevelId, &subgrade.Name); err != nil {
			slog.Error("Getting list view in repo", "error", err)
			return nil, err
		}
		subgrades = append(subgrades, subgrade)
	}
	return subgrades, nil
}
