package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	StartedAt time.Time `json:"started_at"`
}
