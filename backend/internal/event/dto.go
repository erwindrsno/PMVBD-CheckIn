package event

import "github.com/google/uuid"

type EventViewItem struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Status    Status    `json:"status"`
	CreatedAt string    `json:"created_at"`
	StartedAt string    `json:"started_at"`
}
