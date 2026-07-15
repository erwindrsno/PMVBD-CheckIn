package subgrade

import (
	"database/sql"
)

type Subgrade struct {
	Id      int    `json:"id"`
	LevelId int    `json:"level_id"`
	Name    string `json:"name"`
}

type Store interface {
	Insert(sg *Subgrade) error
	GetListView() ([]Subgrade, error)
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (r *SQLiteStore) Insert(g *Subgrade) error {
	return nil
}

func (r *SQLiteStore) GetListView() ([]Subgrade, error) {
	query := `
		SELECT id, level_id, name
		FROM subgrades
		ORDER BY id ASC
	`
	var subgrades []Subgrade
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var subgrade Subgrade
		if err := rows.Scan(&subgrade.Id, &subgrade.LevelId, &subgrade.Name); err != nil {
			return nil, err
		}
		subgrades = append(subgrades, subgrade)
	}
	return subgrades, nil
}
