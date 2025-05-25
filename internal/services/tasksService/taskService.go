package tasksService

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tasksRepo"
)

type TaskService struct {
	repo tasksRepo.TaskRepositoryInterface
}

func NewTaskService(repo tasksRepo.TaskRepositoryInterface) TaskServiceInterface {
	return &TaskService{repo: repo}
}

func (service TaskService) Create(task tasksModel.Task) (tasksModel.Task, error) {
	return service.repo.Create(task)
}
