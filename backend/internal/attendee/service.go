package attendee

import (
	"log/slog"
	"strings"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
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

func (s *Service) Create(req AddAttendeeForm) error {
	privateId, err := uuid.NewV7()
	if err != nil {
		slog.Error("faild to generate uuid", "error", err)
	}
	publicId, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz123456789", 8)
	if err != nil {
		slog.Error("faild to generate nanoid", "error", err)
	}
	a := Attendee{
		Id:                    privateId,
		PublicId:              publicId,
		SchoolId:              req.SchoolId,
		Name:                  strings.ToUpper(req.Name),
		IsMale:                parseGender(req.Gender),
		GradeId:               req.GradeId,
		SubGradeId:            req.SubgradeId,
		ContactNumber:         req.ContactNumber,
		GuardianContactNumber: req.GuardianContactNumber,
	}
	if err := s.st.Insert(a); err != nil {
		slog.Error("failed to insert new attendee", "error", err)
		return ErrInternal
	}
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
