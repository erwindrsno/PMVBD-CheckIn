package subgrade

import "errors"

var (
	ErrInternal             = errors.New("Something went wrong!!!")
	ErrUnsupportedOperation = errors.New("Unsupported Operation")
)
