package event

type EventFilter struct {
	Status *Status // Optional status filter
}

var statusMap = map[string]Status{
	"new":       New,
	"open":      Open,
	"completed": Completed,
	"unknown":   Unknown,
}
