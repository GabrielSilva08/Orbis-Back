package tasksRepo

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/google/uuid"
)

type TaskRepositoryInterface interface {
	Create(task tasksModel.Task) (tasksModel.Task, error) // Cria task passando um objeto task
	ListAllTasks() ([]tasksModel.Task, error)             // Lista todas as tasks
	GetTaskByID(id uuid.UUID) (tasksModel.Task, error)    // Pega uma task pelo ID dela
	DeleteTaskByID(id uuid.UUID) error                    // Deleta uma task pelo ID
}

type TagRepositoryInterface interface {
	Create(tag tasksModel.Tag) (tasksModel.Tag, error) // Cria uma tag
}
