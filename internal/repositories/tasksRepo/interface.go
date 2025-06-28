package tasksRepo

import (
	taskdtos "github.com/GabrielSilva08/Orbis/internal/dtos/taskDtos"
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/google/uuid"
)

type TaskRepositoryInterface interface {
	Create(task models.Task) (models.Task, error) // Cria task passando um objeto task
	ListAllTasks() ([]models.Task, error)             // Lista todas as tasks
	GetTaskByID(id uuid.UUID) (models.Task, error)    // Pega uma task pelo ID dela
	DeleteTaskByID(id uuid.UUID) error                    // Deleta uma task pelo ID
	Update(request taskdtos.UpdateTaskDto) (models.Task, error)
}
