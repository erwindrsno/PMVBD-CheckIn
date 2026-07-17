package event

import (
	"log/slog"

	"github.com/google/uuid"
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
	uuid, err := uuid.NewV7()
	if err != nil {
		slog.Error("Database failure in Service.Create", "error", err)
		return ErrInternal
	}
	payload := Event{
		Id:     uuid,
		Name:   name,
		Status: New,
	}
	// loc, err := time.LoadLocation("Asia/Jakarta")
	// if err != nil {
	// 	fmt.Println(err)
	// 	loc = time.UTC
	// }
	// now := time.Now().In(loc)
	// payload.CreatedAt = now.Truncate(time.Second)
	//
	// payload.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	if err := s.st.Insert(payload); err != nil {
		slog.Error("Database failure in Service.Read", "error", err)
		return err
	}
	return nil
}

func (s *Service) Read(filter EventFilter) ([]EventViewItem, error) {
	res, err := s.st.GetListView(filter)
	if err != nil {
		slog.Error("Database failure in Service.Read", "error", err)
		return nil, ErrInternal
	} else {
		return res, nil
	}
}

func (s *Service) UpdateStatus(id uuid.UUID) error {
	if err := s.st.UpdateStatus(id); err != nil {
		slog.Error("Database failure in Service.UpdateStatus", "error", err)
		return ErrInternal
	}
	return nil
}

func (s *Service) Delete(id uuid.UUID) error {
	if err := s.st.Delete(id); err != nil {
		slog.Error("Database failure in Service.Delete", "error", err)
		return ErrInternal
	}
	return nil
}
