package tasksService

import (
	"github.com/GabrielSilva08/Orbis/internal/models/tasksModel"
	"github.com/GabrielSilva08/Orbis/internal/repositories/tasksRepo"
)

type TaskListService struct {
	repo tasksRepo.TaskListRepositoryInterface
}

func NewTaskListService(repo tasksRepo.TaskListRepositoryInterface) TaskListServiceInterface {
	return &TaskListService{repo: repo}
}

func (service TaskListService) Create(taskList tasksModel.TaskList) (tasksModel.TaskList, error) {
	return service.repo.Create(taskList)
}
