package event

type Status int

const (
	New = iota
	Open
	Closed
	Unknown
)

// String implements the fmt.Stringer interface.
func (s Status) String() string {
	switch s {
	case New:
		return "New"
	case Open:
		return "Open"
	case Closed:
		return "Completed"
	default:
		return "Unknown"
	}
}
