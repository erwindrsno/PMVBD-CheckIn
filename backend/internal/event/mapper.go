package event

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Define a mapper function that knows how to convert the raw database types to your domain model
func mapRowToEvent(id uuid.UUID, name string, createdAtStr string, startedAtNull sql.NullString, statusStr string) Event {
	layout := "2006-01-02 15:04:05Z07:00"

	createdAt, _ := time.Parse(layout, createdAtStr)

	var startedAt time.Time
	if startedAtNull.Valid {
		startedAt, _ = time.Parse(layout, startedAtNull.String)
	}

	// 1. Convert status string to int
	statusInt, _ := strconv.Atoi(statusStr)

	// 2. Cast the int to your custom Status type
	status := Status(statusInt)

	return Event{
		Id:        id,
		Name:      name,
		CreatedAt: createdAt,
		StartedAt: startedAt,
		Status:    status,
	}
}
