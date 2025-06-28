package taskdtos

import (
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/google/uuid"
)

type CreateTaskDto struct {
	Title       string          `json:"title" validate:"required"`
	Description string          `json:"description" validate:"required"`
	DeadLine    string          `json:"deadLine" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Priority    models.Priority `json:"priority" validate:"required,oneof=Low Medium High"`
	Progress    bool            `json:"progress"`
	User        uuid.UUID       `json:"userId" validate:"required"`
	Column      uuid.UUID       `json:"columnId"`
}
