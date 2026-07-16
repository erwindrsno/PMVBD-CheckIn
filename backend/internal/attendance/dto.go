package attendance

import "github.com/google/uuid"

// handler - service DTO
type AttendanceCapturePayload struct {
	AttendeeId string    `json:"attendee_id"`
	EventId    uuid.UUID `json:"event_id"`
}
