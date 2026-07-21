package user

import (
	"log/slog"

	"github.com/erwindrsno/PMVBD-CheckIn/internal/auth"
)

type Service struct {
	st Store
}

func NewService(st Store) *Service {
	s := &Service{
		st: st,
	}
	return s
}

func (s *Service) Create(user *User) error {
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = s.st.Insert(user)
	if err != nil {
		slog.Error("Database failure in Service.Create", "error", err)
		return ErrInternal
	}
	return nil
}
