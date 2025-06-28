package tasksService

import (
	taskdtos "github.com/GabrielSilva08/Orbis/internal/dtos/taskDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/google/uuid"
)

type TaskServiceInterface interface {
	Create(task taskdtos.CreateTaskDto) (taskdtos.CreateTaskResponse, error)
	ListAllTasks() ([]models.Task, error)
	GetTaskByID(id uuid.UUID) (models.Task, error)
	DeleteTaskByID(id uuid.UUID) error
	Update(task taskdtos.UpdateTaskDto) (models.Task, error)
}
