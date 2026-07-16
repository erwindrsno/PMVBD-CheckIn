package attendance

import "errors"

var (
	ErrInternal            = errors.New("Something went wrong!!!")
	ErrDuplicateAttendance = errors.New("Duplicate attendance!")

// UNIQUE constraint failed: attendances.attendee_id, attendances.event_id
)
