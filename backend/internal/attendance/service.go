package attendance

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/erwindrsno/PMVBD-CheckIn/internal/attendee"
	"github.com/google/uuid"
)

// Define an interface for what the AttendanceService needs from AttendeeService
type AttendeeServiceReader interface {
	ReadByPublicId(publicID string) (*attendee.AttendeeViewItem, error)
}

type Service struct {
	st  Store
	asr AttendeeServiceReader
}

func NewService(st Store, asr AttendeeServiceReader) *Service {
	s := &Service{
		st:  st,
		asr: asr,
	}
	return s
}

func (s *Service) Capture(acp AttendanceCapturePayload) (*attendee.AttendeeViewItem, error) {
	//check whether the attendee exist
	attendee, err := s.asr.ReadByPublicId(acp.AttendeeId)
	if err != nil {
		slog.Error("error in capture attendance", "error", err)
		return nil, err
	}

	uuid, err := uuid.NewV7()
	if err != nil {
		slog.Error("failure in generating uuid", "error", err)
		return nil, ErrInternal
	}
	payload := Attendance{
		Id:         uuid,
		AttendeeId: acp.AttendeeId,
		EventId:    acp.EventId,
	}
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err)
		loc = time.UTC
	}
	now := time.Now().In(loc)
	payload.ScannedAt = now.Truncate(time.Second)
	if err := s.st.Insert(payload); err != nil {
		slog.Error("Database failure in capture attendance", "error", err)
		return nil, err
	}
	return attendee, nil
}

func (s *Service) ReadByEventId(req string) ([]AttendanceViewItem, error) {
	eventId, err := uuid.Parse(req)
	if err != nil {
		slog.Error("error during parse string to uuid in attendance read", "error", err.Error())
		return nil, err
	}
	atvis, err := s.st.GetListViewByEventId(eventId)
	if err != nil {
		slog.Error("error in get attendance list view", "error", err)
		return nil, err
	}
	return atvis, nil
}
