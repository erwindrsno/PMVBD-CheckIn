package event

type Status int

const (
	New = iota + 1
	Open
	Completed
	Unknown
)

// String implements the fmt.Stringer interface.
func (s Status) String() string {
	switch s {
	case New:
		return "New"
	case Open:
		return "Open"
	case Completed:
		return "Completed"
	default:
		return "Unknown"
	}
}
