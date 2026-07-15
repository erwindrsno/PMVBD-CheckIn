package attendee

import (
	"log/slog"
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

func (s *Service) Create(name string) error {
	return nil
}

func (s *Service) Read() ([]AttendeeViewItem, error) {
	res, err := s.st.GetListView()
	if err != nil {
		slog.Error("Database failure in Service.Read", "error", err)
		return nil, ErrInternal
	} else {
		return res, nil
	}
}
