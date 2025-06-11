package tasksService

import (
	"github.com/GabrielSilva08/Orbis/internal/models"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tasksRepo"
	"github.com/google/uuid"
)

type TaskService struct {
	repo tasksRepo.TaskRepositoryInterface
}

func NewTaskService(repo tasksRepo.TaskRepositoryInterface) TaskServiceInterface {
	return &TaskService{repo: repo}
}

func (service TaskService) Create(task models.Task) (models.Task, error) {
	return service.repo.Create(task)
}

func (service TaskService) ListAllTasks() ([]models.Task, error) {
	return service.repo.ListAllTasks()
}

func (service TaskService) GetTaskByID(id uuid.UUID) (models.Task, error) {
	return service.repo.GetTaskByID(id)
}

func (service TaskService) DeleteTaskByID(id uuid.UUID) error {
	return service.repo.DeleteTaskByID(id)
}
