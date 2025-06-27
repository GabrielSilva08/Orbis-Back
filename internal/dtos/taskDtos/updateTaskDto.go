package taskdtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/GabrielSilva08/Orbis/internal/models"
)

type UpdateTaskDto struct {
	TaskID      uuid.UUID     `json:"id" validate:"required"`
	Title       *string        `json:"title" validate:"required"`
	Description *string        `json:"description"`
	Deadline    time.Time     `json:"deadLine" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Priority    models.Priority `json:"priority" validate:"required,oneof=Low Medium High"`
	Progress    *bool          `json:"progress"`
	TagID       uuid.UUID     `json:"tagId"`
	UserID      uuid.UUID     `json:"userId" validate:"required"`
	ColumnID    uuid.UUID     `json:"columnId"`
}