package attendance

import "errors"

var (
	ErrInternal            = errors.New("Something went wrong!!!")
	ErrAttendeeNotExist    = errors.New("Attendee not exist!")
	ErrDuplicateAttendance = errors.New("Duplicate attendance!")

// UNIQUE constraint failed: attendances.attendee_id, attendances.event_id
)
