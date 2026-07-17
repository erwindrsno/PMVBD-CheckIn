package event

import (
	"database/sql"

	"github.com/google/uuid"
)

// Define a mapper function that knows how to convert the raw database types to your domain model
func mapRowToEvent(id uuid.UUID, name string, createdAtStr string, startedAtNull sql.NullString, statusInt int) EventViewItem {
	var startedAt string
	if startedAtNull.Valid {
		startedAt = startedAtNull.String
	}
	status := Status(statusInt)

	return EventViewItem{
		Id:        id,
		Name:      name,
		CreatedAt: createdAtStr,
		StartedAt: startedAt,
		Status:    status,
	}
}
