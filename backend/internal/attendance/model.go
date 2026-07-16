package attendance

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	Id         uuid.UUID `json:"id"`
	AttendeeId string    `json:"attendee_id"`
	EventId    uuid.UUID `json:"event_id"`
	ScannedAt  time.Time `json:"scanned_at"`
}
