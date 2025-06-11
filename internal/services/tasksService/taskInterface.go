package tasksService

import (
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/google/uuid"
)

type TaskServiceInterface interface {
	Create(task models.Task) (models.Task, error)
	ListAllTasks() ([]models.Task, error)
	GetTaskByID(id uuid.UUID) (models.Task, error)
	DeleteTaskByID(id uuid.UUID) error
}
