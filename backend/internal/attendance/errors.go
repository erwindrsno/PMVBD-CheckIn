package attendance

import "errors"

var (
	ErrInternal            = errors.New("Something went wrong!!!")
	ErrDuplicateAttendance = errors.New("Duplicate attendance!")
	ErrResourceNotFound    = errors.New("Not found!")

// UNIQUE constraint failed: attendances.attendee_id, attendances.event_id
)
