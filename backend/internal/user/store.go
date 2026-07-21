package user

import (
	"database/sql"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Store interface {
	Insert(user *User) error
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
