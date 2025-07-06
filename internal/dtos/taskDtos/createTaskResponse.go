package taskdtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateTaskResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DeadLine    time.Time `json:"deadLine"`
	Priority    string    `json:"priority"`
	Progress    bool      `json:"progress"`
	TagID       *uuid.UUID `json:"tagId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
