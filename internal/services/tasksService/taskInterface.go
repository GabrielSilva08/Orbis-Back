package tasksService

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/google/uuid"
)

type TaskServiceInterface interface {
	Create(task tasksModel.Task) (tasksModel.Task, error)
	ListAllTasks() ([]tasksModel.Task, error)
	GetTaskByID(id uuid.UUID) (tasksModel.Task, error)
	DeleteTaskByID(id uuid.UUID) error
}
