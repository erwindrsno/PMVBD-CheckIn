package attendee

import "errors"

var (
	ErrInternal         = errors.New("Something went wrong!!!")
	ErrAttendeeNotExist = errors.New("Attendee not exist!")
)
