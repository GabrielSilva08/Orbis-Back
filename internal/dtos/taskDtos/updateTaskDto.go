package taskdtos

import (
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/google/uuid"
)

type UpdateTaskDto struct {
	Title       *string          `json:"title,omitempty"`
	Description *string          `json:"description,omitempty"`
	Deadline    *string          `json:"deadLine,omitempty" validate:"omitempty,datetime=2006-01-02T15:04:05Z07:00"`
	Priority    *models.Priority `json:"priority,omitempty" validate:"omitempty,oneof=Low Medium High"`
	Progress    *bool            `json:"progress,omitempty"`
	TagID       *uuid.UUID       `json:"tagId,omitempty"`
	UserID      uuid.UUID        `json:"userId" validate:"required"`
	ColumnID    *uuid.UUID       `json:"columnId,omitempty"`
}
