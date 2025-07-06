package taskdtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/GabrielSilva08/Orbis/internal/models"
)

type UpdateTaskDto struct {
	TaskID      uuid.UUID        `json:"taskId" validate:"required"`
	Title       *string          `json:"title"`
	Description *string          `json:"description"`
	Deadline    *time.Time       `json:"deadLine"`
	Priority    *models.Priority `json:"priority"`
	Progress    *bool            `json:"progress"`
	TagID       *uuid.UUID       `json:"tagId"`
	UserID      *uuid.UUID       `json:"userId"`
}
