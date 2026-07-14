package school

import "log/slog"

type Service struct {
	st Store
}

func NewService(st Store) *Service {
	s := &Service{
		st: st,
	}
	return s
}

// func (s *Service) Create() ()

func (s *Service) Read() ([]School, error) {
	res, err := s.st.GetAll()
	if err != nil {
		slog.Error("Database failure in Service.Read", "error", err)
		return nil, ErrInternal
	} else {
		return res, nil
	}
}
