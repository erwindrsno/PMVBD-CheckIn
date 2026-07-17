package attendance

import "github.com/google/uuid"

// handler - service DTO
type AttendanceCapturePayload struct {
	AttendeeId string    `json:"attendee_id"`
	EventId    uuid.UUID `json:"event_id"`
}

type AttendanceViewItem struct {
	AttendanceId uuid.UUID `json:"attendance_id"`
	EventId      uuid.UUID `json:"event_id"`
	EventName    string    `json:"event_name"`
	ScannedAt    string    `json:"scanned_at"`
	AttendeeName string    `json:"attendee_name"`
	SchoolName   string    `json:"school_name"`
	GradeLabel   string    `json:"grade_label"`
	SubgradeName string    `json:"subgrade_name"`
}
