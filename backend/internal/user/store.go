package user

import (
	"database/sql"
	"log/slog"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Store interface {
	Insert(user *User) error
	GetByUsername(username string) (*User, error)
}

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func (s *SQLiteStore) Insert(user *User) error {
	query := `
		INSERT INTO users (username, password)
		VALUES(?, ?)
	`
	_, err := s.db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteStore) GetByUsername(username string) (*User, error) {
	query := `
		SELECT username, password
		FROM users u
		WHERE u.username = ?
	`

	rows, err := s.db.Query(query, username)
	if err != nil {
		slog.Error("Getting user by username in repo", "error", err)
		return nil, err
	}
	defer rows.Close()
	var user User
	if rows.Next() {
		if err := rows.Scan(&user.Username, &user.Password); err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		return nil, err
	}
}
