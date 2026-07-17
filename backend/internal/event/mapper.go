package event

import (
	"database/sql"
	// "log/slog"
	"strconv"
	// "time"

	"github.com/google/uuid"
)

// Define a mapper function that knows how to convert the raw database types to your domain model
func mapRowToEvent(id uuid.UUID, name string, createdAtStr string, startedAtNull sql.NullString, statusStr string) EventViewItem {
	var startedAt string
	if startedAtNull.Valid {
		startedAt = startedAtNull.String
	}

	// 1. Convert status string to int
	statusInt, _ := strconv.Atoi(statusStr)

	// 2. Cast the int to your custom Status type
	status := Status(statusInt)

	return EventViewItem{
		Id:        id,
		Name:      name,
		CreatedAt: createdAtStr,
		StartedAt: startedAt,
		Status:    status,
	}
}
