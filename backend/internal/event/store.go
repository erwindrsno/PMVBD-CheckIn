package event

import (
	"database/sql"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id        uuid.UUID
	Name      string
	Status    Status
	CreatedAt time.Time
	StartedAt time.Time
}

type Store interface {
	Insert(payload Event) error
	GetListView() ([]Event, error)
	GetPaginatedListView(limit, offset int) ([]Event, error)
	UpdateStatus(id uuid.UUID, status Status) error
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (s *SQLiteStore) Insert(payload Event) error {
	query := `
		INSERT INTO events (id, name, status, created_at, started_at)
		VALUES(?, ?, ?, ?, ?)
	`
	_, err := s.db.Exec(query, payload.Id.String(), payload.Name, payload.Status, payload.CreatedAt, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteStore) GetPaginatedListView(limit, offset int) ([]Event, error) {
	query := `
		SELECT e.id, e.name, e.created_at, e.started_at, e.status
		FROM events e
		ORDER BY e.created_at -- Always include ORDER BY when using LIMIT
		LIMIT ? OFFSET ?
	`

	var es []Event
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		slog.Error("Getting list view in repo", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id uuid.UUID
		var name, createdAtStr, status string
		var startedAtNull sql.NullString

		if err := rows.Scan(&id, &name, &createdAtStr, &startedAtNull, &status); err != nil {
			return nil, err
		}

		// Store now only handles data retrieval, not transformation
		es = append(es, mapRowToEvent(id, name, createdAtStr, startedAtNull, status))
	}
	return es, nil
}

func (s *SQLiteStore) GetListView() ([]Event, error) {
	query := `
		SELECT e.id, e.name, e.created_at, e.started_at, e.status
		FROM events e
		ORDER BY e.created_at -- Always include ORDER BY when using LIMIT
	`

	var es []Event
	rows, err := s.db.Query(query)
	if err != nil {
		slog.Error("Getting list view in repo", "error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id uuid.UUID
		var name, createdAtStr, status string
		var startedAtNull sql.NullString

		if err := rows.Scan(&id, &name, &createdAtStr, &startedAtNull, &status); err != nil {
			return nil, err
		}

		// Store now only handles data retrieval, not transformation
		es = append(es, mapRowToEvent(id, name, createdAtStr, startedAtNull, status))
	}
	return es, nil
}

func (s *SQLiteStore) UpdateStatus(id uuid.UUID, status Status) error {
	var err error
	if status == Open {
		query := `
			UPDATE events 
			SET status = ?, started_at = ? 
			WHERE id = ?
		`
		_, err = s.db.Exec(query, int(status), time.Now().Format("2006-01-02 15:04:05"), id.String())
	} else {
		query := `
			UPDATE events 
			SET status = ? 
			WHERE id = ?
		`
		_, err = s.db.Exec(query, int(status), id.String())
	}
	return err
}
